package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(question models.UpdateQuestionInput) (int, error) {
	return s.repo.CreateQuestion(question)
}

func (s *QuestionService) GetQuestionById(id int) (models.Question, error) {
	return s.repo.GetQuestionById(id)
}

func (s *QuestionService) GetAllQuestions() ([]models.Question, error) {
	return s.repo.GetAllQuestions()
}

func (s *QuestionService) UpdateQuestion(id int, input models.UpdateQuestionInput) (models.Question, error) {
	return s.repo.UpdateQuestion(id, input)
}

func (s *QuestionService) DeleteQuestion(id int) (int, error) {
	return s.repo.DeleteQuestion(id)
}
