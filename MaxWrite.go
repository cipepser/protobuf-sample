package main

import (
	"io/ioutil"
	"log"
	"math"

	pb "github.com/cipepser/protobuf-sample/max"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &pb.User{
		// Id: math.MaxUint64 - math.MaxUint32,
		// Id: math.MaxInt64 - math.MaxInt32 - 1<<31,
		Id: math.MinInt64 - math.MinInt32,
	}

	if err := write("./max/int.bin", user); err != nil {
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
