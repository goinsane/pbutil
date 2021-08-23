// +build examples

package main

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	fmt.Println(protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("google.protobuf.Duration")))
	//protoregistry.GlobalTypes.RegisterMessage(protoreflect.MessageType())
}
