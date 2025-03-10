# Use Golang with Alpine as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Create a minimal runtime image with Alpine
FROM alpine:latest

# Set working directory in the final container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8000

# Run the application
CMD ["./main"]
