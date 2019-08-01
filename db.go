package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getItems(pType string, minPrice string, sort string) ([]Product, error) {
	// Declare a new DynamoDB instance.
	sess := session.Must(session.NewSession())
	var db = dynamodb.New(sess)

	//Default to 0 if not specified
	if minPrice == "" {
		minPrice = "0"
	}
	scanIndex := false
	if sort != "" && strings.Contains(sort, "a") {
		scanIndex = true
	}

	// Prepare the input for the query.
	input := &dynamodb.QueryInput{
		TableName: aws.String("Product"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {N: aws.String(minPrice)},
			":v2": {S: aws.String(pType)},
		},
		KeyConditionExpression: aws.String("PType = :v2 and Price >= :v1"),
		ScanIndexForward:       &scanIndex,
	}

	// Retrieve the item from DynamoDB. If no matching item is found return nil.
	result, err := db.Query(input)

	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type map[string]*AttributeValue
	// We can use the UnmarshalMap to parse this
	var products []Product
	for _, i := range result.Items {
		product := Product{}
		err := dynamodbattribute.UnmarshalMap(i, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
