FROM node:12.15.0-stretch
LABEL maintainer="Weslen Guerreiro <savi2w@icloud.com>"

RUN apt-get update
RUN apt-get install -y protobuf-compiler

WORKDIR /petssenger
COPY package.json ./
COPY yarn.lock ./

WORKDIR /petssenger/services/ride
COPY services/ride/package.json ./

WORKDIR /petssenger
RUN yarn --pure-lockfile

WORKDIR /petssenger/services/ride
COPY services/ride/. .

WORKDIR /petssenger/protos
COPY protos/pricing.proto ./
COPY protos/user.proto ./
RUN protoc *.proto \
  --plugin="protoc-gen-ts=../node_modules/.bin/protoc-gen-ts" \
  --plugin="protoc-gen-grpc=../node_modules/.bin/grpc_tools_node_protoc_plugin" \
  --js_out="import_style=commonjs,binary:../services/ride/src/protos" \
  --ts_out="service=grpc-node:../services/ride/src/protos" \
  --grpc_out="../services/ride/src/protos"

WORKDIR /petssenger/services/ride
RUN yarn build

EXPOSE 3000
CMD ["yarn", "start"]
