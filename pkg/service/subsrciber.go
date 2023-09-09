package service

import (
	. "context"
	"encoding/json"
	"fmt"
	"github.com/SoorajKothari/Messaging_Pub/pkg/context"
	. "github.com/SoorajKothari/Messaging_Pub/pkg/model"
	"log"
)

var broadcast = make(chan Message)

func Listen(context *context.Context) {
	pubsub := context.Client.Subscribe(Background(), "main")
	channel := pubsub.Channel()

	go func() {
		for msg := range channel {
			var message Message
			err := json.Unmarshal([]byte(msg.Payload), &message)
			if err != nil {
				log.Println("Invalid message")
			}
			log.Println("Message received on Channel: ", msg.Channel)
			broadcast <- message
		}
	}()

	go func() {
		for {
			received := <-broadcast
			received.Content = fmt.Sprintf("Hello, How Can I Help You? %s", received.Content)
			marshal, err := json.Marshal(received)
			if err != nil {
				log.Println("Error Parsing", err)
				return
			}
			result, err := context.Client.Publish(Background(), "reply", marshal).Result()
			if err != nil {
				log.Println("Error sending reply", err)
				return
			}
			log.Println("Published reply", result)
		}
	}()
}
