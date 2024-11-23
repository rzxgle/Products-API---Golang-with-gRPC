package repository

import (
	"apiproducts/src/pb/products"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const filename string = "products.txt"

func (pr *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("error on read file. error: %+v", err)
	}

	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("error on unmarshal. error: %+v", err)
	}

	return productList, nil
}

func (pr *ProductRepository) saveDate(productsList products.ProductList) error {
	data, err := proto.Marshal(&productsList)
	if err != nil {
		return fmt.Errorf("error on marshal. error: %+v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Errorf("error on write file. error:  %+v", err)
	}

	return nil
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := pr.loadData()
	if err != nil {
		return product, err
	}

	product.Id = int32(len(productList.Products) + 1)
	productList.Products = append(productList.Products, &product)
	err = pr.saveDate(productList)

	return product, err
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	return pr.loadData()
}
