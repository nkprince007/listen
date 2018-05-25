//+build !test

package main

import (
	"fmt"

	"gitlab.com/gitmate-micro/listen/provider"

	"github.com/micro/go-log"
	"github.com/micro/go-web"
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
		web.Name("gitmate.micro.web.listen"),
		web.Version("latest"),
	)

	github := &provider.GitHub{}
	gitlab := &provider.GitLab{}

	// register call handlers
	service.HandleFunc("/github", RejectOtherMethods("POST", github))
	service.HandleFunc("/gitlab", RejectOtherMethods("POST", gitlab))

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
