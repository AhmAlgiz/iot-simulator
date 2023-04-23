package main

import (
	"log"
	"os"
	"simulator/client"

	"github.com/urfave/cli/v2"
)

func main() {
	var broker, id, username, pass string
	var port int

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "broker",
				Value:       "dev.rightech.io",
				Usage:       "MQTT broker address",
				Destination: &broker,
				DefaultText: "dev.rightech.io",
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       1883,
				Usage:       "MQTT broker port",
				Destination: &port,
				DefaultText: "1883",
			},
			&cli.StringFlag{
				Name:        "id",
				Value:       "",
				Usage:       "Client's id",
				Required:    true,
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "u",
				Value:       "",
				Usage:       "User's name",
				Required:    true,
				Destination: &username,
			},
			&cli.StringFlag{
				Name:        "p",
				Value:       "",
				Usage:       "User's password",
				Required:    true,
				Destination: &pass,
			},
		},
		Action: func(cCtx *cli.Context) error {
			opts := client.CreateClientOptions(broker, port, id, username, pass)
			client := client.CreateClient(opts)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
