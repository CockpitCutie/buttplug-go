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
}

// Create a new Buttplugio client with a given name
func New(name string) *Client {
	return &Client{
		name:       name,
		connector:  nil,
		msg_recv:   make(map[uint32]chan message.Message),
		serverName: "",
	}
}

// Connect to a Buttplug server using the provided Connector
func (c *Client) Connect(connector Connector) error {
	c.connector = connector
	err := connector.Connect(c.msg_recv)
	if err != nil {
		return err
	}
	return c.onConnect()
}

func (c *Client) onConnect() error {
	recv, err := c.connector.SendRecv(&message.RequestServerInfo{
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

// Check if the client is connected to a Buttplug server
func (c Client) Connected() bool {
	return c.connector != nil && c.Connected()
}

// Disconnect from the Buttplug server
func (c *Client) Disconnect() error {
	return c.connector.Disconnect()
}

// Start scanning for devices
func (c *Client) StartScanning() error {
	_, err := c.connector.SendRecv(&message.StartScanning{})
	if err != nil {
		return err
	}
	return nil
}

// Stop scanning for devices
func (c *Client) StopScanning() error {
	_, err := c.connector.SendRecv(&message.StopScanning{})
	if err != nil {
		return err
	}
	return nil
}

// Stop all connected devices
func (c *Client) StopAllDevices() error {
	_, err := c.connector.SendRecv(&message.StopAllDevices{})
	if err != nil {
		return err
	}
	return nil
}

// Retrieve the list of connected devices from the Buttplug server
func (c *Client) Devices() ([]device.Device, error) {
	devicelist, err := c.connector.SendRecv(&message.RequestDeviceList{})
	if err != nil {
		return nil, err
	}

	if devicelist, ok := devicelist.(*message.DeviceList); ok {
		fmt.Printf("%+v", devicelist)
		var devices []device.Device
		for _, d := range devicelist.Devices {
			devices = append(devices, device.FromMessage(d, c.connector))
		}
		return devices, nil
	}
	return nil, fmt.Errorf("expected DeviceList, found %T", devicelist)
}

// Ping the Buttplug server to maintain a keep-alive, or check connectivity
func (c *Client) Ping() error {
	_, err := c.connector.SendRecv(&message.Ping{})
	if err != nil {
		return err
	}
	return nil
}

// Get the server name of the connected Buttplug server
func (c Client) ServerName() string {
	return c.serverName
}
