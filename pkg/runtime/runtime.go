package runtime

import (
	"errors"

	"github.com/britbus/notify/pkg/config"
	"github.com/britbus/notify/pkg/slack"
	"github.com/britbus/notify/pkg/util"
	"github.com/rs/zerolog/log"
)

func ProcessEvent(event *config.EventConfig, userData map[string]interface{}) error {
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

	// Convert the array of env variables into an addressable map
	environmentVariables := util.GetEnvironmentVariables()

	// Template data contains both env variables and payload user data
	data := map[string]interface{}{
		"env":  environmentVariables,
		"data": userData,
	}

	// Render the provider config when initing it
	// TODO: Init should really be a one time per provider step instead of every run
	RenderTemplate(providerConfig, &data)

	provider.Init(providerConfig)

	// Render the event template and then send it to the provider to process
	providerEventTemplate, _ := provider.CreateProviderEventTemplate(event)
	RenderTemplate(providerEventTemplate, &data)
	err := provider.ProcessEvent(providerEventTemplate, &data)

	if err != nil {
		return err
	}

	log.Info().Interface("message", providerEventTemplate).Msgf("Event %s triggered with provider %s", event.Name, event.Provider)

	return nil
}
