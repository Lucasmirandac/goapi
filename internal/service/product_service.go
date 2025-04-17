package service

import (
	"github.com/Lucasmirandac/go_api_ecommerce/internal/database"
	"github.com/Lucasmirandac/go_api_ecommerce/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	// Convert []entity.Product to []*entity.Product
	var productsPtr []*entity.Product
	for i := range products {
		productsPtr = append(productsPtr, &products[i])
	}
	return productsPtr, nil
}

func (ps *ProductService) GetProductByID(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, price, description, categoryID, imageURL)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(id string, name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, price, description, categoryID, imageURL)
	_, err := ps.ProductDB.UpdateProduct(id, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) DeleteProduct(id string) error {
	return ps.ProductDB.DeleteProduct(id)
}
