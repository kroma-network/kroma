# Build stage
FROM golang:1.21.1-alpine3.18

COPY ./cmd /app/cmd
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app
RUN go build -o /usr/local/bin/migration-verifier ./cmd/main.go

