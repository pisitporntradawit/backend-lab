# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

# Install ca-certificates (ถ้าใช้ HTTPS)
RUN apk --no-cache add ca-certificates

# Add non-root user
RUN adduser -D dev
USER dev

WORKDIR /home/dev

# Copy binary from builder
COPY --from=builder /app/main ./



# Expose port (ตามความเหมาะสม)
EXPOSE 30606

# Health check (optional)
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD [ "test", "-f", "./main" ] || exit 1

CMD ["./main"]