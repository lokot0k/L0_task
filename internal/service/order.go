package service

import "L0_task/internal/models"

type OrderRepository interface {
	GetByID(uid string) (*models.Order, error)
	Create(order *models.Order) error
}

func GetOrderByID(uid string, repository OrderRepository) (*models.Order, error) {
	return repository.GetByID(uid)
}
