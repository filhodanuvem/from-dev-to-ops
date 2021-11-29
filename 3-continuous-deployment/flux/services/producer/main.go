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
	sc, _ := nats.Connect(nats_url)
	defer sc.Close()

	log.Println("Running crazy producer...")

	messages := make(chan int)
	go func(sc *nats.Conn) {
		for {
			m := <-messages
			sc.Publish(nats_subject, []byte(string(m)))
			log.Println(m)
		}
	}(sc)

	for {
		number := rand.Intn(100)
		messages <- number
		time.Sleep(3 * time.Second)
	}
}
