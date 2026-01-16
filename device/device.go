package device

import "github.com/CockpitCutie/buttplug-go/message"

type Device struct {
	Name             string
	Index            uint
	MessageTimingGap uint
	DisplayName      string
	Inputs           []Input
	Outputs          []Output
	MsgSender        MessageSender
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

func newDevice(msg message.Device, sender MessageSender) (Device, error) {
	inputs, err := inputsFromFeatures(msg.Features, sender)
	if err != nil {
		return Device{}, err
	}
	outputs, err := outputsFromFeatures(msg.Features, sender)
	if err != nil {
		return Device{}, err
	}
	return Device{
		Name:             msg.DeviceName,
		Index:            msg.DeviceIndex,
		MessageTimingGap: msg.DeviceMessageTimingGap,
		DisplayName:      msg.DeviceDisplayName,
		Inputs:           inputs,
		Outputs:          outputs,
		MsgSender:        sender,
	}, nil
}

type Feature interface {
	Description() string
	Index() uint32
	Sender() MessageSender
}

type feature struct {
	description string
	index       uint32
	sender      MessageSender
}

func (f feature) Description() string {
	return f.description
}

func (f feature) Index() uint32 {
	return f.index
}

func (f feature) Sender() MessageSender {
	return f.sender
}
