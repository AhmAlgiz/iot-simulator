package device

import (
	"math/rand"
	"simulator/client"
	"strconv"
)

// measuring device
type Device struct {
	client client.Client
	base   int
	point  int
	rate   int
	topic  string
}

func CreateDevice(client client.Client, base, point, rate int, topic string) *Device {
	return &Device{
		client: client,
		base:   base,
		point:  point,
		rate:   rate,
		topic:  topic,
	}
}

func (d *Device) Publish() {
	client.Publish(d.client, d.topic, strconv.Itoa(d.base))
}

func (d *Device) Generate() {
	d.base = d.generateBaseVal() + d.generateCondVal()
}

func (d *Device) generateBaseVal() int {
	return d.base + rand.Intn(3*d.rate) - d.rate
}

func (d *Device) generateCondVal() int {
	var out int
	delta := d.base - d.point
	if delta > 0 {
		out = -1 * d.rate
	} else if delta == 0 {
		out = 0
	} else {
		out = d.rate
	}
	return out
}
