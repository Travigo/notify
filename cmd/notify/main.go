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

					// runtime.ProcessEvent(config.Global.Events[0], map[string]interface{}{
					// 	"stop_groups": map[string]interface{}{
					// 		"inserts": float64(4),
					// 		"updates": float64(3),
					// 	},
					// 	"stops": map[string]interface{}{
					// 		"inserts": float64(11),
					// 		"updates": float64(18),
					// 	},
					// })

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	/*
		api := slack.New("xoxb-2896435494805-2892761762070-VUF7XabghpBauGMYB5hLFlrM")
		attachment := slack.Attachment{
			Title:   "wowowow",
			Pretext: "some pretext",
			Text:    "some text",
			Fields: []slack.AttachmentField{
				{
					Title: "What the hell",
					Value: "no",
				},
			},
		}

		channelID, timestamp, err := api.PostMessage(
			"C02T53DLZ4G",
			slack.MsgOptionText("Some text", false),
			slack.MsgOptionAttachments(attachment),
			slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
		)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	*/
}
