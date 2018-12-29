package controller

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/pb-go-server/io/ws"
	"github.com/just1689/pb-go-server/model/incoming"
)
import _ "github.com/go-sql-driver/mysql"

func HandleIncoming(client *ws.Client, msg []byte) {
	var message incoming.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		fmt.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	HandleTermSearch(client, message)

}
