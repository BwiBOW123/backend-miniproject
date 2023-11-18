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
	Card       string    `gorm:"size:255;" json:"card"`
	Products   []Product `gorm:"foreignKey:MemberID" json:"products"`
}

// Product table
type Product struct {
	gorm.Model             // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name           string  `gorm:"size:255;not null" json:"name"`
	Description    string  `gorm:"type:text;not null" json:"description"`
	Price          float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID     int     `json:"category_id"`                        // References Category
	MemberID       int     `json:"member_id"`                          // References Member
	CartID         int     `json:"cart_id"`                            // References Member
	UrlDownload    string  `json:"url_download"`                       // References Member
	ProductPayment string  `json:"product_Payment"`                    // References Member                              // References Member
	Images         []Image `gorm:"foreignKey:ProductID" json:"images"` // One-to-many relationship with Image
	Files          []File  `gorm:"foreignKey:ProductID" json:"files"`  // One-to-many relationship with File
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
	TotalCost   float64   `gorm:"type:decimal(10,2);" json:"total_cost"`
	MemberID    uint      `json:"member_id"`
	MemberEmail string    `json:"member_email"`                      // References Member
	Payments    []Payment `gorm:"foreignKey:CartID" json:"payments"` // One-to-many relationship with Payment
	Products    []Product `gorm:"foreignKey:CartID" json:"products"` // One-to-many relationship with Payment
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
	ID             int
	Name           string
	Price          float64
	Description    string
	Data           []byte
	ProductPayment string
}

type CartProducts struct {
	gorm.Model        // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	CartID     uint   `gorm:"not null" json:"cart_id"`    // References Cart
	ProductID  uint   `gorm:"not null" json:"product_id"` // References Product
	Quantity   int    `gorm:"not null" json:"quantity"`   // Quantity of the product in the cart
	Email      string `gorm:"not null" json:"email"`
}

type CartProductDetails struct {
	ID             float64 `gorm:"column:id"`
	ProductID      float64 `gorm:"column:ProductID"`
	TotalCost      float64 `gorm:"column:total_cost"`
	Quantity       int     `gorm:"column:quantity"`
	ProductName    string  `gorm:"column:name"`
	ProductDesc    string  `gorm:"column:description"`
	ProductPrice   float64 `gorm:"column:price"`
	ProductImage   string  `gorm:"column:image_data"`
	ProductPayment string  `gorm:"column:product_payment"`
}
