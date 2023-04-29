package device

import (
	"simulator/client"
)

// executive device
type Appliance struct {
	client client.Client
	topic  []string
}

func CreateAppliance(client client.Client, topic []string) *Appliance {
	return &Appliance{
		client: client,
		topic:  topic,
	}
}

func (a *Appliance) Subscribe() {
	for _, t := range a.topic {
		client.Subscribe(a.client, t)
	}
}
