package message

type Ok struct {
	message
}

type Error struct {
	message
	Message string
	Code ErrorCode
}

type ErrorCode int

const (
	UnknownError ErrorCode = iota
	InitError
	PingError
	MsgError
	DeviceError
)

type Ping struct {
	message
}