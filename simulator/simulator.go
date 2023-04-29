package simulator

import (
	"simulator/client"
	"simulator/device"
	"time"
)

func Simulate(client client.Client) {
	hygrometer := device.CreateDevice(client, 40, 40, "base/state/humidity")
	termometer := device.CreateDevice(client, 25, 25, "base/state/temperature")

	heaterTopics := [2]string{"base/state/heater", "base/state/heater-temp"}
	condTopics := [3]string{"base/state/conditioner", "base/state/conditioner-temp", "base/state/conditioner-hym"}
	heater := device.CreateAppliance(client, heaterTopics[:])
	conditioner := device.CreateAppliance(client, condTopics[:])
	heater.Subscribe()
	conditioner.Subscribe()

	for {
		hygrometer.Generate()
		termometer.Generate()

		go hygrometer.Publish()
		go termometer.Publish()
		time.Sleep(time.Second * 30)
	}
}
