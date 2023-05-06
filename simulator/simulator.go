package simulator

import (
	"encoding/json"
	"fmt"
	"os"
	"simulator/client"
	"simulator/device"
	"time"
)

func Simulate(c client.Client) {
	//json options initialization
	opts := client.JsonFile{CondTemp: 20, CondHum: 40, CondStatus: true,
		HeaterTemp: 40, HeaterStatus: false}
	file, err := json.Marshal(opts)
	if err != nil {
		fmt.Printf("\nOptions marshal error: %e", err)
		return
	}
	err = os.WriteFile("options.json", file, 0666)
	if err != nil {
		fmt.Printf("\nOptions write error: %e", err)
		return
	}

	//devices initialization
	hygrometer := device.CreateHygrometer(c, 40, "base/state/humidity")
	termometer := device.CreateTermometer(c, 25, "base/state/temperature")

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
	client.Publish(c, "base/state/cond-hum", "40")
	client.Publish(c, "base/state/cond-temp", "20")
	client.Publish(c, "base/state/heater-temp", "40")

	//generation loop
	for {
		hygrometer.Generate()
		termometer.Generate()

		termometer.Publish()
		hygrometer.Publish()
		time.Sleep(time.Second * 15)
	}
}
