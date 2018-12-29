package incoming

type Message struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
