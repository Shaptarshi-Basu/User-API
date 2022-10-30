package main

import (
	"api-proj/routes"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(routes.Handler)
}
