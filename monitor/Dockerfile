# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git protobuf

# Install go-based dependencies and tools
RUN go get \
        github.com/cortesi/modd/cmd/modd \
        google.golang.org/grpc \
        github.com/golang/protobuf/protoc-gen-go \
        github.com/gogo/protobuf/protoc-gen-gogofast \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# Add code.
COPY . /go/src/go.nlx.io/nlx
WORKDIR /go/src/go.nlx.io/nlx/monitor

# Ignore failing make to allow this stage to complete even when there are errors in the sourcecode.
RUN make || true 


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/monitor/dist/bin/nlx-monitor /usr/local/bin/nlx-monitor
