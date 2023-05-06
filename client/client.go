package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//client's options and base functionality

type Client mqtt.Client

type JsonFile struct {
	CondTemp     int
	CondHum      int
	HeaterTemp   int
	CondStatus   bool
	HeaterStatus bool
}

func ReadJson(name string) (*JsonFile, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	out := &JsonFile{}
	err = json.Unmarshal(file, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func WriteJson(name string, data *JsonFile) error {
	file, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(name, file, 0666)
	if err != nil {
		return err
	}
	return nil
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("\nReceived message: %s from topic: %s", msg.Payload(), msg.Topic())
	deviceOptions, err := ReadJson("options.json")
	if err != nil {
		fmt.Printf("\nOptions read error: %e", err)
		return
	}

	switch msg.Topic() {
	case "base/state/conditioner":
		deviceOptions.CondStatus, _ = strconv.ParseBool(string(msg.Payload()))
	case "base/state/heater":
		deviceOptions.HeaterStatus, _ = strconv.ParseBool(string(msg.Payload()))
	case "base/relay/cond-temp":
		deviceOptions.CondTemp, _ = strconv.Atoi(string(msg.Payload()))
	case "base/relay/cond-hum":
		deviceOptions.CondHum, _ = strconv.Atoi(string(msg.Payload()))
	case "base/relay/heater-temp":
		deviceOptions.HeaterTemp, _ = strconv.Atoi(string(msg.Payload()))

	}
	err = WriteJson("options.json", deviceOptions)
	if err != nil {
		fmt.Printf("\nOptions marshal error: %e", err)
		return
	}
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
