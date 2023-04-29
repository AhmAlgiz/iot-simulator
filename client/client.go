package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//client's options and base functionality

type Client mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("\nReceived message: %s from topic: %s", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("\nConnected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("\nConnect lost: %v", err)
}

func CreateClientOptions(broker string, port int, id, username, pass string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(fmt.Sprint(id))
	opts.SetUsername(fmt.Sprint(username))
	opts.SetPassword(fmt.Sprint(pass))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}

func Subscribe(client Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("\nSubscribed to topic %s", topic)
}

func Publish(client Client, topic, data string) {
	token := client.Publish(topic, 0, false, data)
	token.Wait()
	fmt.Printf("\nPublished to topic: %s ; data: %s", topic, data)
}

func CreateClient(opts *mqtt.ClientOptions) Client {
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
