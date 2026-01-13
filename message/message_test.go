package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	jsonMessage := `[
  {
    "Ok": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*Ok); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ok")
	}

}

func TestError(t *testing.T) {
	jsonMessage := `[
  {
    "Error": {
      "Id": 0,
      "ErrorMessage": "Server received invalid JSON.",
      "ErrorCode": 3
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*Error); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Server received invalid JSON.", msg.Message, "Expected ErrorMessage 'Server received invalid JSON.' found '%s'", msg.Message)
		assert.Equalf(t, MsgError, msg.Code, "Expected ErrorCode MsgError found %d", msg.Code)
	} else {
		t.Errorf("Deserialized message is not of type Error")
	}
}

func TestPing(t *testing.T) {
	jsonMessage := `[
  {
    "Ping": {
      "Id": 5
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*Ping); ok {
		assert.Equalf(t, uint32(5), msg.ID(), "Expected Id 5 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ping")
	}
}
