package device

import (
	"github.com/CockpitCutie/buttplug-go/message"
)

type Input interface {
	Feature
	InputType() InputType
	ReadRange() [2]int32
}

type InputType string

const (
	BatteryInput  InputType = "Battery"
	RSSIInput     InputType = "RSSI"
	PressureInput InputType = "Pressure"
	ButtonInput   InputType = "Button"
)

func inputsFromFeatures(features message.DeviceFeatures, sender MessageSender) ([]Input, error) {
	var inputs []Input
	for _, featureMsg := range features {
		feature := feature{
			description: featureMsg.FeatureDescription,
			index: featureMsg.FeatureIndex,
			sender: sender,
		}
		if featureMsg.Input == nil {
			continue
		}
		for kind, properties := range featureMsg.Input {
			_, _ = kind, properties
		}
		inputs = append(inputs, Battery{feature: feature})
	}
	return inputs, nil
}

func makeInput(kind InputType, properties message.DeviceInput)

type Battery struct {
	feature
}

func (b Battery) InputType() InputType {
	return BatteryInput
}

func (b Battery) ReadRange() [2]int32 {
	return [2]int32{0, 100}
}

func (b Battery) Read() uint8 {
	return 0 // TODO
}

func (b Battery) Percentage() uint8 {
	return b.Read()
}

type RSSI struct {
	feature
	readRange [2]int32
}

func (r RSSI) InputType() InputType {
	return RSSIInput
}

func (r RSSI) ReadRange() [2]int32 {
	return r.readRange
}

func (r RSSI) Read() int8 {
	return 0 // TODO
}

type Pressure struct {
	feature
	readRange [2]int32
}

func (p Pressure) InputType() InputType {
	return PressureInput
}

func (p Pressure) ReadRange() [2]int32 {
	return p.readRange
}

func (p Pressure) Read() uint32 {
	return 0 // TODO
}

type Button struct {
	feature
}

func (b Button) InputType() InputType {
	return ButtonInput
}

func (b Button) ReadRange() [2]int32 {
	return [2]int32{0, 1}
}

func (b Button) Read() uint8 {
	return 0
}

func (b Button) IsPressed() bool {
	return b.Read() == 1
}
