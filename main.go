package main

import (
	"user-api/routes"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(routes.Handler)
}
