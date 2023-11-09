package domain

import "time"

type Member struct {
	ID_USER  uint      `gorm:"primaryKey" json:"id_user"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Pay_CAST float32   `json:"pay_cast"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Payment struct {
	Payment_ID  uint      `gorm:"primaryKey" json:"payment_id"`
	Amount      float64   `json:"amount"`
	PaymentType string    `json:"payment_type"`
	CreateAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Product struct {
	P_ID        uint      `gorm:"primaryKey" json:"p_id"`
	P_Name      string    `json:"p_name"`
	Description string    `gorm:"type:text" json:"description"`
	Image       []byte    `gorm:"type:blob" json:"image"`
	P_Price     float64   `json:"p_price"`
	File        []byte    `gorm:"type:blob" json:"file"`
	CategoryID  uint      `json:"category_id"`
	CreateAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Category struct {
	Category_ID uint      `gorm:"primaryKey" json:"category_id"`
	Name        string    `json:"name"`
	CreateAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Cart struct {
	Cart_ID   uint      `gorm:"primaryKey" json:"cart_id"`
	Quantity  int       `json:"quantity"`
	TotalCost float64   `json:"total_cost"`
	ProductID uint      `json:"product_id"`
	PaymentID uint      `json:"payment_id"`
	UserID    uint      `json:"user_id"`
	CreateAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
