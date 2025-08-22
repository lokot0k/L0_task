package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

// CreateConsumerGroup используем группу потребителей для автоматического управления оффсетами
func CreateConsumerGroup(kafkaBrokers string, groupID string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // при неимении коммитнутого оффсета - читаем все что есть
	// для исключения повторных чтений необходимо знать рпс, пока что оставляем по дефолту
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup([]string{kafkaBrokers}, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}
	return consumerGroup, nil
}
