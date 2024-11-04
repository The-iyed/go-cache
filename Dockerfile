FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o i-cache ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/i-cache .

EXPOSE 6379

CMD ["./i-cache"]
