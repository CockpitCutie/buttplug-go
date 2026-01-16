package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserializeOk(t *testing.T) {
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
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ok")
	}
}

func TestDeserializeError(t *testing.T) {
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
	if msg, ok := msg[0].(*Error); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Server received invalid JSON.", msg.Message, "Expected ErrorMessage 'Server received invalid JSON.' found '%s'", msg.Message)
		assert.Equalf(t, MsgError, msg.Code, "Expected ErrorCode MsgError found %d", msg.Code)
	} else {
		t.Errorf("Deserialized message is not of type Error")
	}
}

func TestDeserializePing(t *testing.T) {
	jsonMessage := `[
  {
    "Ping": {
      "Id": 5
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*Ping); ok {
		assert.Equalf(t, uint32(5), msg.ID(), "Expected Id 5 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type Ping")
	}
}

func TestDeserializeRequestServerInfo(t *testing.T) {
	jsonMessage := `[
  {
    "RequestServerInfo": {
      "Id": 1,
      "ClientName": "Test Client",
      "ProtocolVersionMajor": 4,
      "ProtocolVersionMinor": 0
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*RequestServerInfo); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Client", msg.ClientName, "Expected ClientName 'Test Client' found '%s'", msg.ClientName)
		assert.Equalf(t, uint(4), msg.ProtocolVersionMajor, "Expected ProtoMajor 1 found %d", msg.ProtocolVersionMajor)
		assert.Equalf(t, uint(0), msg.ProtocolVersionMinor, "Expected ProtoMinor 1 found %d", msg.ProtocolVersionMinor)
	} else {
		t.Errorf("Deserialized message is not of type RequestServerInfo")
	}
}

func TestServerInfo(t *testing.T) {
	jsonMessage := `[
  {
    "ServerInfo": {
      "Id": 1,
      "ServerName": "Test Server",
      "MaxPingTime": 100,
      "ProtocolVersionMajor": 4,
      "ProtocolVersionMinor": 0
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*ServerInfo); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Server", msg.ServerName, "Expected ServerName 'Test Server' found '%s'", msg.ServerName)
		assert.Equalf(t, uint(100), msg.MaxPingTime, "Expected MaxPingTime 100 found %d", msg.MaxPingTime)
		assert.Equalf(t, uint(4), msg.ProtocolVersionMajor, "Expected ProtoMajor 4 found %d", msg.ProtocolVersionMajor)
		assert.Equalf(t, uint(0), msg.ProtocolVersionMinor, "Expected ProtoMinor 0 found %d", msg.ProtocolVersionMinor)
	} else {
		t.Errorf("Deserialized message is not of type ServerInfo")
	}
}

func TestDeserializeStartScanning(t *testing.T) {
	jsonMessage := `[
  {
    "StartScanning": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StartScanning); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StartScanning")
	}
}

func TestDeserializeStopScanning(t *testing.T) {
	jsonMessage := `[
  {
    "StopScanning": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopScanning); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StopScanning")
	}
}

func TestDeserializeScanningFinished(t *testing.T) {
	jsonMessage := `[
  {
    "ScanningFinished": {
      "Id": 0
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*ScanningFinished); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 0 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type ScanningFinished")
	}
}

func TestDeserializeRequestDeviceList(t *testing.T) {
	jsonMessage := `[
  {
    "RequestDeviceList": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*RequestDeviceList); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type RequestDeviceList")
	}
}

func TestDeserializeDeviceList(t *testing.T) {
	jsonMessage := `[
  {
    "DeviceList": {
      "Id": 1,
      "Devices": {
        "0": {
          "DeviceName": "Test Vibrator",
          "DeviceIndex": 0,
          "Features": {
            "0": {
              "FeatureIndex": 0,
              "FeatureDescription": "Clitoral Stimulator",
              "Output": {
                "Vibrate": {
                  "Value": [0, 20]
                }
              }
            },
            "1": {
              "FeatureIndex": 1,
              "FeatureDescription": "Insertable Stimulator",
              "Output": {
                "Vibrate": {
                  "Value": [0, 20]
                }
              }
            },
            "2": {
              "FeatureIndex": 2,
              "FeatureDescription": "Rotating Head with Directional Control",
              "Output": {
                "Vibrate": {
                  "Value": [-20, 20]
                }
              }
            },
            "3": {
              "FeatureIndex": 3,
              "FeatureDescription": "Battery",
              "Input": {
                "Battery": {
                  "Value": [0, 100],
                  "Command": ["Read"]
                }
              }
            }
          }
        },
        "1": {
          "DeviceName": "Test Stroker",
          "DeviceIndex": 1,
          "DeviceMessageTimingGap": 100,
          "DeviceDisplayName": "User set name",
          "Features": {
            "0": {
              "FeatureIndex": 0,
              "FeatureDescription": "Stroker",
              "Output": {
                "PositionWithDuration": {
                  "Position": [0, 100],
                  "Duration": [0, 100000]
                },
                "Position": {
                  "Position": [0, 100]
                }
              },
              "Input": {
                "Position": {
                  "Value": [0, 100],
                  "Command": ["Read", "Subscribe"]
                }
              }
            },
            "2": {
              "FeatureIndex": 2,
              "FeatureDescription": "Bluetooth Radio RSSI",
              "Input": {
                "RSSI": {
                  "Value": [-10, -100],
                  "Command": ["Read"]
                }
              }
            }
          }
        }
      }
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*DeviceList); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		if dev0, ok := msg.Devices["0"]; assert.True(t, ok) {
			assert.Equal(t, uint(0), dev0.DeviceIndex)
			assert.Equal(t, "Test Vibrator", dev0.DeviceName)
			assert.Equal(t, "", dev0.DeviceDisplayName)
			assert.Equal(t, uint(0), dev0.DeviceMessageTimingGap)
			if feat, ok := dev0.Features["0"]; assert.True(t, ok) {
				assert.Equal(t, uint32(0), feat.FeatureIndex)
				assert.Equal(t, "Clitoral Stimulator", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: &[2]int{0, 20}}}, feat.Output)
			}
			if feat, ok := dev0.Features["1"]; assert.True(t, ok) {
				assert.Equal(t, uint32(1), feat.FeatureIndex)
				assert.Equal(t, "Insertable Stimulator", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: &[2]int{0, 20}}}, feat.Output)
			}
			if feat, ok := dev0.Features["2"]; assert.True(t, ok) {
				assert.Equal(t, uint32(2), feat.FeatureIndex)
				assert.Equal(t, "Rotating Head with Directional Control", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: &[2]int{-20, 20}}}, feat.Output)
			}
			if feat, ok := dev0.Features["3"]; assert.True(t, ok) {
				assert.Equal(t, uint32(3), feat.FeatureIndex)
				assert.Equal(t, "Battery", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceInput{"Battery": {Value: &[2]int{0, 100}, Command: []string{"Read"}}}, feat.Input)
			}
		}
		if dev1, ok := msg.Devices["1"]; assert.True(t, ok) {
			assert.Equal(t, uint(1), dev1.DeviceIndex)
			assert.Equal(t, "Test Stroker", dev1.DeviceName)
			assert.Equal(t, "User set name", dev1.DeviceDisplayName)
			assert.Equal(t, uint(100), dev1.DeviceMessageTimingGap)
			if feat, ok := dev1.Features["0"]; assert.True(t, ok) {
				assert.Equal(t, uint32(0), feat.FeatureIndex)
				assert.Equal(t, "Stroker", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{
					"PositionWithDuration": {Position: &[2]uint{0, 100}, Duration: &[2]uint{0, 100000}},
					"Position":             {Position: &[2]uint{0, 100}}}, feat.Output)
			}
			if feat, ok := dev1.Features["2"]; assert.True(t, ok) {
				assert.Equal(t, uint32(2), feat.FeatureIndex)
				assert.Equal(t, "Bluetooth Radio RSSI", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceInput{"RSSI": {Value: &[2]int{-10, -100}, Command: []string{"Read"}}}, feat.Input)
			}
		}
	} else {
		t.Errorf("Deserialized message is not of type DeviceList")
	}
}

func TestDeserializeStopDeviceCmd(t *testing.T) {
	jsonMessage := `[
  {
    "StopDeviceCmd": {
      "Id": 1,
      "DeviceIndex": 0,
      "Inputs": true,
      "Outputs": true
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopDeviceCmd); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, uint(0), msg.DeviceIndex)
		assert.Equal(t, true, *msg.Inputs)
		assert.Equal(t, true, *msg.Outputs)
	} else {
		t.Errorf("Deserialized message is not of type StopDeviceCmd")
	}
}

func TestDeserializeStopAllDevices(t *testing.T) {
	jsonMessage := `[
  {
    "StopAllDevices": {
      "Id": 1,
      "Inputs": true,
      "Outputs": true
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*StopAllDevices); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equal(t, true, *msg.Inputs)
		assert.Equal(t, true, *msg.Outputs)
	} else {
		t.Errorf("Deserialized message is not of type StopAllDevices")
	}
}

func TestDeserializeOutputCmdVibrate(t *testing.T) {
	jsonMessage := `  [{
    "OutputCmd": {
      "Id": 1,
      "DeviceIndex": 0,
      "FeatureIndex": 0,
      "Command": {
        "Vibrate": {
          "Value": 10
        }
      }
    }
  }]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoError(t, err)
	if msg, ok := msg[0].(*OutputCmd); assert.True(t, ok) {
		assert.Equal(t, uint32(1), msg.Id)
		assert.Equal(t, uint(0), msg.DeviceIndex)
		assert.Equal(t, uint(0), msg.FeatureIndex)
		assert.Equal(t, uint32(10), msg.Command["Vibrate"].Value)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

func TestDeserializeOutputCmdRotationWithDirection(t *testing.T) {
	jsonMessage := `  [{
    "OutputCmd": {
      "Id": 1,
      "DeviceIndex": 0,
      "FeatureIndex": 0,
      "Command": {
        "RotateWithDirection": {
          "Value": 10,
          "Clockwise": false
        }
      }
    }
  }]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoError(t, err)
	if msg, ok := msg[0].(*OutputCmd); assert.True(t, ok) {
		assert.Equal(t, uint32(1), msg.Id)
		assert.Equal(t, uint(0), msg.DeviceIndex)
		assert.Equal(t, uint(0), msg.FeatureIndex)
		assert.Equal(t, uint32(10), msg.Command["RotateWithDirection"].Value)
		assert.False(t, *msg.Command["RotateWithDirection"].Clockwise)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

func TestDeserializeOutputCmdPositionWithDuration(t *testing.T) {
	jsonMessage := `  [{
    "OutputCmd": {
      "Id": 1,
      "DeviceIndex": 0,
      "FeatureIndex": 0,
      "Command": {
        "PositionWithDuration": {
          "Value": 85,
          "Duration": 15
        }
      }
    }
  }]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoError(t, err)
	if msg, ok := msg[0].(*OutputCmd); assert.True(t, ok) {
		assert.Equal(t, uint32(1), msg.Id)
		assert.Equal(t, uint(0), msg.DeviceIndex)
		assert.Equal(t, uint(0), msg.FeatureIndex)
		assert.Equal(t, uint32(85), msg.Command["PositionWithDuration"].Value)
		assert.Equal(t, uint32(15), *msg.Command["PositionWithDuration"].Duration)
	} else {
		t.Errorf("Deserialized message is not of type OutputCmd")
	}
}

func TestDeserializeMultipleMessages(t *testing.T) {
	jsonMessage := `[
  {
    "OutputCmd": {
      "Id": 1,
      "DeviceIndex": 0,
      "FeatureIndex": 0,
      "Command": {
        "Vibrate": {
          "Value": 10
        }
      }
    }
  },
  {
    "OutputCmd": {
      "Id": 2,
      "DeviceIndex": 1,
      "FeatureIndex": 0,
      "Command": {
        "PositionWithDuration": {
          "Position": 91,
          "Duration": 150
        }
      }
    }
  }
]`
	msgs, err := Deserialize([]byte(jsonMessage))
	assert.NoError(t, err)
	assert.Len(t, msgs, 2)
}
