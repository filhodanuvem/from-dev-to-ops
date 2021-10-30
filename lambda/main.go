package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Sprintf("printing this %s!", name.Name)
	return fmt.Sprintf("Lambda running %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
