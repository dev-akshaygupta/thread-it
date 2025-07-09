# Stage 1: Build the Go app
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go mod files and download
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (so it can see all packages)
COPY . .

# Build the binary from cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o post-service ./cmd/post-service

# # Stage 2: Run the built binary in a lightweight image
# FROM debian:bullseye-slim

# WORKDIR /app

# # Copy the binary from builder stage
# COPY --from=builder /app/server .

# EXPOSE 8080

# # Run the app
# CMD ["./server"]

FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o post-service ./cmd/post-service

EXPOSE 8080
CMD ["./post-service"]