FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata netcat-openbsd
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Create wait script that uses netcat to check MySQL port
RUN echo '#!/bin/sh' > wait-for-mysql.sh && \
    echo 'echo "Waiting for MySQL at $DB_HOST:$DB_PORT..."' >> wait-for-mysql.sh && \
    echo 'while ! nc -z "$DB_HOST" "$DB_PORT"; do' >> wait-for-mysql.sh && \
    echo '  echo "MySQL is not ready yet - sleeping 2 seconds..."' >> wait-for-mysql.sh && \
    echo '  sleep 2' >> wait-for-mysql.sh && \
    echo 'done' >> wait-for-mysql.sh && \
    echo 'echo "MySQL is ready! Starting application..."' >> wait-for-mysql.sh && \
    echo 'sleep 2' >> wait-for-mysql.sh && \
    echo 'exec "$@"' >> wait-for-mysql.sh && \
    chmod +x wait-for-mysql.sh

EXPOSE 8080

ENTRYPOINT ["./wait-for-mysql.sh"]
CMD ["./main"]