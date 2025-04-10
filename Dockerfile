FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /pvz-service cmd/main.go

# Final lightweight image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /pvz-service .

EXPOSE 8080

CMD ["./pvz-service"]