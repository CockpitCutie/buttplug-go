package message

type StopDeviceCmd struct {
	message
	DeviceIndex uint
}

type StopAllDevices struct {
	message
}

type ScalarCmd struct {
	message
	DeviceIndex uint
	Scalars []Scalar
}

type Scalar struct {
	Index uint
	Scalar float64
	ActuatorType string
}

type LinearCmd struct {
	message
	DeviceIndex uint
	Vectors []Vector
}

type Vector struct {
	Index uint
	Duration uint
	Position float64
}

type RotateCmd struct {
	message
	DeviceIndex uint
	Rotations []Rotation
}

type Rotation struct {
	Index uint
	Speed float64
	Clockwise bool
}