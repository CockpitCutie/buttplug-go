package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const RequestServerInfoDocExample = `[{"RequestServerInfo":{"Id":1,"ClientName":"Test Client","ProtocolVersionMajor":4,"ProtocolVersionMinor":0}}]`

func TestSerialize_RequestServerInfo_DocExample(t *testing.T) {
	msg := RequestServerInfo{
		message: message{
			Id: 1,
		},
		ClientName:           "Test Client",
		ProtocolVersionMajor: 4,
		ProtocolVersionMinor: 0,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := RequestServerInfoDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_RequestServerInfo_DocExample(t *testing.T) {
	jsonMessage := RequestServerInfoDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*RequestServerInfo); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Client", msg.ClientName, "Expected ClientName 'Test Client' found '%s'", msg.ClientName)
		assert.Equalf(t, 4, msg.ProtocolVersionMajor, "Expected ProtoMajor 1 found %d", msg.ProtocolVersionMajor)
		assert.Equalf(t, 0, msg.ProtocolVersionMinor, "Expected ProtoMinor 1 found %d", msg.ProtocolVersionMinor)
	} else {
		t.Errorf("Deserialized message is not of type RequestServerInfo")
	}
}

const ServerInfoDocExample = `[{"ServerInfo":{"Id":1,"ServerName":"Test Server","MaxPingTime":100,"ProtocolVersionMajor":4,"ProtocolVersionMinor":0}}]`

func TestSerialize_ServerInfo_DocExample(t *testing.T) {
	msg := ServerInfo{
		message: message{
			Id: 1,
		},
		ServerName:           "Test Server",
		MaxPingTime:          100,
		ProtocolVersionMajor: 4,
		ProtocolVersionMinor: 0,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := ServerInfoDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_ServerInfo_DocExample(t *testing.T) {
	jsonMessage := ServerInfoDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*ServerInfo); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Server", msg.ServerName, "Expected ServerName 'Test Server' found '%s'", msg.ServerName)
		assert.Equalf(t, 100, msg.MaxPingTime, "Expected MaxPingTime 100 found %d", msg.MaxPingTime)
		assert.Equalf(t, 4, msg.ProtocolVersionMajor, "Expected ProtoMajor 4 found %d", msg.ProtocolVersionMajor)
		assert.Equalf(t, 0, msg.ProtocolVersionMinor, "Expected ProtoMinor 0 found %d", msg.ProtocolVersionMinor)
	} else {
		t.Errorf("Deserialized message is not of type ServerInfo")
	}
}

const DisconnectDocExample = `[{"Disconnect":{"Id":1}}]`

func TestSerialize_Disconnect_DocExample(t *testing.T) {
	msg := Disconnect{
		message: message{
			Id: 1,
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	expectedJson := DisconnectDocExample
	assert.Equalf(t, expectedJson, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_Disconnect_DocExample(t *testing.T) {
	jsonMessage := DisconnectDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Disconnect); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Disconnect")
	}
}