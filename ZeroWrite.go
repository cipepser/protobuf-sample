// ZeroWrite.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/zero"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Name: &pb.Name{Value: "Alice"},
		Age:  &pb.Age{Value: 20},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	if err := ioutil.WriteFile("./zero/person.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write:", err)
	}
}
