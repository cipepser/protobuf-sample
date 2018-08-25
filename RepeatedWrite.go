package main

import (
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/repeated"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &pb.User{
		Name: []string{"Alice", "Bob"},
	}

	if err := write("./repeated/02.bin", user); err != nil {
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
