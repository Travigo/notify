package notify_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type NotificationClient struct {
	Endpoint  string
	WaitGroup *sync.WaitGroup
}

func (c *NotificationClient) Setup(endpoint string) {
	c.Endpoint = endpoint
	c.WaitGroup = &sync.WaitGroup{}
}

func (c *NotificationClient) SendEvent(eventName string, payload interface{}) {
	if c.Endpoint == "" {
		return
	}

	c.WaitGroup.Add(1)

	go func() {
		jsonValue, err := json.Marshal(payload)

		if err != nil {
			log.Error().Err(err).Msg("Failed to marshall notification")

			c.WaitGroup.Done()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		request, err := http.NewRequestWithContext(ctx, "POST", c.Endpoint+"/event/"+eventName, bytes.NewBuffer(jsonValue))
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Error().Err(err).Msg("Failed to create notification request")

			c.WaitGroup.Done()
			return
		}

		client := http.DefaultClient
		response, err := client.Do(request)

		if err != nil {
			log.Error().Err(err).Msg("Failed to send notification")

			c.WaitGroup.Done()
			return
		}
		defer response.Body.Close()

		io.Copy(os.Stdout, response.Body)
		c.WaitGroup.Done()
	}()
}

func (c *NotificationClient) Await() {
	c.WaitGroup.Wait()
}
