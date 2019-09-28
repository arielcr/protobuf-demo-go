package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	simplepb "github.com/arielcr/protobuf-demo-go/src/simple"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)

	jsonDemo(sm)
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "John",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "Renamed name!"

	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println("To JSON: ", smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Succesfully created proto struct from JSON:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Cant convert JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Could not unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin", sm2)

	fmt.Println("Read the content: ", sm2)
}

func writeToFile(name string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant serialize bytes", err)
		return err
	}

	if err := ioutil.WriteFile(name, out, 0644); err != nil {
		log.Fatalln("Cant write to file", err)
		return err
	}

	fmt.Println("Data has been written!")

	return nil
}

func readFromFile(name string, pb proto.Message) error {
	in, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln("Something went wrong reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Could not put the bytes into the protocol buffer struct", err)
		return err
	}

	return nil
}
