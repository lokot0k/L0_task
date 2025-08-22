package kafka

import (
	"L0_task/internal/models"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

// ProduceOrdersToKafka Функция для отправки Order в брокер
func ProduceOrdersToKafka(producer sarama.SyncProducer, kafkaTopic string, data []*models.Order) error {
	for _, order := range data {
		jsonData, err := json.Marshal(order)
		if err != nil {
			return fmt.Errorf("failed to encode: %w", err)
		}

		msg := &sarama.ProducerMessage{
			Topic: kafkaTopic,
			Value: sarama.ByteEncoder(jsonData),
			Key:   sarama.StringEncoder(order.ID),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			return fmt.Errorf("failed to send: %w", err)
		}

		log.Printf("Sent message: ID=%s\n", order.ID)
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
