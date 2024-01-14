package main

import (
	"leoho.io/go-mqtt-kafka-example/env"
	"leoho.io/go-mqtt-kafka-example/kafka"
	"leoho.io/go-mqtt-kafka-example/mqtt"
)

func main() {
	configs := env.ReadEnvConfig()

	mqttBcs := configs.GetMQTTBrokerConfigurations()
	mqtt.NewClient(mqttBcs[0])

	kafkaBcs := configs.GetKafkaBrokerConfigurations()
	kafka.NewProducer(kafkaBcs[0])
}
