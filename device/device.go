package device

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Device interface {
	GenerateValue(mqtt.Client) int
	GetValue(mqtt.Client) int
	Publish(mqtt.Client, string)
}
