package producer

import (
	"fmt"
	gitKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func NewKafkaProducer() *gitKafka.Producer {
	configMap := &gitKafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	producer, err := gitKafka.NewProducer(configMap)
	if err != nil {
		log.Fatal("Error connecting process")
	}
	return producer
}

func Publish(msg string, topic string, producer *gitKafka.Producer, deliveryChan chan gitKafka.Event) error {
	message := &gitKafka.Message{
		TopicPartition: gitKafka.TopicPartition{Topic: &topic, Partition: gitKafka.PartitionAny},
		Value: []byte(msg),
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		log.Fatal("Error")
	}
	return nil
}

func DeliveryReport(deliveryChan chan gitKafka.Event)  {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *gitKafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}