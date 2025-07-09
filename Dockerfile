# Stage 1: Build the Go app
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go mod files and download
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (so it can see all packages)
COPY . .

# Build the binary from cmd/server/main.go
RUN go build -o server ./cmd/server

# Stage 2: Run the built binary in a lightweight image
FROM debian:bullseye-slim

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/server .

# Run the app
CMD ["./server"]