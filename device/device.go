package device

import "github.com/CockpitCutie/buttplug-go/message"

type Device struct {
	Name             string
	Index            uint
	MessageTimingGap uint
	DisplayName      string
	Inputs           []Input
	Outputs          []Output
	msgSender        MessageSender
}

type MessageSender interface {
	SendRecv(message.Message) (message.Message, error)
}

func FromDeviceList(devlist *message.DeviceList, sender MessageSender) ([]Device, error) {
	var devices []Device
	for _, msg := range devlist.Devices {
		dev, err := newDevice(msg, sender)
		if err != nil {
			return nil, err
		}
		devices = append(devices, dev)
	}
	return devices, nil
}

func newDevice(msg message.Device, msgSender MessageSender) (Device, error) {
	device := Device{
		Name:             msg.DeviceName,
		Index:            msg.DeviceIndex,
		MessageTimingGap: msg.DeviceMessageTimingGap,
		DisplayName:      msg.DeviceDisplayName,
		msgSender:        msgSender,
	}
	err := device.registerOutputs(msg.Features)
	if err != nil {
		return Device{}, err
	}
	err = device.registerInputs(msg.Features)
	if err != nil {
		return Device{}, err
	}
	
	return device, nil
}

type Feature interface {
	Description() string
	Index() uint32
	Device() *Device
}

type feature struct {
	description string
	index       uint32
	device      *Device
}

func (f feature) Description() string {
	return f.description
}

func (f feature) Index() uint32 {
	return f.index
}

func (f feature) Device() *Device {
	return f.device
}
