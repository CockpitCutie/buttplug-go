package device

type Sensor struct {
	Index      uint
	Descriptor string
	Type       SensorType
	Range      [][2]int
	Command    Command
}

type SensorType string

const (
	UnknownSensor  SensorType = "Unknown"
	BatterySensor  SensorType = "Battery"
	RSSISensor     SensorType = "RSSI"
	ButtonSensor   SensorType = "Button"
	PressureSensor SensorType = "Pressure"
)
