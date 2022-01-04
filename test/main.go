package main

import (
	"github.com/britbus/notify/pkg/notify_client"
)

type Updates struct {
	Stops       UpdatesStatus
	Stop_Groups UpdatesStatus
}
type UpdatesStatus struct {
	Inserts string
	Updates string
}

func main() {
	client := notify_client.NotificationClient{}
	client.Setup("http://localhostfake:8081")

	client.SendEvent("britbus/traveline/import", Updates{
		Stops: UpdatesStatus{
			Inserts: "123",
			Updates: "456",
		},
		Stop_Groups: UpdatesStatus{
			Inserts: "321",
			Updates: "654",
		},
	})

	client.Await()
}
