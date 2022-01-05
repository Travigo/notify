package server

import (
	"github.com/britbus/notify/pkg/config"
	"github.com/britbus/notify/pkg/runtime"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func SetupServer(listen string) {
	webApp := fiber.New()

	webApp.Get("/event/+", runEvent)
	webApp.Post("/event/+", runEvent)

	webApp.Listen(listen)
}

func runEvent(c *fiber.Ctx) error {
	eventID := c.Params("+")

	for _, event := range config.Global.Events {
		if event.Name == eventID {
			payload := map[string]interface{}{}
			if err := c.BodyParser(&payload); err != nil {
				return err
			}

			err := runtime.ProcessEvent(event, payload)

			if err != nil {
				log.Error().Err(err).Msg("Failed to process event")

				c.SendStatus(503)
				return c.JSON(fiber.Map{
					"error":   "Failed to process event",
					"details": err,
				})
			} else {
				return c.JSON(fiber.Map{
					"success": true,
				})
			}
		}
	}

	c.SendStatus(404)
	return c.JSON(fiber.Map{
		"error": "Could not find Event",
	})

}
