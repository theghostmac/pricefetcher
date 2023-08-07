FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o pricefetcher cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/pricefetcher .

EXPOSE 8080

CMD ["./pricefetcher"]