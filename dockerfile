FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY .. .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-exec ./cmd/api/main.go

FROM alpine:latest

ARG DOCKER_API_USER

WORKDIR /app

COPY  --from=builder /api-exec .

EXPOSE ${DOCKER_API_PORT}

ENTRYPOINT ["./api-exec"]