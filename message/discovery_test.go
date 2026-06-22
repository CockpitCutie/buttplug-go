package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const StartScanningDocExample = `[{"StartScanning":{"Id":1}}]`

func TestSerialize_StartScanning_DocExample(t *testing.T) {
	msg := StartScanning{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := StartScanningDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_StartScanning_DocExample(t *testing.T) {
	jsonMessage := StartScanningDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StartScanning); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StartScanning")
	}
}

const StopScanningDocExample = `[{"StopScanning":{"Id":1}}]`

func TestSerialize_StopScanning_DocExample(t *testing.T) {
	msg := StopScanning{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := StopScanningDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_StopScanning_DocExample(t *testing.T) {
	jsonMessage := StopScanningDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopScanning); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StopScanning")
	}
}

const ScanningFinishedDocExample = `[{"ScanningFinished":{"Id":0}}]`

func TestSerialize_ScanningFinished_DocExample(t *testing.T) {
	msg := ScanningFinished{
		message: message{
			Id: 0,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := ScanningFinishedDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_ScanningFinished_DocExample(t *testing.T) {
	jsonMessage := ScanningFinishedDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*ScanningFinished); ok {
		assert.Equalf(t, 0, msg.ID(), "Expected Id 0 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type ScanningFinished")
	}
}

const RequestDeviceListDocExample = `[{"RequestDeviceList":{"Id":1}}]`

func TestSerialize_RequestDeviceList_DocExample(t *testing.T) {
	msg := RequestDeviceList{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := RequestDeviceListDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_RequestDeviceList_DocExample(t *testing.T) {
	jsonMessage := RequestDeviceListDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*RequestDeviceList); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type RequestDeviceList")
	}
}
