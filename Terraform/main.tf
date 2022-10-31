terraform {
 required_providers {
   aws = {
     source = "hashicorp/aws"
   }
 }
}


provider "aws" {
  region = "eu-north-1"
  access_key = ""
  secret_key = ""
}

//Create User table with primary key as email
resource "aws_dynamodb_table" "User" {
  name             = "User"
  hash_key         = "email"
  billing_mode     = "PROVISIONED"
  read_capacity    = 5
  write_capacity   = 5

  attribute {
    name = "email"
    type = "S"
  }
  ttl {
    attribute_name = "TimeToExist"
    enabled        = false
  }
}

//Create policy so that lambda can add, list update and delete to table User in DynamoDB
resource "aws_iam_role_policy" "lambda_policy" {
  name = "lambda_policy"
  role = aws_iam_role.role_for_LDC.id

  policy = file("policy.json")
}


resource "aws_iam_role" "role_for_LDC" {
  name = "myrole"
  assume_role_policy = file("assume_role_policy.json")

}


resource "aws_lambda_function" "terraform_lambda_func" {
  filename      = "api.zip"
  function_name = "User_Api"
  role          = aws_iam_role.role_for_LDC.arn
  handler       = "main"
  runtime       = "go1.x"
}


resource "aws_api_gateway_rest_api" "apiLambda" {
  name        = "User-API"
}

resource "aws_api_gateway_resource" "resource1" {
   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   parent_id   = aws_api_gateway_rest_api.apiLambda.root_resource_id
   path_part   = "users"
}
resource "aws_api_gateway_resource" "resource2" {
   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   parent_id   = aws_api_gateway_rest_api.apiLambda.root_resource_id
   path_part   = "user"
}

resource "aws_api_gateway_method" "getUsersMethod" {
   rest_api_id   = aws_api_gateway_rest_api.apiLambda.id
   resource_id   = aws_api_gateway_resource.resource1.id
   http_method   = "GET"
   authorization = "NONE"
}
resource "aws_api_gateway_method" "updateUserMethod" {
   rest_api_id   = aws_api_gateway_rest_api.apiLambda.id
   resource_id   = aws_api_gateway_resource.resource2.id
   http_method   = "PUT"
   authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda-intgr-get" {
   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   resource_id = aws_api_gateway_method.getUsersMethod.resource_id
   http_method = aws_api_gateway_method.getUsersMethod.http_method

   integration_http_method = "POST"
   type                    = "AWS_PROXY"
   uri                     = aws_lambda_function.terraform_lambda_func.invoke_arn
}



resource "aws_api_gateway_integration" "lambda-intgr-put" {
   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   resource_id = aws_api_gateway_method.updateUserMethod.resource_id
   http_method = aws_api_gateway_method.updateUserMethod.http_method

   integration_http_method = "POST"
   type                    = "AWS_PROXY"
   uri                     = aws_lambda_function.terraform_lambda_func.invoke_arn
}

resource "aws_api_gateway_deployment" "api-deploy-put" {
   depends_on = [
     aws_api_gateway_integration.lambda-intgr-put,
   ]

   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   stage_name  = "test"
}
resource "aws_api_gateway_deployment" "api-deploy-get" {
   depends_on = [

     aws_api_gateway_integration.lambda-intgr-get,
   ]

   rest_api_id = aws_api_gateway_rest_api.apiLambda.id
   stage_name  = "test"
}


resource "aws_lambda_permission" "apigw" {
   statement_id  = "AllowAPIGatewayInvoke"
   action        = "lambda:InvokeFunction"
   function_name = aws_lambda_function.terraform_lambda_func.function_name
   principal     = "apigateway.amazonaws.com"

   # The "/*/*" portion grants access from any method on any resource
   # within the API Gateway REST API.
   source_arn = "${aws_api_gateway_rest_api.apiLambda.execution_arn}/*/*"
}

