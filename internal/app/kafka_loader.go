package app

import (
	"L0_task/internal/config"
	"L0_task/internal/database"
	"L0_task/internal/handler/kafka_loader"
	"context"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

func RunOrderConsumer(consumerGroup sarama.ConsumerGroup, kafkaTopic string, cfg *config.Config) {
	handler := kafka_loader.NewOrderConsumerGroupHandler(database.MustLoad(cfg))
	ctx := context.WithoutCancel(context.Background()) // контекст для консумер группы
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() { // бесконечный шорт-полл топика
		defer wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, []string{kafkaTopic}, handler); err != nil {
				log.Printf("Can't consume: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			handler.Ready = make(chan bool) // пересоздаем на случай реконнекта
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
	<-handler.Ready
	<-ctx.Done()
}
