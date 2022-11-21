package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type FeedbackService struct {
	repo        repository.Feedback
	productRepo repository.AuthPostgres
}

func NewFeedbackService(repo repository.Feedback) *FeedbackService {
	return &FeedbackService{repo: repo}
}

func (s *FeedbackService) CreateFeedback(feedback models.Feedback) (int, error) {
	return s.repo.CreateFeedback(feedback)
}

func (s *FeedbackService) GetFeedbackById(id int) (models.Feedback, error) {
	return s.repo.GetFeedbackById(id)
}

func (s *FeedbackService) GetAllFeedbacks() ([]models.Feedback, error) {
	return s.repo.GetAllFeedbacks()
}

func (s *FeedbackService) UpdateFeedback(id int, input models.UpdateFeedbackInput) (models.Feedback, error) {
	return s.repo.UpdateFeedback(id, input)
}

func (s *FeedbackService) DeleteFeedback(id int) (int, error) {
	return s.repo.DeleteFeedback(id)
}
