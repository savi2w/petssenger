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
RUN protoc *.proto \
  --plugin="protoc-gen-ts=../node_modules/.bin/protoc-gen-ts" \
  --plugin="protoc-gen-grpc=../node_modules/.bin/grpc_tools_node_protoc_plugin" \
  --js_out="import_style=commonjs,binary:." \
  --ts_out="service=grpc-node:." \
  --grpc_out="."

# run
WORKDIR /petssenger/services/ride
COPY services/ride/. .

EXPOSE 3000
CMD ["yarn", "start"]