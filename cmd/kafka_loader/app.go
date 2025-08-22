package main

import (
	"L0_task/internal/app"
	"L0_task/internal/config"
	"L0_task/tools/kafka"
	"github.com/IBM/sarama"
)

func main() {
	cfg := config.LoadConfig()
	kafkaConsumerGroup, err := kafka.CreateConsumerGroup(cfg.KafkaBrokers, cfg.KafkaConsumerGroup)
	if err != nil {
		panic(err)
	}
	defer func(kafkaConsumer sarama.ConsumerGroup) {
		err := kafkaConsumer.Close()
		if err != nil {
			panic(err)
		}
	}(kafkaConsumerGroup)
	app.RunOrderConsumer(kafkaConsumerGroup, cfg.KafkaTopic, cfg)
}
