# syntax=docker/dockerfile:1
# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o catchall ./cmd/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/catchall .

EXPOSE 8080
CMD [ "/app/catchall" ]