package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(product models.ProductRequest) (int, error) {

	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	queryCreateProduct := `INSERT INTO products (product_name, category_id, price, color, count,image_url) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	row := tx.QueryRow(queryCreateProduct, product.ProductName, product.CategoryId, product.Price, product.Color, product.Count, product.ImageUrl)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryCreateSizes := `INSERT INTO product_sizes (product_id, size_id) VALUES ($1, $2)`

	for _, value := range product.Sizes {
		_, err := tx.Exec(queryCreateSizes, id, value)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *ProductPostgres) GetProductById(id int) (models.Product, error) {
	var resp models.Product
	var availableSizes []int
	tx, err := r.db.Begin()

	if err != nil {
		return models.Product{}, errors.New("transaction couldn`t begin")
	}

	queryGetProductById := `SELECT id, product_name,category_id, (select name from categories as cat where cat.id=p.category_id), price, color, count
	FROM products as p WHERE id=$1`
	queryGetAvailableSizes := `SELECT size_id FROM product_sizes WHERE product_id=$1`
	queryGetSizeNums := `SELECT size_num from sizes WHERE id=$1`

	row, err := tx.Query(queryGetAvailableSizes, id)
	if err != nil {
		tx.Rollback()
		return models.Product{}, err
	}

	for row.Next() {
		var sizeId int
		var num int
		err := row.Scan(
			&sizeId,
		)
		if err != nil {
			tx.Rollback()
			return models.Product{}, err
		}

		err = r.db.Get(&num, queryGetSizeNums, sizeId)
		if err != nil {
			tx.Rollback()
			return models.Product{}, err
		}
		availableSizes = append(availableSizes, num)

	}

	row, err = r.db.Query(queryGetProductById, id)
	if err != nil {
		tx.Rollback()
		return models.Product{}, err
	}
	for row.Next() {
		err := row.Scan(
			&resp.ID,
			&resp.ProductName,
			&resp.CategoryId,
			&resp.CategoryName,
			&resp.Price,
			&resp.Color,
			&resp.Count,
		)
		if err != nil {
			tx.Rollback()
			return models.Product{}, err
		}
	}
	resp.Sizes = availableSizes
	return resp, nil
}

func (r *ProductPostgres) GetAllProducts(queryParams models.QueryParams) ([]models.Product, error) {
	var (
		resp   []models.Product
		filter string
		params = make(map[string]interface{})
	)

	if queryParams.Search != "" {
		filter += " AND p.product_name ILIKE '%' || :search || '%' "
		params["search"] = queryParams.Search
	}
	countQuery := `SELECT count(1) FROM products as p WHERE true ` + filter

	var count int
	q, arr := ReplaceQueryParams(countQuery, params)
	err := r.db.QueryRow(q, arr...).Scan(
		&count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	queryGetAllProducts := `SELECT id, product_name,category_id, (select name from categories as cat where cat.id=p.category_id), price, color, count, image_url 
		FROM products as p WHERE true` + filter

	queryGetAllProducts += " LIMIT :limit OFFSET :offset"
	params["limit"] = queryParams.Limit
	params["offset"] = queryParams.Offset

	q, arr = ReplaceQueryParams(queryGetAllProducts, params)
	row, err := r.db.Query(q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	queryGetProductSizes := `SELECT size_id FROM product_sizes WHERE product_id=$1`
	queryGetSizeNums := `SELECT size_num from sizes WHERE id=$1`

	// row, err := r.db.Query(queryGetAllProducts)

	if err != nil {
		return []models.Product{}, err
	}
	for row.Next() {
		var size int
		var sizeNum int
		var availableSizes []int

		var product models.Product
		err := row.Scan(
			&product.ID,
			&product.ProductName,
			&product.CategoryId,
			&product.CategoryName,
			&product.Price,
			&product.Color,
			&product.Count,
			&product.ImageUrl,
		)
		if err != nil {
			return []models.Product{}, err
		}
		row, err := r.db.Query(queryGetProductSizes, product.ID)

		if err != nil {
			return []models.Product{}, err
		}

		for row.Next() {
			err := row.Scan(
				&size,
			)
			if err != nil {
				return []models.Product{}, err
			}

			err = r.db.Get(&sizeNum, queryGetSizeNums, size)
			if err != nil {
				return []models.Product{}, err
			}
			availableSizes = append(availableSizes, sizeNum)
		}
		product.Sizes = availableSizes
		resp = append(resp, product)
	}

	return resp, nil
}

func (r *ProductPostgres) UpdateProduct(id int, input models.ProductRequest) (int, error) {

	var resp int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.New("transacion couldn`t begin")
	}
	queryUpdateProduct := `UPDATE products SET product_name=$1, category_id=$2, price=$3, color=$4, count=$5, image_url=$6 WHERE id=$7 RETURNING id`
	queryDeleteProductSizes := `DELETE FROM product_sizes WHERE product_id=$1`
	queryInsertProductSizes := `INSERT INTO product_sizes (product_id, size_id) VALUES ($1,$2)`

	_, err = tx.Exec(queryDeleteProductSizes, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, value := range input.Sizes {
		_, err := tx.Exec(queryInsertProductSizes, id, value)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	row, err := tx.Query(queryUpdateProduct, input.ProductName, input.CategoryId, input.Price, input.Color, input.Count, input.ImageUrl, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for row.Next() {
		err := row.Scan(
			&resp,
		)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return resp, tx.Commit()
}

func (r *ProductPostgres) DeleteProduct(id int) (int, error) {

	queryDeleteProduct := `DELETE FROM products WHERE id=$1 RETURNING id`
	queryDeleteProductSizes := `DELETE FROM product_sizes WHERE product_id=$1`

	var resp int
	_, err := r.db.Exec(queryDeleteProductSizes, id)

	if err != nil {
		return 0, err
	}

	err = r.db.Get(&resp, queryDeleteProduct, id)
	if err != nil {
		return 0, err
	}

	return resp, nil
}
func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
