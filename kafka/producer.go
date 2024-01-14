package kafka

import (
	"fmt"
	"os"
	"time"

	kafkaGo "github.com/confluentinc/confluent-kafka-go/kafka"
	"leoho.io/go-mqtt-kafka-example/env"
	"leoho.io/go-mqtt-kafka-example/models"
)

func NewProducer(kbc env.KafkaBrokerConfiguration) {
	p, err := kafkaGo.NewProducer(&kafkaGo.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", kbc.Host, kbc.Port),
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafkaGo.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf(
						"Delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic,
						ev.TopicPartition.Partition,
						ev.TopicPartition.Offset,
					)
				}
			case kafkaGo.Error:
				fmt.Printf("Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	Produce(kbc, p)
}

func Produce(kbc env.KafkaBrokerConfiguration, p *kafkaGo.Producer) {
	message := models.GetMessage()
	err := p.Produce(&kafkaGo.Message{
		TopicPartition: kafkaGo.TopicPartition{
			Topic:     &kbc.Topic,
			Partition: kafkaGo.PartitionAny,
		},
		Value: message.Body,
	}, nil)
	if err != nil {
		if err.(kafkaGo.Error).Code() == kafkaGo.ErrQueueFull {
			fmt.Println("Producer queue full")
			time.Sleep(time.Second)
		}
		fmt.Printf("Failed to produce message: %v\n", err)
	}

	for p.Flush(10000) > 0 {
		fmt.Println("Still waiting to flush outstanding messages")
	}
	p.Close()
}
