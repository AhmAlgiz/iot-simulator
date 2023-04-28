package simulator

import (
	"simulator/client"
	"simulator/device"
	"time"
)

func Simulate(client client.Client) {
	hygrometer := device.CreateDevice(client, 40, 40, "base/state/humidity")
	termometer := device.CreateDevice(client, 25, 25, "base/state/temperature")
	for {
		hygrometer.Generate()
		termometer.Generate()

		go hygrometer.Publish()
		go termometer.Publish()
		time.Sleep(time.Second * 30)
	}
}
