package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"leoho.io/go-mqtt-kafka-example/models"
)

func PublishHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received MQTT Message: %s From Topic: %s\n", msg.Payload(), msg.Topic())
	models.CacheMessage(msg.Payload())
}

func OnConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to MQTT Broker")
}

func OnConnectLostHandler(client MQTT.Client, err error) {
	fmt.Printf("Connection to MQTT Broker lost: %v\n", err)
}
