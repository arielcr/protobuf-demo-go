package main

import (
	"fmt"

	example_simple "github.com/arielcr/protobuf-demo-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "John",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "Renamed name!"

	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())

}
