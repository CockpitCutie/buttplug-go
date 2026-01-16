package message

type StopDeviceCmd struct {
	message
	DeviceIndex uint
	Inputs      *bool
	Outputs     *bool
}

type StopAllDevices struct {
	message
	Inputs  *bool
	Outputs *bool
}

type OutputCmd struct {
	message
	DeviceIndex  uint
	FeatureIndex uint
	Command      OutputValue
}

type OutputValue map[string]struct {
	Value     uint32
	Clockwise *bool
	Duration  *uint32
}

type InputCmd struct {
	message
	DeviceIndex  uint
	FeatureIndex uint
	Type         string
	Command      string
}

type InputReading struct {
	message
	DeviceIndex  uint
	FeatureIndex uint
	Reading      InputData
}

type InputData struct {
	Battery *struct {
		Value uint8
	}
	RSSI *struct {
		Value int8
	}
	Pressure *struct {
		Value uint32
	}
	Button *struct {
		Value uint8
	}
}
