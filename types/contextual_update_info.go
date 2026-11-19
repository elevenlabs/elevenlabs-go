package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ContextualUpdateInfo struct {
	// Client-supplied identifier grouping related contextual updates.
	ContextId string `json:"context_id" xml:"context_id" required:"true"`
	// True when this contextual update has been replaced by a newer update with the same context_id.
	IsSuperseded    *bool `json:"is_superseded,omitempty" xml:"is_superseded,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContextualUpdateInfo) GetContextId() (value string) {
	if c == nil {
		return
	}
	return c.ContextId
}

func (c *ContextualUpdateInfo) GetIsSuperseded() (value *bool) {
	if c == nil {
		return
	}
	return c.IsSuperseded
}

func (c *ContextualUpdateInfo) SetContextId(contextId string) {
	c.ContextId = contextId
}

func (c *ContextualUpdateInfo) SetIsSuperseded(isSuperseded *bool) {
	c.IsSuperseded = isSuperseded
}

func (c *ContextualUpdateInfo) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContextualUpdateInfo) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContextualUpdateInfo to string"
	}
	return string(jsonData)
}

func (c *ContextualUpdateInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ContextualUpdateInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ContextualUpdateInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContextualUpdateInfo) MarshalJSON() ([]byte, error) {
	type embed ContextualUpdateInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
