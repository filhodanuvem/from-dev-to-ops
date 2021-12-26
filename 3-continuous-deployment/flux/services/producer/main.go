package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/filhodanuvem/producer/metric"
	tracex "github.com/filhodanuvem/producer/trace"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const tracerName = "producer"
const eventType = "PAYMENT_ORDER_CREATED"

var nats_url = os.Getenv("NATS_URL")
var nats_subject = os.Getenv("NATS_SUBJECT")

var bmetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "payment_order_time_in_seconds",
	Help: "Duration time of an order creation",
}, metric.Labels)

type message struct {
	Amount    int               `json:"amount"`
	PaymentID string            `json:"payment_id"`
	Headers   map[string]string `json:"headers"`
}

func main() {
	prometheus.Register(bmetric)
	tp, err := tracex.NewProvider("http://jaeger-collector:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	defer func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)
	otel.SetTracerProvider(tp)

	sc, err := nats.Connect(nats_url)
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}
	defer sc.Close()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2222", nil)
	}()

	js, err := sc.JetStream()
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}

	log.Println("Running crazy producer...")

	numbers := make(chan int)
	go func(js nats.JetStreamContext) {
		for {
			amount := <-numbers
			publishPayment(ctx, js, amount)
		}
	}(js)

	for {
		number := rand.Intn(100)
		numbers <- number
		time.Sleep(3 * time.Second)
	}
}

func publishPayment(ctx context.Context, js nats.JetStreamContext, amount int) {
	ctx, span := otel.Tracer(tracex.ServiceName).Start(ctx, "Run")
	defer span.End()

	u1 := uuid.NewV4()
	m := message{
		Amount:    amount,
		PaymentID: u1.String(),
		Headers:   map[string]string{},
	}
	span.SetAttributes(
		attribute.Key("payment-id").String(m.PaymentID),
	)

	labels := metric.NewLabels(
		strconv.Itoa(m.Amount),
		m.Headers["x-trace-id"],
		eventType,
	)
	recorder := metric.NewRecorder().WithTimer(bmetric, labels)

	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("Error on publishing to nats: %s\n", err)
		return
	}

	if _, err := js.Publish(nats_subject, b); err != nil {
		log.Printf("Error on publishing to nats: %s\n", err)
		return
	}

	recorder.RecordDuration()

	log.Println(m)
}
