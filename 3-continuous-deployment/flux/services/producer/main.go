package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
)

const tracerName = "producer"
const eventType = "PAYMENT_ORDER_CREATED"

var nats_url = os.Getenv("NATS_URL")
var nats_subject = os.Getenv("NATS_SUBJECT")

type message struct {
	Amount  int               `json:"amount"`
	Headers map[string]string `json:"headers"`
}

func main() {
	sc, err := nats.Connect(nats_url)
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}
	defer sc.Close()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2222", nil)
	}()

	metric := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "payment_order_created",
		Help: "Order created",
	}, []string{"amount", "x_trace_id", "event_type"})

	js, err := sc.JetStream()
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}

	log.Println("Running crazy producer...")

	numbers := make(chan int)
	go func(js nats.JetStreamContext) {
		for {
			u1 := uuid.NewV4()
			amount := <-numbers
			m := message{
				Amount: amount,
				Headers: map[string]string{
					"x-trace-id": u1.String(),
				},
			}

			b, err := json.Marshal(m)
			if err != nil {
				log.Printf("Error on publishing to nats: %s\n", err)
				continue
			}

			if _, err := js.Publish(nats_subject, b); err != nil {
				log.Printf("Error on publishing to nats: %s\n", err)
				continue
			}

			labels := prometheus.Labels{
				"amount":     strconv.Itoa(m.Amount),
				"x_trace_id": m.Headers["x-trace-id"],
				"event_type": eventType,
			}

			timer := prometheus.NewTimer(metric.With(labels))
			timer.ObserveDuration()

			log.Println(m)
		}
	}(js)

	for {
		number := rand.Intn(100)
		numbers <- number
		time.Sleep(3 * time.Second)
	}
}
