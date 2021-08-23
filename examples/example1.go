// +build examples

package main

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	_ "github.com/goinsane/pbutil/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	fmt.Println(protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("google.protobuf.Duration")))
	//protoregistry.GlobalTypes.RegisterMessage(protoreflect.MessageType())
}
