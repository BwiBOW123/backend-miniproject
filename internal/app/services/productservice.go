package service

import (
	"BwiBOW123/backend-miniproject/internal/app/repository"
	"BwiBOW123/backend-miniproject/internal/domain"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repo.CreateProduct(product)
}
