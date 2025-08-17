package repository

import (
	"L0_task/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository { return &OrderRepository{db} }

func (r *OrderRepository) GetByID(uid string) (*models.Order, error) {
	var order models.Order

	err := r.db.
		Preload("Delivery").
		Preload("Payment").
		Preload("Items").
		Where("id = ?", uid).
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) Create(order *models.Order) error {
	err := r.db.Create(order).Error
	return err
}
