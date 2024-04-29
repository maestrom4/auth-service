![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![GitHub Contributions](https://img.shields.io/github/contributors/maestrom4/time-management-microservices)
![GitHub Stars](https://img.shields.io/github/stars/maestrom4/auth-service)
[![codecov](https://codecov.io/gh/maestrom4/auth-service/graph/badge.svg?token=DCWOO8T8NQ)](https://codecov.io/gh/maestrom4/auth-service)

# Auth Service

## Overview
The Auth Service is part of a microservices architecture focused on handling authentication and authorization. It facilitates user registration and credential verification, leveraging a range of modern technologies for robust, scalable, and secure implementations.

## Technologies
- **Go**: A statically typed, compiled programming language designed at Google. Known for its simplicity, efficiency, and excellent support for concurrent operations and microservices.
- **Gin Framework**: A high-performance web framework for Go that provides a robust set of features for building web applications and microservices.
- **GraphQL**: A query language for APIs and a runtime for executing those queries with your existing data. GraphQL provides a more efficient, powerful, and flexible alternative to REST.
- **Docker**: A set of platform-as-a-service products that use OS-level virtualization to deliver software in packages called containers.
- **Testify**: A Go testing toolkit with common assertions and mocks that plays nicely with the standard library.
- **Logrus**: A structured logger for Go, completely API compatible with the standard library logger.
- 

## Installation

### Prerequisites
- Docker
- Go 1.15 or higher (for local development and testing)

This version should provide a more comprehensive guide for anyone looking to understand or use your Auth Service.

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


### Explanation of Key Directories:

- **`cmd/auth-service`**: Contains the main application entry point. This is where the service is initialized and run.
- **`pkg/mongodb`**: Includes all database interaction logic, specifically with MongoDB. For Pure CRUD interaction, but if with logic services folder can be utilized.
- **`internal`**: Houses application-specific code that shouldn't be exposed externally. This includes business logic, configuration, and routing.
- **`internal/middleware`**: Middleware functions are located here, handling tasks like authentication and logging across HTTP requests.
- **`utils`**: Provides utility functions that are widely used throughout the application, such as helpers for authentication and context management.
- **`tests` and `internal/middleware/tests`**: Contain unit and integration tests ensuring that individual units and the integrated components work correctly.

This structure supports a clean separation of concerns, making the application scalable and maintainable. Each component has a specific role, ensuring that the service remains robust and easy to manage as it grows.

