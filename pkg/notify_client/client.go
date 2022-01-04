package notify_client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var Global *NotificationClient

type NotificationClient struct {
	Endpoint  string
	WaitGroup *sync.WaitGroup
}

func Setup(endpoint string) {
	Global = &NotificationClient{
		Endpoint:  endpoint,
		WaitGroup: &sync.WaitGroup{},
	}
	log.Info().Msgf("Notification client setup pointing to %s", endpoint)
}

func SendEvent(eventName string, payload interface{}) {
	if Global == nil {
		return
	}

	Global.WaitGroup.Add(1)

	go func() {
		jsonValue, err := json.Marshal(payload)

		if err != nil {
			log.Error().Err(err).Msg("Failed to marshall notification")

			Global.WaitGroup.Done()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		request, err := http.NewRequestWithContext(ctx, "POST", Global.Endpoint+"/event/"+eventName, bytes.NewBuffer(jsonValue))
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Error().Err(err).Msg("Failed to create notification request")

			Global.WaitGroup.Done()
			return
		}

		client := http.DefaultClient
		response, err := client.Do(request)

		if err != nil {
			log.Error().Err(err).Msg("Failed to send notification")

			Global.WaitGroup.Done()
			return
		}
		defer response.Body.Close()

		Global.WaitGroup.Done()
	}()
}

func Await() {
	if Global == nil {
		return
	}

	Global.WaitGroup.Wait()
}
