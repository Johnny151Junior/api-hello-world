# Use the official Golang image to create a build stage
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN go build -o main .  # Ensure binary is built inside /app

# Use a minimal base image for deployment
FROM alpine:latest

# Set working directory to /app
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Ensure execution permission
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Run the application explicitly from /app
CMD ["/app/main"]
