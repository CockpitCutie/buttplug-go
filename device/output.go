package device

import (
	"fmt"
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

func (d *Device) registerOutputs(features message.DeviceFeatures) error {
	for _, featureMsg := range features {
		feature := feature{
			description: featureMsg.FeatureDescription,
			index:       featureMsg.FeatureIndex,
			device:      d,
		}
		_ = feature
		if featureMsg.Output == nil {
			continue
		}
		for kind, properties := range featureMsg.Output {
			output, err := outputFromProps(OutputType(kind), properties, feature)
			if err != nil {
				return err
			}
			d.Outputs = append(d.Outputs, output)
		}
	}
	return nil
}

func outputFromProps(kind OutputType, properties message.DeviceOutput, feature feature) (Output, error) {
	switch kind {
	case VibrateOutput:
		return Vibrator{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case RotateOutput:
		return Rotator{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case RotationWithDirectionOutput:
		return RotatorWithDirection{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case OscillateOutput:
		return Oscillator{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case ConstrictOutput:
		return Constrictor{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case HeaterOutput:
		return Heater{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case LEDOutput:
		return LED{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case PositionOutput:
		return Position{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	case PositionWithDirectionOutput:
		return PositionWithDuration{
			feature:   feature,
			stepCount: uint32(properties.Value[1]),
		}, nil
	default:
		return nil, fmt.Errorf("unknown device type %s", kind)
	}
}

type Vibrator struct {
	feature
	stepCount uint32
}

func (v Vibrator) OutputType() OutputType {
	return VibrateOutput
}

func (v Vibrator) Activate(speedStep uint32) error {
	return nil
}

func (v Vibrator) Vibrate(speedStep uint32) error {
	return v.Activate(speedStep)
}

type Rotator struct {
	feature
	stepCount uint32
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

func (r Rotator) Activate(speedStep uint32) error {
	return nil
}

func (r Rotator) Rotate(speedStep uint32) error {
	return r.Rotate(speedStep)
}

type RotatorWithDirection struct {
	feature
	stepCount uint32
}

func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

func (r RotatorWithDirection) Activate(speedStep uint32, clockwise bool) error {
	return nil
}
func (r RotatorWithDirection) RotateDirection(speedStep uint32, clockwise bool) error {
	return r.Activate(speedStep, clockwise)
}

type Oscillator struct {
	feature
	stepCount uint32
}

func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

func (o Oscillator) Activate(speedStep uint32) error {
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

func (h Heater) Activate(heatLevel uint32) error {
	return nil
}

type LED struct {
	feature
	stepCount uint32
}

func (l LED) OutputType() OutputType {
	return LEDOutput
}

func (l LED) Activate(brightnessStep uint32) error {
	return nil
}

type Position struct {
	feature
	stepCount uint32
}

func (p Position) OutputType() OutputType {
	return PositionOutput
}

func (p Position) Activate(positionStep uint32) error {
	return nil
}

type PositionWithDuration struct {
	feature
	stepCount uint32
}

func (p PositionWithDuration) OutputType() OutputType {
	return PositionWithDirectionOutput
}

func (p PositionWithDuration) Activate(positionStep uint32, duration time.Time) error {
	return nil
}
