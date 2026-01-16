package device

import "github.com/CockpitCutie/buttplug-go/message"

type Output interface {
	Feature
	OutputType() OutputType
}

type OutputType string

func outputsFromFeatures(msg message.DeviceFeatures, sender MessageSender) ([]Output, error) {
	return nil, nil
}
