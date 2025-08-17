package models

type Item struct {
	ID          string `json:"-" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ChrtID      int    `json:"chrt_id" gorm:"not null"`
	OrderID     string `json:"-"`
	TrackNumber string `json:"track_number" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Rid         string `json:"rid" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Sale        int    `json:"sale" gorm:"not null"`
	Size        string `json:"size" gorm:"not null"`
	TotalPrice  int    `json:"total_price" gorm:"not null"`
	NmID        int    `json:"nm_id" gorm:"not null"`
	Brand       string `json:"brand" gorm:"not null"`
	Status      int    `json:"status" gorm:"not null"`
}
