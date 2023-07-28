package main

import (
	"log"

	"github.com/devfajar/rtc/internal/server"
)

func main()  {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
