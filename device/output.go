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
	stepRange [2]int
}

func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

func (r Rotator) StepRange() [2]int {
	return r.stepRange
}

func (r Rotator) Activate(speedStep int) error {
	stepIsInRange := speedStep == 0 || (speedStep > r.stepRange[0] && speedStep < r.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("speed step %d is out of range for this vibrator", speedStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  r.Device().Index,
		FeatureIndex: r.Index(),
		Command: message.OutputValue{
			"Vibrate": {Value: speedStep},
		},
	}
	_, err := r.device.msgSender.SendRecv(&msg)
	return err
}

func (r Rotator) Rotate(speedStep int) error {
	return r.Rotate(speedStep)
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

func (r RotatorWithDirection) Activate(speedStep int, clockwise bool) error {
	stepIsInRange := speedStep == 0 || (speedStep > r.stepRange[0] && speedStep < r.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("speed step %d is out of range for this vibrator", speedStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  r.Device().Index,
		FeatureIndex: r.Index(),
		Command: message.OutputValue{
			"RotationWithDirection": {
				Value:     speedStep,
				Clockwise: &clockwise,
			},
		},
	}
	_, err := r.device.msgSender.SendRecv(&msg)
	return err
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

func (o Oscillator) Activate(speedStep int) error {
	stepIsInRange := speedStep == 0 || (speedStep > o.stepRange[0] && speedStep < o.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("speed step %d is out of range for this vibrator", speedStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  o.Device().Index,
		FeatureIndex: o.Index(),
		Command: message.OutputValue{
			"Oscillate": {Value: speedStep},
		},
	}
	_, err := o.device.msgSender.SendRecv(&msg)
	return err
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
	stepIsInRange := step == 0 || (step > c.stepRange[0] && step < c.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("step %d is out of range for this constrictor", step)
	}
	msg := message.OutputCmd{
		DeviceIndex:  c.Device().Index,
		FeatureIndex: c.Index(),
		Command: message.OutputValue{
			"Constrict": {Value: step},
		},
	}
	_, err := c.device.msgSender.SendRecv(&msg)
	return err
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

func (h Heater) Activate(heatLevel int) error {
	stepIsInRange := heatLevel == 0 || (heatLevel > h.stepRange[0] && heatLevel < h.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("heat level %d is out of range for this heater", heatLevel)
	}
	msg := message.OutputCmd{
		DeviceIndex:  h.Device().Index,
		FeatureIndex: h.Index(),
		Command: message.OutputValue{
			"Heat": {Value: heatLevel},
		},
	}
	_, err := h.device.msgSender.SendRecv(&msg)
	return err
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

func (l LED) Activate(brightnessStep int) error {
	stepIsInRange := brightnessStep == 0 || (brightnessStep > l.stepRange[0] && brightnessStep < l.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("brightness step %d is out of range for this LED", brightnessStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  l.Device().Index,
		FeatureIndex: l.Index(),
		Command: message.OutputValue{
			"LED": {Value: brightnessStep},
		},
	}
	_, err := l.device.msgSender.SendRecv(&msg)
	return err
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

func (p Position) Activate(positionStep int) error {
	stepIsInRange := positionStep == 0 || (positionStep > p.stepRange[0] && positionStep < p.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("position step %d is out of range for this position output", positionStep)
	}
	msg := message.OutputCmd{
		DeviceIndex:  p.Device().Index,
		FeatureIndex: p.Index(),
		Command: message.OutputValue{
			"Position": {Value: positionStep},
		},
	}
	_, err := p.device.msgSender.SendRecv(&msg)
	return err
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

func (p PositionWithDuration) Activate(positionStep int, duration time.Duration) error {
	stepIsInRange := positionStep == 0 || (positionStep > p.stepRange[0] && positionStep < p.stepRange[1])
	if !stepIsInRange {
		return fmt.Errorf("position step %d is out of range for this position output", positionStep)
	}
	durationMillis := int(duration.Milliseconds())
	msg := message.OutputCmd{
		DeviceIndex:  p.Device().Index,
		FeatureIndex: p.Index(),
		Command: message.OutputValue{
			"PositionWithDuration": {
				Value:    positionStep,
				Duration: &durationMillis,
			},
		},
	}
	_, err := p.device.msgSender.SendRecv(&msg)
	return err
}
