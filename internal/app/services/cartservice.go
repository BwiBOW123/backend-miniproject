package service

import (
	"BwiBOW123/backend-miniproject/internal/app/repository"
	"BwiBOW123/backend-miniproject/internal/domain"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(cart *domain.Cart) error {
	return s.repo.CreateCart(cart)
}
func (s *CartService) CreateCartProduct(cartproduct *domain.CartProducts) error {
	return s.repo.CreateCartProducts(cartproduct)
}
func (s *CartService) GetCarts() ([]domain.Cart, error) {
	return s.repo.GetCarts()
}
func (s *CartService) GetCart(id string) (*domain.Cart, error) {
	return s.repo.GetCart(id)
}
func (s *CartService) GetCartProducts(id string) ([]domain.CartProductDetails, error) {
	return s.repo.GetCartProductsBybyProductID(id)
}
func (s *CartService) GetCartProductsEmail(email string) ([]domain.CartProductDetails, error) {
	return s.repo.GetCartProductsBybyEmail(email)
}
func (s *CartService) DeleteCartProducts(id string) error {
	return s.repo.DeleteCartProduct(id)
}
func (s *CartService) UpdateTotal() error {
	return s.repo.UpdateTotal()
}
