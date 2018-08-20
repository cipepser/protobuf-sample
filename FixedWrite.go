package main

import (
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/fixed"
	"github.com/golang/protobuf/proto"
	"fmt"
	//"math"
)

func main() {
	user := &pb.User{
		//Id: 12345678901234567890, // uint64
		//Id: 1234567890,
		//Id: math.MaxUint64 - math.MaxUint32,
		//Id: math.MaxUint32,
		//Id: math.MaxUint64,
		Id: -1,
		//Id: math.MaxInt64 - math.MaxInt32 - 1<<31,
		//Id: math.MinInt64 - math.MinInt32,
	}

	fmt.Printf("%x\n", user.Id)

	if err := write("./fixed/sfixed64.bin", user); err != nil {
		log.Fatal(err)
	}
}

func write(file string, user *pb.User) error {
	out, err := proto.Marshal(user)
	if err != nil {
		return err
	}
	fmt.Printf("% x\n", out)
	if err := ioutil.WriteFile(file, out, 0644); err != nil {
		return err
	}
	return nil
}
