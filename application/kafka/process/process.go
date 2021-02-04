package process

import (
	"log"

	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	ConsumerTopics 		[]string
	BootstrapServers 	string
	ConsumerGroup 		string
}

func NewKafkaProcessor() *KafkaProcessor {
	return &KafkaProcessor{}
}

func (kafkaProcessor *KafkaProcessor) Consumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaProcessor.BootstrapServers,
		"group.id": kafkaProcessor.ConsumerGroup,
		"auto.offset.reset":"earliest",
	})

	if err != nil {
		log.Fatal("Error consuming")
	}

	consumer.SubscribeTopics(kafkaProcessor.ConsumerTopics, nil)

	log.Printf("Kafka producer has been started on host %s", kafkaProcessor.BootstrapServers)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}

	consumer.Close()
}