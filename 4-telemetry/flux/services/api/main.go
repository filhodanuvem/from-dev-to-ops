package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/filhodanuvem/log-api/metric"
	tracex "github.com/filhodanuvem/log-api/trace"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const serviceName = "payment-api"
const successEventType = "PAYMENT_ORDER_SUCCEEDED"
const failureEventType = "PAYMENT_ORDER_FAILED"

var prop = propagation.TraceContext{}
var successMetric = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "payment_order_completed",
	Help: "Order completed",
}, []string{"amount", "x_trace_id", "event_type"})

var failMetric = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "payment_order_failed",
	Help: "Order failed",
}, []string{"amount", "x_trace_id", "event_type"})

type requestBody struct {
	Amount    int               `json:"amount"`
	PaymentID string            `json:"payment_id"`
	Headers   map[string]string `json:"headers"`
}

func paymentHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Header)
	ctx := context.Background()
	ctx = prop.Extract(ctx, propagation.HeaderCarrier(req.Header))
	ctx, span := otel.Tracer(tracex.ServiceName).Start(ctx, "Run")
	defer span.End()
	body, _ := ioutil.ReadAll(req.Body)
	var request requestBody
	if err := json.Unmarshal(body, &request); err != nil {
		span.RecordError(err)
		log.Printf("failed to unmarshal %s\n", body)
		return
	}

	status := "success"
	bmetric := successMetric
	eventType := successEventType
	if request.Amount%3 == 0 {
		eventType = failureEventType
		bmetric = failMetric
		status = "fail"
	}

	labels := metric.NewLabels(
		strconv.Itoa(request.Amount),
		request.Headers["x-trace-id"],
		eventType,
	)

	metric.Record(bmetric, labels)

	log.Printf("status=%s, Amount=%d Headers=%s\n", status, request.Amount, req.Header["traceparent"])
}

func main() {
	userPassword := os.Getenv("DB_PASSWORD")
	log.Printf("Running service with DB_PASSWORD=%s", userPassword)
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

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2222", nil)
	}()

	http.HandleFunc("/", paymentHandler)
	log.Println("Server running on 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
