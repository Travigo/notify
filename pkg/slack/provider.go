package slack

import (
	"github.com/britbus/notify/pkg/config"
	"github.com/slack-go/slack"
)

type SlackProvider struct {
	Config map[string]string
}

func (slackProvider *SlackProvider) Init(config *config.ProviderConfig) error {
	slackProvider.Config = config.Config

	return nil
}

func (slackProvider *SlackProvider) CreateProviderEventTemplate(eventConfig *config.EventConfig) (interface{}, error) {
	slackTemplate := SlackMessageTemplate{}
	slackTemplate.LoadFromEventConfig(eventConfig)

	return &slackTemplate, nil
}

func (slackProvider *SlackProvider) ProcessEvent(templateInterface interface{}, payload *map[string]interface{}) error {
	slackTemplate := templateInterface.(*SlackMessageTemplate)

	api := slack.New(slackProvider.Config["token"])
	attachment := slack.Attachment{
		Title:   slackTemplate.Attachment.Title,
		Pretext: slackTemplate.Attachment.Pretext,
		Text:    slackTemplate.Attachment.Text,
		Fields:  slackTemplate.Attachment.Fields,
	}

	_, _, err := api.PostMessage(
		slackTemplate.Channel,
		slack.MsgOptionText(slackTemplate.Text, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		return err
	}

	return nil
}
