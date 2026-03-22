FROM golang:${DOCKER_GOLANG_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY .. ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-exec ./cmd/api/main.go

FROM alpine:latest

RUN adduser -D ${DOCKER_API_USER}
USER ${DOCKER_API_PASSWORD}

WORKDIR /app

EXPOSE ${DOCKER_API_PORT}

COPY  --from=builder /api-exec .

ENTRYPOINT ["./api-exec"]