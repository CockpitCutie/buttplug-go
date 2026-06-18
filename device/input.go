package device

import (
	"fmt"

	"github.com/CockpitCutie/buttplug-go/message"
)

// Input represents a common interface for device inputs, like Buttons, Battery,
// RSSI, and Pressures.
//
// The Input interface does not provide a common method for reading from inputs,
// as they all have different return types. Instead, each input type has its own
// Read method, and Inputs must be type asserted to their specific type to access
// it.
type Input interface {
	Feature
	InputType() InputType
	ReadRange() [2]int
}

// InputType represents the type of an input feature, such as Battery, RSSI, Pressure,
// or Button. This list of input types is taken from the ButtplugIO v4 specification.
type InputType string

const (
	BatteryInput  InputType = "Battery"
	RSSIInput     InputType = "RSSI"
	PressureInput InputType = "Pressure"
	ButtonInput   InputType = "Button"
)

// ----- Device methods for dealing with inputs -----

// registerInputs takes a list of device features from a DeviceFeatures message and adds
// any features with input capabilities to the device's Inputs map. It returns an error
// if any of the input features have an unknown type or invalid properties.
func (d *Device) registerInputs(features message.DeviceFeatures) error {
	for _, featureMsg := range features {
		feature := feature{
			description: featureMsg.FeatureDescription,
			index:       featureMsg.FeatureIndex,
			device:      d,
		}
		if featureMsg.Input == nil {
			continue
		}
		for kind, properties := range featureMsg.Input {
			input, err := makeInput(InputType(kind), properties, feature)
			if err != nil {
				return err
			}
			d.Inputs[input.Index()] = input
		}

	}
	return nil
}

// makeInput is a helper function that takes an input type, its properties from a DeviceInput message,
// and the parent feature, and returns an instance of the corresponding Input type. It returns an error
// if the input type is unknown or if the properties are invalid for that input type.
func makeInput(kind InputType, properties message.DeviceInput, feature feature) (Input, error) {
	switch kind {
	case BatteryInput:
		return Battery{feature: feature}, nil
	case RSSIInput:
		return RSSI{feature: feature, readRange: properties.Value}, nil
	case PressureInput:
		return Pressure{feature: feature, readRange: properties.Value}, nil
	case ButtonInput:
		return Button{feature: feature}, nil
	default:
		return nil, fmt.Errorf("unknown input type: %s", kind)
	}
}

// ReadInput is a helper method on Device that takes an Input and sends a message
// to the device to read the current value of that input. It returns the InputReading
// message received from the device, or an error if the message sending or receiving
// fails, or the received message is not of the expected type.
func (d Device) readInput(input Input) (*message.InputReading, error) {
	msg := message.InputCmd{
		DeviceIndex:  d.Index,
		FeatureIndex: input.Index(),
		Type:         string(input.InputType()),
		Command:      "Read",
	}
	reading, err := d.msgSender.SendRecv(&msg)
	if err != nil {
		return nil, err
	}
	if reading, ok := reading.(*message.InputReading); ok {
		return reading, nil
	}
	return nil, fmt.Errorf("unexpected message type for %s input reading: %T", input.InputType(), reading)
}

// ----- Battery Input -----

// Battery represents a battery level input feature of a device. It provides methods
// to read the battery level as a percentage.
type Battery struct {
	feature
}

// InputType returns the type of the input, which is always BatteryInput ("Battery").
func (b Battery) InputType() InputType {
	return BatteryInput
}

// ReadRange returns the range of valid values that the input can return. For Battery
// inputs, this is always [0, 100], representing a percentage.
func (b Battery) ReadRange() [2]int {
	return [2]int{0, 100}
}

// Read returns the current battery level as a percentage. The value is between 0
// and 100, inclusive.
func (b Battery) Read() (int, error) {
	reading, err := b.Device().readInput(b)
	if err != nil {
		return 0, err
	}
	if reading.Reading.Battery != nil {
		return reading.Reading.Battery.Value, nil
	}
	return 0, fmt.Errorf("unexpected input reading type for battery input: %v", reading.Reading)
}

// Convenience method for Read, to make it more clear that this accesses battery
// level as a percentage.
func (b Battery) Percentage() (int, error) {
	return b.Read()
}

// ----- RSSI Input -----

// RSSI represents a signal strength input feature of a device. It provides methods
// to read the current signal strength level.
type RSSI struct {
	feature
	readRange [2]int
}

// InputType returns the type of the input, which is always RSSIInput ("RSSI").
func (r RSSI) InputType() InputType {
	return RSSIInput
}

// ReadRange returns the range of valid values that the input can return. The specific
// range of values can vary between devices, but the result is always negative, and
// usually between -10 and -100, where values closer to 0 represent stronger signals.
func (r RSSI) ReadRange() [2]int {
	return r.readRange
}

// Read returns the current signal strength level. The value is between the minimum and
// maximum values specified in the ReadRange.
func (r RSSI) Read() (int, error) {
	reading, err := r.Device().readInput(r)
	if err != nil {
		return 0, err
	}
	if reading.Reading.RSSI != nil {
		return reading.Reading.RSSI.Value, nil
	}
	return 0, fmt.Errorf("unexpected input reading type for RSSI input: %v", reading.Reading)
}

// ----- Pressure Input -----

// Pressure represents a pressure level input feature of a device. It provides methods
// to read the current pressure level.
type Pressure struct {
	feature
	readRange [2]int
}

// InputType returns the type of the input, which is always PressureInput ("Pressure").
func (p Pressure) InputType() InputType {
	return PressureInput
}

// ReadRange returns the range of valid values that the input can return. The specific
// range of values for pressure inputs can vary widely between devices, and there is
// no standardized unit or scale for pressure levels.
func (p Pressure) ReadRange() [2]int {
	return p.readRange
}

// Read returns the current pressure level. The value is between the minimum and
// maximum values specified in the ReadRange.
//
// IMPORTANT: there is no standardized level between manufacturers for pressure
// inputs, so the values returned by this method may not be consistent across
// different devices. Calibration and unit interpretation must be handled at
// the application level.
func (p Pressure) Read() (int, error) {
	reading, err := p.Device().readInput(p)
	if err != nil {
		return 0, err
	}
	if reading.Reading.Pressure != nil {
		return reading.Reading.Pressure.Value, nil
	}
	return 0, fmt.Errorf("unexpected input reading type for pressure input: %v", reading.Reading)
}

// ----- Button Input -----

// Button represents a button input feature of a device. It provides methods to read
// the current state of the button (pressed or not pressed).
type Button struct {
	feature
}

// InputType returns the type of the input, which is always ButtonInput ("Button").
func (b Button) InputType() InputType {
	return ButtonInput
}

// ReadRange returns the range of valid values that the input can return. For Button
// inputs, this is always [0, 1], where 0 represents not pressed and 1 represents pressed.
func (b Button) ReadRange() [2]int {
	return [2]int{0, 1}
}

// Read returns the current state of the button, where 0 represents not pressed
// and 1 represents pressed.
func (b Button) Read() (int, error) {
	reading, err := b.Device().readInput(b)
	if err != nil {
		return 0, err
	}
	if reading.Reading.Button != nil {
		return reading.Reading.Button.Value, nil
	}
	return 0, fmt.Errorf("unexpected input reading type for button input: %v", reading.Reading)
}

// IsPressed returns true if the button is currently pressed, and false otherwise.
// If there is an error reading the button state, it returns false by default.
func (b Button) IsPressed() bool {
	data, err := b.Read()
	if err != nil {
		return false
	}
	return data == 1
}
