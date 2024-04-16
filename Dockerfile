FROM golang:1.19 AS builder

WORKDIR /app

# Copy the Go Modules manifests and download the modules
COPY go.* ./
RUN go mod download

# Copy the entire module directory
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/auth-service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

# Copy the compiled binary from the builder stage
COPY --from=builder /auth-service /auth-service

EXPOSE 8080

CMD ["/auth-service"]
