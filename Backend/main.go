package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handle(ctx context.Context, name MyEvent) (string, error) {
	log.Print("Called by ", name)
	return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
	lambda.Start(handle)
}