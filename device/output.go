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
			d.Outputs[output.Index()] = output
		}
	}
	return nil
}

func outputFromProps(kind OutputType, properties message.DeviceOutput, feature feature) (Output, error) {
	switch kind {
	case VibrateOutput:
		return Vibrator{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case RotateOutput:
		return Rotator{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case RotationWithDirectionOutput:
		return RotatorWithDirection{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case OscillateOutput:
		return Oscillator{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case ConstrictOutput:
		return Constrictor{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case HeaterOutput:
		return Heater{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case LEDOutput:
		return LED{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case PositionOutput:
		return Position{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	case PositionWithDirectionOutput:
		return PositionWithDuration{
			feature:   feature,
			stepCount: properties.Value[1],
		}, nil
	default:
		return nil, fmt.Errorf("unknown device type %s", kind)
	}
}

type Vibrator struct {
	feature
	stepRange [2]int
}

func (v Vibrator) OutputType() OutputType {
	return VibrateOutput
}

func (v Vibrator) Activate(speedStep int) error {
	stepIsInRange := speedStep == 0 || (speedStep > v.stepRange[0] && speedStep < v.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("speed step %d is out of range for this vibrator", speedStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  v.Device().Index,
		FeatureIndex: v.Index(),
		Command: message.OutputValue{
			"Vibrate": {Value: speedStep},
		},
	}
	_, err := v.device.msgSender.SendRecv(&msg)
	return err
}

func (v Vibrator) Vibrate(speedStep int) error {
	return v.Activate(speedStep)
}

type Rotator struct {
	feature
	stepCount int
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

func (r Rotator) Activate(speedStep int) error {
	return nil
}

func (r Rotator) Rotate(speedStep int) error {
	return r.Rotate(speedStep)
}

type RotatorWithDirection struct {
	feature
	stepCount int
}

func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

func (r RotatorWithDirection) Activate(speedStep int, clockwise bool) error {
	return nil
}
func (r RotatorWithDirection) RotateDirection(speedStep int, clockwise bool) error {
	return r.Activate(speedStep, clockwise)
}

type Oscillator struct {
	feature
	stepCount int
}

func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

func (o Oscillator) Activate(speedStep int) error {
	return nil
}

type Constrictor struct {
	feature
	stepCount int
}

func (c Constrictor) OutputType() OutputType {
	return ConstrictOutput
}

func (c Constrictor) Activate(step int) error {
	return nil
}

type Heater struct {
	feature
	stepCount int
}

func (h Heater) OutputType() OutputType {
	return HeaterOutput
}

func (h Heater) Activate(heatLevel int) error {
	return nil
}

type LED struct {
	feature
	stepCount int
}

func (l LED) OutputType() OutputType {
	return LEDOutput
}

func (l LED) Activate(brightnessStep int) error {
	return nil
}

type Position struct {
	feature
	stepCount int
}

func (p Position) OutputType() OutputType {
	return PositionOutput
}

func (p Position) Activate(positionStep int) error {
	return nil
}

type PositionWithDuration struct {
	feature
	stepCount int
}

func (p PositionWithDuration) OutputType() OutputType {
	return PositionWithDirectionOutput
}

func (p PositionWithDuration) Activate(positionStep int, duration time.Time) error {
	return nil
}
