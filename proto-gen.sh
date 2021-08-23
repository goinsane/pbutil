#! /usr/bin/env bash

set -e pipefail

cd "$(dirname "$0")/."
mkdir -p target/

go build -mod readonly -o target/ \
  google.golang.org/protobuf/cmd/protoc-gen-go
PATH="target/:$PATH"

rm -f -- *.pb.go
protoc --go_out=./ --go_opt=module=github.com/goinsane/pbutil -I proto/ proto/goinsane/pbutil/*.proto

find types -depth -type f -name \*.pb.go -delete
find types -type d -empty -delete
protoc --go_out=./ --go_opt=module=google.golang.org/protobuf -I proto/ proto/google/protobuf/*.proto
protoc --go_out=./ --go_opt=module=google.golang.org/protobuf -I proto/ proto/google/protobuf/compiler/*.proto
