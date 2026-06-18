package message

type DeviceList struct {
	message
	Devices map[string]Device
}

type Device struct {
	DeviceName             string
	DeviceIndex            int
	DeviceMessageTimingGap int    `json:",omitempty"`
	DeviceDisplayName      string `json:",omitempty"`
	DeviceFeatures         DeviceFeatures
}

type DeviceFeatures map[string]DeviceFeature

type DeviceFeature struct {
	FeatureDescription string
	FeatureIndex       int
	Output             map[string]DeviceOutput
	Input              map[string]DeviceInput
}

type DeviceOutput struct {
	Value    [2]int
	Position []int
	Duration []int
}

type DeviceInput struct {
	Value   [2]int
	Command []string
}
