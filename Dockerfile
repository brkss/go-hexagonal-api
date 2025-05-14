# Build stage
FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dextrace-server ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/dextrace-server .

COPY .env .env

EXPOSE 8080

CMD ["./dextrace-server"]
