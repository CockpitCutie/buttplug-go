package message

type DeviceList struct {
	message
	Devices map[string]Device
}

type Device struct {
	DeviceName             string
	DeviceIndex            uint
	DeviceMessageTimingGap uint   `json:",omitempty"`
	DeviceDisplayName      string `json:",omitempty"`
	Features               DeviceFeatures
}

type DeviceFeatures map[string]DeviceFeature

type DeviceFeature struct {
	FeatureDescription string
	FeatureIndex       uint32
	Output             DeviceOutput
	Input              DeviceInput
}

type DeviceOutput map[string]struct {
	Value    *[2]int
	Position *[2]uint
	Duration *[2]uint
}

type DeviceInput map[string]struct {
	Value   *[2]int
	Command []string
}
