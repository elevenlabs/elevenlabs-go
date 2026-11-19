package tokens

import (
	big "math/big"
)

var (
	createSingleUseRequestFieldMaxSessionDurationSeconds = big.NewInt(1 << 0)
)

type CreateSingleUseRequest struct {
	MaxSessionDurationSeconds *int `json:"-" url:"max_session_duration_seconds,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateSingleUseRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetMaxSessionDurationSeconds sets the MaxSessionDurationSeconds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSingleUseRequest) SetMaxSessionDurationSeconds(maxSessionDurationSeconds *int) {
	c.MaxSessionDurationSeconds = maxSessionDurationSeconds
	c.require(createSingleUseRequestFieldMaxSessionDurationSeconds)
}
