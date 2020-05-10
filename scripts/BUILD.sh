protoc protos/*.proto \
  --go_out=plugins=grpc:services/pricing \
  --go_out=plugins=grpc:services/user

yarn install
protoc protos/*.proto \
  --plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
  --plugin="protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin" \
  --js_out="import_style=commonjs,binary:services/ride/src" \
  --ts_out="service=grpc-node:services/ride/src" \
  --grpc_out="services/ride/src"
