package samples

import (
	big "math/big"
)

var (
	getAudioRequestFieldRemoveBackgroundNoise = big.NewInt(1 << 0)
)

type GetAudioRequest struct {
	RemoveBackgroundNoise *bool `json:"-" url:"remove_background_noise,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAudioRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetRemoveBackgroundNoise sets the RemoveBackgroundNoise field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudioRequest) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	g.RemoveBackgroundNoise = removeBackgroundNoise
	g.require(getAudioRequestFieldRemoveBackgroundNoise)
}
