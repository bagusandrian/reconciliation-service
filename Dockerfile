# Start from the official Go image
FROM golang:1.22
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# make directory for storing the config file
RUN mkdir -p /etc/reconciliation-service/
# copy config file
COPY files/etc/reconciliation-service/reconciliation-service.development.yaml /etc/reconciliation-service/
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Set environment variable for the views directory
ENV SysEnv=testing
# Build the Go app
RUN cd ./cmd/http/ && go build && cd ../..
# Expose port 8080 to the outside world
EXPOSE 9000
# Run the executable
CMD ["/app/cmd/http/http"]