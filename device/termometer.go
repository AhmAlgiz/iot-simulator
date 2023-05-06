package device

import (
	"fmt"
	"math/rand"
	"simulator/client"
	"strconv"
)

type Termometer Meter

func CreateTermometer(client client.Client, base int, topic string) *Termometer {
	return &Termometer{
		client: client,
		base:   base,
		topic:  topic,
	}
}

func (t *Termometer) Generate() {
	t.base = t.generateBaseVal() + t.generateCondVal()
}

func (t *Termometer) generateBaseVal() int {
	return t.base + rand.Intn(3) - 1
}

func (t *Termometer) generateCondVal() int {
	deviceOptions, err := client.ReadJson("options.json")
	if err != nil {
		fmt.Printf("\nOptions read error: %e", err)
		return 0
	}
	var out int
	if deviceOptions.HeaterStatus {
		out = (deviceOptions.HeaterTemp - t.base) / 15
	} else {
		out = (deviceOptions.CondTemp - t.base) / 2
	}
	return out
}

func (t *Termometer) Publish() {
	client.Publish(t.client, t.topic, strconv.Itoa(t.base))
}
