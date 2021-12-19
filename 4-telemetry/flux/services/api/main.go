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
	"go.opentelemetry.io/otel/sdk/trace"
)

const successEventType = "PAYMENT_ORDER_SUCCEEDED"
const failureEventType = "PAYMENT_ORDER_FAILED"

var successMetric = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "payment_order_completed",
	Help: "Order completed",
}, []string{"amount", "x_trace_id", "event_type"})

var failMetric = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "payment_order_failed",
	Help: "Order failed",
}, []string{"amount", "x_trace_id", "event_type"})

type requestBody struct {
	Amount  int               `json:"amount"`
	Headers map[string]string `json:"headers"`
}

func paymentHandler(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var request requestBody
	if err := json.Unmarshal(body, &request); err != nil {
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

	log.Printf("status=%s, Amount=%d TraceID=%s\n", status, request.Amount, request.Headers["x-trace-id"])
}

func main() {
	f, err := os.Create("traces.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	exp, err := tracex.NewExporter(f)
	if err != nil {
		log.Fatal(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(trace.NewResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
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
