package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(name string) (int, error) {
	return s.repo.CreateCategory(name)
}

func (s *CategoryService) GetCategoryById(id int) (models.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) UpdateCategory(id int, name string) (int, error) {
	return s.repo.UpdateCategory(id, name)
}

func (s *CategoryService) DeleteCategory(id int) (int, error) {
	return s.repo.DeleteCategory(id)
}
