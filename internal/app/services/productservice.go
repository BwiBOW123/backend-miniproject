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

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	return s.repo.GetProduct(id)
}
func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.repo.GetProducts()
}
func (s *ProductService) GetProductwithImage() ([]domain.ProductWithImageData, error) {
	return s.repo.GetProductwithImage()
}
func (s *ProductService) GetProductwithImageByCat(Cat_id string) ([]domain.ProductWithImageData, error) {
	return s.repo.GetProductwithImageByCat(Cat_id)
}
func (s *ProductService) GetProductwithImageById(id string) ([]domain.ProductWithImageData, error) {
	return s.repo.GetProductwithImageById(id)
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repo.CreateProduct(product)
}
func (s *ProductService) UploadImage(image *domain.Image) error {
	return s.repo.UploadImage(image)
}
func (s *ProductService) UploadFile(file *domain.File) error {
	return s.repo.UploadFile(file)
}
