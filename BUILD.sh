# Utility script to gen proto files

protoc protos/*.proto --go_out=plugins=grpc:.

yarn install
protoc protos/*.proto \
  --plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
  --plugin="protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin" \
  --js_out="import_style=commonjs,binary:." \
  --ts_out="service=grpc-node:." \
  --grpc_out="."
