package message

type StartScanning struct {
	message
}

type StopScanning struct {
	message
}

type ScanningFinished struct {
	message
}

type RequestDevicelist struct {
	message
}

type DeviceList struct {
	message
	Devices []Device
}

type Device struct {
	DeviceName string
	DeviceIndex uint
	DeviceMessageTimingGap *uint
	DeviceDisplayName *string
	DeviceMessages map[string]any
}

type Attributes struct {
	FeatureDescriptor *string
	StepCount *uint
	ActuatorType *string
	SensorType *string
	SensorRange [][2]int
	EndPoints []string
}

type DeviceRemoved struct {
	message
	DeviceIndex uint
}