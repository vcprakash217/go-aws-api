package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type APITest struct {
	request events.APIGatewayProxyRequest
	expect  int
}

func TestRouter(t *testing.T) {
	test := APITest{events.APIGatewayProxyRequest{HTTPMethod: ""}, 405}
	response, _ := router(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}

func TestShowNoQueryParams(t *testing.T) {
	test := APITest{events.APIGatewayProxyRequest{HTTPMethod: "GET"}, 400}
	response, _ := show(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}

func TestShowNoPType(t *testing.T) {
	parameters := make(map[string]string)
	parameters["minPrice"] = "123"
	test := APITest{events.APIGatewayProxyRequest{QueryStringParameters: parameters}, 400}
	response, _ := show(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}

func TestShowInvalidminPrice(t *testing.T) {
	parameters := make(map[string]string)
	parameters["pType"] = "Type1"
	parameters["minPrice"] = "a123."
	test := APITest{events.APIGatewayProxyRequest{QueryStringParameters: parameters}, 400}
	response, _ := show(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}

func TestShowInvalidSortDigits(t *testing.T) {
	parameters := make(map[string]string)
	parameters["pType"] = "Type1"
	parameters["sort"] = "123"
	test := APITest{events.APIGatewayProxyRequest{QueryStringParameters: parameters}, 400}
	response, _ := show(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}

func TestShowInvalidSortChars(t *testing.T) {
	parameters := make(map[string]string)
	parameters["pType"] = "Type1"
	parameters["sort"] = "asd"
	test := APITest{events.APIGatewayProxyRequest{QueryStringParameters: parameters}, 400}
	response, _ := show(test.request)
	assert.Equal(t, test.expect, response.StatusCode)
}
