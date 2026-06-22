package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const OkDocExample = `[{"Ok":{"Id":1}}]`

func TestSerialize_Ok_DocExample(t *testing.T) {
	msg := Ok{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := OkDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_Ok_DocExample(t *testing.T) {
	jsonMessage := OkDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Ok); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ok")
	}
}

const ErrorDocExample = `[{"Error":{"Id":0,"ErrorMessage":"Server received invalid JSON.","ErrorCode":3}}]`

func TestSerialize_Error_DocExample(t *testing.T) {
	msg := Error{
		message: message{
			Id: 0,
		},
		Message: "Server received invalid JSON.",
		Code:    MsgError,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := ErrorDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_Error_DocExample(t *testing.T) {
	jsonMessage := ErrorDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Error); ok {
		assert.Equalf(t, 0, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Server received invalid JSON.", msg.Message, "Expected ErrorMessage 'Server received invalid JSON.' found '%s'", msg.Message)
		assert.Equalf(t, MsgError, msg.Code, "Expected ErrorCode MsgError found %d", msg.Code)
	} else {
		t.Errorf("Deserialized message is not of type Error")
	}
}

const PingDocExample = `[{"Ping":{"Id":5}}]`

func TestSerialize_Ping_DocExample(t *testing.T) {
	msg := Ping{
		message: message{
			Id: 5,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := PingDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_Ping_DocExample(t *testing.T) {
	jsonMessage := PingDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Ping); ok {
		assert.Equalf(t, 5, msg.ID(), "Expected Id 5 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ping")
	}
}
