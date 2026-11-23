package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ContextualUpdate struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Context text to inject into the conversation without interrupting.
	Text string `json:"text" xml:"text" required:"true"`
	// Optional client-supplied identifier. When a newer update with the same context_id arrives, the older one supersedes it: the older update is dropped from the LLM view and visually marked as superseded in the rendered transcript.
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContextualUpdate) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ContextualUpdate) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *ContextualUpdate) GetContextId() (value *string) {
	if c == nil {
		return
	}
	return c.ContextId
}

func (c *ContextualUpdate) SetType(type_ *string) {
	c.Type = type_
}

func (c *ContextualUpdate) SetText(text string) {
	c.Text = text
}

func (c *ContextualUpdate) SetContextId(contextId *string) {
	c.ContextId = contextId
}

func (c *ContextualUpdate) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContextualUpdate) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContextualUpdate to string"
	}
	return string(jsonData)
}

func (c *ContextualUpdate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ContextualUpdate
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ContextualUpdate(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContextualUpdate) MarshalJSON() ([]byte, error) {
	type embed ContextualUpdate
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
