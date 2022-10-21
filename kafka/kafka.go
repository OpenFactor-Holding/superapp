package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

const (
	FcmbKafkaAddressKey   = "bootstrap.servers"
	FcmbKafkaGroupIdKey = "group.id"
	FcmbKafkaOffsetKey  = "auto.offset.reset"

	FcmbKafkaAddress  = "KAFKA_ADDRESS"
	FcmbKafkaGroupId = "KAFKA_GROUP_ID"
	FcmbKafkaOffset  = "KAFKA_OFFSET"
)
func Publish(data interface{}, topic string)  {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		FcmbKafkaAddressKey: os.Getenv(FcmbKafkaAddress),
		FcmbKafkaGroupIdKey: os.Getenv(FcmbKafkaGroupId),
		FcmbKafkaOffsetKey:  os.Getenv(FcmbKafkaOffset),
	})

	if err != nil {
		log.Println(err)
	}

	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	byteData, err := json.Marshal(data)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Value:          byteData,
	}, nil)

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)
}