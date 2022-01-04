package slack

import (
	"github.com/britbus/notify/pkg/config"
	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)

type SlackMessageTemplate struct {
	Channel    string
	Text       string
	Attachment struct {
		Title   string
		Pretext string
		Text    string
		Fields  []slack.AttachmentField
	}
}

func (s *SlackMessageTemplate) LoadFromEventConfig(eventConfig *config.EventConfig) {
	// This is absoloutely awful and should be redone as something nice
	for _, item := range eventConfig.Template {
		if item.Key == "channel" {
			s.Channel = item.Value.(string)
		} else if item.Key == "text" {
			s.Text = item.Value.(string)
		} else if item.Key == "attachment" {
			for _, item := range item.Value.(yaml.MapSlice) {
				if item.Key == "title" {
					s.Attachment.Title = item.Value.(string)
				} else if item.Key == "pretext" {
					s.Attachment.Pretext = item.Value.(string)
				} else if item.Key == "text" {
					s.Attachment.Text = item.Value.(string)
				} else if item.Key == "fields" {
					s.Attachment.Fields = []slack.AttachmentField{}

					for _, fields := range item.Value.([]interface{}) {
						field := slack.AttachmentField{}
						for _, item := range fields.(yaml.MapSlice) {
							if item.Key == "title" {
								field.Title = item.Value.(string)
							} else if item.Key == "value" {
								field.Value = item.Value.(string)
							}
						}
						s.Attachment.Fields = append(s.Attachment.Fields, field)
					}
				}
			}
		}
	}
}
