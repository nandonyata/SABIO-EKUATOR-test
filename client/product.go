package main

import (
	"context"
	"io"
	"log"
	pb "sabio-ekuator/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func productCreate(c pb.ProductServiceClient, p *pb.Product) {
	res, err := c.Create(context.Background(), p)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Print(res.Message)
}

func productGetOne(c pb.ProductServiceClient, p *pb.ProductId) {
	res, err := c.GetOne(context.Background(), p)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("Fetch one: %v\n", res)

}

func productGetAll(c pb.ProductServiceClient) {
	res, err := c.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		stream, err := res.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Err receiving server: %v\n", err)
		}

		log.Printf("Success: %v", stream)
	}
}
