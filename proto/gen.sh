#! /usr/bin/env bash

set -e pipefail

cd "$(dirname "$0")/.."
mkdir -p target/

PKG="github.com/goinsane/pbutil"

go build -mod readonly -o target/ \
  google.golang.org/protobuf/cmd/protoc-gen-go
PATH="target/:$PATH"

mkdir -p target/src/
protoc --go_out=./target/src/ --go_opt= -I proto/ proto/google/protobuf/*.proto
protoc --go_out=./target/src/ --go_opt= -I proto/ proto/google/protobuf/compiler/*.proto

find types -depth -type f -name \*.pb.go -delete
find types -type d -empty -delete
protoc --go_out=./ --go_opt=module="$PKG" -I proto/ proto/goinsane/pbutil/protobuf/*.proto

rm -f -- examplespb/*.pb.go
protoc --go_out=./ --go_opt=module="$PKG" -I proto/ proto/goinsane/pbutil/examples/*.proto
