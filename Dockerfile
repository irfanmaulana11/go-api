# Use the official Go base image with version 1.18.5
FROM golang:1.18.5

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application source code into the container
COPY . .

# Build your Go application
RUN go build -o myapp

# Expose the port your application will listen on
EXPOSE 8080

# Command to run your application
CMD ["./myapp"]