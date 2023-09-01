package server

import (
	. "github.com/SoorajKothari/Messaging_Sub/pkg/appContext"
	. "github.com/SoorajKothari/Messaging_Sub/pkg/service"
	"log"
)

func Start(context *AppContext) {
	log.Println("Starting Server...")
	Listen(context)
	log.Println("Server Started...")
	select {}
}
