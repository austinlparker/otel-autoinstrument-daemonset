# Use the official Golang image to create a build artifact.
FROM golang:1.16 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go source
COPY main.go main.go

# Build the Go app as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o optimistic-message main.go

# Use the official lightweight Alpine image for a lean production container
FROM alpine:3.14

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the prebuilt binary file from the builder
COPY --from=builder /app/optimistic-message .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./optimistic-message"]
