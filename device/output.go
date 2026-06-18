package device

import (
	"fmt"
	"time"

	"github.com/CockpitCutie/buttplug-go/message"
)

type Output interface {
	Feature
	OutputType() OutputType
	StepRange() [2]int
}

type OutputType string

const (
	VibrateOutput               OutputType = "Vibrate"
	RotateOutput                OutputType = "Rotate"
	RotationWithDirectionOutput OutputType = "RotationWithDirection"
	OscillateOutput             OutputType = "Oscillate"
	ConstrictOutput             OutputType = "Constrict"
	HeaterOutput                OutputType = "Heater"
	LEDOutput                   OutputType = "Led"
	PositionOutput              OutputType = "Position"
	PositionWithDurationOutput  OutputType = "PositionWithDuration"
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

func (d *Device) activateOutput(output Output, value message.OutputValue) error {
	outputValue, ok := value[string(output.OutputType())]
	if !ok {
		return fmt.Errorf("output value for type %s not found", output.OutputType())
	}
	step := outputValue.Value
	stepIsInRange := step == 0 || (step >= output.StepRange()[0] && step <= output.StepRange()[1])
	if !stepIsInRange {
		return fmt.Errorf("step %d is out of range for this %s", step, output.OutputType())
	}
	msg := message.OutputCmd{
		DeviceIndex:  d.Index,
		FeatureIndex: output.Index(),
		Command:      value,
	}
	_, err := d.msgSender.SendRecv(&msg)
	return err
}

func outputFromProps(kind OutputType, properties message.DeviceOutput, feature feature) (Output, error) {
	switch kind {
	case VibrateOutput:
		return Vibrator{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case RotateOutput:
		return Rotator{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case RotationWithDirectionOutput:
		return RotatorWithDirection{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case OscillateOutput:
		return Oscillator{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case ConstrictOutput:
		return Constrictor{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case HeaterOutput:
		return Heater{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case LEDOutput:
		return LED{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case PositionOutput:
		return Position{
			feature:   feature,
			stepRange: properties.Value,
		}, nil
	case PositionWithDurationOutput:
		return PositionWithDuration{
			feature:   feature,
			stepRange: properties.Value,
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

func (v Vibrator) StepRange() [2]int {
	return v.stepRange
}

func (v Vibrator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(v.OutputType()): {Value: step},
	}
	return v.device.activateOutput(v, outputValue)
}

func (v Vibrator) Vibrate(speedStep int) error {
	return v.Activate(speedStep)
}

type Rotator struct {
	feature
	stepRange [2]int
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

func (r Rotator) StepRange() [2]int {
	return r.stepRange
}

func (r Rotator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(r.OutputType()): {Value: step},
	}
	return r.device.activateOutput(r, outputValue)
}

func (r Rotator) Rotate(speedStep int) error {
	return r.Activate(speedStep)
}

type RotatorWithDirection struct {
	feature
	stepRange [2]int
}

func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

func (r RotatorWithDirection) StepRange() [2]int {
	return r.stepRange
}

func (r RotatorWithDirection) Activate(step int, clockwise bool) error {
	outputValue := message.OutputValue{
		string(r.OutputType()): {
			Value:     step,
			Clockwise: &clockwise,
		},
	}
	return r.device.activateOutput(r, outputValue)
}
func (r RotatorWithDirection) RotateDirection(speedStep int, clockwise bool) error {
	return r.Activate(speedStep, clockwise)
}

type Oscillator struct {
	feature
	stepRange [2]int
}

func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

func (o Oscillator) StepRange() [2]int {
	return o.stepRange
}

func (o Oscillator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(o.OutputType()): {Value: step},
	}
	return o.device.activateOutput(o, outputValue)
}

type Constrictor struct {
	feature
	stepRange [2]int
}

func (c Constrictor) OutputType() OutputType {
	return ConstrictOutput
}

func (c Constrictor) StepRange() [2]int {
	return c.stepRange
}

func (c Constrictor) Activate(step int) error {
	outputValue := message.OutputValue{
		string(c.OutputType()): {Value: step},
	}
	return c.device.activateOutput(c, outputValue)
}

type Heater struct {
	feature
	stepRange [2]int
}

func (h Heater) OutputType() OutputType {
	return HeaterOutput
}

func (h Heater) StepRange() [2]int {
	return h.stepRange
}

func (h Heater) Activate(step int) error {
	outputValue := message.OutputValue{
		string(h.OutputType()): {Value: step},
	}
	return h.device.activateOutput(h, outputValue)
}

type LED struct {
	feature
	stepRange [2]int
}

func (l LED) OutputType() OutputType {
	return LEDOutput
}

func (l LED) StepRange() [2]int {
	return l.stepRange
}

func (l LED) Activate(step int) error {
	outputValue := message.OutputValue{
		string(l.OutputType()): {Value: step},
	}
	return l.device.activateOutput(l, outputValue)
}

type Position struct {
	feature
	stepRange [2]int
}

func (p Position) OutputType() OutputType {
	return PositionOutput
}

func (p Position) StepRange() [2]int {
	return p.stepRange
}

func (p Position) Activate(step int) error {
	outputValue := message.OutputValue{
		string(p.OutputType()): {Value: step},
	}
	return p.device.activateOutput(p, outputValue)
}

type PositionWithDuration struct {
	feature
	stepRange [2]int
}

func (p PositionWithDuration) OutputType() OutputType {
	return PositionWithDurationOutput
}

func (p PositionWithDuration) StepRange() [2]int {
	return p.stepRange
}

func (p PositionWithDuration) Activate(step int, duration time.Duration) error {
	durationMillis := int(duration.Milliseconds())
	outputValue := message.OutputValue{
		string(p.OutputType()): {
			Value:    step,
			Duration: &durationMillis,
		},
	}
	return p.device.activateOutput(p, outputValue)
}
