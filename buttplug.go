package buttplug

type Client struct {
	
}

func New(name string) *Client {
	return nil
}

func (c *Client) Connect(connector Connector) error {
	return nil
}

func (c Client) Connected() bool {
	return false
}

func (c *Client) Disconnect() error {
	return nil
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
	return nil
}

type Device struct {}