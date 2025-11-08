package buttplug

import "github.com/CockpitCutie/buttplug-go/message"

type Connector interface {
	Connect(msg_recv chan message.Message) error
	Connected() bool
	Disconnect() error
	Send(msg message.Message)
}