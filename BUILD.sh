protoc protos/*.proto --go_out=plugins=grpc:.

yarn install
protoc protos/*.proto --plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" --js_out="import_style=commonjs,binary:." --ts_out="."