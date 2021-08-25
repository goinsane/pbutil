# pbutil

[![Go Reference](https://pkg.go.dev/badge/github.com/goinsane/pbutil.svg)](https://pkg.go.dev/github.com/goinsane/pbutil)

Package pbutil provides utilities for protobuf.
Please see [godoc](https://pkg.go.dev/github.com/goinsane/pbutil).

## Proto packages

The `proto` (`github.com/goinsane/pbutil/proto`) folder provides those major proto packages:
* **goinsane.pbutil** in `goinsane/pbutil/`
* **goinsane.pbutil.protobuf** (pbutil's protobuf) in `goinsane/pbutil/protobuf/`
* **google.protobuf** (Google's protobuf) in `google/protobuf/`

The `proto` folder can be used directly as include path for `protoc`. 

### goinsane.pbutil

Proto package **goinsane.pbutil** in `goinsane/pbutil/`, provides proto files for pbutil.

### goinsane.pbutil.protobuf

Proto package **goinsane.pbutil.protobuf** in `goinsane/pbutil/protobuf/`, provides
some extended types derived from **google.protobuf** (Google's protobuf).

Compiled output of **goinsane.pbutil.protobuf**, is the `types` folder (`github.com/goinsane/pbutil/types`)
that is same structure with `google.golang.org/protobuf/types`.

For example, you might use the proto file 
`goinsane/pbutil/protobuf/timestamp.proto` (*github.com/goinsane/pbutil/types/known/timestamppb*)
instead of
`google/protobuf/timestamp.proto` (*google.golang.org/protobuf/types/known/timestamppb*).
In this example, you must use *github.com/goinsane/pbutil/types/known/timestamppb* as Golang import path.

### google.protobuf

Proto package **google.protobuf** (Google's protobuf) in `google/protobuf/` provides proto files of Google's protobuf.

## Examples

To run any example, please use the command like the following:

    cd examples/
    go run example1.go
