package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func sendReceiptHandler(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	log.Println(string(body))
}

func main() {
	http.HandleFunc("/", sendReceiptHandler)

	log.Println("Server running on 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
