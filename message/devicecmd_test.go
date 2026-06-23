package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const StopCmdDocExample = `[{"StopCmd":{"Id":1,"DeviceIndex":0,"Inputs":true,"Outputs":true}}]`

func TestSerialize_StopCmd_DocExample(t *testing.T) {
	trueVal := true
	msg := StopCmd{
		message:     message{Id: 1},
		DeviceIndex: 0,
		Inputs:      &trueVal,
		Outputs:     &trueVal,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, StopCmdDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_StopCmd_DocExample(t *testing.T) {
	jsonMessage := StopCmdDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopCmd); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, true, *msg.Inputs)
		assert.Equal(t, true, *msg.Outputs)
	} else {
		t.Errorf("Deserialized message is not of type StopCmd")
	}
}

const StopAllDevicesDocExample = `[{"StopAllDevices":{"Id":1,"Inputs":true,"Outputs":true}}]`

func TestSerialize_StopAllDevices_DocExample(t *testing.T) {
	trueVal := true
	msg := StopAllDevices{
		message: message{Id: 1},
		Inputs:  &trueVal,
		Outputs: &trueVal,
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, StopAllDevicesDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_StopAllDevices_DocExample(t *testing.T) {
	jsonMessage := StopAllDevicesDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopAllDevices); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, true, *msg.Inputs)
		assert.Equal(t, true, *msg.Outputs)
	} else {
		t.Errorf("Deserialized message is not of type StopAllDevices")
	}
}

const OutputCmdVibrateDocExample = `[{"OutputCmd":{"Id":1,"DeviceIndex":0,"FeatureIndex":0,"Command":{"Vibrate":{"Value":10,"Clockwise":null,"Duration":null}}}}]`

func TestSerialize_OutputCmdVibrate_DocExample(t *testing.T) {
	msg := OutputCmd{
		message:      message{Id: 1},
		DeviceIndex:  0,
		FeatureIndex: 0,
		Command: OutputValue{
			"Vibrate": {Value: 10},
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, OutputCmdVibrateDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_OutputCmdVibrate_DocExample(t *testing.T) {
	jsonMessage := OutputCmdVibrateDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*OutputCmd); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, 0, msg.FeatureIndex)
		assert.Equal(t, 10, msg.Command["Vibrate"].Value)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

const OutputCmdRotateWithDirectionDocExample = `[{"OutputCmd":{"Id":1,"DeviceIndex":0,"FeatureIndex":0,"Command":{"RotateWithDirection":{"Value":10,"Clockwise":false,"Duration":null}}}}]`

func TestSerialize_OutputCmdRotateWithDirection_DocExample(t *testing.T) {
	falseVal := false
	msg := OutputCmd{
		message:      message{Id: 1},
		DeviceIndex:  0,
		FeatureIndex: 0,
		Command: OutputValue{
			"RotateWithDirection": {Value: 10, Clockwise: &falseVal},
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, OutputCmdRotateWithDirectionDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_OutputCmdRotateWithDirection_DocExample(t *testing.T) {
	jsonMessage := OutputCmdRotateWithDirectionDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*OutputCmd); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, 0, msg.FeatureIndex)
		assert.Equal(t, 10, msg.Command["RotateWithDirection"].Value)
		assert.Equal(t, false, *msg.Command["RotateWithDirection"].Clockwise)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

const OutputCmdPositionWithDurationDocExample = `[{"OutputCmd":{"Id":1,"DeviceIndex":0,"FeatureIndex":0,"Command":{"PositionWithDuration":{"Value":85,"Clockwise":null,"Duration":15}}}}]`

func TestSerialize_OutputCmdPositionWithDuration_DocExample(t *testing.T) {
	duration := 15
	msg := OutputCmd{
		message:      message{Id: 1},
		DeviceIndex:  0,
		FeatureIndex: 0,
		Command: OutputValue{
			"PositionWithDuration": {Value: 85, Duration: &duration},
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, OutputCmdPositionWithDurationDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_OutputCmdPositionWithDuration_DocExample(t *testing.T) {
	jsonMessage := OutputCmdPositionWithDurationDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*OutputCmd); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, 0, msg.FeatureIndex)
		assert.Equal(t, 85, msg.Command["PositionWithDuration"].Value)
		assert.Equal(t, 15, *msg.Command["PositionWithDuration"].Duration)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

const InputCmdDocExample = `[{"InputCmd":{"Id":1,"DeviceIndex":0,"FeatureIndex":1,"Type":"Battery","Command":"Read"}},{"InputCmd":{"Id":2,"DeviceIndex":1,"FeatureIndex":0,"Type":"Pressure","Command":"Subscribe"}}]`

func TestSerialize_InputCmd_DocExample(t *testing.T) {
	msgs := []*InputCmd{
		{
			message:      message{Id: 1},
			DeviceIndex:  0,
			FeatureIndex: 1,
			Type:         "Battery",
			Command:      "Read",
		},
		{
			message:      message{Id: 2},
			DeviceIndex:  1,
			FeatureIndex: 0,
			Type:         "Pressure",
			Command:      "Subscribe",
		},
	}
	jsonMsg, err := Serialize(msgs...)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, InputCmdDocExample, jsonMsg, "Serialized message does not match expected JSON")	
}

func TestDeserialize_InputCmd_DocExample(t *testing.T) {
	jsonMessage := InputCmdDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*InputCmd); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, 1, msg.FeatureIndex)
		assert.Equal(t, "Battery", msg.Type)
		assert.Equal(t, "Read", msg.Command)
	} else {
		t.Errorf("Deserialized message is not of type InputCmd")
	}
	if msg, ok := msg[1].(*InputCmd); ok {
		assert.Equalf(t, 2, msg.ID(), "Expected Id 2 found %d", msg.ID())
		assert.Equal(t, 1, msg.DeviceIndex)
		assert.Equal(t, 0, msg.FeatureIndex)
		assert.Equal(t, "Pressure", msg.Type)
		assert.Equal(t, "Subscribe", msg.Command)
	} else {
		t.Errorf("Deserialized message is not of type InputCmd")
	}	
}

const InputReadingDocExample = `[{"InputReading":{"Id":1,"DeviceIndex":0,"FeatureIndex":0,"Reading":{"Battery":{"Value":50},"RSSI":null,"Pressure":null,"Button":null}}}]`

func TestSerialize_InputReading_DocExample(t *testing.T) {
	msg := InputReading{
		message:      message{Id: 1},
		DeviceIndex:  0,
		FeatureIndex: 0,
		Reading: InputData{
			Battery: &struct{ Value int }{Value: 50},
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	assert.Equalf(t, InputReadingDocExample, jsonMsg, "Serialized message does not match expected JSON")
}

func TestDeserialize_InputReading_DocExample(t *testing.T) {
	jsonMessage := InputReadingDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*InputReading); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, 0, msg.DeviceIndex)
		assert.Equal(t, 0, msg.FeatureIndex)
		assert.NotNil(t, msg.Reading.Battery)
		assert.Equal(t, 50, msg.Reading.Battery.Value)
		assert.Nil(t, msg.Reading.RSSI)
		assert.Nil(t, msg.Reading.Pressure)
		assert.Nil(t, msg.Reading.Button)
	} else {
		t.Errorf("Deserialized message is not of type InputReading")
	}
}
