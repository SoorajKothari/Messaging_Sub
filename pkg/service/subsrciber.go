package service

import (
	. "context"
	. "github.com/SoorajKothari/Messaging_Sub/pkg/appContext"
	"log"
)

func Listen(context *AppContext) {
	pubsub := context.Client.Subscribe(Background(), "main")
	channel := pubsub.Channel()

	go func() {
		for msg := range channel {
			log.Println("Received message: ", msg.Payload)
		}
	}()
}
