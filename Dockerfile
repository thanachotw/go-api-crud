# Start from the official Golang image
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .
# Copy .env file to image
# COPY .env .env

# Build the application
RUN go build -o wallet-api ./cmd/main.go

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/wallet-api .

# Expose the port (change if needed)
EXPOSE 8080

# Set environment variables (can be overridden)
ENV APP_PORT=8080
ENV APP_DEBUG=false

# Run the binary
CMD ["./wallet-api", "serve-rest"]
