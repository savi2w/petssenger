FROM golang:stretch
LABEL maintainer="weslenng"

# system setup
RUN apt-get update
RUN apt-get install -y protobuf-compiler
RUN apt-get install -y golang-protobuf-extensions-dev

# proto files
WORKDIR /petssenger/protos
COPY protos/pricing.proto ./
RUN protoc *.proto --go_out=plugins=grpc:.

# dependencies
WORKDIR /petssenger
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build
WORKDIR /petssenger/services/pricing
COPY services/pricing/. .
RUN go build

CMD ["./pricing"]
