package server

import (
	"github.com/SoorajKothari/Messaging_Pub/pkg/context"
	. "github.com/SoorajKothari/Messaging_Sub/pkg/service"
	"log"
)

func Start(context *context.Context) {
	log.Println("Starting Server...")
	Listen(context)
	log.Println("Server Started...")
	select {}
}
