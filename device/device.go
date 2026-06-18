package device

import (
	"github.com/CockpitCutie/buttplug-go/message"
)

// Device represents a physical device that can be controlled by the Buttplug server. 
// It contains information about the device's name, index, display name, and its 
// features. Hardware controls can be performed by accessing specific 
// inputs and outputs of the device, such as vibrators, rotators, buttons, etc.
type Device struct {
	Name             string
	Index            int
	MessageTimingGap int
	DisplayName      string
	Inputs           map[int]Input
	Outputs          map[int]Output
	msgSender        messageSender
}

// messageSender is a minimal interface for sending messages to the server, 
// used internally by Device for sending commands.
type messageSender interface {
	SendRecv(message.Message) (message.Message, error)
}

// ----- Device Constructors -----

// FromDeviceList creates a list of Device instances from a DeviceList message 
// received from the server.
func FromDeviceList(devlist *message.DeviceList, sender messageSender) ([]Device, error) {
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

// newDevice creates a Device instance from a Device message received from the server.
func newDevice(msg message.Device, msgSender messageSender) (Device, error) {
	device := Device{
		Name:             msg.DeviceName,
		Index:            msg.DeviceIndex,
		MessageTimingGap: msg.DeviceMessageTimingGap,
		DisplayName:      msg.DeviceDisplayName,
		msgSender:        msgSender,
		Inputs:           make(map[int]Input),
		Outputs:          make(map[int]Output),
	}
	err := device.registerOutputs(msg.DeviceFeatures)
	if err != nil {
		return Device{}, err
	}
	err = device.registerInputs(msg.DeviceFeatures)
	if err != nil {
		return Device{}, err
	}

	return device, nil
}

// ----- Stop Commands -----

// Stop stops all inputs and outputs for the device. 
// It returns an error if the stop command fails or if sending a message to the
// server fails (for example if the connection was dropped).
func (d Device) Stop() error {
	t := true // to get an address for *bool
	stopMsg := message.StopCmd{
		DeviceIndex: d.Index,
		Inputs: &t,
		Outputs: &t,
	}
	_, err := d.msgSender.SendRecv(&stopMsg)
	return err
}

// ----- Output Accessors -----

// Vibrators returns a list of all output features with Vibrate capabilities.
// It returns a nil slice if the device contains no vibrating outputs.
func (d Device) Vibrators() []Vibrator {
	var outputs []Vibrator
	for _, output := range d.Outputs {
		if output, ok := output.(Vibrator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Rotators returns a list of all output features with Rotate capabilities.
// It returns a nil slice if the device contains no rotating outputs.
func (d Device) Rotators() []Rotator {
	var outputs []Rotator
	for _, output := range d.Outputs {
		if output, ok := output.(Rotator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// RotatorsWithDirection returns a list of all output features with RotateWithDirection
// capabilities. It returns a nil slice if the device contains no rotating outputs 
// with direction.
func (d Device) RotatorsWithDirection() []RotatorWithDirection {
	var outputs []RotatorWithDirection
	for _, output := range d.Outputs {
		if output, ok := output.(RotatorWithDirection); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Oscillators returns a list of all output features with Oscillate capabilities.
// It returns a nil slice if the device contains no oscillating outputs.
func (d Device) Oscillators() []Oscillator {
	var outputs []Oscillator
	for _, output := range d.Outputs {
		if output, ok := output.(Oscillator); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Constrictors returns a list of all output features with Constrict capabilities.
// It returns a nil slice if the device contains no constricting outputs.
func (d Device) Constrictors() []Constrictor {
	var outputs []Constrictor
	for _, output := range d.Outputs {
		if output, ok := output.(Constrictor); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Heaters returns a list of all output features with Heater capabilities.
// It returns a nil slice if the device contains no heating outputs.
func (d Device) Heaters() []Heater {
	var outputs []Heater
	for _, output := range d.Outputs {
		if output, ok := output.(Heater); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// LEDs returns a list of all output features with LED capabilities.
// It returns a nil slice if the device contains no lighting outputs.
func (d Device) LEDs() []LED {
	var outputs []LED
	for _, output := range d.Outputs {
		if output, ok := output.(LED); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Positioners returns a list of all output features with Position capabilities.
// It returns a nil slice if the device contains no positioning outputs.
func (d Device) Positioners() []Position {
	var outputs []Position
	for _, output := range d.Outputs {
		if output, ok := output.(Position); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// PositionersWithDuration returns a list of all output features with PositionWithDuration
// capabilities. It returns a nil slice if the device contains no positioning 
// outputs with duration.
func (d Device) PositionersWithDuration() []PositionWithDuration {
	var outputs []PositionWithDuration
	for _, output := range d.Outputs {
		if output, ok := output.(PositionWithDuration); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// ----- Input Accessors -----

// Batteries returns a list of all input features with Battery capabilities.
// It returns a nil slice if the device contains no battery level inputs.
func (d Device) Batteries() []Battery {
	var inputs []Battery
	for _, input := range d.Inputs {
		if input, ok := input.(Battery); ok {
			inputs = append(inputs, input)
		}
	}
	return inputs
}

// RSSIs returns a list of all input features with RSSI capabilities.
// It returns a nil slice if the device contains no RSSI inputs.
func (d Device) RSSIs() []RSSI {
	var inputs []RSSI
	for _, input := range d.Inputs {
		if input, ok := input.(RSSI); ok {
			inputs = append(inputs, input)
		}
	}
	return inputs
}

// Pressures returns a list of all input features with Pressure capabilities.
// It returns a nil slice if the device contains no pressure inputs.
func (d Device) Pressures() []Pressure {
	var inputs []Pressure
	for _, input := range d.Inputs {
		if input, ok := input.(Pressure); ok {
			inputs = append(inputs, input)
		}
	}
	return inputs
}

// Buttons returns a list of all input features with Button capabilities.
// It returns a nil slice if the device contains no button inputs.
func (d Device) Buttons() []Button {
	var inputs []Button
	for _, input := range d.Inputs {
		if input, ok := input.(Button); ok {
			inputs = append(inputs, input)
		}
	}
	return inputs
}

// ----- Device Features

// Feature represents a common interface for a hardware controllable component
// of a device, such as a vibrator, rotator, button, etc. It is common between
// inputs and outputs. It provides methods to access the feature's description, 
// index, and parent device. 
// 
// This interface is intended to share behavior and functionality between inputs
// and outputs, not necessarily to be used on its own.
type Feature interface {
	Description() string
	Index() int
	Device() *Device
}

// feature represents a concrete implementation of Feature that can be embedded
// in other structs.
type feature struct {
	description string
	index       int
	device      *Device
}

// Description returns a human readable description of the feature, such as 
// "Clitoral Stimulator" or "Battery".
func (f feature) Description() string {
	return f.description
}

// Index returns the feature index of the feature, which is used to identify 
// the feature when paired with the device index, and is used for sending 
// commands to control a specific feature.
func (f feature) Index() int {
	return f.index
}

// Device returns a pointer to the parent device of the feature. This can be used
// to access the device's name, index, and other Device methods.
// 
// This is used internally for sending commands to the server, since the device
// contains the handle for message sending.
func (f feature) Device() *Device {
	return f.device
}
