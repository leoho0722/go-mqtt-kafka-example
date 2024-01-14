package mqtt

import (
	"fmt"

	mqttGo "github.com/eclipse/paho.mqtt.golang"
	"leoho.io/go-mqtt-kafka-example/env"
)

func NewClient(bc env.MQTTBrokerConfiguration) {
	options := mqttGo.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%d", bc.Host, bc.Port))
	options.SetDefaultPublishHandler(PublishHandler)
	options.OnConnect = OnConnectHandler
	options.OnConnectionLost = OnConnectLostHandler

	client := mqttGo.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	Subscribe(client, bc)
}
