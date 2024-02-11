# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go get -d -v ./...

# Install your application
RUN go install -v ./...

# Command to run your application
CMD ["app"]
