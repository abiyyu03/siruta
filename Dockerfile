# Build stage
FROM golang:1.24-alpine AS builder

# Install git and other dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Setelah WORKDIR /app
COPY keys ./keys

# Download dependencies (opsional, tapi disarankan)
RUN go mod tidy

# Build Go binary
RUN go build -o satuwarga-prod main.go

# Run stage (smaller image)
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/satuwarga-prod .

# Expose port (ubah sesuai port aplikasi Go-mu)
EXPOSE 8080

# Command to run the app
CMD ["./satuwarga-prod"]
