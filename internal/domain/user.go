package domain

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string    `json:"username"`
	Password string    `json:"password"`
	FullName string    `json:"full_name"`
	Email    string    `gorm:"unique" json:"email"`
	Role     string    `json:"role"`
	Payments []Payment `gorm:"foreignKey:MemberID" json:"payments,omitempty"`
}

type Category struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"unique" json:"category_name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

type Product struct {
	gorm.Model          // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `json:"product_name"`
	Description string  `json:"description"`
	Image       []byte  `json:"image"` // Byte slice for binary image data
	File        []byte  `json:"file"`  // Byte slice for binary file data
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	MemberEmail string  `gorm:"not null" json:"member_email"`
	Member      Member  `gorm:"foreignKey:Email;references:MemberEmail" json:"member,omitempty"`
	Cart        []Cart  `gorm:"foreignKey:ProductID" json:"cart,omitempty"`
}

type Payment struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount      float64   `json:"amount"`
	PaymentType string    `json:"payment_type"`
	CreatedAt   time.Time `json:"created_at"`
	MemberID    uint      `json:"member_id"`
}

type Cart struct {
	gorm.Model         // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Quantity   int     `json:"quantity"`
	TotalCost  float64 `json:"total_cost"`
	ProductID  uint    `json:"product_id"`
	MemberID   uint    `json:"member_id"`
}
