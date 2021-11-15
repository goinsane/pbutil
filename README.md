# pbutil

[![Go Reference](https://pkg.go.dev/badge/github.com/goinsane/pbutil.svg)](https://pkg.go.dev/github.com/goinsane/pbutil)

Package pbutil provides utilities for protobuf.
Please see [godoc](https://pkg.go.dev/github.com/goinsane/pbutil).

## Proto packages

The `proto` (`github.com/goinsane/pbutil/proto`) folder provides those proto packages:
* **goinsane.pbutil** in `goinsane/pbutil/`
* **google.protobuf** (Google's protobuf) in `google/protobuf/`

The `proto` folder can be used directly as include path for `protoc`. 

### goinsane.pbutil

Proto package **goinsane.pbutil** in `goinsane/pbutil/`, provides proto files for pbutil.

### google.protobuf

Proto package **google.protobuf** (Google's protobuf) in `google/protobuf/` provides proto files of Google's protobuf.

## Examples

To run any example, please use the command like the following:

    cd examples/
    go run example1.go

## Tests

To run all tests, please use the following command:

    go test -v

To run all examples, please use the following command:

    go test -v -run=^Example

To run all benchmarks, please use the following command:

    go test -v -run=^Benchmark -bench=.
