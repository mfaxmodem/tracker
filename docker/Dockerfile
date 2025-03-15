FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Compile with optimized settings
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage for execution
FROM alpine:latest
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]