
FROM golang:1.19 AS builder

WORKDIR /app


COPY go.* ./
RUN go mod download


COPY ./cmd/apigateway/ ./cmd/apigateway/


RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/auth-service/main.go


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /


COPY --from=builder /auth-service /auth-service


EXPOSE 8080


CMD ["/auth-service"]
