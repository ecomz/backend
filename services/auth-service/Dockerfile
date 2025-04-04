FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ./go.work ./go.work
COPY ./go.work.sum ./go.work.sum
COPY ./libs ./libs
COPY ./services ./services

# Fetch Dependencies
WORKDIR /app/libs
RUN go mod download

WORKDIR /app/services/auth-service
RUN go mod download

RUN mkdir -p /usr/bin && CGO_ENABLED=0 go build -o /usr/bin/auth-service ./cmd

FROM builder

COPY --from=builder /usr/bin/auth-service /usr/bin/auth-service

CMD ["auth-service"]