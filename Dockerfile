FROM golang:1.19 AS builder


WORKDIR /app

COPY go.* ./
RUN go mod download

COPY ./cmd/auth-service/main.go ./cmd/auth-service/


RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/auth-service/main.go


FROM alpine:latest

WORKDIR /

COPY --from=builder /auth-service /auth-service

EXPOSE 8080

CMD ["/auth-service"]
