package utils

import (
	"L0_task/internal/models"
	"fmt"
	"github.com/go-faker/faker/v4"
)

// GenerateMockOrders вспомогательная функция для генерирования мока Order
func GenerateMockOrders(count int) []*models.Order {
	testdata := make([]*models.Order, count)
	for i := 0; i < count; i++ {
		var order models.Order
		var item models.Item
		var secondItem models.Item
		err := faker.FakeData(&order)
		if err != nil {
			panic(fmt.Errorf("GenerateMockOrders error: %w", err))
		}
		err = faker.FakeData(&item)
		err = faker.FakeData(&secondItem)
		item.TrackNumber = order.TrackNumber
		secondItem.TrackNumber = order.TrackNumber
		order.Items = []models.Item{item, secondItem}
		order.ID = faker.UUIDHyphenated()
		order.Payment.Transaction = order.ID
		testdata[i] = &order
	}
	return testdata
}
