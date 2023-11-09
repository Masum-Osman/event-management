FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the application files to the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port your Go service is running on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
