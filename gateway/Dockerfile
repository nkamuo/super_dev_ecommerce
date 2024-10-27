# Dockerfile for API Gateway
FROM golang:1.22-alpine AS builder

# Install Git for fetching dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o /go/bin/api-gateway

# Final stage for a smaller runtime image
FROM alpine:latest
COPY --from=builder /go/bin/api-gateway /api-gateway

# Expose the port the service will run on
EXPOSE 8080

# Run the application
ENTRYPOINT ["/api-gateway"]



# # Dockerfile for API Gateway
# # Use the official Golang image to build the Go application
# # FROM golang:1.19-alpine AS builder
# FROM golang:1.22-alpine AS builder

# # Set the working directory inside the container
# WORKDIR /app

# # Copy all Go files
# COPY . .

# # Download dependencies
# RUN go mod tidy

# # Build the application
# RUN go build -o api-gateway .

# # Start a new image to run the API Gateway
# FROM alpine:latest

# WORKDIR /app

# # Copy the built binary from the builder stage
# COPY --from=builder /app/api-gateway /app/api-gateway

# # Expose the API Gateway port
# EXPOSE 8080

# # Run the API Gateway binary
# CMD ["./api-gateway"]
