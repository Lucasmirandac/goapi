package database

import (
	"database/sql"

	"github.com/Lucasmirandac/go_api_ecommerce/internal/entity"
)

type ProductDB struct {
	DB *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (pd *ProductDB) GetProducts() ([]entity.Product, error) {
	rows, err := pd.DB.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []entity.Product{}
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := pd.DB.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}

func (pd *ProductDB) GetProductByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.DB.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.DB.Query("SELECT id, name, price, description, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*entity.Product{}
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) UpdateProduct(id string, product *entity.Product) (string, error) {
	_, err := pd.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, image_url = ? WHERE id = ?", product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL, id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (pd *ProductDB) DeleteProduct(id string) error {
	_, err := pd.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
