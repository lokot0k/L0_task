package kafka_loader

import (
	"L0_task/internal/models"
	"L0_task/internal/repository"
	"L0_task/internal/service"
	"encoding/json"
	"github.com/IBM/sarama"
	"gorm.io/gorm"
	"log"
)

// OrderConsumerGroupHandler выполняет контракт интерфейса для работы с ConsumerGroup
type OrderConsumerGroupHandler struct {
	db    *gorm.DB
	Ready chan bool // канал для сигнала о готовности консьюмера
}

func NewOrderConsumerGroupHandler(db *gorm.DB) *OrderConsumerGroupHandler {
	return &OrderConsumerGroupHandler{db, make(chan bool)}
}

func (consumer *OrderConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(consumer.Ready) // закрывая канал, сигнализируем о готовности консумера
	return nil
}

func (consumer *OrderConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	// для нашего кейса нечего очищать - оставляем stub для выполнения контракта
	return nil
}

func (consumer *OrderConsumerGroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for message := range claim.Messages() {
		session.MarkMessage(message, "") // помечаем сразу, чтобы никогда больше не процессить невалидные данные
		var order models.Order
		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Printf("Failed to parse data to Order: %v", err)
			continue // пропускаем невалидные сообщения
		}
		err := service.CreateOrder(&order, repository.NewOrderRepository(consumer.db))
		if err != nil {
			log.Printf("Failed to create order: %v", err)
			continue
		}
		log.Printf("Received order with uuid: %s", order.ID)
	}
	return nil
}
