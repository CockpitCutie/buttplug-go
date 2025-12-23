package message

type SensorReadCmd struct {
	message
	DeviceIndex uint
	SensorIndex uint
	SensorType  string
}

type SensorReading struct {
	message
	DeviceIndex uint
	SensorIndex uint
	SensorType  string
	Data        []int
}

type SensorSubscribeCmd struct {
	message
	DeviceIndex uint
	SensorIndex uint
	SensorType  string
}

type SensorUnsubscribeCmd struct {
	message
	DeviceIndex uint
	SensorIndex uint
	SensorType  string
}
