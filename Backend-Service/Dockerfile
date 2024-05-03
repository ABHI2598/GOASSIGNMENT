# Use the official Golang image as base
FROM golang:1.17-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Use a small image to run the application
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
