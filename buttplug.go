package buttplug

import "github.com/CockpitCutie/buttplug-go/message"

type Client struct {
	name string
	connector Connector
	msg_recv chan message.Message
	serverName string
}

func New(name string) *Client {
	return &Client{
		name: name,
		connector: nil,
		msg_recv: make(chan message.Message),
		serverName: "",
	}
}

func (c *Client) Connect(connector Connector) error {
	return connector.Connect(c.msg_recv)
}

func (c Client) Connected() bool {
	return c.connector != nil && c.Connected()
}

func (c *Client) Disconnect() error {
	return c.Disconnect()
}

func (c *Client) StartScanning() error {
	return nil
}

func (c *Client) StopScanning() error {
	return nil
}

func (c *Client) StopAllDevices() error {
	return nil
}

func (c *Client) Devices() []Device {
	return nil
}

func (c *Client) Ping() error {
	return nil
}

func (c *Client) ServerName() *string {
	return &c.serverName
}

type Device struct {}