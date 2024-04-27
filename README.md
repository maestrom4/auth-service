# Auth Service

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![GitHub Contributions](https://img.shields.io/github/contributors/maestrom4/time-management-microservices)
![GitHub Stars](https://img.shields.io/github/stars/maestrom4/auth-service)

## Overview
The Auth Service provides functionality for Authentication, Authorization, and user registration.

## Technologies Used
- Golang
- GraphQL
- Gin Framework
- Docker Compose
- Testify for unit testing
- Logrus for logging

## Setup with Docker
To run this project locally, follow these steps:

1. Clone the repository and navigate to the project directory.
2. Build the Docker image:
    ```bash
    docker build -t auth-service .
    ```
3. Run the Docker container:
    ```bash
    docker run -p 8085:8080 --name auth-service auth-service
    ```
   If you're using a local registry:
    ```bash
    docker build -t localhost:5000/auth-service:latest .
    docker run -p 8085:8080 localhost:5000/auth-service:latest
    ```
4. Test the API:
    ```bash
    curl http://localhost:8080/api/health
    ```
   Expected response:
    ```json
    { "message": "Hello from the API!" }
    ```
   Remember to check the logs for more information.

## Postman Testing
![Postman Testing](https://github.com/maestrom4/auth-service/blob/develop/postmanTesting.png)

1. Click on collections.
2. Select the GraphQL collection.
3. Click the reload button highlighted in the screenshot to refresh the collection.
4. Tick/untick the checkbox for query/mutation schema.
5. Click Query to execute the selected operation.
