package device

import "github.com/CockpitCutie/buttplug-go/message"

type Device struct {
	Name             string
	Index            uint
	MessageTimingGap uint
	DisplayName      string
	Inputs           map[uint32]Input
	Outputs          map[uint32]Output
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
		Inputs:           make(map[uint32]Input),
		Outputs:          make(map[uint32]Output),
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

func (d Device) Vibrators() []Vibrator {
	var outputs []Vibrator
	for _, output := range d.Outputs {
		if output, ok := output.(Vibrator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) Rotators() []Rotator {
	var outputs []Rotator
	for _, output := range d.Outputs {
		if output, ok := output.(Rotator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) RotatorsWithDirection() []RotatorWithDirection {
	var outputs []RotatorWithDirection
	for _, output := range d.Outputs {
		if output, ok := output.(RotatorWithDirection); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) Oscillators() []Oscillator {
	var outputs []Oscillator
	for _, output := range d.Outputs {
		if output, ok := output.(Oscillator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) Constrictors() []Constrictor {
	var outputs []Constrictor
	for _, output := range d.Outputs {
		if output, ok := output.(Constrictor); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) Heaters() []Heater {
	var outputs []Heater
	for _, output := range d.Outputs {
		if output, ok := output.(Heater); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) LEDs() []LED {
	var outputs []LED
	for _, output := range d.Outputs {
		if output, ok := output.(LED); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) Positioners() []Position {
	var outputs []Position
	for _, output := range d.Outputs {
		if output, ok := output.(Position); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func (d Device) PositionersWithDuration() []PositionWithDuration {
	var outputs []PositionWithDuration
	for _, output := range d.Outputs {
		if output, ok := output.(PositionWithDuration); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
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
