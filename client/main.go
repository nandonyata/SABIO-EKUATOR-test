package main

import (
	"log"

	pb "sabio-ekuator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3003"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed conecting address: %v\n", err)
	}

	defer conn.Close()
	clientProduct := pb.NewProductServiceClient(conn)

	// productCreate(clientProduct, &pb.Product{
	// 	Name:  "Samsung TV 32 inch",
	// 	Price: 1300000,
	// 	Stock: 2,
	// })

	// productGetOne(clientProduct, &pb.ProductId{Id: 1})

	productGetAll(clientProduct)
}
