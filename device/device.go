package device

import "github.com/CockpitCutie/buttplug-go/message"

type Device struct {
	Name             string
	Index            uint
	MessageTimingGap *uint
	DisplayName      *string
	actuators        []Actuator
	sensors          []Sensor
}

func FromMessage(msg message.Device) *Device {
	device := &Device{
		Name:             msg.DeviceName,
		Index:            msg.DeviceIndex,
		MessageTimingGap: msg.DeviceMessageTimingGap,
		DisplayName:      msg.DeviceDisplayName,
	}
	for cmdLabel, attrs := range msg.DeviceMessages {
		cmd := Command(cmdLabel)
		isActuator := cmd == ScalarCmd || cmd == LinearCmd || cmd == RotateCmd
		if isActuator {
			device.addActuators(cmd, attrs)
		} else {
			device.addSensors(cmd, attrs)
		}
	}
	return device
}

func (d *Device) addActuators(cmd Command, attrs []message.Attributes) {
	for i, attrs := range attrs {
		d.actuators = append(d.actuators, Actuator{
			Index:      uint(i),
			Descriptor: *attrs.FeatureDescriptor,
			Type:       ActuatorType(*attrs.ActuatorType),
			StepCount:  *attrs.StepCount,
			Command:    cmd,
		})
	}
}

func (d *Device) addSensors(cmd Command, attrs []message.Attributes) {
	for i, attrs := range attrs {
		d.sensors = append(d.sensors, Sensor{
			Index:      uint(i),
			Descriptor: *attrs.FeatureDescriptor,
			Type:       SensorType(*attrs.SensorType),
			Range:      attrs.SensorRange,
			Command:    cmd,
		})
	}
}

func (d *Device) Actuators() []Actuator {
	return nil
}

func (d *Device) Sensors() []Sensor {
	return nil
}

func (d *Device) Vibrate(intensity float64) error {
	return nil
}

func (d *Device) Rotate(intensity float64) error {
	return nil
}

func (d *Device) Oscillate(intensity float64) error {
	return nil
}

func (d *Device) Constrict(intensity float64) error {
	return nil
}

func (d *Device) Inflate(intensity float64) error {
	return nil
}

func (d *Device) Position(intensity float64) error {
	return nil
}

type Command string

const (
	ScalarCmd          Command = "ScalarCmd"
	RotateCmd          Command = "RotateCmd"
	LinearCmd          Command = "LinearCmd"
	SensorReadCmd      Command = "SensorReadCmd"
	SensorSubscribeCmd Command = "SensorSubscribeCmd"
)
