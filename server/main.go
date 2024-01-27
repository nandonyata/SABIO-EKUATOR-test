package main

import (
	"fmt"
	"log"
	"net"
	_ "sabio-ekuator/server/config"

	pb "sabio-ekuator/proto"

	"google.golang.org/grpc"
)

var address string = "localhost:3003"

type Server struct {
	pb.ProductServiceServer
}

func main() {
	list, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Err listening address: %v\n", err)
	}

	fmt.Println("Listening ", address)

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Err serving: %v\n", err)
	}
}
