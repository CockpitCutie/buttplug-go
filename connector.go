package buttplug

import (
	"net/http"

	"github.com/CockpitCutie/buttplug-go/message"
	"github.com/gorilla/websocket"
)

type Connector interface {
	Connect(msgRecv map[uint32]chan message.Message) error
	Connected() bool
	Disconnect() error
	Send(msg message.Message)
}

type WebsocketConnector struct {
	url      string
	isOpen bool
	conn     *websocket.Conn
	msgRecv map[uint32]chan message.Message
}

func NewWsConnector(url string) *WebsocketConnector {
	return &WebsocketConnector{
		url:      url,
		isOpen: false,
		conn:     nil,
		msgRecv: nil,
	}
}

func (w *WebsocketConnector) Connect(msgRecv map[uint32]chan message.Message) error {
	w.msgRecv = msgRecv
	conn, _, err := websocket.DefaultDialer.Dial(w.url, http.Header{})
	if err != nil {
		return err
	}
	w.conn = conn
	w.isOpen = true
	w.conn.SetCloseHandler(func(code int, text string) error {
		w.isOpen = false
		return w.conn.WriteMessage(websocket.CloseMessage, []byte{})
	})
	return nil
}

func (w *WebsocketConnector) Connected() bool {
	return w.isOpen
}

func (w *WebsocketConnector) Disconnect() error {
	return w.conn.Close()
}
