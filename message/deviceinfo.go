package message

import "encoding/json"

type DeviceList struct {
	message
	Devices []Device
}

type Device struct {
	DeviceName             string
	DeviceIndex            uint
	DeviceMessageTimingGap *uint
	DeviceDisplayName      *string
	DeviceMessages         map[string]DeviceAttrs
}

type DeviceAttrs struct {
	Attrs []Attributes
}

func (d *DeviceAttrs) UnmarshalJSON(b []byte) error {
	if string(b) == "{}" {
		return nil
	}
	return json.Unmarshal(b, &d.Attrs)
}

type Attributes struct {
	FeatureDescriptor *string
	StepCount         *uint
	ActuatorType      *string
	SensorType        *string
	SensorRange       [][2]int
	EndPoints         []string
}

type DeviceAdded struct {
	message
	DeviceName             string
	DeviceIndex            uint
	DeviceMessageTimingGap *uint
	DeviceDisplayName      *string
	DeviceMessages         map[string]DeviceAttrs
}

type DeviceRemoved struct {
	message
	DeviceIndex uint
}
