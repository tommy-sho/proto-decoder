package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/tommy-sho/proto-decoder/_proto"
)

func main() {
	p := &pb.Person{
		Name: &pb.Name{
			Value: "Alice",
		},
		Age: &pb.Age{
			Value: 132,
		},
	}
	if err := write("./person/alice.bin", p); err != nil {
		log.Fatal(err)
	}
}

func write(file string, p *pb.Person) error {
	out, err := proto.Marshal(p)
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	if err := ioutil.WriteFile(file, out, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
