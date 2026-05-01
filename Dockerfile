# Use 1.23 or 'latest' to ensure you have the most recent toolchain
FROM golang:1.23-alpine AS builder

WORKDIR /app

# This allows Go to download a newer toolchain if go.mod requires it
ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir -p /data

WORKDIR /root/
COPY --from=builder /app/main .

ENV DB_PATH=/data/jokes.db
ENV GIN_MODE=release

# Expose port
EXPOSE 8080

CMD ["./main"]