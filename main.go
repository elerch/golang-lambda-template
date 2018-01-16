package main

import (
	"context"
	"errors"
	// Used for things like API Gateway Integration
	//"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log"
)

// Request type must be public so lambda.start can marshall from json to the type
// using encoding/json package.
//
// Supported intrinsics include:
//   bool
//   float64
//   string
//   []interface{}  // JSON arrays
//   map[string]interface{} // JSON objects
//   nill
type Request struct {
	ID		float64		`json:"id"`
	Value	string		`json:"value"`
}

// Response type also must be public for lambda.start to marshall to json
type Response struct {
	Message	string		`json:"message"`
	Ok		bool		`json:"ok"`
}

var (
	// ErrValueNotProvided is thrown when a value is not provided
	ErrValueNotProvided = errors.New("No value provided in request JSON")
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(context context.Context, request Request) (Response, error) {
    awscontext, _ := lambdacontext.FromContext(context)
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", awscontext.AwsRequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Value) < 1 {
		return Response { Message: "", Ok: false },
				ErrValueNotProvided
	}

	return Response{
		Message: "Hello " + request.Value,
		Ok: true,
	}, nil
}

// main is the initial entry point
// unlike other runtimes, the specific function is **not** specified in the
// configuration. Instead the handler configuration is simply the binary name,
// from there main will be entered by the go runtime and the handler will be
// registered.
func main() {
	// Unlike other runtimes, we explicitly call into the library that works
	// with the underlying service. lambda.Start will register our handler
	//
	// Note: This allows us to do setup "off the clock" and away from AWS
	// billing, much like Python globals
	lambda.Start(Handler)
}
