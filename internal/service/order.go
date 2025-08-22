package service

import (
	"L0_task/internal/models"
)

type orderRepository interface {
	GetByID(uid string) (*models.Order, error)
	Create(order *models.Order) error
}

func GetOrderByID(uid string, repository orderRepository) (*models.Order, error) {
	return repository.GetByID(uid)
}

func CreateOrder(order *models.Order, repository orderRepository) error {
	return repository.Create(order)
}
