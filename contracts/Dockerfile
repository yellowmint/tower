FROM golang:1.18.3-alpine3.16 as golang-plugins

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
RUN go install github.com/envoyproxy/protoc-gen-validate@v0.6.7

FROM bufbuild/buf:1.5.0

# install protoc
RUN apk update && apk add --no-cache protobuf

# copy golang plugins
COPY --from=golang-plugins /go/bin/ /usr/local/go/bin/
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /workspace
