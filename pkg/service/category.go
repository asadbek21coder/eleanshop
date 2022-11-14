package service

import "github.com/asadbek21coder/eleanshop/pkg/repository"

type Category interface {
}



type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

