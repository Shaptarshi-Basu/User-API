# REST API

The REST API for AWS Lambda. It enable to create/update and list users

## Get list of users

### Request

`GET /users`


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

`PUT /user`

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
    