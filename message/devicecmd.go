package message

type StopDeviceCmd struct {
	message
	DeviceIndex int
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
	DeviceIndex  int
	FeatureIndex int
	Command      OutputValue
}

type OutputValue map[string]struct {
	Value     int
	Clockwise *bool
	Duration  *int
}

type InputCmd struct {
	message
	DeviceIndex  int
	FeatureIndex int
	Type         string
	Command      string
}

type InputReading struct {
	message
	DeviceIndex  int
	FeatureIndex int
	Reading      InputData
}

type InputData struct {
	Battery *struct {
		Value int
	}
	RSSI *struct {
		Value int
	}
	Pressure *struct {
		Value int
	}
	Button *struct {
		Value int
	}
}
