package device

import (
	"github.com/CockpitCutie/buttplug-go/message"
)

type Device struct {
	Name             string
	Index            uint
	MessageTimingGap *uint
	DisplayName      *string
	actuators        []Actuator
	sensors          []Sensor
	messageSender    MessageSender
}

func FromMessage(msg message.Device, sender MessageSender) Device {
	device := Device{
		Name:             msg.DeviceName,
		Index:            msg.DeviceIndex,
		MessageTimingGap: msg.DeviceMessageTimingGap,
		DisplayName:      msg.DeviceDisplayName,
		messageSender:    sender,
	}
	for cmdLabel, attrs := range msg.DeviceMessages {
		cmd := Command(cmdLabel)
		isActuator := cmd == ScalarCmd || cmd == LinearCmd || cmd == RotateCmd
		isSensor := cmd == SensorReadCmd || cmd == SensorSubscribeCmd
		if isActuator {
			device.addActuators(cmd, attrs.Attrs)
		} else if isSensor {
			device.addSensors(cmd, attrs.Attrs)
		}
	}
	return device
}

func (d *Device) addActuators(cmd Command, attrs []message.Attributes) {
	for i, attrs := range attrs {
		d.actuators = append(d.actuators, Actuator{
			Index:         uint(i),
			Descriptor:    *attrs.FeatureDescriptor,
			Type:          ActuatorType(*attrs.ActuatorType),
			StepCount:     *attrs.StepCount,
			Command:       cmd,
			DeviceIndex:   d.Index,
			messageSender: d.messageSender,
		})
	}
}

func (d *Device) addSensors(cmd Command, attrs []message.Attributes) {
	for i, attrs := range attrs {
		d.sensors = append(d.sensors, Sensor{
			Index:         uint(i),
			Descriptor:    *attrs.FeatureDescriptor,
			Type:          SensorType(*attrs.SensorType),
			Range:         attrs.SensorRange,
			Command:       cmd,
			messageSender: d.messageSender,
		})
	}
}

func (d *Device) Actuators() []Actuator {
	return d.actuators
}

func (d *Device) Sensors() []Sensor {
	return d.sensors
}

func (d *Device) Vibrate(intensity float64) error {
	var scalars []message.Scalar
	for _, actuator := range d.actuators {
		if actuator.Type == VibrateActuator {
			actuator.Vibrate(intensity)
		}
	}
	_, err := d.messageSender.SendRecv(&message.ScalarCmd{
		DeviceIndex: d.Index,
		Scalars:     scalars,
	})
	return err
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
	StopDeviceCmd      Command = ""
)

type MessageSender interface {
	SendRecv(message.Message) (message.Message, error)
}
