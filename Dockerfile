# Use official lightweight Golang image
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main ./cmd/main.go

# Use a minimal base image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port 3000
EXPOSE 3000

# Start the application
CMD ["./main"]
