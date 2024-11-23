package main

import (
	"apiproducts/src/pb/products"
	"apiproducts/src/repository"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		return product, err
	}

	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	productList, err := s.productRepo.FindAll()
	return &productList, err
}

func main() {
	fmt.Println("Starting grpc server")
	srv := server{productRepo: &repository.ProductRepository{}}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("error on create listener. error: ", err)
	}

	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on server. error: ", err)
	}
}
