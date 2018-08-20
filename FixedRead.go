package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/fixed"
	"github.com/golang/protobuf/proto"
)

func main() {
	if err := read("./fixed/fixed32.bin"); err != nil {
		log.Fatal(err)
	}
	if err := read("./fixed/fixed64.bin"); err != nil {
		log.Fatal(err)
	}
	if err := read("./fixed/sfixed32.bin"); err != nil {
		log.Fatal(err)
	}
	if err := read("./fixed/sfixed64.bin"); err != nil {
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
	fmt.Printf("%d\n", user.Id)
	return nil

}
