package buttplug

import (
	"fmt"

	"github.com/CockpitCutie/buttplug-go/device"
	"github.com/CockpitCutie/buttplug-go/message"
)

// Client represents a Buttplugio client that can connect to a Buttplug server.
type Client struct {
	name       string
	connector  Connector
	msg_recv   map[uint32]chan message.Message
	serverName string
	idCounter  uint32
}

func New(name string) *Client {
	return &Client{
		name:       name,
		connector:  nil,
		msg_recv:   make(map[uint32]chan message.Message),
		serverName: "",
		idCounter:  1,
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
	recv, err := c.sendRecv(&message.RequestServerInfo{
		ClientName:     c.name,
		MessageVersion: 3,
	})
	if err != nil {
		return err
	}
	if serverInfo, ok := recv.(*message.ServerInfo); ok {
		c.serverName = serverInfo.ServerName
	}
	return nil
}

func (c Client) Connected() bool {
	return c.connector != nil && c.Connected()
}

func (c *Client) Disconnect() error {
	return c.connector.Disconnect()
}

func (c *Client) StartScanning() error {
	_, err := c.sendRecv(&message.StartScanning{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StopScanning() error {
	_, err := c.sendRecv(&message.StopScanning{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StopAllDevices() error {
	_, err := c.sendRecv(&message.StopAllDevices{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Devices() ([]device.Device, error) {
	devicelist, err := c.sendRecv(&message.RequestDevicelist{})
	if err != nil {
		return nil, err
	}
	if devicelist, ok := devicelist.(*message.DeviceList); ok {
		var devices []device.Device
		for _, d := range devicelist.Devices {
			devices = append(devices, device.FromMessage(d))
		}
		return devices, nil
	}
	return nil, fmt.Errorf("expected DeviceList, found %T", devicelist)
	
}

func (c *Client) Ping() error {
	_, err := c.sendRecv(&message.Ping{})
	if err != nil {
		return err
	}
	return nil
}

func (c Client) ServerName() string {
	return c.serverName
}

func (c *Client) sendRecv(m message.Message) (message.Message, error) {
	id := c.idCounter
	c.idCounter++
	m.SetID(id)
	err := c.connector.Send(m)
	if err != nil {
		return nil, err
	}
	recv := <-c.msg_recv[id]
	if err, ok := recv.(*message.Error); !ok {
		return err, err.Error()
	}
	delete(c.msg_recv, id)
	return recv, nil
}
