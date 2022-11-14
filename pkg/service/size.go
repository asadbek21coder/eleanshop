package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type SizeService struct {
	repo repository.Size
}

func NewSizeService(repo repository.Size) *SizeService {
	return &SizeService{repo: repo}
}

type Size interface {
	CreateSize(models.SizeInput) (models.Size, error)
	GetAllSize() ([]models.Size, error)
	DeleteSize(int) error
}

func (r *SizeService) CreateSize(input models.SizeInput) (models.Size, error) {
	return r.repo.CreateSize(input)
}

func (r *SizeService) GetAllSize() ([]models.Size, error) {
	return r.repo.GetAllSize()

}

func (r *SizeService) DeleteSize(id int) error {
	return r.repo.DeleteSize(id)
}
