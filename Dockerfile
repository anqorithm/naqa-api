# Base golang image with minimal footprint
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy and download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o main ./cmd/api

# App port
EXPOSE 3000

# Run application
CMD ["./main"]