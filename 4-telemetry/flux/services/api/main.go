package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
	log.Printf("Received request")
	body, _ := ioutil.ReadAll(req.Body)
	var request requestBody
	if err := json.Unmarshal(body, &request); err != nil {
		log.Printf("failed to unmarshal %s\n", body)
		return
	}

	labels := prometheus.Labels{
		"amount":     strconv.Itoa(request.Amount),
		"x_trace_id": request.Headers["x-trace-id"],
		"event_type": successEventType,
	}

	status := "success"
	metric := successMetric
	if request.Amount%3 == 0 {
		labels["event_type"] = failureEventType
		metric = failMetric
		status = "fail"
	}

	metric.With(labels)
	timer := prometheus.NewTimer(metric.With(labels))
	timer.ObserveDuration()

	log.Printf("status=%s, Amount=%d TraceID=%s\n", status, request.Amount, request.Headers["x-trace-id"])
}

func main() {

	http.HandleFunc("/", paymentHandler)

	log.Println("Server running on 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
