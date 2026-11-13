package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Platform limit with usage computed externally. Example: PVCs use the count_owned_pro_voices_in_billing_group function to compute the usage.
type ComputedUsagePlatformLimit struct {
	Limit           *int `json:"limit,omitempty" xml:"limit,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ComputedUsagePlatformLimit) GetLimit() (value *int) {
	if c == nil {
		return
	}
	return c.Limit
}

func (c *ComputedUsagePlatformLimit) SetLimit(limit *int) {
	c.Limit = limit
}

func (c *ComputedUsagePlatformLimit) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ComputedUsagePlatformLimit) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ComputedUsagePlatformLimit to string"
	}
	return string(jsonData)
}

func (c *ComputedUsagePlatformLimit) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ComputedUsagePlatformLimit) MarshalJSON() ([]byte, error) {
	type embed ComputedUsagePlatformLimit
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
