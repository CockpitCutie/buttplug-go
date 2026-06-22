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
				Value: [2]int{0, 20},
			},
		},
	},
	"1": {
		FeatureIndex:       1,
		FeatureDescription: "Insertable Stimulator",
		Output: map[string]message.DeviceOutput{
			"Vibrate": {
				Value: [2]int{0, 20},
			},
		},
	},
	"2": {
		FeatureIndex:       2,
		FeatureDescription: "Rotating Head with Directional Control",
		Output: map[string]message.DeviceOutput{
			"RotationWithDirection": {
				Value: [2]int{-20, 20},
			},
		},
	},
	"3": {
		FeatureIndex:       3,
		FeatureDescription: "Battery",
		Input: map[string]message.DeviceInput{
			"Battery": {
				Value:   [2]int{0, 100},
				Command: []string{"Read"},
			},
		},
	},
}

func TestRegisterOutputs(t *testing.T) {
	d := &Device{
		name:  "Test Vibrator",
		index: 0,
	}
	err := d.registerOutputs(featuresMsg)
	assert.NoError(t, err)
	assert.Len(t, d.Outputs(), 3)

	out0 := d.GetFeatureById(0).(Output)
	assert.IsType(t, Vibrator{}, out0)
	assert.Equal(t, 0, out0.Index())
	assert.Equal(t, "Clitoral Stimulator", out0.Description())
	assert.Equal(t, VibrateOutput, out0.OutputType())

	out1 := d.GetFeatureById(1).(Output)
	assert.IsType(t, Vibrator{}, out1)
	assert.Equal(t, 1, out1.Index())
	assert.Equal(t, "Insertable Stimulator", out1.Description())
	assert.Equal(t, VibrateOutput, out1.OutputType())

	out2 := d.GetFeatureById(2).(Output)
	assert.IsType(t, RotatorWithDirection{}, out2)
	assert.Equal(t, 2, out2.Index())
	assert.Equal(t, "Rotating Head with Directional Control", out2.Description())
	assert.Equal(t, RotationWithDirectionOutput, out2.OutputType())
}

func TestRegisterInputs(t *testing.T) {
	d := &Device{
		name:  "Test Vibrator",
		index: 0,
	}
	err := d.registerInputs(featuresMsg)
	assert.NoError(t, err)
	assert.Len(t, d.inputs, 1)

	in3 := d.GetFeatureById(3).(Input)
	assert.IsType(t, Battery{}, in3)
	assert.Equal(t, 3, in3.Index())
	assert.Equal(t, "Battery", in3.Description())
	assert.Equal(t, BatteryInput, in3.InputType())
}
