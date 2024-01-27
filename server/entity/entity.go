package entity

import (
	pb "sabio-ekuator/proto"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Order represents the order entity.
type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ResultOrder(o *Order) *pb.Order {
	createdAt := timestamppb.New(o.CreatedAt)
	updatedAt := timestamppb.New(o.UpdatedAt)

	return &pb.Order{
		Id:        int32(o.ID),
		ProductId: int32(o.ProductID),
		Quantity:  int32(o.Quantity),
		Total:     float32(o.Total),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// Product represents the product entity.
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ResultProduct(p *Product) *pb.Product {
	createdAt := timestamppb.New(p.CreatedAt)
	updatedAt := timestamppb.New(p.UpdatedAt)

	return &pb.Product{
		Id:        int32(p.ID),
		Name:      p.Name,
		Price:     p.Price,
		Stock:     int32(p.Stock),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// Customer represents the customer entity.
type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ResultCustomer(c *Customer) *pb.Customer {
	createdAt := timestamppb.New(c.CreatedAt)
	updatedAt := timestamppb.New(c.UpdatedAt)

	return &pb.Customer{
		Id:        int32(c.ID),
		Name:      c.Name,
		Email:     c.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
