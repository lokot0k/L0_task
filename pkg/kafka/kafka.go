package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
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
