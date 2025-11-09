package buttplug

import (
	"github.com/CockpitCutie/buttplug-go/message"
)

type Client struct {
	name       string
	connector  Connector
	msg_recv   map[uint32]chan message.Message
	serverName string
}

func New(name string) *Client {
	return &Client{
		name:       name,
		connector:  nil,
		msg_recv:   make(map[uint32]chan message.Message),
		serverName: "",
	}
}

func (c *Client) Connect(connector Connector) error {
	c.connector = connector
	err := connector.Connect(c.msg_recv)
	if err != nil {
		return err
	}
	return c.onConnect()
}

func (c *Client) onConnect() error {
	msg := &message.RequestServerInfo{
		ClientName:     c.name,
		MessageVersion: 3,
	}
	msg.SetID(1)
	c.connector.Send(msg)
	res := <-c.msg_recv[1]

	if serverInfo, ok := res.(*message.ServerInfo); ok {
		c.serverName = serverInfo.ServerName
	}
	println(c.ServerName())
	return nil
}

func (c Client) Connected() bool {
	return c.connector != nil && c.Connected()
}

func (c *Client) Disconnect() error {
	return c.connector.Disconnect()
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

func (c Client) ServerName() string {
	return c.serverName
}

type Device struct{}
