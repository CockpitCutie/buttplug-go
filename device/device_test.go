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
		Name:    "Test Vibrator",
		Index:   0,
		Inputs:  make(map[int]Input),
		Outputs: make(map[int]Output),
	}
	err := d.registerOutputs(featuresMsg)
	assert.NoError(t, err)
	assert.Len(t, d.Outputs, 3)

	assert.IsType(t, Vibrator{}, d.Outputs[0])
	assert.Equal(t, 0, d.Outputs[0].Index())
	assert.Equal(t, "Clitoral Stimulator", d.Outputs[0].Description())
	assert.Equal(t, VibrateOutput, d.Outputs[0].OutputType())

	assert.IsType(t, Vibrator{}, d.Outputs[1])
	assert.Equal(t, uint32(1), d.Outputs[1].Index())
	assert.Equal(t, "Insertable Stimulator", d.Outputs[1].Description())
	assert.Equal(t, VibrateOutput, d.Outputs[1].OutputType())

	assert.IsType(t, RotatorWithDirection{}, d.Outputs[2])
	assert.Equal(t, uint32(2), d.Outputs[2].Index())
	assert.Equal(t, "Rotating Head with Directional Control", d.Outputs[2].Description())
	assert.Equal(t, RotationWithDirectionOutput, d.Outputs[2].OutputType())
}

func TestRegisterInputs(t *testing.T) {
	d := &Device{
		Name:    "Test Vibrator",
		Index:   0,
		Inputs:  make(map[uint32]Input),
		Outputs: make(map[uint32]Output),
	}
	err := d.registerInputs(featuresMsg)
	assert.NoError(t, err)
	assert.Len(t, d.Inputs, 1)

	assert.IsType(t, Battery{}, d.Inputs[3])
	assert.Equal(t, uint32(3), d.Inputs[3].Index())
	assert.Equal(t, "Battery", d.Inputs[3].Description())
	assert.Equal(t, BatteryInput, d.Inputs[3].InputType())
}
