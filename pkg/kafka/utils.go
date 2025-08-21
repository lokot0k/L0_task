package kafka

import (
	"L0_task/internal/models"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

func CreateProducer(kafkaBrokers string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Compression = sarama.CompressionSnappy

	producer, err := sarama.NewSyncProducer([]string{kafkaBrokers}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return producer, nil
}

func SendOrdersToKafka(producer sarama.SyncProducer, kafkaTopic string, data []*models.Order) error {
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
