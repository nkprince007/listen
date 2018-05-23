package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"gitlab.com/nkprince007/listen/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.listen"),
		web.Version("latest"),
	)

	// register call handler
	service.HandleFunc("/", handler.Echo)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
