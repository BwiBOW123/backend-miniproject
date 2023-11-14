package domain

import (
	"gorm.io/gorm"
)

// Member table
type Member struct {
	gorm.Model           // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Password   string    `gorm:"size:255;not null" json:"password"`
	FullName   string    `gorm:"size:255;not null" json:"full_name"`
	Email      string    `gorm:"size:255;not null;unique" json:"email"`
	PayCost    float64   `gorm:"type:decimal(10,2)" json:"pay_cost"`
	Role       string    `gorm:"size:100;not null" json:"role"`
	Products   []Product `gorm:"foreignKey:MemberID" json:"products"` // One-to-many relationship with Product
}

// Product table
type Product struct {
	gorm.Model          // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `gorm:"size:255;not null" json:"name"`
	Description string  `gorm:"type:text;not null" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID  int     `json:"category_id"`                        // References Category
	MemberID    int     `json:"member_id"`                          // References Member
	CartID      int     `json:"Cart_id"`                            // References Member
	Images      []Image `gorm:"foreignKey:ProductID" json:"images"` // One-to-many relationship with Image
	Files       []File  `gorm:"foreignKey:ProductID" json:"files"`  // One-to-many relationship with File
}

// Payment model
type Payment struct {
	gorm.Model
	PaymentType string  `gorm:"size:100;not null" json:"payment_type"`
	Amount      float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	MemberID    uint    `json:"member_id"` // References Member
	CartID      uint    `json:"cart_id"`   // References Cart
}

// Cart model
type Cart struct {
	gorm.Model
	TotalCost float64   `gorm:"type:decimal(10,2);not null" json:"total_cost"`
	MemberID  uint      `json:"member_id"`                         // References Member
	Payments  []Payment `gorm:"foreignKey:CartID" json:"payments"` // One-to-many relationship with Payment
	Products  []Product `gorm:"foreignKey:CartID" json:"products"` // One-to-many relationship with Payment
}

// Category table
type Category struct {
	gorm.Model           // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string    `gorm:"size:255;not null" json:"name"`
	Products   []Product `gorm:"foreignKey:CategoryID" json:"products"` // One-to-many relationship with Product
}

// Image table
type Image struct {
	gorm.Model        // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	ProductID  uint   `json:"product_id"` // References Product
	Data       string `json:"data"`       // Image data
}

// File table
type File struct {
	gorm.Model        // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	ProductID  uint   `json:"product_id"` // References Product
	Data       string `json:"data"`       // File data
}

type ProductWithImageData struct {
	Name        string
	Price       float64
	Description string
	Data        []byte // Assuming image data is stored as a byte array
}
