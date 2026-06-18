package device

import (
	"fmt"
	"time"

	"github.com/CockpitCutie/buttplug-go/message"
)

// Output represents a common interface for device outputs, like Vibrate, Rotate,
// Oscillate, etc. Each output has a specific type and a range of valid step values
// that can be used to determine bounds for output intensities.
//
// The Output interface does not provide a common method for activating outputs,
// as they are not interchangeable, and Activate signatures differ. Instead,
// each output type has its own Activate method, and Outputs must be type asserted
// to their concrete type to access it.
type Output interface {
	Feature
	OutputType() OutputType
	StepRange() [2]int
}

// OutputType represents the type of an output feature, such as Vibrate, Rotate, Oscillate,
// etc. This list of output types is taken from the ButtplugIO v4 specification.
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

// registerOutputs takes a list of device features from a DeviceFeatures message and adds
// any features with Output capabilities to the device's Outputs map. It returns an error
// if any of the output features have an unknown type or invalid properties.
func (d *Device) registerOutputs(features message.DeviceFeatures) error {
	for _, featureMsg := range features {
		feature := feature{
			description: featureMsg.FeatureDescription,
			index:       featureMsg.FeatureIndex,
			device:      d,
		}
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

// outputFromProps is a helper function that takes an output type, its properties from a DeviceOutput message,
// and the parent feature, and returns an instance of the corresponding Output type. It returns an error
// if the output type is unknown or if the properties are invalid for that output type.
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
		return nil, fmt.Errorf("unknown output type %s", kind)
	}
}

// activateOutput is a helper method on Device that takes an Output and sends a
// message to the device to activate the given output. This can be common because
// most of the logic for activating outputs is the same between kinds. activateOutput
// returns an error if the provided step is not in StepRange or 0 (stop) or if
// message sending or receiving fails.
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

// Vibrator represents a vibrating output feature of a device. It provides methods
// to activate the vibrator.
type Vibrator struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always VibrateOutput ("Vibrate").
func (v Vibrator) OutputType() OutputType {
	return VibrateOutput
}

// StepRange returns the range of valid step values that can be used to activate the vibrator.
func (v Vibrator) StepRange() [2]int {
	return v.stepRange
}

// Activate sends a message to the device to activate the vibrator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the vibrator off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (v Vibrator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(v.OutputType()): {Value: step},
	}
	return v.device.activateOutput(v, outputValue)
}

// Vibrate sends a message to the device to activate the vibrator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the vibrator off. Vibrate returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// Vibrate is a convenience method that calls Activate with the given speed step. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (v Vibrator) Vibrate(speedStep int) error {
	return v.Activate(speedStep)
}

// Rotator represents a rotating output feature of a device. It provides methods
// to activate the rotator.
type Rotator struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always RotateOutput ("Rotate").
func (r Rotator) OutputType() OutputType {
	return RotateOutput
}

// StepRange returns the range of valid step values that can be used to activate the rotator.
func (r Rotator) StepRange() [2]int {
	return r.stepRange
}

// Activate sends a message to the device to activate the rotator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the rotator off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (r Rotator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(r.OutputType()): {Value: step},
	}
	return r.device.activateOutput(r, outputValue)
}

// Rotate sends a message to the device to activate the rotator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the rotator off. Rotate returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// Rotate is a convenience method that calls Activate with the given speed step. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (r Rotator) Rotate(speedStep int) error {
	return r.Activate(speedStep)
}

// RotatorWithDirection represents a rotating output feature of a device that also
// has directional control. It provides methods to activate the rotator with a
// specified direction.
type RotatorWithDirection struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always RotationWithDirectionOutput
// ("RotationWithDirection").
func (r RotatorWithDirection) OutputType() OutputType {
	return RotationWithDirectionOutput
}

// StepRange returns the range of valid step values that can be used to activate the
// rotator with direction.
func (r RotatorWithDirection) StepRange() [2]int {
	return r.stepRange
}

// Activate sends a message to the device to activate the rotator with direction at the given
// step and direction. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the rotator off. The direction is determined by the clockwise parameter,
// where true represents clockwise rotation and false represents counterclockwise rotation.
// Activate returns an error if the step is out of range, or if the message sending or receiving fails.
func (r RotatorWithDirection) Activate(step int, clockwise bool) error {
	outputValue := message.OutputValue{
		string(r.OutputType()): {
			Value:     step,
			Clockwise: &clockwise,
		},
	}
	return r.device.activateOutput(r, outputValue)
}

// RotateDirection sends a message to the device to activate the rotator with direction at the given
// step and direction. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the rotator off. The direction is determined by the clockwise parameter,
// where true represents clockwise rotation and false represents counterclockwise rotation.
// RotateDirection returns an error if the step is out of range, or if the message sending or receiving fails.
//
// RotateDirection is a convenience method that calls Activate with the given speed step and direction. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (r RotatorWithDirection) RotateDirection(speedStep int, clockwise bool) error {
	return r.Activate(speedStep, clockwise)
}

// Oscillator represents an oscillating output feature of a device. It provides methods
// to activate the oscillator.
type Oscillator struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always OscillateOutput ("Oscillate").
func (o Oscillator) OutputType() OutputType {
	return OscillateOutput
}

// StepRange returns the range of valid step values that can be used to activate the oscillator.
func (o Oscillator) StepRange() [2]int {
	return o.stepRange
}

// Activate sends a message to the device to activate the oscillator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the oscillator off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (o Oscillator) Activate(step int) error {
	outputValue := message.OutputValue{
		string(o.OutputType()): {Value: step},
	}
	return o.device.activateOutput(o, outputValue)
}

// Oscillate sends a message to the device to activate the oscillator at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the oscillator off. Oscillate returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// Oscillate is a convenience method that calls Activate with the given speed step. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (o Oscillator) Oscillate(speedStep int) error {
	return o.Activate(speedStep)
}

// Constrictor represents a constricting output feature of a device. It provides methods
// to activate the constrictor.
type Constrictor struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always ConstrictOutput ("Constrict").
func (c Constrictor) OutputType() OutputType {
	return ConstrictOutput
}

// StepRange returns the range of valid step values that can be used to activate the constrictor.
func (c Constrictor) StepRange() [2]int {
	return c.stepRange
}

// Activate sends a message to the device to activate the constrictor at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the constrictor off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (c Constrictor) Activate(step int) error {
	outputValue := message.OutputValue{
		string(c.OutputType()): {Value: step},
	}
	return c.device.activateOutput(c, outputValue)
}

// Constrict sends a message to the device to activate the constrictor at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the constrictor off. Constrict returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// Constrict is a convenience method that calls Activate with the given constriction
// step. It is provided for readability and to match the common terminology for this
// type of output, but it does not add any new functionality.
func (c Constrictor) Constrict(constriction int) error {
	return c.Activate(constriction)
}

// Heater represents a heating output feature of a device. It provides methods
// to activate the heater.
type Heater struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always HeaterOutput ("Heater").
func (h Heater) OutputType() OutputType {
	return HeaterOutput
}

// StepRange returns the range of valid step values that can be used to activate the heater.
func (h Heater) StepRange() [2]int {
	return h.stepRange
}

// Activate sends a message to the device to activate the heater at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the heater off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (h Heater) Activate(step int) error {
	outputValue := message.OutputValue{
		string(h.OutputType()): {Value: step},
	}
	return h.device.activateOutput(h, outputValue)
}

// Heat sends a message to the device to activate the heater at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the heater off. Heat returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// Heat is a convenience method that calls Activate with the given heat step. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (h Heater) Heat(heatStep int) error {
	return h.Activate(heatStep)
}

// LED represents an LED output feature of a device. It provides methods to activate the LED.
type LED struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always LEDOutput ("Led").
func (l LED) OutputType() OutputType {
	return LEDOutput
}

// StepRange returns the range of valid step values that can be used to activate the LED.
func (l LED) StepRange() [2]int {
	return l.stepRange
}

// Activate sends a message to the device to activate the LED at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the LED off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (l LED) Activate(step int) error {
	outputValue := message.OutputValue{
		string(l.OutputType()): {Value: step},
	}
	return l.device.activateOutput(l, outputValue)
}

// SetBrightness sends a message to the device to activate the LED at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the LED off. SetBrightness returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// SetBrightness is a convenience method that calls Activate with the given
// brightness step. It is provided for readability and to match the common
// terminology for this type of output, but it does not add any new functionality.
func (l LED) SetBrightness(brightness int) error {
	return l.Activate(brightness)
}

// Position represents a position output feature of a device. It provides methods
// to activate the position.
type Position struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always PositionOutput ("Position").
func (p Position) OutputType() OutputType {
	return PositionOutput
}

// StepRange returns the range of valid step values that can be used to activate the position output.
func (p Position) StepRange() [2]int {
	return p.stepRange
}

// Activate sends a message to the device to activate the position output at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the position output off. Activate returns an error if the
// step is out of range, or if the message sending or receiving fails.
func (p Position) Activate(step int) error {
	outputValue := message.OutputValue{
		string(p.OutputType()): {Value: step},
	}
	return p.device.activateOutput(p, outputValue)
}

// MoveTo sends a message to the device to activate the position output at the given
// step. The step must be between the values returned by StepRange, inclusive,
// or it can be 0 to turn the position output off. MoveTo returns an error if the
// step is out of range, or if the message sending or receiving fails.
//
// MoveTo is a convenience method that calls Activate with the given position step.
// It is provided for readability and to match the common terminology for this type
// of output, but it does not add any new functionality.
func (p Position) MoveTo(position int) error {
	return p.Activate(position)
}

// PositionWithDuration represents a position output feature of a device that also has
// duration control. It provides methods to activate the position with a specified duration.
type PositionWithDuration struct {
	feature
	stepRange [2]int
}

// OutputType returns the type of the output, which is always PositionWithDurationOutput
// ("PositionWithDuration").
func (p PositionWithDuration) OutputType() OutputType {
	return PositionWithDurationOutput
}

// StepRange returns the range of valid step values that can be used to activate the
// position with duration output.
func (p PositionWithDuration) StepRange() [2]int {
	return p.stepRange
}

// Activate sends a message to the device to activate the position output at the given
// step for the specified duration. The step must be between the values returned by
// StepRange, inclusive, or it can be 0 to turn the position output off. The duration
// is specified as a time.Duration, and is converted to milliseconds for the message.
// Activate returns an error if the step is out of range, or if the message sending
// or receiving fails.
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

// MoveToFor sends a message to the device to activate the position output at the given
// step for the specified duration. The step must be between the values returned by
// StepRange, inclusive, or it can be 0 to turn the position output off. The duration
// is specified as a time.Duration, and is converted to milliseconds for the message.
// MoveToFor returns an error if the step is out of range, or if the message sending
// or receiving fails.
//
// MoveToFor is a convenience method that calls Activate with the given position step and duration. It is provided
// for readability and to match the common terminology for this type of output, but it does not
// add any new functionality.
func (p PositionWithDuration) MoveToFor(position int, duration time.Duration) error {
	return p.Activate(position, duration)
}
