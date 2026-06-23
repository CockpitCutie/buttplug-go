package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const DeviceListDocExample = `[
 {
   "DeviceList": {
     "Id": 1,
     "Devices": {
       "0": {
         "DeviceName": "Test Vibrator",
         "DeviceIndex": 0,
         "DeviceFeatures": {
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
         "DeviceFeatures": {
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

func TestSerialize_DeviceList_DocExample(t *testing.T) {
	msg := DeviceList{
		message: message{Id: 1},
		Devices: map[string]Device{
			"0": {
				DeviceName:  "Test Vibrator",
				DeviceIndex: 0,
				DeviceFeatures: DeviceFeatures{
					"0": {
						FeatureIndex:       0,
						FeatureDescription: "Clitoral Stimulator",
						Output: map[string]DeviceOutput{
							"Vibrate": {Value: [2]int{0, 20}},
						},
					},
				},
			},
		},
	}
	jsonMsg, err := Serialize(&msg)
	assert.NoErrorf(t, err, "Error serializing message")
	msgs, err := Deserialize([]byte(jsonMsg))
	assert.NoErrorf(t, err, "Error deserializing serialized message")
	if dl, ok := msgs[0].(*DeviceList); ok {
		assert.Equalf(t, 1, dl.ID(), "Expected Id 1 found %d", dl.ID())
		assert.Equal(t, 1, len(dl.Devices))
		dev := dl.Devices["0"]
		assert.Equal(t, "Test Vibrator", dev.DeviceName)
		assert.Equal(t, 0, dev.DeviceIndex)
		assert.Equal(t, "Clitoral Stimulator", dev.DeviceFeatures["0"].FeatureDescription)
		assert.Equal(t, [2]int{0, 20}, dev.DeviceFeatures["0"].Output["Vibrate"].Value)
	} else {
		t.Errorf("Deserialized message is not of type DeviceList")
	}
}

func TestDeserialize_DeviceList_DocExample(t *testing.T) {
	jsonMessage := DeviceListDocExample
	msg, err := Deserialize([]byte(jsonMessage))
	assert.NoErrorf(t, err, "Error deserializing message")
	if msg, ok := msg[0].(*DeviceList); ok {
		assert.Equalf(t, 1, msg.ID(), "Expected Id 1 found %d", msg.ID())
		if dev0, ok := msg.Devices["0"]; assert.True(t, ok) {
			assert.Equal(t, 0, dev0.DeviceIndex)
			assert.Equal(t, "Test Vibrator", dev0.DeviceName)
			assert.Equal(t, "", dev0.DeviceDisplayName)
			assert.Equal(t, 0, dev0.DeviceMessageTimingGap)
			if feat, ok := dev0.DeviceFeatures["0"]; assert.True(t, ok) {
				assert.Equal(t, 0, feat.FeatureIndex)
				assert.Equal(t, "Clitoral Stimulator", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: [2]int{0, 20}}}, feat.Output)
			}
			if feat, ok := dev0.DeviceFeatures["1"]; assert.True(t, ok) {
				assert.Equal(t, 1, feat.FeatureIndex)
				assert.Equal(t, "Insertable Stimulator", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: [2]int{0, 20}}}, feat.Output)
			}
			if feat, ok := dev0.DeviceFeatures["2"]; assert.True(t, ok) {
				assert.Equal(t, 2, feat.FeatureIndex)
				assert.Equal(t, "Rotating Head with Directional Control", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{"Vibrate": {Value: [2]int{-20, 20}}}, feat.Output)
			}
			if feat, ok := dev0.DeviceFeatures["3"]; assert.True(t, ok) {
				assert.Equal(t, 3, feat.FeatureIndex)
				assert.Equal(t, "Battery", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceInput{"Battery": {Value: [2]int{0, 100}, Command: []string{"Read"}}}, feat.Input)
			}
		}
		if dev1, ok := msg.Devices["1"]; assert.True(t, ok) {
			assert.Equal(t, 1, dev1.DeviceIndex)
			assert.Equal(t, "Test Stroker", dev1.DeviceName)
			assert.Equal(t, "User set name", dev1.DeviceDisplayName)
			assert.Equal(t, 100, dev1.DeviceMessageTimingGap)
			if feat, ok := dev1.DeviceFeatures["0"]; assert.True(t, ok) {
				assert.Equal(t, 0, feat.FeatureIndex)
				assert.Equal(t, "Stroker", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceOutput{
					"PositionWithDuration": {Position: []int{0, 100}, Duration: []int{0, 100000}},
					"Position":             {Position: []int{0, 100}}}, feat.Output)
			}
			if feat, ok := dev1.DeviceFeatures["2"]; assert.True(t, ok) {
				assert.Equal(t, 2, feat.FeatureIndex)
				assert.Equal(t, "Bluetooth Radio RSSI", feat.FeatureDescription)
				assert.Equal(t, map[string]DeviceInput{"RSSI": {Value: [2]int{-10, -100}, Command: []string{"Read"}}}, feat.Input)
			}
		}
	} else {
		t.Errorf("Deserialized message is not of type DeviceList")
	}
}
