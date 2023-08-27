package service

import (
	. "Messaging_sub.go/pkg/appContext"
	. "context"
	"log"
)

func Listen(context *AppContext) {
	pubsub := context.Client.Subscribe(Background(), "main")
	channel := pubsub.Channel()

	go func() {
		for msg := range channel {
			log.Println("Received message: %s\n", msg.Payload)
		}
	}()
}
