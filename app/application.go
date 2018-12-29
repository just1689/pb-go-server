package main

import (
	"flag"
	"github.com/just1689/pb-go-server/controller"
	"github.com/just1689/pb-go-server/io/ws"
	"log"
)

var addr = flag.String("addr", "127.0.0.1:8000", "http service address")

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")

	//Blocking call
	ws.StartServer(addr, controller.HandleIncoming)

}
