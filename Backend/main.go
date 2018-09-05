package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type myReturn struct {
	Response string `json:"response"`
}

type RequestBody struct {
	Name string `json:"name"`
}

func handle(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Called by ", name)
	log.Print("context ", ctx)
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}

	var body RequestBody
	jsonParseError := json.Unmarshal([]byte(name.Body), &body)
	if jsonParseError != nil {
		log.Println(jsonParseError)
		return events.APIGatewayProxyResponse{500, headers, "Internal Server Error", false}, nil
	}

	code := 200
	response, jsonBuildError := json.Marshal(myReturn{Response: "Hello, " + body.Name})
	if jsonBuildError != nil {
		log.Println(jsonBuildError)
		response = []byte("Internal Server Error")
		code = 500
	}

	return events.APIGatewayProxyResponse{code, headers, string(response), false}, nil
}

func main() {
	lambda.Start(handle)
}
