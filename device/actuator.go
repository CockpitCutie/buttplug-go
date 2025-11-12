package device

type Actuator struct {
	Index     uint
	Descriptor      string
	Type      ActuatorType
	StepCount uint
	Command   Command
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
