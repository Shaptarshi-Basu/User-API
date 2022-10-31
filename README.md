# REST API

The REST API using AWS Lambda, API gateway and DynamoDB. It enable to create/update and list users.



## Get list of users

### Request

`GET` https://oiddgxkwij.execute-api.eu-north-1.amazonaws.com/test/users

### Response
    - HTTP/1.1 200 StatusOk   
       [
            {
                "email": "test@example.com", 
                "first_name": "first name",
                "subject_id": 1,
            }, 
            {
                "email": "test2@example.com", 
                "first_name": "first name2",
                "last_name": "last name2",
                "address": "real address2",
                "subject_id": 2
            }
        ]
    
    - HTTP/1.1 500 InternalSeverError
        {"Error message in detail"}  


## Create/ Updates user

### Request

`PUT` https://oiddgxkwij.execute-api.eu-north-1.amazonaws.com/test/user

### Request
    {
        "email": "test@example.com", 
        "first_name": "first name",
        "last_name": "last name",
        "address": "real address" 
    }

### Response

    - HTTP/1.1 200 StatusOk
    {
    "email": "test@example.com", 
    "first_name": "first name",
    "last_name": "last name",
    "address": "real address" 
    }
    
    - HTTP/1.1 500 InternalSeverError
    {"Error message in detail"}
    
    - HTTP/1.1 400 BadRequest
    {"Error message in detail"}

    
## Terraform infrastructure setup

### commands
    GOOS=linux go build main.go
    zip api.zip main
    
    //copy the zip to the Terraform dir
    
    terraform init
    terraform apply
    
