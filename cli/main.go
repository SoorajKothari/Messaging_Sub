package main

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/context"
	. "github.com/SoorajKothari/Messaging_Sub/pkg/server"
	"log"
)

func main() {
	log.Println("Starting App")
	context, err := GetContext()
	if err != nil {
		log.Println("Fail to start app", err)
		return
	}
	Start(context)
	log.Println("App Started", context.Name)
}
