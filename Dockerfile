# Use an official Go runtime as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory to the container's working directory
COPY . .

# Navigate into the 'web' directory
WORKDIR /app/simple_go_web

# Build the Go application inside the 'web' directory
RUN go build -o myapp

# Set the command to run your application
CMD ["./myapp"]
