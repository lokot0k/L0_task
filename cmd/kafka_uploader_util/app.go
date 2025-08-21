package main

import (
	"L0_task/pkg/utils"
	"fmt"
)

const (
	kafkaBrokers = "0.0.0.0:9092"
	kafkaTopic   = "test"
	uploadCount  = 10
)

func main() {
	//producer, err := CreateProducer
	//if err != nil {
	//	panic(fmt.Errorf("failed to create producer: %v", err))
	//}
	//
	//defer func(producer sarama.SyncProducer) {
	//	err := producer.Close()
	//	if err != nil {
	//		panic("Failed to close producer")
	//	}
	//}(producer)
	//
	//err = produceJSONs(producer, generateTestData(15))
	//if err != nil {
	//	log.Fatalf("Failed to produce messages: %v", err)
	//}
	//
	//fmt.Println("Done")
	fmt.Printf("%+v", utils.GenerateMockOrders(1)[0])
}
