package repository

import (
	"fmt"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategory(name string) (int, error) {

	queryCreateSize := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.Get(&id, queryCreateSize, name)

	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *CategoryPostgres) GetCategoryById(id int) (models.Category, error) {
	var resp models.Category
	queryCreateSize := `SELECT * FROM categories WHERE id=$1`
	row := r.db.QueryRow(queryCreateSize, id)

	err := row.Scan(
		&resp.ID,
		&resp.CategoryName,
	)
	if err != nil {
		return models.Category{}, err
	}

	fmt.Println(resp)

	return resp, nil

}

func (r *CategoryPostgres) GetAllCategories() ([]models.Category, error) {
	var resp []models.Category

	queryGetAllCategories := `SELECT * FROM categories`

	row, err := r.db.Query(queryGetAllCategories)
	if err != nil {
		return []models.Category{}, err
	}

	for row.Next() {
		var category models.Category
		err := row.Scan(
			&category.ID,
			&category.CategoryName,
		)
		if err != nil {
			return []models.Category{}, err
		}
		resp = append(resp, category)

	}

	return resp, nil
}

func (r *CategoryPostgres) UpdateCategory(id int, name string) (int, error) {
	queryCreateSize := `UPDATE categories SET name=$1 WHERE id=$2 RETURNING id`
	err := r.db.Get(&id, queryCreateSize, name, id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CategoryPostgres) DeleteCategory(id int) (int, error) {
	var resp int
	queryCreateSize := `DELETE FROM categories WHERE id=$1 RETURNING id`
	err := r.db.Get(&resp, queryCreateSize, id)

	if err != nil {
		return 0, err
	}

	return resp, nil

}
