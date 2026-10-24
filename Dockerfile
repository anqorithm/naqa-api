# Build stage
FROM golang:1.26-alpine AS build

WORKDIR /src

# Copy and download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build a static binary
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/main ./cmd/api

# Runtime stage
FROM alpine:3.22

RUN apk add --no-cache ca-certificates wget && adduser -D -u 10001 app

WORKDIR /app

# The seeder reads the yearly JSON files from ./sources
COPY --from=build /out/main ./main
COPY sources ./sources

USER app

# App port
EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s CMD wget -qO- http://127.0.0.1:3000/api/v1/health || exit 1

CMD ["./main"]
