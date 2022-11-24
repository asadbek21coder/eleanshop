package repository

import (
	"fmt"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.UserFull, error) {
	var user models.UserFull
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err

}

func (r *AuthPostgres) SetAdmin(input models.SetAdmin) error {

	query := `UPDATE users SET is_admin=$1 WHERE id=$2`
	_, err := r.db.Exec(query, input.IsAdmin, input.UserId)
	if err != nil {
		return err
	}

	return nil
}
