# 🏗 Stage 1: Build the Go Application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy module files first (caching dependencies)
COPY ../go.mod ../go.sum ./

RUN go mod tidy

# Copy entire project
COPY ../ ./

# Build the Go application
RUN go build -o gpsd-map-mgmt-svc ./cmd/main.go

# 🏗 Stage 2: Use Debian Bookworm (latest GLIBC version)
FROM debian:bookworm

WORKDIR /app

# Install necessary libraries including latest GLIBC
RUN apt-get update && apt-get install -y ca-certificates libc6

# Copy the compiled binary from the builder stage
COPY --from=builder /app/gpsd-map-mgmt-svc .

# Copy configuration files if needed
COPY ../config/ config/

# Expose the application port
EXPOSE 7000

# Start the application
CMD ["./gpsd-map-mgmt-svc"]
