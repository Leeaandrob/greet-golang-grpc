package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Leeaandrob/greet-golang-grpc/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("vim-go")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
