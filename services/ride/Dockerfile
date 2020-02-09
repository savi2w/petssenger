FROM node:stretch
LABEL maintainer="weslenng"

# system setup
RUN apt-get update
RUN apt-get install -y protobuf-compiler

# dependencies
WORKDIR /petssenger
COPY package.json ./
COPY yarn.lock ./
RUN mkdir -p /petssenger/services/ride
COPY services/ride/package.json ./services/ride/
RUN yarn --pure-lockfile

# proto files
WORKDIR /petssenger/protos
COPY protos/pricing.proto ./
RUN protoc *.proto --plugin="protoc-gen-ts=../node_modules/.bin/protoc-gen-ts" --js_out="import_style=commonjs,binary:." --ts_out="."

# build
WORKDIR /petssenger/services/ride
COPY services/ride/. .
CMD ["yarn", "start"]