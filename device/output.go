package device

import "github.com/CockpitCutie/buttplug-go/message"

type Output interface {
	Feature
	OutputType() OutputType
}

type OutputType string

const (
	VibrateOutput               OutputType = "Vibrate"
	RotateOutput                OutputType = "RotateOutput"
	RotationWithDirectionOutput OutputType = "RotationWithDirection"
	OscillateOutput             OutputType = "Oscillate"
	ConstrictOutput             OutputType = "Constrict"
	HeaterOutput                OutputType = "Heater"
	LEDOutput                   OutputType = "Led"
	PositionOutput              OutputType = "Position"
	PositionWithDirectionOutput OutputType = "PositionWithDirection"
)

func outputsFromFeatures(features message.DeviceFeatures, sender MessageSender) ([]Output, error) {
	var outputs []Output
	for _, featureMsg := range features {
		feature := feature{
			description: featureMsg.FeatureDescription,
			index:       featureMsg.FeatureIndex,
			sender:      sender,
		}
		_ = feature
		if featureMsg.Input == nil {
			continue
		}
		for kind, properties := range featureMsg.Input {
			_, _ = kind, properties
		}
		outputs = append(outputs, nil)
	}
	return outputs, nil
}

type Vibrator struct {
	feature
	stepCount uint32
}

func (v Vibrator) OutputType() OutputType {
	return VibrateOutput
}

type Rotator struct {
	feature
	stepCount uint32
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

type RotatorWithDirection struct {
	feature
	stepCount uint32
	clockwise bool
}

func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

type Oscillator struct {
	feature
	stepCount uint32
}

func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

type Constrictor struct {
	feature
	stepCount uint32
}

func (c Constrictor) OutputType() OutputType {
	return ConstrictOutput
}

type Heater struct {
	feature
	stepCount uint32
}

func (h Heater) OutputType() OutputType {
	return HeaterOutput
}

type LED struct {
	feature
	stepCount uint32
}

func (l LED) OutputType() OutputType {
	return LEDOutput
}

type Position struct {
	feature
	stepCount uint32
}

func (p Position) OutputType() OutputType {
	return PositionOutput
}

type PositionWithDuration struct {
	stepCount uint32
}

func (r PositionWithDuration) OutputType() OutputType {
	return PositionWithDirectionOutput
}
