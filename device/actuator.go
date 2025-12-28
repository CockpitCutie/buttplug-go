package device

import "github.com/CockpitCutie/buttplug-go/message"

type Actuator struct {
	Index         uint
	Descriptor    string
	Type          ActuatorType
	StepCount     uint
	Command       Command
	DeviceIndex   uint
	messageSender MessageSender
}

func (a Actuator) Vibrate(intensity float64) error {
	a.messageSender.SendRecv(&message.ScalarCmd{
		DeviceIndex: a.DeviceIndex,
		Scalars: []message.Scalar{{
			Index:        a.Index,
			Scalar:       intensity,
			ActuatorType: string(VibrateActuator),
		}},
	})
	return nil
}

type ActuatorType string

const (
	UnknownActuator   ActuatorType = "Unknown"
	VibrateActuator   ActuatorType = "Vibrate"
	RotateActuator    ActuatorType = "Rotate"
	OscillateActuator ActuatorType = "Oscillate"
	ConstrictActuator ActuatorType = "Constrict"
	InflateActuator   ActuatorType = "Inflate"
	PositionActuator  ActuatorType = "Position"
)
