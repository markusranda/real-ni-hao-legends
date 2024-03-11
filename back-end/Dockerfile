# First stage: build the Go application
FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o nihao_app .

# Second stage: run the application
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/nihao_app .

CMD ["./nihao_app"]
