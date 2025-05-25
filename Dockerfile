FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY subprocessoids/snifflomatron/go.mod subprocessoids/snifflomatron/go.sum ./
RUN go mod download

# Copy full source code
COPY subprocessoids/snifflomatron/ ./
# Build the server from the correct entry point
RUN go build -o server ./cmd

# Runtime image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]