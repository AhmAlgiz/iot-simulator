package device

import (
	"fmt"
	"math/rand"
	"simulator/client"
	"strconv"
)

type Hygrometer Meter

func CreateHygrometer(client client.Client, base int, topic string) *Hygrometer {
	return &Hygrometer{
		client: client,
		base:   base,
		topic:  topic,
	}
}

func (h *Hygrometer) Generate() {
	h.base = h.generateBaseVal() + h.generateCondVal()
}

func (h *Hygrometer) generateBaseVal() int {
	return h.base + rand.Intn(9) - 4
}

func (h *Hygrometer) generateCondVal() int {
	deviceOptions, err := client.ReadJson("options.json")
	if err != nil {
		fmt.Printf("\nOptions read error: %e", err)
		return 0
	}
	out := (deviceOptions.CondHum - h.base) / 5
	return out
}

func (h *Hygrometer) Publish() {
	client.Publish(h.client, h.topic, strconv.Itoa(h.base))
}
