package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "sabio-ekuator/proto"
	"sabio-ekuator/server/config"
	"sabio-ekuator/server/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Create(ctx context.Context, in *pb.Product) (*pb.ResponseStringProduct, error) {
	log.Printf("Create product was invoked with: %v\n", in)

	sqlQuery := `
		INSERT INTO "Product" (name, price, stock)
		VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(sqlQuery, in.Name, in.Price, in.Stock)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Error inserting data: %v", err),
		)
	}

	return &pb.ResponseStringProduct{Message: "Success create product"}, nil
}

func (s *Server) GetOne(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	log.Printf("Get one product was invoked with: %v\n", in)

	sqlQuery := `
		SELECT * FROM Product
		WHERE id = $1
	`

	product := &entity.Product{}
	err := config.DB.QueryRow(sqlQuery, in.Id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Product not found")
	}
	if err == nil {
		return nil, status.Error(codes.Internal, "Error get one")
	}

	return entity.ResultProduct(product), nil
}

func (s *Server) GetAll(_ *emptypb.Empty, stream pb.ProductService_GetAllServer) error {

	sqlQuery := `
		SELECT * FROM product
		ORDER BY id DESC
	`
	row, err := config.DB.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var product entity.Product

		if err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			log.Fatal(err)
			return err
		}

		stream.Send(&pb.Product{
			Name: product.Name,
		})
	}

	return nil
}
