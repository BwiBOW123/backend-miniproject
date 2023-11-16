package repository

import (
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {

	result := r.db.Create(product) // Use the product pointer directly

	if result.Error != nil {
		logs.Error("Failed to create product")
		return result.Error
	}

	return nil
}

func (r *ProductRepository) GetProduct(id string) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	return &product, result.Error
}

func (r *ProductRepository) GetProductwithImage() ([]domain.ProductWithImageData, error) {
	var products []domain.ProductWithImageData
	result := r.db.Raw("SELECT p.id, p.name,p.price,p.description, i.data FROM products p RIGHT JOIN (SELECT product_id, MAX(id) AS image_id FROM images GROUP BY product_id ) AS img ON p.id = img.product_id LEFT JOIN images i ON img.image_id = i.id;").Scan(&products)
	return products, result.Error
}
func (r *ProductRepository) GetProductwithImageByCat(cat_id string) ([]domain.ProductWithImageData, error) {
	var products []domain.ProductWithImageData
	result := r.db.Raw(`SELECT p.id, p.name, p.price, p.description, i.data FROM products p RIGHT JOIN (SELECT product_id, MAX(id) AS image_id FROM images GROUP BY product_id) AS img ON p.id = img.product_id LEFT JOIN images i ON img.image_id = i.id WHERE p.category_id = ?;`, cat_id).Scan(&products)
	return products, result.Error
}
func (r *ProductRepository) GetProductwithImageById(id string) ([]domain.ProductWithImageData, error) {
	var products []domain.ProductWithImageData
	result := r.db.Raw(`SELECT
    p.id,
    p.name,
    p.price,
    p.description,
    i.data,
	p.product_payment
FROM
    products p
LEFT JOIN
    images i ON p.id = i.product_id
WHERE
    p.id = ?;`, id).Scan(&products)
	return products, result.Error
}

func (r *ProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) UploadImage(image *domain.Image) error {

	result := r.db.Create(image)

	if result.Error != nil {
		logs.Error("Failed to Upload Image")
		return result.Error
	}

	return nil
}
func (r *ProductRepository) UploadFile(file *domain.File) error {

	result := r.db.Create(file)

	if result.Error != nil {
		logs.Error("Failed to Upload Image")
		return result.Error
	}

	return nil
}
