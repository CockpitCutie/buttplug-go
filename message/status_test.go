package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeOkDocExample(t *testing.T) {
	msg := Ok{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := `[{"Ok":{"Id":1}}]`
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserializeOkDocExample(t *testing.T) {
	jsonMessage := `[
  {
    "Ok": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Ok); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ok")
	}
}

func TestSerializeErrorDocExample(t *testing.T) {
	msg := Error{
		message: message{
			Id: 0,
		},
		Message: "Server received invalid JSON.",
		Code:    MsgError,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := `[{"Error":{"Id":0,"ErrorMessage":"Server received invalid JSON.","ErrorCode":3}}]`
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserializeErrorDocExample(t *testing.T) {
	jsonMessage := `[{"Error":{"Id":0,"ErrorMessage":"Server received invalid JSON.","ErrorCode":3}}]`
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

func TestSerializePingDocExample(t *testing.T) {
	msg := Ping{
		message: message{
			Id: 5,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := `[{"Ping":{"Id":5}}]`
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserializePingDocExample(t *testing.T) {
	jsonMessage := `[{"Ping":{"Id":5}}]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Ping); ok {
		assert.Equalf(t, 5, msg.ID(), "Expected Id 5 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ping")
	}
}
