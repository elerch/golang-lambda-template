package main

import (
	// Used for things like API Gateway Integration 
	// "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request Request
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: Request{ID: 1.1, Value: "World"},
			expect:  "Hello World",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: Request{Value: ""},
			expect:  "",
			err:     ErrValueNotProvided,
		},
	}

	for _, test := range tests {
		response, err := Handler(lambdacontext.LambdaContext{AwsRequestID: "test"}, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Message)
	}
}
