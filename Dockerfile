FROM golang:1.16-alpine
LABEL MAINTAINER="xcbeyond"

# setting go proxy.
RUN go env -w GO111MODULE=on; \
   go env -w GOPROXY=https://goproxy.cn,direct

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o go-httpbin

EXPOSE 8080

# Run the application service on container startup.
CMD ["/app/go-httpbin"]