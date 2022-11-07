# User API

Rest API implmentation using AWS Lambda, API gateway and DynamoDB. It enable to create/update and list users.



## Get list of users

### Request

`GET` /users

### Response
    - HTTP/1.1 200 StatusOk   
       [
            {
                "email": "test1@example.com", 
                "first_name": "first name1",
                "last_name": "last name2",
                "address": "real address1"
            }, 
            {
                "email": "test2@example.com", 
                "first_name": "first name2",
                "last_name": "last name2",
                "address": "real address2"
            }
        ]
    
    - HTTP/1.1 500 InternalSeverError
        {"Error message in detail"}  


## Create/ Updates user

### Request

`PUT` /user

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

### Design for Authentication and Authorization

- Auth0 could be used as a jwt token provider. This would provide as
  the first basis of authorization. 
AuthO could be configured to provide us the bearer token

- Next thing would be to implement a Lambda function which would be the custom authorizer

- In s3 we can add a JSON  where roles can we mapped to
endpoints and stored.

### example
    {
        "role1": {
            "GET": ["/some/path", "some/other/path"],
            "PUT": ["some/path"]
        },
        "role2": {
            "GET": ["/some/path2", "some/other/path2"],
            "POST": ["some/path"]
        }
    }
- The Lambda authorizer would read the json from s3 and then will be parsing the jwt token to fetch the role value from there and then match
#### if the request method and path are in the role then 
 
    {
    "Version": "2012-10-17",
    "Statement": [
        {
        "Action": "execute-api:Invoke",
        "Effect": "Allow",
        "Resource": "arn:aws:execute-api:us-east-1:123456789012:ivdtdhp7b5/ESTestInvoke-stage/GET/"
        }
    ]
    }

#### else return 

    {
    "Version": "2012-10-17",
    "Statement": [
        {
        "Action": "execute-api:Invoke",
        "Effect": "Deny",
        "Resource": "arn:aws:execute-api:us-east-1:123456789012:ivdtdhp7b5/ESTestInvoke-stage/GET/"
        }
    ]
    }

### TODO items
-  Implement docker based tests for the db operations
    https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html#docker
- Add tests events to lambda
- implement lambda based authorizer
- upload zip to s3 and use s3 to upload zip to lambda
