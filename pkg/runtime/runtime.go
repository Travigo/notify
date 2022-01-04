package runtime

import (
	"errors"

	"github.com/britbus/notify/pkg/config"
	"github.com/britbus/notify/pkg/slack"
	"github.com/rs/zerolog/log"
)

func ProcessEvent(event *config.EventConfig, payload map[string]interface{}) error {
	var providerConfig *config.ProviderConfig
	for _, providerCheck := range config.Global.Providers {
		if providerCheck.Name == event.Provider {
			providerConfig = providerCheck
		}
	}

	if providerConfig == nil {
		return errors.New("Could not find provider for event")
	}

	var provider config.Provider

	switch providerConfig.Type {
	case "slack":
		provider = &slack.SlackProvider{}
	default:
		return errors.New("Provider could not be found")
	}

	provider.Init(providerConfig)

	providerTemplate, _ := provider.CreateProviderTemplate(event)
	RenderTemplate(providerTemplate, &payload)
	err := provider.ProcessEvent(providerTemplate, &payload)

	if err != nil {
		return err
	}

	log.Info().Interface("message", providerTemplate).Msgf("Event %s triggered with provider %s", event.Name, event.Provider)

	return nil
}
