# Use official Go image
FROM golang:1.21-alpine

# Install git (needed for go get) and ca-certificates
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN go build -o main ./cmd

# Expose port
EXPOSE 8080

# Start the app
CMD ["./main"]