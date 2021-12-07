package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	dtrace "go.opentelemetry.io/otel/trace"
)

const tracerName = "producer"

var nats_url = os.Getenv("NATS_URL")
var nats_subject = os.Getenv("NATS_SUBJECT")

type message struct {
	Number  int               `json:"number"`
	Headers map[string]string `json:"headers"`
}

func main() {
	sc, err := nats.Connect(nats_url)
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}
	defer sc.Close()

	js, err := sc.JetStream()
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %s", nats_url, err)
	}

	log.Println("Running crazy producer...")

	exp, err := newExporter(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	numbers := make(chan int)
	go func(js nats.JetStreamContext) {
		for {
			var span dtrace.Span
			_, span = otel.Tracer(tracerName).Start(context.Background(), "Produce")

			u1 := uuid.NewV4()
			number := <-numbers
			m := message{
				Number: number,
				Headers: map[string]string{
					"x-trace-id": u1.String(),
				},
			}

			b, err := json.Marshal(m)
			if err != nil {
				log.Printf("Error on publishing to nats: %s\n", err)
				span.RecordError(err)
				span.End()
				continue
			}

			if _, err := js.Publish(nats_subject, b); err != nil {
				log.Printf("Error on publishing to nats: %s\n", err)
				span.RecordError(err)
				span.End()
				continue
			}

			span.SetAttributes(attribute.Int("number", m.Number))
			span.End()
			log.Println(m)
		}
	}(js)

	for {
		number := rand.Intn(100)
		numbers <- number
		time.Sleep(3 * time.Second)
	}
}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("PRODUCER"),
			semconv.ServiceVersionKey.String("v0.1.0"),
		),
	)
	return r
}
