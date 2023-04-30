FROM golang:1.14rc1-stretch
LABEL maintainer="Weslen Guerreiro <savi2w@icloud.com>"

RUN apt-get update
RUN apt-get install -y protobuf-compiler golang-protobuf-extensions-dev

WORKDIR /petssenger/services/pricing
COPY services/pricing/go.mod ./
COPY services/pricing/go.sum ./
RUN go mod download

COPY services/pricing/. .

WORKDIR /petssenger/protos
COPY protos/pricing.proto ./
RUN protoc *.proto --go_out=plugins=grpc:../services/pricing/protos

WORKDIR /petssenger/services/pricing
RUN go build

CMD ["./pricing"]
