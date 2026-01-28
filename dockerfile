# Use your Go version
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o employee-api ./cmd/main.go

# Final minimal image
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/employee-api .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./employee-api"]