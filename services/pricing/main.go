package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting...")
	lst, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	if err := s.Serve(lst); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	fmt.Println("Started!")
	defer fmt.Println("Ending...")
}
