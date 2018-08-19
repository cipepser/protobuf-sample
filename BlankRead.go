package main

import (
	"fmt"
	"log"

	pb "github.com/cipepser/protobuf-sample/blank"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &pb.User{}
	in := []byte{}

	if err := proto.Unmarshal(in, user); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID: ", user.Id)
	fmt.Println("Name: ", user.Name)
	fmt.Println("Age: ", user.Age)
	fmt.Println("Contact: ", user.Contact)

	// fmt.Println("Phone: ", user.Contact.Phone)
	// panic: runtime error: invalid memory address or nil pointer dereference
}
