package models

type Delivery struct {
	ID      string `json:"-" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string `json:"name" gorm:"not null"`
	Phone   string `json:"phone" gorm:"not null"`
	Zip     string `json:"zip" gorm:"not null"`
	City    string `json:"city" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	Region  string `json:"region" gorm:"not null"`
	Email   string `json:"email" gorm:"not null"`
	// считаем, что связь Delivery к Order - один ко многим,
	// потому что информация куда доставить может быть одинаковой для разных заказов
	Orders []Order `json:"-" gorm:"foreignKey:DeliveryID"`
}
