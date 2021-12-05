package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

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

	numbers := make(chan int)
	go func(js nats.JetStreamContext) {
		for {
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
				continue
			}

			if _, err := js.Publish(nats_subject, b); err != nil {
				log.Printf("Error on publishing to nats: %s\n", err)
				continue
			}

			log.Println(m)
		}
	}(js)

	for {
		number := rand.Intn(100)
		numbers <- number
		time.Sleep(3 * time.Second)
	}
}
