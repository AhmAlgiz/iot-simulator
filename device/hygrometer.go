package device

import (
	"fmt"
	"math/rand"
	"simulator/client"
	"strconv"
)

type Hygrometer Meter

func CreateHygrometer(client client.Client, value int, topic string) *Hygrometer {
	return &Hygrometer{
		client: client,
		value:  value,
		topic:  topic,
	}
}

func (h *Hygrometer) Generate() {
	h.value = h.generatevalueVal() + h.generateCondVal()
}

func (h *Hygrometer) generatevalueVal() int {
	return h.value + rand.Intn(9) - 4
}

func (h *Hygrometer) generateCondVal() int {
	deviceOptions, err := client.ReadJson("options.json")
	if err != nil {
		fmt.Printf("\nOptions read error: %e", err)
		return 0
	}
	out := (deviceOptions.CondHum - h.value) / 5
	return out
}

func (h *Hygrometer) Publish() {
	client.Publish(h.client, h.topic, strconv.Itoa(h.value))
}
