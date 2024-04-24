# Auth Service

## Overview
This service handles Authentication, Authorization, registration.

## Setup with docker
To run this project, install it locally using npm:

```bash
npm install
npm start

docker build -t auth-service .
docker run -p 8080:8080 --name file-gen-service file-gen-service

# if to local registry
docker build -t localhost:5000/auth-service:latest .
docker run -p 8080:8080 localhost:5000/auth-service:latest
```
Test the api 
```bash
curl http://localhost:8080/api/health
```
```bash
response { "message": "Hello from the API!" }
```
''check the logs also''

### Postman testing
![GitHub Logo](https://github.com/maestrom4/Bundy-management-microservice/blob/develop/postmanTesting.png?raw=true)

- Click collections
- Click new select Graphql
- Always click the reload button highlighted in the screenshot.
- Tick/untick the checkbox query/mutation schema 
- Click Query
