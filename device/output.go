package device

import (
	"time"

	"github.com/CockpitCutie/buttplug-go/message"
)

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

func (v Vibrator) Activate(step uint32) error {
	return nil
}

func (v Vibrator) Vibrate(step uint32) error {
	return v.Activate(step)
}

type Rotator struct {
	feature
	stepCount uint32
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

func (r Rotator) Activate(step uint32) error {
	return nil
}

func (r Rotator) Rotate(step uint32) error {
	return r.Rotate(step)
}

type RotatorWithDirection struct {
	feature
	stepCount uint32
}

func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

func (r RotatorWithDirection) Activate(step uint32, clockwise bool) error {
	return nil
}
func (r RotatorWithDirection) RotateDirection(step uint32, clockwise bool) error {
	return r.Activate(step, clockwise)
}

type Oscillator struct {
	feature
	stepCount uint32
}

func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

func (o Oscillator) Activate(step uint32) error {
	return nil
}

type Constrictor struct {
	feature
	stepCount uint32
}

func (c Constrictor) OutputType() OutputType {
	return ConstrictOutput
}

func (c Constrictor) Activate(step uint32) error {
	return nil
}

type Heater struct {
	feature
	stepCount uint32
}

func (h Heater) OutputType() OutputType {
	return HeaterOutput
}

func (h Heater) Activate(step uint32) error {
	return nil
}

type LED struct {
	feature
	stepCount uint32
}

func (l LED) OutputType() OutputType {
	return LEDOutput
}

func (l LED) Activate(step uint32) error {
	return nil
}

type Position struct {
	feature
	stepCount uint32
}

func (p Position) OutputType() OutputType {
	return PositionOutput
}

func (p Position) Activate(step uint32) error {
	return nil
}

type PositionWithDuration struct {
	feature
	stepCount uint32
}

func (p PositionWithDuration) OutputType() OutputType {
	return PositionWithDirectionOutput
}

func (p PositionWithDuration) Activate(step uint32, duration time.Time) error {
	return nil
}
