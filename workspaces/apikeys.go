package workspaces

import (
	big "math/big"
)

var (
	disableApiKeysRequestFieldApiKeyName = big.NewInt(1 << 0)
)

type DisableApiKeysRequest struct {
	ApiKeyName string `json:"-" url:"api_key_name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DisableApiKeysRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetApiKeyName sets the ApiKeyName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DisableApiKeysRequest) SetApiKeyName(apiKeyName string) {
	d.ApiKeyName = apiKeyName
	d.require(disableApiKeysRequestFieldApiKeyName)
}
