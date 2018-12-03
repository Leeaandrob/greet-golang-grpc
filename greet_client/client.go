package main

import (
	"fmt"
	"log"

	"github.com/Leeaandrob/greet-golang-grpc/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("vim-go")

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)

	}

	defer con.Close()
	c := greetpb.NewGreetServiceClient(con)

	fmt.Printf("Created client: %f", c)
}
