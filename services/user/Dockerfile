FROM golang:1.14rc1-stretch
LABEL maintainer="weslenng"

# system setup
RUN apt-get update
RUN apt-get install -y protobuf-compiler golang-protobuf-extensions-dev

# proto files
WORKDIR /petssenger/protos
COPY protos/user.proto ./
RUN protoc *.proto --go_out=plugins=grpc:.

# dependencies
WORKDIR /petssenger
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build
WORKDIR /petssenger/services/user
COPY services/user/. .
RUN go build

CMD ["./user"]