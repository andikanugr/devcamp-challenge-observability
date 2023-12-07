# Stage 1: Building the application
FROM golang:1.19 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o todo-app cmd/http/main.go

# Stage 2: Build a lightweight image
FROM alpine:3.14

WORKDIR /app

# Copy the binary from the builder stage to the current working directory
COPY --from=builder /app/todo-app .

# Create the logs directory
# ideally the logs is on var/log but for the sake of simplicity
RUN mkdir -p files/logs

# Command to run the executable
CMD ["./todo-app"]
