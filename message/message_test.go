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

func TestRequestServerInfo(t *testing.T) {
	jsonMessage := `[
  {
    "RequestServerInfo": {
      "Id": 1,
      "ClientName": "Test Client",
      "MessageVersion": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*RequestServerInfo); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Client", msg.ClientName, "Expected ClientName 'Test Client' found '%s'", msg.ClientName)
		assert.Equalf(t, uint(1), msg.MessageVersion, "Expected MessageVersion 1 found %d", msg.MessageVersion)
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
      "MessageVersion": 1,
      "MaxPingTime": 100
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*ServerInfo); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, "Test Server", msg.ServerName, "Expected ServerName 'Test Server' found '%s'", msg.ServerName)
		assert.Equalf(t, uint(1), msg.MessageVersion, "Expected MessageVersion 1 found %d", msg.MessageVersion)
		assert.Equalf(t, uint(100), msg.MaxPingTime, "Expected MaxPingTime 100 found %d", msg.MaxPingTime)
	} else {
		t.Errorf("Deserialized message is not of type ServerInfo")
	}
}

func TestStartScanning(t *testing.T) {
	jsonMessage := `[
  {
    "StartScanning": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*StartScanning); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StartScanning")
	}
}

func TestStopScanning(t *testing.T) {
	jsonMessage := `[
  {
    "StopScanning": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*StopScanning); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type StopScanning")
	}
}

func TestScanningFinished(t *testing.T) {
	jsonMessage := `[
  {
    "ScanningFinished": {
      "Id": 0
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*ScanningFinished); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 0 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type ScanningFinished")
	}
}

func TestRequestDeviceList(t *testing.T) {
	jsonMessage := `[
  {
    "RequestDeviceList": {
      "Id": 1
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*RequestDeviceList); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
	} else {
		t.Errorf("Deserialized message is not of type RequestDeviceList")
	}
}

func TestDeviceList(t *testing.T) {
	jsonMessage := `[
  {
    "DeviceList": {
      "Id": 1,
      "Devices": [
        {
          "DeviceName": "Test Vibrator",
          "DeviceIndex": 0,
          "DeviceMessages": {
            "ScalarCmd": [
              {
                "StepCount": 20,
                "FeatureDescriptor": "Clitoral Stimulator",
                "ActuatorType": "Vibrate"
              },
              {
                "StepCount": 20,
                "FeatureDescriptor": "Insertable Vibrator",
                "ActuatorType": "Vibrate"
              }
            ],
            "StopDeviceCmd": {}
          }
        },
        {
          "DeviceName": "Test Stroker",
          "DeviceIndex": 1,
          "DeviceMessageTimingGap": 100,
          "DeviceDisplayName": "User set name",
          "DeviceMessages": {
            "LinearCmd": [ {
              "StepCount": 100,
              "FeatureDescriptor": "Stroker",
              "ActuatorType": "Linear"
            } ],
            "StopDeviceCmd": {}
          }
        }
      ]
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*DeviceList); ok {
		assert.Equalf(t, uint32(1), msg.ID(), "Expected Id 1 found %d", msg.ID())
		assert.Equalf(t, 2, len(msg.Devices), "Expected 2 devices found %d", len(msg.Devices))
	} else {
		t.Errorf("Deserialized message is not of type DeviceList")
	}
}

func TestDeviceAdded(t *testing.T) {
	jsonMessage := `[
  {
    "DeviceAdded": {
      "Id": 0,
      "DeviceName": "Test Vibrator",
      "DeviceIndex": 0,
      "DeviceMessageTimingGap": 100,
      "DeviceDisplayName": "Rabbit Vibrator",
      "DeviceMessages": {
        "ScalarCmd": [
          {
            "StepCount": 20,
            "FeatureDescriptor": "Clitoral Stimulator",
            "ActuatorType": "Vibrate"
          },
          {
            "StepCount": 20,
            "FeatureDescriptor": "Insertable Vibrator",
            "ActuatorType": "Vibrate"
          }
        ],
        "StopDeviceCmd": {}
       }
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*DeviceAdded); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 0 found %d", msg.ID())
		assert.Equalf(t, "Test Vibrator", msg.DeviceName, "Expected DeviceName 'Test Vibrator' found '%s'", msg.DeviceName)
		assert.Equalf(t, uint(0), msg.DeviceIndex, "Expected DeviceIndex 0 found %d", msg.DeviceIndex)
		assert.Equalf(t, uint(100), *msg.DeviceMessageTimingGap, "Expected DeviceMessageTimingGap 100 found %d", *msg.DeviceMessageTimingGap)
		assert.Equalf(t, "Rabbit Vibrator", *msg.DeviceDisplayName, "Expected DeviceDisplayName 'Rabbit Vibrator' found '%s'", *msg.DeviceDisplayName)
	} else {
		t.Errorf("Deserialized message is not of type DeviceAdded")
	}
}

func TestDeviceRemoved(t *testing.T) {
	jsonMessage := `[
  {
    "DeviceRemoved": {
      "Id": 0,
      "DeviceIndex": 0
    }
  }
]`
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg.(*DeviceRemoved); ok {
		assert.Equalf(t, uint32(0), msg.ID(), "Expected Id 0 found %d", msg.ID())
		assert.Equalf(t, uint(0), msg.DeviceIndex, "Expected DeviceIndex 0 found %d", msg.DeviceIndex)
	} else {
		t.Errorf("Deserialized message is not of type DeviceRemoved")
	}
}
