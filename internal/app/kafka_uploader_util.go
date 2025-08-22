package app

import (
	"L0_task/internal/models"
	"L0_task/tools/kafka"
	"github.com/IBM/sarama"
)

func UploadOrders(producer sarama.SyncProducer, kafkaTopic string, data []*models.Order) {
	err := kafka.ProduceOrdersToKafka(producer, kafkaTopic, data)
	if err != nil {
		panic(err)
	}
}
