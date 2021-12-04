package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var nats_url = os.Getenv("NATS_URL")
var nats_subject = os.Getenv("NATS_SUBJECT")

func main() {
	sc, err := nats.Connect(nats_url)
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %w", nats_url, err)
	}
	defer sc.Close()

	js, err := sc.JetStream()
	if err != nil {
		log.Fatalf("Couldn't connect to nats %s, err: %w", nats_url, err)
	}

	log.Println("Running crazy producer...")

	messages := make(chan int)
	go func(js nats.JetStreamContext) {
		for {
			m := <-messages
			if _, err := js.Publish(nats_subject, []byte(string(m))); err != nil {
				log.Printf("Error on publishing to nats: %w\n", err)
			}

			log.Println(m)
		}
	}(js)

	for {
		number := rand.Intn(100)
		messages <- number
		time.Sleep(3 * time.Second)
	}
}
