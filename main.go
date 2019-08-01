package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Allows digits and '.' is optional
var priceRegexp = regexp.MustCompile(`^[0-9]+.?[0-9]+$`)

// Allows a/A/d/D
var sortRegexp = regexp.MustCompile(`^[adAD]{1,1}$`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	PType string `json:"pType"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	// TODO
	// case "POST":
	//     return create(req)
	// case "PUT":
	//     return update(req)
	// case "DELETE":
	//     return delete(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(req.QueryStringParameters) == 0 {
		return clientError(http.StatusBadRequest)
	}
	minPrice := req.QueryStringParameters["minPrice"]
	pType := req.QueryStringParameters["pType"]
	sort := req.QueryStringParameters["sort"]

	// pType is required field. minPrice and sort are optional.
	if pType == "" {
		return clientError(http.StatusBadRequest)
	}
	if minPrice != "" && !priceRegexp.MatchString(minPrice) {
		return clientError(http.StatusBadRequest)
	}
	if sort != "" && !sortRegexp.MatchString(sort) {
		return clientError(http.StatusBadRequest)
	}

	return fetch(pType, minPrice, sort)
}

func fetch(pType string, minPrice string, sort string) (events.APIGatewayProxyResponse, error) {
	// Fetch the product records from the database
	products, err := getItems(pType, minPrice, sort)
	if err != nil {
		return serverError(err)
	}
	if products == nil {
		return clientError(http.StatusBadRequest)
	}

	// The APIGatewayProxyResponse.Body field needs to be a string, so
	// we marshal the product record into JSON.
	js, err := json.Marshal(&Products{Products: products})
	if err != nil {
		return serverError(err)
	}

	// Return a response with a 200 OK status and the JSON product record(s)
	// as the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

// This logs any error and returns a 500 Internal Server Error response
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Sends responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
