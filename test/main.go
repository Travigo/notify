package main

import "github.com/britbus/notify/pkg/notify_client"

type Updates struct {
	Operators       UpdatesStatus
	Operator_Groups UpdatesStatus
}
type UpdatesStatus struct {
	Inserts string
	Updates string
}

func main() {
	// notify_client.Setup("http://localhost:8081")
	notify_client.Setup()

	notify_client.SendEvent("britbus/traveline/import", Updates{
		Operators: UpdatesStatus{
			Inserts: "123",
			Updates: "456",
		},
		Operator_Groups: UpdatesStatus{
			Inserts: "321",
			Updates: "654",
		},
	})

	notify_client.Await()
}
