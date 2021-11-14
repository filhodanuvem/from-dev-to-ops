package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://checkip.amazonaws.com/", nil)
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		return "Request error to https://checkip.amazonaws.com/", err
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error to read response", err
	}

	return fmt.Sprintf("Lambda running %s! Ip: %s ", name.Name, ip), nil
}

func main() {
	lambda.Start(HandleRequest)
}
