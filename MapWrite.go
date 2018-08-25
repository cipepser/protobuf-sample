package main

import (
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/map"
	"github.com/golang/protobuf/proto"
)

func main() {
	m := map[string]int32{}
	m["Alice"] = 20
	m["Bob"] = 25
	user := &pb.User{
		Name2Age: m,
	}

	if err := write("./map/02.bin", user); err != nil {
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
