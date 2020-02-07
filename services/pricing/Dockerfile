FROM golang:alpine
LABEL maintainer="weslenng"

WORKDIR /petssenger/pricing

COPY services/pricing/go.mod ./
COPY services/pricing/go.sum ./

RUN go mod download

COPY services/pricing/. .

RUN go build

CMD ["./pricing"]
