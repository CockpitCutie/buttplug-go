package device

import (
	"testing"

	"github.com/CockpitCutie/buttplug-go/message"
	"github.com/stretchr/testify/assert"
)

var featuresMsg = message.DeviceFeatures{
	"0": {
		FeatureIndex:       0,
		FeatureDescription: "Clitoral Stimulator",
		Output: map[string]message.DeviceOutput{
			"Vibrate": {
				Value: []int{0, 20},
			},
		},
	},
	"1": {
		FeatureIndex:       1,
		FeatureDescription: "Insertable Stimulator",
		Output: map[string]message.DeviceOutput{
			"Vibrate": {
				Value: []int{0, 20},
			},
		},
	},
	"2": {
		FeatureIndex:       2,
		FeatureDescription: "Rotating Head with Directional Control",
		Output: map[string]message.DeviceOutput{
			"Vibrate": {
				Value: []int{-20, 20},
			},
		},
	},
	"3": {
		FeatureIndex:       3,
		FeatureDescription: "Battery",
		Input: map[string]message.DeviceInput{
			"Battery": {
				Value:   [2]int32{0, 100},
				Command: []string{"Read"},
			},
		},
	},
}

func TestRegisterOutputs(t *testing.T) {
	d := &Device{
		Name:  "Test Vibrator",
		Index: 0,
	}
	err := d.registerOutputs(featuresMsg)
	assert.NoError(t, err)
	assert.Len(t, d.Outputs, 3)

	assert.IsType(t, Vibrator{}, d.Outputs[0])
	assert.Equal(t, uint32(0), d.Outputs[0].Index())
	assert.Equal(t, "Clitoral Stimulator", d.Outputs[0].Description())
	assert.Equal(t, VibrateOutput, d.Outputs[0].OutputType())
	
	assert.IsType(t, Vibrator{}, d.Outputs[0])
	assert.Equal(t, uint32(1), d.Outputs[0].Index())
	assert.Equal(t, "Insertable Stimulator", d.Outputs[0].Description())
	assert.Equal(t, VibrateOutput, d.Outputs[0].OutputType())
}
