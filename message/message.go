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

func Deserialize(b []byte) ([]Message, error) {
	rawMsgs := []map[string]json.RawMessage{}
	err := json.Unmarshal(b, &rawMsgs)
	if err != nil || len(rawMsgs) == 0 {
		return nil, err
	}
	var msgs []Message
	for _, msg := range rawMsgs {
		for key, value := range msg {
			switch key {
			case "Ok":
				var ok Ok
				err = json.Unmarshal(value, &ok)
				msgs = append(msgs, &ok)
			case "Error":
				var errMsg Error
				err = json.Unmarshal(value, &errMsg)
				msgs = append(msgs, &errMsg)
			case "Ping":
				var ping Ping
				err = json.Unmarshal(value, &ping)
				msgs = append(msgs, &ping)
			case "RequestServerInfo":
				var reqInfo RequestServerInfo
				err = json.Unmarshal(value, &reqInfo)
				msgs = append(msgs, &reqInfo)
			case "ServerInfo":
				var serverInfo ServerInfo
				err = json.Unmarshal(value, &serverInfo)
				msgs = append(msgs, &serverInfo)
			case "StartScanning":
				var startScan StartScanning
				err = json.Unmarshal(value, &startScan)
				msgs = append(msgs, &startScan)
			case "StopScanning":
				var stopScan StopScanning
				err = json.Unmarshal(value, &stopScan)
				msgs = append(msgs, &stopScan)
			case "ScanningFinished":
				var scanFinished ScanningFinished
				err = json.Unmarshal(value, &scanFinished)
				msgs = append(msgs, &scanFinished)
			case "RequestDeviceList":
				var reqDevList RequestDeviceList
				err = json.Unmarshal(value, &reqDevList)
				msgs = append(msgs, &reqDevList)
			case "DeviceList":
				var devList DeviceList
				err = json.Unmarshal(value, &devList)
				msgs = append(msgs, &devList)
			case "StopDeviceCmd":
				var stopDev StopDeviceCmd
				err = json.Unmarshal(value, &stopDev)
				msgs = append(msgs, &stopDev)
			case "StopAllDevices":
				var stopAll StopAllDevices
				err = json.Unmarshal(value, &stopAll)
				msgs = append(msgs, &stopAll)
			case "OutputCmd":
				var outputCmd OutputCmd
				err = json.Unmarshal(value, &outputCmd)
				msgs = append(msgs, &outputCmd)
			default:
				return nil, fmt.Errorf("unknown message type: %s", key)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return msgs, nil
}
