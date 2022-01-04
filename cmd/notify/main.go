package main

import (
	"os"
	"time"

	"github.com/britbus/notify/pkg/config"
	"github.com/britbus/notify/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	app := &cli.App{
		Name: "notify",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "run notify api server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "listen",
						Value: ":8081",
						Usage: "listen target for the notify web server",
					},
					&cli.StringFlag{
						Name:  "config",
						Usage: "config file for notify",
					},
				},
				Action: func(c *cli.Context) error {
					listenAddress := c.String("listen")
					configFilePath := c.String("config")

					if configFilePath == "" {
						log.Fatal().Msg("Config file location paramater must be set")
						return nil
					}

					config.LoadFromFile(configFilePath)

					server.SetupServer(listenAddress)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
