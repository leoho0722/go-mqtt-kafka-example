package env

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadEnvConfig() *Configs {
	envFile, err := os.Open("config.json")
	defer func(envConfig *os.File) {
		err := envConfig.Close()
		if err != nil {
			panic(err)
		}
	}(envFile)

	if err != nil {
		panic(err)
	}
	fmt.Println("config.json open successfully")

	envConfigs, _ := io.ReadAll(envFile)
	var configs Configs
	err = json.Unmarshal(envConfigs, &configs)
	if err != nil {
		panic(err)
	}
	fmt.Println("config.json unmarshal successfully")

	return &configs
}

type Configs struct {
	MQTT  []MQTTBrokerConfiguration  `json:"mqtt"`
	Kafka []KafkaBrokerConfiguration `json:"kafka"`
}

type MQTTBrokerConfiguration struct {

	// Host is the host of the MQTT broker
	Host string `json:"host"`

	// Port is the port of the MQTT broker
	Port int `json:"port"`

	// Topic is the topic to subscribe to
	Topic string `json:"topic"`

	// QosLevel is the quality of MQTT Broker level
	QosLevel byte `json:"qosLevel"`
}

func (cs *Configs) GetMQTTBrokerConfigurations() []MQTTBrokerConfiguration {
	var brokerConfigurations []MQTTBrokerConfiguration
	for _, broker := range cs.MQTT {
		bc := MQTTBrokerConfiguration{
			Host:     broker.Host,
			Port:     broker.Port,
			Topic:    broker.Topic,
			QosLevel: broker.QosLevel,
		}
		brokerConfigurations = append(brokerConfigurations, bc)
	}
	return brokerConfigurations
}

type KafkaBrokerConfiguration struct {

	// Host is the host of the Kafka broker
	Host string `json:"host"`

	// Port is the port of the Kafka broker
	Port int `json:"port"`

	// Topic is the topic to subscribe to
	Topic string `json:"topic"`
}

func (cs *Configs) GetKafkaBrokerConfigurations() []KafkaBrokerConfiguration {
	var brokerConfigurations []KafkaBrokerConfiguration
	for _, broker := range cs.Kafka {
		kbc := KafkaBrokerConfiguration{
			Host:  broker.Host,
			Port:  broker.Port,
			Topic: broker.Topic,
		}
		brokerConfigurations = append(brokerConfigurations, kbc)
	}
	return brokerConfigurations
}
