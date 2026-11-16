package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CloseContext struct {
	ContextId       string `json:"context_id" xml:"context_id" required:"true"`
	CloseContext    bool   `json:"close_context" xml:"close_context" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CloseContext) GetContextId() (value string) {
	if c == nil {
		return
	}
	return c.ContextId
}

func (c *CloseContext) GetCloseContext() (value bool) {
	if c == nil {
		return
	}
	return c.CloseContext
}

func (c *CloseContext) SetContextId(contextId string) {
	c.ContextId = contextId
}

func (c *CloseContext) SetCloseContext(closeContext bool) {
	c.CloseContext = closeContext
}

func (c *CloseContext) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CloseContext) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CloseContext to string"
	}
	return string(jsonData)
}

func (c *CloseContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CloseContext
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CloseContext(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CloseContext) MarshalJSON() ([]byte, error) {
	type embed CloseContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
