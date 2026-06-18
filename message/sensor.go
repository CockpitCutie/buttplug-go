package message

type SensorReadCmd struct {
	message
	DeviceIndex int
	SensorIndex int
	SensorType  string
}

type SensorReading struct {
	message
	DeviceIndex int
	SensorIndex int
	SensorType  string
	Data        []int
}

type SensorSubscribeCmd struct {
	message
	DeviceIndex int
	SensorIndex int
	SensorType  string
}

type SensorUnsubscribeCmd struct {
	message
	DeviceIndex int
	SensorIndex int
	SensorType  string
}
