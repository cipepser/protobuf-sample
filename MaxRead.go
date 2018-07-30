package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/max"
	"github.com/golang/protobuf/proto"
)

func main() {
	if err := read("./max/int.bin"); err != nil {
		log.Fatal(err)
	}
}

func read(file string) error {
	in, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	user := &pb.User{}

	if err := proto.Unmarshal(in, user); err != nil {
		return err
	}
	fmt.Printf("0d%d\n0b%b\n", user.Id, user.Id)
	return nil
}
