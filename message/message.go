package message

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Message interface {
	ID() uint32
	SetID(id uint32)
	IsServerEvent() bool
}

type message struct {
	Id uint32
}

func (m message) ID() uint32 {
	return m.Id
}

func (m *message) SetID(id uint32) {
	m.Id = id
}

func (m message) IsServerEvent() bool {
	return m.Id == 0
}

func Serialize(m Message) (string, error) {
	// get name of specific message type for json key
	messageKind := reflect.TypeOf(m).Elem().Name()
	serialized, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`[{"%s":%s}]`, messageKind, string(serialized)), nil
}

func Deserialize(b []byte) (Message, error) {
	println(string(b))
	msgs := []map[string]json.RawMessage{}
	err := json.Unmarshal(b, &msgs)
	if err != nil || len(msgs) == 0 {
		return nil, err
	}
	msg := msgs[0]
	for key, value := range msg {
		switch key {
		case "Ok":
			var ok Ok
			err = json.Unmarshal(value, &ok)
			return &ok, err
		case "Error":
			var errMsg Error
			err = json.Unmarshal(value, &errMsg)
			return &errMsg, err
		case "Ping":
			var ping Ping
			err = json.Unmarshal(value, &ping)
			return &ping, err
		case "RequestServerInfo":
			var reqInfo RequestServerInfo
			err = json.Unmarshal(value, &reqInfo)
			return &reqInfo, err
		case "ServerInfo":
			var serverInfo ServerInfo
			err = json.Unmarshal(value, &serverInfo)
			return &serverInfo, err
		case "StartScanning":
			var startScan StartScanning
			err = json.Unmarshal(value, &startScan)
			return &startScan, err
		case "StopScanning":
			var stopScan StopScanning
			err = json.Unmarshal(value, &stopScan)
			return &stopScan, err
		case "ScanningFinished":
			var scanFinished ScanningFinished
			err = json.Unmarshal(value, &scanFinished)
			return &scanFinished, err
		case "RequestDeviceList":
			var reqDevList RequestDeviceList
			err = json.Unmarshal(value, &reqDevList)
			return &reqDevList, err
		case "DeviceList":
			var devList DeviceList
			err = json.Unmarshal(value, &devList)
			return &devList, err
		case "StopAllDevices":
			var stopAll StopAllDevices
			err = json.Unmarshal(value, &stopAll)
			return &stopAll, err
		case "DeviceAdded":
			var devAdded DeviceAdded
			err = json.Unmarshal(value, &devAdded)
			return &devAdded, err
		case "DeviceRemoved":
			var devRemoved DeviceRemoved
			err = json.Unmarshal(value, &devRemoved)
			return &devRemoved, err
		default:
			return nil, fmt.Errorf("unknown message type: %s", key)
		}
	}
	return nil, fmt.Errorf("no message found")
}
