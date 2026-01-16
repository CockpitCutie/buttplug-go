package device

import "github.com/CockpitCutie/buttplug-go/message"

type Input interface {
	Feature
	InputType() InputType
}

type InputType string

func inputsFromFeatures(features message.DeviceFeatures, sender MessageSender) ([]Input, error) {
	return nil, nil
}
