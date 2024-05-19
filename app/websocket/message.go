package websocket

type Message struct {
	Data   []byte
	Client *Client
}

func NewMessage(data []byte, client *Client) *Message {
	return &Message{
		Data:   data,
		Client: client,
	}
}
