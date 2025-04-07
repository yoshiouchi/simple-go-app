# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64  go build -o main .

# Expose port 8080 to the outside world (if you have a web server or similar)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


