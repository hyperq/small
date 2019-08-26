package main

import (
	"log"
	"small/router"

	"github.com/micro/go-micro/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()
	r := router.GetRouter()
	service.Handle("/", r)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
