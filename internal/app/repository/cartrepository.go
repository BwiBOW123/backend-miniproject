package repository

import (
	"BwiBOW123/backend-miniproject/internal/domain"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) CreateCart(cart *domain.Cart) error {
	return r.db.Create(cart).Error
}
func (r *CartRepository) CreateCartProducts(cartProduct *domain.CartProducts) error {
	return r.db.Create(cartProduct).Error
}
func (r *CartRepository) UpdateTotal() error {
	r.db.Exec(`DELETE FROM cart_products WHERE cart_products.deleted_at is not null`)
	err := r.db.Exec(`UPDATE carts
	SET total_cost = (
		SELECT COALESCE(SUM(products.Price * cart_products.Quantity), 0)
		FROM cart_products
		INNER JOIN products ON cart_products.product_id = products.id
		WHERE cart_products.cart_id = carts.id
	)`).Error
	return err
}
func (r *CartRepository) GetCarts() ([]domain.Cart, error) {
	var carts []domain.Cart
	result := r.db.Find(&carts)
	return carts, result.Error
}
func (r *CartRepository) GetCart(id string) (*domain.Cart, error) {
	var cart domain.Cart
	result := r.db.First(&cart, id)
	return &cart, result.Error
}
func (r *CartRepository) GetCartProductsBybyProductID(id string) ([]domain.CartProductDetails, error) {
	var cartProduct []domain.CartProductDetails
	result := r.db.Raw(`SELECT cart_products.id, products.id as ProductID, carts.total_cost, cart_products.quantity, products.name, products.description,  products.price, (SELECT images.data FROM images WHERE images.product_id = products.id LIMIT 1) as image_data FROM carts JOIN cart_products ON carts.id = cart_products.cart_id JOIN products ON cart_products.product_id = products.id WHERE carts.member_id = ? and cart_products.deleted_at is null`, id).Scan(&cartProduct)
	return cartProduct, result.Error
}
func (r *CartRepository) GetCartProductsBybyEmail(email string) ([]domain.CartProductDetails, error) {
	var cartProduct []domain.CartProductDetails
	result := r.db.Raw(`SELECT cart_products.id, products.id as ProductID, carts.total_cost, cart_products.quantity, products.name, products.description,  products.price, (SELECT images.data FROM images WHERE images.product_id = products.id LIMIT 1) as image_data,products.product_payment FROM carts JOIN cart_products ON carts.id = cart_products.cart_id JOIN products ON cart_products.product_id = products.id WHERE carts.member_email = ? and cart_products.deleted_at is null`, email).Scan(&cartProduct)
	return cartProduct, result.Error
}

func (r *CartRepository) DeleteCartProduct(id string) error {
	var cartP domain.CartProducts
	result := r.db.Where("id = ?", id).Delete(&cartP)
	return result.Error
}
