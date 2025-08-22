package order

import (
	"L0_task/internal/models"
	"fmt"
	"github.com/go-faker/faker/v4"
	"log"
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
			log.Printf(fmt.Sprintf("GenerateMockOrders error: %w", err))
			continue // на случай каких-то неудачных данных, печатаем ошибку и все
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
