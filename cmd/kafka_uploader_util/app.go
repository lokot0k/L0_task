package main

import (
	"L0_task/internal/app"
	"L0_task/internal/config"
	"L0_task/pkg/kafka"
	"L0_task/tools/order"
	"github.com/IBM/sarama"
)

const (
	uploadCount = 3 // подразумевается, что это утилитарный тестовый сервис, поэтому в .env не выносим
)

func main() {
	cfg := config.LoadConfig()
	kafkaProducer, err := kafka.CreateProducer(cfg.KafkaBrokers)
	if err != nil {
		panic(err)
	}
	defer func(kafkaProducer sarama.SyncProducer) {
		err := kafkaProducer.Close()
		if err != nil {
			panic(err)
		}
	}(kafkaProducer)
	orders := order.GenerateMockOrders(uploadCount)
	app.UploadOrders(kafkaProducer, cfg.KafkaTopic, orders)
}
