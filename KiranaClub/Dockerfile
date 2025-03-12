# Use the official Golang image as a builder stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app .

# Use a lightweight image to run the app
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./app"]
