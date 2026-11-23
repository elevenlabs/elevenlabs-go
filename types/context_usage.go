package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ContextUsage struct {
	ContextUsageEvent *ContextUsageContextUsageEvent `json:"context_usage_event" xml:"context_usage_event" required:"true"`
	Type              *string                        `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ContextUsage) GetContextUsageEvent() (value *ContextUsageContextUsageEvent) {
	if c == nil {
		return
	}
	return c.ContextUsageEvent
}

func (c *ContextUsage) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ContextUsage) SetContextUsageEvent(contextUsageEvent *ContextUsageContextUsageEvent) {
	c.ContextUsageEvent = contextUsageEvent
}

func (c *ContextUsage) SetType(type_ *string) {
	c.Type = type_
}

func (c *ContextUsage) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContextUsage) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContextUsage to string"
	}
	return string(jsonData)
}

func (c *ContextUsage) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ContextUsage
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ContextUsage(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContextUsage) MarshalJSON() ([]byte, error) {
	type embed ContextUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
