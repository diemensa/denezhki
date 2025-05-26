FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git bash

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o denezhki ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/denezhki .

EXPOSE 8080

CMD ["./denezhki"]
