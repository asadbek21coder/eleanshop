package repository

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type Size interface {
	CreateSize(models.SizeInput) (models.Size, error)
	GetAllSize() ([]models.Size, error)
	GetSizesById(int) (models.Size, error)
	DeleteSize(int) error
}

type SizePostgres struct {
	db *sqlx.DB
}

func NewSizePostgres(db *sqlx.DB) *SizePostgres {
	return &SizePostgres{db: db}
}

func (r *SizePostgres) CreateSize(size models.SizeInput) (models.Size, error) {
	var resp models.Size

	queryCreateSize := `INSERT INTO sizes (size_num) VALUES ($1) RETURNING *`

	row, err := r.db.Query(queryCreateSize, size.SizeNum)

	if err != nil {
		return models.Size{}, err
	}

	for row.Next() {
		err := row.Scan(
			&resp.ID,
			&resp.SizeNum,
		)
		if err != nil {
			return models.Size{}, err
		}
	}

	if err != nil {
		return models.Size{}, err
	}
	return resp, nil
}

func (r *SizePostgres) GetAllSize() ([]models.Size, error) {
	var resp []models.Size
	queryGetAllSizes := `SELECT * FROM sizes`
	row, err := r.db.Query(queryGetAllSizes)

	if err != nil {
		return []models.Size{}, nil
	}
	for row.Next() {
		var size models.Size
		err := row.Scan(
			&size.ID,
			&size.SizeNum,
		)

		if err != nil {
			return []models.Size{}, nil
		}
		resp = append(resp, size)
	}

	return resp, nil

}

func (r *SizePostgres) DeleteSize(id int) error {

	queryDeleteSize := `DELETE FROM sizes WHERE id=$1`

	_, err := r.db.Exec(queryDeleteSize, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SizePostgres) GetSizesById(id int) (models.Size, error) {
	var resp models.Size
	queryDeleteSize := `SELECT * FROM sizes WHERE id=$1`

	row := r.db.QueryRow(queryDeleteSize, id)
	err := row.Scan(&resp.ID, &resp.SizeNum)
	if err != nil {
		return models.Size{}, err
	}

	return resp, nil
}
