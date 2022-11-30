package repository

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.UserFull, error)
	SetAdmin(input models.SetAdmin) error
}

type Category interface {
	CreateCategory(string) (int, error)
	GetCategoryById(int) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategory(int, string) (int, error)
	DeleteCategory(int) (int, error)
}

type Product interface {
	CreateProduct(product models.ProductRequest) (int, error)
	GetProductById(id int) (models.Product, error)
	GetAllProducts(models.QueryParams) ([]models.Product, error)
	UpdateProduct(id int, product models.ProductRequest) (int, error)
	DeleteProduct(id int) (int, error)
}

type Feedback interface {
	CreateFeedback(feedback models.UpdateFeedbackInput, userId int) (int, error)
	GetFeedbackById(id int) (models.Feedback, error)
	GetAllFeedbacks() ([]models.Feedback, error)
	UpdateFeedback(id int, input models.UpdateFeedbackInput, userId *int) (models.Feedback, error)
	DeleteFeedback(id int, userId int) (int, error)
}
type Question interface {
	CreateQuestion(question models.UpdateQuestionInput) (int, error)
	GetQuestionById(id int) (models.Question, error)
	GetAllQuestions() ([]models.Question, error)
	UpdateQuestion(id int, input models.UpdateQuestionInput) (models.Question, error)
	DeleteQuestion(id int) (int, error)
}

type Repository struct {
	Authorization
	Product
	Feedback
	Question
	Category
	Size
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductPostgres(db),
		Feedback:      NewFeedbackPostgres(db),
		Question:      NewQuestionPostgres(db),
		Category:      NewCategoryPostgres(db),
		Size:          NewSizePostgres(db),
	}
}
