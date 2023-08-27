package main

import (
	. "Messaging_sub.go/pkg/appContext"
	. "Messaging_sub.go/pkg/server"
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
	log.Println("App Started", context.Version)
}
