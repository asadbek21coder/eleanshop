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
	query := "INSERT INTO feedbacks (name, phone_number, email, text) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := r.db.QueryRow(query, feedback.Name, feedback.PhoneNumber, feedback.Email, feedback.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *FeedbackPostgres) GetFeedbackById(id int) (models.Feedback, error) {
	var fb models.Feedback
	query := "SELECT * FROM feedbacks WHERE id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&fb.ID, &fb.Name, &fb.PhoneNumber, &fb.Email, &fb.Text)
	return fb, err
}

func (r *FeedbackPostgres) GetAllFeedbacks() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	query := "SELECT id, name, phone_number, email, text FROM feedbacks;"
	row, err := r.db.Query(query)
	for row.Next() {
		var fb models.Feedback
		_ = row.Scan(&fb.ID, &fb.Name, &fb.PhoneNumber, &fb.Email, &fb.Text)
		feedbacks = append(feedbacks, fb)
	}
	return feedbacks, err
}

func (r *FeedbackPostgres) UpdateFeedback(id int, input models.UpdateFeedbackInput) (models.Feedback, error) {
	var fb models.Feedback
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE feedbacks SET %s RETURNING id, name, phone_number, email, text;", setQuery)
	row := r.db.QueryRow(query, args...)
	err := row.Scan(&fb.ID, &fb.Name, &fb.PhoneNumber, &fb.Email, &fb.Text)
	return fb, err
}

func (r *FeedbackPostgres) DeleteFeedback(id int) (int, error) {
	query := "DELETE FROM feedbacks WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return id, err
}
