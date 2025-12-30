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
	go w.readFromServer()
	return nil
}

func (w *WebsocketConnector) readFromServer() {
	for {
		if !w.isOpen {
			break
		}
		kind, buf, err := w.conn.ReadMessage()
		if err != nil {
			w.isOpen = false
		}
		if kind != websocket.TextMessage {
			continue
		}
		msg, err := message.Deserialize(buf)
		if err != nil {
			continue
		}
		go w.sortMessage(msg)
	}
}

func (w *WebsocketConnector) sortMessage(msg message.Message) {
	id := msg.ID()
	if _, ok := w.msgRecv[id]; !ok {
		w.msgRecv[id] = make(chan message.Message)
	}
	w.msgRecv[id] <- msg
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
	if w.conn == nil {
		return fmt.Errorf("not connected")
	}
	w.msgRecv[msg.ID()] = make(chan message.Message)
	serialized, err := message.Serialize(msg)
	println(serialized)
	if err != nil {
		return err
	}
	return w.conn.WriteMessage(websocket.TextMessage, []byte(serialized))
}

func (w *WebsocketConnector) SendRecv(m message.Message) (message.Message, error) {
	m.SetID(w.nextID())
	err := w.Send(m)
	if err != nil {
		return nil, err
	}
	return w.recv(m.ID())
}

func (w *WebsocketConnector) recv(id uint32) (message.Message, error) {
	msg := <-w.msgRecv[id]
	if err, ok := msg.(*message.Error); ok {
		return err, err.Error()
	}
	delete(w.msgRecv, id)
	return msg, nil
}

func (w *WebsocketConnector) nextID() uint32 {
	id := w.idCounter
	w.idCounter++
	return id
}