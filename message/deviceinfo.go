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
	Output             map[string]DeviceOutput
	Input              map[string]DeviceInput
}

type DeviceOutput struct {
	Value    []int
	Position []uint
	Duration []uint
}

type DeviceInput struct {
	Value   [2]int32
	Command []string
}
