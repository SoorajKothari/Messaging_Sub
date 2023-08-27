package appContext

import (
	. "Messaging_sub.go/pkg/messaging"
	"github.com/go-redis/redis/v8"
	"log"
)

type AppContext struct {
	Client  *redis.Client
	Version string
}

func GetContext() (*AppContext, error) {
	context := &AppContext{
		Version: "1",
	}

	client, err := InitializeRedis("localhost:6379")
	if err != nil {
		return nil, err
	}
	context.Client = client
	log.Println("Client Initialized Successfully")
	return context, nil
}
