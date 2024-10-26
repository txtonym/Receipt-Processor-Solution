# Use Go 1.23.2 with Alpine for a lightweight image
FROM golang:1.23.2-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o Receipt-Processor-Solution .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./Receipt-Processor-Solution"]
