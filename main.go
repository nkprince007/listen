package main

import (
	"os"

	_ "github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"gitlab.com/nkprince007/listen/handler"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// grab port from environment
	port := getEnv("PORT", "8000")

	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.listen"),
		web.Version("latest"),
		web.Address(":"+port),
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
