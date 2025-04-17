package database

import (
	"database/sql"
	"log"

	"github.com/Lucasmirandac/go_api_ecommerce/internal/entity"
)

type CategoryDB struct {
	DB *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{DB: db}
}

func (cd *CategoryDB) GetCategories() ([]entity.Category, error) {
	rows, err := cd.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // defer is used to close the rows after the function is executeds

	categories := []entity.Category{}
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (cd *CategoryDB) GetCategoryByID(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	log.Println("Creating category", category.Name)
	log.Println("Category ID", category.ID)
	_, err := cd.DB.Exec("INSERT INTO categories (id, name) VALUES (?,?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
