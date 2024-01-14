package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"leoho.io/go-mqtt-kafka-example/env"
)

func Subscribe(client MQTT.Client, bc env.MQTTBrokerConfiguration) {
	token := client.Subscribe(bc.Topic, bc.QosLevel, nil)
	token.Wait()
	fmt.Printf("Subscribed to MQTT Topic: %s\n", bc.Topic)
}
