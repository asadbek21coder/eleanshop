package repository

import (
	"fmt"
	"strings"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type QuestionPostgres struct {
	db *sqlx.DB
}

func NewQuestionPostgres(db *sqlx.DB) *QuestionPostgres {
	return &QuestionPostgres{db: db}
}

func (r *QuestionPostgres) CreateQuestion(question models.UpdateQuestionInput) (int, error) {
	var id int
	query := "INSERT INTO questions (name, phone_number, time, text) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := r.db.QueryRow(query, question.Name, question.PhoneNumber, question.Time, question.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *QuestionPostgres) GetQuestionById(id int) (models.Question, error) {
	var qt models.Question
	query := "SELECT * FROM questions WHERE id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&qt.ID, &qt.Name, &qt.PhoneNumber, &qt.Time, &qt.Text)
	return qt, err
}

func (r *QuestionPostgres) GetAllQuestions() ([]models.Question, error) {
	var questions []models.Question
	query := "SELECT id, name, phone_number, time, text FROM questions;"
	row, err := r.db.Query(query)
	for row.Next() {
		var qt models.Question
		_ = row.Scan(&qt.ID, &qt.Name, &qt.PhoneNumber, &qt.Time, &qt.Text)
		questions = append(questions, qt)
	}
	return questions, err
}

func (r *QuestionPostgres) UpdateQuestion(id int, input models.UpdateQuestionInput) (models.Question, error) {
	var qt models.Question
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 2

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

	if input.Time != nil {
		setValues = append(setValues, fmt.Sprintf("time=$%d", argID))
		args = append(args, *input.Time)
		argID++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argID))
		args = append(args, *input.Text)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE questions SET %s RETURNING id, name, phone_number, time, text where id=$1", setQuery)
	row := r.db.QueryRow(query, args...)
	err := row.Scan(&qt.ID, &qt.Name, &qt.PhoneNumber, &qt.Time, &qt.Text)
	return qt, err
}

func (r *QuestionPostgres) DeleteQuestion(id int) (int, error) {
	query := "DELETE FROM questions WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return id, err
}
