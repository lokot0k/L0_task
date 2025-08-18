package models

import (
	"time"
)

type Order struct {
	ID                string    `json:"order_uid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	DeliveryID        string    `json:"-" gorm:"not null"` // поле для связи с доставкой
	TrackNumber       string    `json:"track_number" gorm:"not null;unique"`
	Entry             string    `json:"entry" gorm:"not null"`
	Delivery          Delivery  `json:"delivery" gorm:"foreignKey:DeliveryID"`
	Payment           Payment   `json:"payment" gorm:"foreignKey:Transaction"`
	Items             []Item    `json:"items" gorm:"foreignKey:TrackNumber;references:TrackNumber"`
	Locale            string    `json:"locale" gorm:"not null"`
	InternalSignature string    `json:"internal_signature" gorm:"not null"`
	CustomerID        string    `json:"customer_id" gorm:"not null"`
	DeliveryService   string    `json:"delivery_service" gorm:"not null"`
	Shardkey          string    `json:"shardkey" gorm:"not null"`
	SmID              int       `json:"sm_id" gorm:"not null"`
	DateCreated       time.Time `json:"date_created" gorm:"not null"`
	OofShard          string    `json:"oof_shard" gorm:"not null"`
}

type CachableOrder struct {
	*Order
	LastUsed time.Time
}

func (o *CachableOrder) Key() string {
	return o.ID
}

func (o *CachableOrder) Value() interface{} {
	return o.Order
}

func (o *CachableOrder) UpdatePriority() {
	o.LastUsed = time.Now()
}

func (o *CachableOrder) Priority() int64 {
	return o.LastUsed.UnixNano()
}
