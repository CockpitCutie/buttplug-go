package buttplug

import "github.com/CockpitCutie/buttplug-go/message"

type Connector interface {
	Connect(msg_recv chan message.Message) error
	Disconnect() error
	Send(msg message.Message)
}