# Use the official Go image to create a build artifact.
# This is based on Debian and includes the Go toolchain.
FROM golang:1.21 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o myapp specifies the output binary name, change as needed.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o myapp

# Use the official lightweight Scratch image.
# https://hub.docker.com/_/scratch
FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp /myapp

# Expose port (default to 1111)
ARG PORT=1111
ENV PORT=${PORT}
EXPOSE ${PORT}

# Run the web service on container startup.
CMD ["/myapp"]
