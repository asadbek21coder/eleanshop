package service

import (
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (r *ProductService) CreateProduct(product models.ProductRequest) (int, error) {
	return r.repo.CreateProduct(product)
}

func (r *ProductService) GetProductById(id int) (models.Product, error) {

	return r.repo.GetProductById(id)
}

func (r *ProductService) GetAllProducts() ([]models.Product, error) {

	return r.repo.GetAllProducts()
}

func (r *ProductService) UpdateProduct(id int, input models.ProductRequest) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}
	return r.repo.UpdateProduct(id, input)
}

func (r *ProductService) DeleteProduct(id int) (int, error) {

	return r.repo.DeleteProduct(id)
}
