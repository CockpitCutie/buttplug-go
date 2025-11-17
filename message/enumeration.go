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

type RequestDeviceList struct {
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
	DeviceMessages map[string][]Attributes
}

type Attributes struct {
	FeatureDescriptor *string
	StepCount *uint
	ActuatorType *string
	SensorType *string
	SensorRange [][2]int
	EndPoints []string
}

type DeviceAdded struct {
	message
	DeviceName string
	DeviceIndex uint
	DeviceMessageTimingGap *uint
	DeviceDisplayName *string
	DeviceMessages map[string]Attributes
}

type DeviceRemoved struct {
	message
	DeviceIndex uint
}