#! /usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.."
mkdir -p target/

PKG="github.com/goinsane/pbutil"

go build -mod readonly -o target/ \
  google.golang.org/protobuf/cmd/protoc-gen-go
PATH="target/:$PATH"

#mkdir -p target/src/
#protoc --go_out=./target/src/ --go_opt= -I proto/ proto/google/protobuf/*.proto
#protoc --go_out=./target/src/ --go_opt= -I proto/ proto/google/protobuf/compiler/*.proto

rm -f -- mongopb/*.pb.go
protoc --go_out=./ --go_opt=module="$PKG" -I proto/ proto/goinsane/pbutil/mongo.proto
