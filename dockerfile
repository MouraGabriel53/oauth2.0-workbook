ARG DOCKER_GOLANG_VERSION=1.26.1
FROM golang:${DOCKER_GOLANG_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY .. .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-exec ./cmd/api/main.go

FROM alpine:latest

ARG DOCKER_API_USER=golang
ARG DOCKER_API_PORT=8000

RUN adduser -D ${DOCKER_API_USER}

WORKDIR /app

COPY  --from=builder /api-exec .

RUN chmod +x ./api-exec && chown ${DOCKER_API_USER}:${DOCKER_API_USER} ./api-exec

USER ${DOCKER_API_USER}

EXPOSE ${DOCKER_API_PORT}

ENTRYPOINT ["./api-exec"]