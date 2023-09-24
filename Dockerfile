# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o myapp

# Expose a port that the application will listen on
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./myapp"]
