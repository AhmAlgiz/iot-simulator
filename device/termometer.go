package device

import (
	"fmt"
	"math/rand"
	"simulator/client"
	"strconv"
)

type Termometer Meter

func CreateTermometer(client client.Client, value int, topic string) *Termometer {
	return &Termometer{
		client: client,
		value:  value,
		topic:  topic,
	}
}

func (t *Termometer) Generate() {
	t.value = t.generatevalueVal() + t.generateCondVal()
}

func (t *Termometer) generatevalueVal() int {
	return t.value + rand.Intn(3) - 1
}

func (t *Termometer) generateCondVal() int {
	deviceOptions, err := client.ReadJson("options.json")
	if err != nil {
		fmt.Printf("\nOptions read error: %e", err)
		return 0
	}
	var out int
	if deviceOptions.HeaterStatus {
		out = (deviceOptions.HeaterTemp - t.value) / 15
	} else {
		out = (deviceOptions.CondTemp - t.value) / 2
	}
	return out
}

func (t *Termometer) Publish() {
	client.Publish(t.client, t.topic, strconv.Itoa(t.value))
}
