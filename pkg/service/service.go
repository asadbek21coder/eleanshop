package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
	Product
	Feedback
	Question
	Category
	Size
}

type Product interface {
	CreateProduct(product models.ProductRequest) (int, error)
	GetProductById(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(id int, product models.ProductRequest) (int, error)
	DeleteProduct(id int) (int, error)
}

type Feedback interface {
	CreateFeedback(feedback models.Feedback) (int, error)
	GetFeedbackById(id int) (models.Feedback, error)
	GetAllFeedbacks() ([]models.Feedback, error)
	UpdateFeedback(id int, input models.UpdateFeedbackInput) (models.Feedback, error)
	DeleteFeedback(id int) (int, error)
}
type Question interface {
	CreateQuestion(question models.Question) (int, error)
	GetQuestionById(id int) (models.Question, error)
	GetAllQuestions() ([]models.Question, error)
	UpdateQuestion(id int, input models.UpdateQuestionInput) (models.Question, error)
	DeleteQuestion(id int) (int, error)
}

type Category interface {
	CreateCategory(string) (int, error)
	GetCategoryById(int) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategory(int, string) (int, error)
	DeleteCategory(int) (int, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
		Feedback:      NewFeedbackService(repos.Feedback),
		Question:      NewQuestionService(repos.Question),
		Category:      NewCategoryService(repos.Category),
		Size:          NewSizeService(repos.Size),
	}
}
