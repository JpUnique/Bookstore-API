# Start from official Go image
FROM golang:1.20-alpine

# Create app directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o bookstore-api main.go

# Expose port (customize if needed)
EXPOSE 8080

# Start the app
CMD ["./bookstore-api"]
