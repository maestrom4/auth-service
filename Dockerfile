FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /auth-service ./cmd/auth-service/


FROM alpine:latest
WORKDIR /root/


COPY --from=builder /auth-service .


CMD ["./auth-service"]

# FROM golang:1.18 AS builder

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /uth-service ./cmd/uth-service/


# FROM alpine:latest
# WORKDIR /root/


# COPY --from=builder /auth-service .


# CMD ["./auth-service"]
