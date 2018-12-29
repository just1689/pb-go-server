package controller

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/pb-go-server/io/ws"
	"github.com/just1689/pb-go-server/model/incoming"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

const (
	RouteTermSearch = "term-search"
)

func HandleIncoming(client *ws.Client, msg []byte) {
	var message incoming.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	if message.Message == RouteTermSearch {
		HandleTermSearch(client, msg)
	} else {
		log.Println("Unknown route: ", message.Message)
	}

}
