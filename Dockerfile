# Use official Go image
FROM golang:1.22-alpine AS builder

# Set up working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go app binary
 RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o post-service ./cmd/post-service

# Use a lightweight final image
FROM alpine:latest

# Create non-root user for safety
RUN adduser -D appuser

# Set working directory
WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/post-service .

# Use non-root user
USER appuser

# Set entrypoint
CMD ["./post-service"]
