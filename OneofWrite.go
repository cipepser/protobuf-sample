package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/oneof"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &pb.User{
		//Result: &pb.User_Ok{Ok: "Alice"},
		Result: &pb.User_Err{Err: "Alice"},
	}
	user.
	fmt.Println(user.GetResult())

	if err := write("./oneof/03.bin", user); err != nil {
		log.Fatal(err)
	}
}

func write(file string, user *pb.User) error {
	out, err := proto.Marshal(user)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(file, out, 0644); err != nil {
		return err
	}
	return nil
}
