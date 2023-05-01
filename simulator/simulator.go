package simulator

import (
	"simulator/client"
	"simulator/device"
	"time"
)

func Simulate(c client.Client) {
	hygrometer := device.CreateDevice(c, 40, 40, 5, "base/state/humidity")
	termometer := device.CreateDevice(c, 25, 25, 1, "base/state/temperature")

	heaterTopics := [2]string{"base/state/heater", "base/relay/heater-temp"}
	condTopics := [3]string{"base/state/conditioner", "base/relay/cond-temp", "base/relay/cond-hum"}
	heater := device.CreateAppliance(c, heaterTopics[:])
	conditioner := device.CreateAppliance(c, condTopics[:])
	heater.Subscribe()
	conditioner.Subscribe()

	//publish default values for testing
	client.Publish(c, "base/state/temperature-point", "25")
	client.Publish(c, "base/state/humidity-point", "40")
	client.Publish(c, "base/state/conditioner", "true")
	client.Publish(c, "base/state/heater", "false")
	client.Publish(c, "base/state/heater-rate", "5")
	client.Publish(c, "base/state/conditioner-hum", "40")
	client.Publish(c, "base/state/conditioner-temp", "20")
	client.Publish(c, "base/state/heater-temp", "40")

	for {
		hygrometer.Generate()
		termometer.Generate()

		termometer.Publish()
		hygrometer.Publish()
		time.Sleep(time.Second * 15)
	}
}
