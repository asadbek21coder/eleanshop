package repository

import (
	"fmt"
	"strings"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type FeedbackPostgres struct {
	db *sqlx.DB
}

func NewFeedbackPostgres(db *sqlx.DB) *FeedbackPostgres {
	return &FeedbackPostgres{db: db}
}

func (r *FeedbackPostgres) CreateFeedback(feedback models.Feedback) (int, error) {
	var id int
	query := "INSERT INTO feedbacks (user_id, phone_number, email, text, product_id) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

	row := r.db.QueryRow(query, feedback.UserId, feedback.PhoneNumber, feedback.Email, feedback.Text, feedback.ProductId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *FeedbackPostgres) GetFeedbackById(id int) (models.Feedback, error) {
	var fb models.Feedback
	query := "SELECT * FROM feedbacks WHERE id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&fb.ID, &fb.UserId, &fb.PhoneNumber, &fb.Email, &fb.Text, &fb.ProductId)
	return fb, err
}

func (r *FeedbackPostgres) GetAllFeedbacks() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	query := "SELECT id, user_id, phone_number, email, text, product_id FROM feedbacks;"
	row, err := r.db.Query(query)
	for row.Next() {
		var fb models.Feedback
		_ = row.Scan(&fb.ID, &fb.UserId, &fb.PhoneNumber, &fb.Email, &fb.Text, &fb.ProductId)
		feedbacks = append(feedbacks, fb)
	}
	return feedbacks, err
}

func (r *FeedbackPostgres) UpdateFeedback(id int, input models.UpdateFeedbackInput) (models.Feedback, error) {
	var fb models.Feedback
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 2
	args = append(args, id)

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argID))
		args = append(args, input.UserId)
		argID++
	}

	if input.PhoneNumber != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argID))
		args = append(args, *input.PhoneNumber)
		argID++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Email)
		argID++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argID))
		args = append(args, *input.Text)
		argID++
	}

	if input.ProductId != nil {
		setValues = append(setValues, fmt.Sprintf("product_id=$%d", argID))
		args = append(args, *input.ProductId)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE feedbacks SET %s where id=$1 RETURNING id, user_id, phone_number, email, text, product_id;", setQuery)
	row := r.db.QueryRow(query, args...)
	err := row.Scan(&fb.ID, &fb.UserId, &fb.PhoneNumber, &fb.Email, &fb.Text, &fb.ProductId)
	return fb, err
}

func (r *FeedbackPostgres) DeleteFeedback(id int) (int, error) {
	query := "DELETE FROM feedbacks WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return id, err
}
