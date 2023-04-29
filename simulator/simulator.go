package simulator

import (
	"simulator/client"
	"simulator/device"
	"time"
)

func Simulate(c client.Client) {
	hygrometer := device.CreateDevice(c, 40, 40, "base/state/humidity")
	termometer := device.CreateDevice(c, 25, 25, "base/state/temperature")

	heaterTopics := [2]string{"base/state/heater", "base/state/heater-temp"}
	condTopics := [3]string{"base/state/conditioner", "base/state/conditioner-temp", "base/state/conditioner-hym"}
	heater := device.CreateAppliance(c, heaterTopics[:])
	conditioner := device.CreateAppliance(c, condTopics[:])
	heater.Subscribe()
	conditioner.Subscribe()

	//publish default values for testing
	client.Publish(c, "base/state/temperature-point", "25")
	client.Publish(c, "base/state/humidity-point", "40")
	client.Publish(c, "base/state/conditioner", "1")
	client.Publish(c, "base/state/heater", "0")

	for {
		hygrometer.Generate()
		termometer.Generate()

		go hygrometer.Publish()
		go termometer.Publish()
		time.Sleep(time.Second * 30)
	}
}
