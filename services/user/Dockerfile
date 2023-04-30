FROM golang:1.14rc1-stretch
LABEL maintainer="Weslen Guerreiro <savi2w@icloud.com>"

RUN apt-get update
RUN apt-get install -y protobuf-compiler golang-protobuf-extensions-dev

WORKDIR /petssenger/services/user
COPY services/user/go.mod ./
COPY services/user/go.sum ./
RUN go mod download

COPY services/user/. .

WORKDIR /petssenger/protos
COPY protos/user.proto ./
RUN protoc *.proto --go_out=plugins=grpc:../services/user/protos

WORKDIR /petssenger/services/user
RUN go build

EXPOSE 3002
CMD ["./user"]
