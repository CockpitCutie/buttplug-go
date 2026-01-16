package message

type DeviceList struct {
	message
	Devices map[string]Device
}

type Device struct {
	DeviceName             string
	DeviceIndex            uint
	DeviceMessageTimingGap *uint
	DeviceDisplayName      *string
	Features         map[string]DeviceFeatures
}

type DeviceFeatures struct {
	FeatureDescription string
	FeatureIndex       uint32
	Output             map[string]DeviceOutput
	Input              map[string]DeviceInput
}

type DeviceOutput struct {
	Value    *[2]int
	Position *[2]uint
	Duration *[2]uint
}

type DeviceInput struct {
	Value   *[2]int
	Command []string
}