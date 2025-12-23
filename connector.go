package buttplug

import (
	"fmt"
	"net/http"

	"github.com/CockpitCutie/buttplug-go/message"
	"github.com/gorilla/websocket"
)

type Connector interface {
	Connect(msgRecv map[uint32]chan message.Message) error
	Connected() bool
	Disconnect() error
	Send(msg message.Message) error
	SendRecv(msg message.Message) (message.Message, error)
}

type WebsocketConnector struct {
	url       string
	isOpen    bool
	conn      *websocket.Conn
	msgRecv   map[uint32]chan message.Message
	idCounter uint32
}

func NewWsConnector(url string) *WebsocketConnector {
	return &WebsocketConnector{
		url:       url,
		isOpen:    false,
		conn:      nil,
		msgRecv:   nil,
		idCounter: 1,
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
	go w.listenLoop()
	return nil
}

func (w *WebsocketConnector) listenLoop() {
	for {
		kind, buf, err := w.conn.ReadMessage()
		if err != nil {
			println("failed to read")
			break
		}
		fmt.Printf("recv: `%s`\n", buf)
		if kind != websocket.TextMessage {
			continue
		}
		deserialized, err := message.Deserialize(buf)
		if err != nil {
			println(err.Error())
			continue
		}
		println("deserialized")
		if _, ok := w.msgRecv[deserialized.ID()]; !ok {
			w.msgRecv[deserialized.ID()] = make(chan message.Message)
		}
		w.msgRecv[deserialized.ID()] <- deserialized
	}
}

func (w *WebsocketConnector) Connected() bool {
	return w.isOpen
}

func (w *WebsocketConnector) Disconnect() error {
	if w.conn == nil {
		return nil
	}
	return w.conn.Close()
}

func (w *WebsocketConnector) Send(msg message.Message) error {
	w.msgRecv[msg.ID()] = make(chan message.Message)
	serialized, err := message.Serialize(msg)
	fmt.Printf("sending: `%s`\n", serialized)
	if err != nil {
		return err
	}
	return w.conn.WriteMessage(websocket.TextMessage, []byte(serialized))
}

func (w *WebsocketConnector) SendRecv(m message.Message) (message.Message, error) {
	id := w.idCounter
	w.idCounter++
	m.SetID(id)
	err := w.Send(m)
	if err != nil {
		return nil, err
	}
	recv := <-w.msgRecv[id]
	if err, ok := recv.(*message.Error); ok {
		return err, err.Error()
	}
	delete(w.msgRecv, id)
	return recv, nil
}
