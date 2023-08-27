package server

import (
	. "Messaging_sub.go/pkg/appContext"
	. "Messaging_sub.go/pkg/service"
	"log"
)

func Start(context *AppContext) {
	log.Println("Starting Server...")
	Listen(context)
	log.Println("Server Started...")
	select {}
}
