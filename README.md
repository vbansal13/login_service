# Login Service

This is a Token based login service written in Golang and supports following features:

* Supports User Registration
* Supports User Login
* Supports User Profile
* Use Token Based Authentication
* Tokens is valid for only 1 hour
* Supports Account locking on multiple consecutive unsuccessful login attempts.
* Tracks Login activity by publishing messages to Kafka topic
* Swagger documentation

Refer to Postman collection under /docs/LoginService.postman_collection.json

Here is a basic software reference architecture of this application

![Architecture](docs/LoginService%20High%20Level%20Diagram.png)

