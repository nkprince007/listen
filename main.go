package main

import (
	"fmt"

	_ "github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"gitlab.com/nkprince007/listen/handler"
)

func main() {
	fmt.Print(`
        ___      __
       / (_)____/ /____  ____
      / / / ___/ __/ _ \/ __ \
     / / (__  ) /_/  __/ / / /
    /_/_/____/\__/\___/_/ /_/  server catches webhooks :P


`)
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.listen"),
		web.Version("latest"),
	)

	// register call handler
	service.HandleFunc("/", handler.Capture)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
