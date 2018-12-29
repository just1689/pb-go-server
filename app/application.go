package main

import (
	"flag"
	"github.com/just1689/pb-go-server/api"
	"github.com/just1689/pb-go-server/io"
	"log"
)

var addr = flag.String("addr", "127.0.0.1:8000", "http service address")

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")

	//Blocking call
	io.StartServer(addr, api.HandleIncoming)

}
