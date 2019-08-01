# go-aws-api
## AWS DynamoDB API
REST API to retrieve Product information 

## Endpoints
/products/?pType=xx – returns products of type ‘pType’. Required parameter

## Optional additional parameters
minPrice – returns subset of ‘pType’ products that have price greater than or equal to minimum price specified. Should be a number; decimals are allowed

sort – sorts the products in ascending or descending order of price. A/a – ascending, D/d – descending. Defaults to descending if not specified

## Access 
	Amazon API Gateway was used to expose the API. It can be accessed at below url
https://odmg3v9db2.execute-api.us-west-2.amazonaws.com/staging/products?pType=Type1

## Code
API was developed using GO language and DynamoDB was used for data persistence

## Files
main.go – Routes, validates and responds to the requests
db.go – Connects and fetches data from DynamoDB
main_test.go – unit tests the main.go file

## Database
	Tablename: Product
	Partition Key: PType
	Sort Key: Price
https://stackoverflow.com/questions/45581744/how-does-dynamodb-partition-key-works

## Setup
1.	Install go
2.	Install AWS CLI
3.	Configure AWS access keys, region and format using ‘aws configure’

## Testing
Run unit tests in main_test.go file using ‘go test’ command

## TODO
1.	Add more endpoints
2.	Handle POST, PUT and DELETE requests

## References
1.	https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda#deploying-the-api
2.	https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html
3.	https://blog.mwaysolutions.com/2014/06/05/10-best-practices-for-better-restful-api/
4.	https://golang.org/doc/effective_go.html
