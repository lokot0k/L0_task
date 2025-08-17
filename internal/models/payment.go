package models

type Payment struct {
	ID           string `json:"transaction" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RequestID    string `json:"request_id" gorm:"not null"`
	Currency     string `json:"currency" gorm:"not null"`
	Provider     string `json:"provider" gorm:"not null"`
	Amount       int    `json:"amount" gorm:"not null"`
	PaymentDt    int    `json:"payment_dt" gorm:"not null"`
	Bank         string `json:"bank" gorm:"not null"`
	DeliveryCost int    `json:"delivery_cost" gorm:"not null"`
	GoodsTotal   int    `json:"goods_total" gorm:"not null"`
	CustomFee    int    `json:"custom_fee" gorm:"not null"`
	OrderID      string `json:"-"` // связь с Order 1 к 1
}
