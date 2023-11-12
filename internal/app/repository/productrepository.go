package repository

import (
	"BwiBOW123/backend-miniproject/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	query := `
        INSERT INTO products (P_Name, Description,  P_Price, Member_ID, Category_ID)
        VALUES (?, ?, ?, ?, ?)
    `
	result := r.db.Exec(query)
	return result.Error
}
func (r *ProductRepository) UploadFile(memberEmail string, fileData []byte, fileName string) error {
	return nil
}
func (r *ProductRepository) UploadImage(memberEmail string, imageData []byte, imageName string) error {
	// Find the product by member email
	var product domain.Product
	result := r.db.Where("member_email = ?", memberEmail).First(&product)
	if result.Error != nil {
		return result.Error
	}

	// Update the image field
	product.Image = imageData
	result = r.db.Save(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
