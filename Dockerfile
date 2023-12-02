# Stage 1: Build the Go application
FROM golang:1.19.0-alpine3.15 as builder

WORKDIR /app

COPY . .

# Build the Go application and name the executable as "app"
RUN go build -o app

# Stage 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /app

# Install necessary packages including Go
RUN apk update && apk add --no-cache ca-certificates && apk add --no-cache bash

# Copy the compiled Go application from the builder stage
COPY --from=builder /app/db/migrations /app/db/migrations
COPY --from=builder /app/app /app/app
COPY wait-for-it.sh /app/wait-for-it.sh
COPY startup.sh /app/startup.sh
COPY monsterfellowship.json /app/monsterfellowship.json

# Make scripts executable
RUN chmod +x /app/wait-for-it.sh
RUN chmod +x /app/startup.sh

EXPOSE $PORT

CMD ["./startup.sh"]  # Use startup.sh as the entry point