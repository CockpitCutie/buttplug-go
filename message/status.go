package message

import "fmt"

type Ok struct {
	message
}

type Error struct {
	message
	Message string    `json:"ErrorMessage"`
	Code    ErrorCode `json:"ErrorCode"`
}

func (e Error) Error() error {
	return fmt.Errorf("buttplug %s on message %d: %s", e.Code.String(), e.Id, e.Message)
}

type ErrorCode int

const (
	UnknownError ErrorCode = iota
	InitError
	PingError
	MsgError
	DeviceError
)

func (ec ErrorCode) String() string {
	switch ec {
	case UnknownError:
		return "UnknownError"
	case InitError:
		return "InitError"
	case PingError:
		return "PingError"
	case MsgError:
		return "MsgError"
	case DeviceError:
		return "DeviceError"
	default:
		return "Unrecognized Error Code"
	}
}

type Ping struct {
	message
}
