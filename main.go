package main

import (
	"fmt"

	"github.com/ntttrang/t-log-service/hello"
	"github.com/ntttrang/t-selling-service/controller"
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/startup/config"
)

func init() {
	config.Init()
	messagehandler.Init()
}
func main() {
	fmt.Println("Welcome to t-selling-service!")
	hello.Hi()
	hello.Welcome()
	cont := controller.NewEchoHTTPManager()
	cont.RegisterRouter()
	cont.Start()
}
