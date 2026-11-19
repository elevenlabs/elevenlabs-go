package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ContextualUpdateInputItem struct {
	type_ string
	// Context for the agent that the user did not write, such as the client's current state.
	Text string `json:"text" xml:"text" required:"true" minLength:"1"`
	// What the update describes. The agent sees only the latest update for each `context_id`.
	ContextId string `json:"context_id" xml:"context_id" required:"true" minLength:"1"`
	// Where the LLM view renders a contextual update.
	Placement       *ContextualUpdatePlacement `json:"placement,omitempty" xml:"placement,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContextualUpdateInputItem) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *ContextualUpdateInputItem) GetContextId() (value string) {
	if c == nil {
		return
	}
	return c.ContextId
}

func (c *ContextualUpdateInputItem) GetPlacement() (value *ContextualUpdatePlacement) {
	if c == nil {
		return
	}
	return c.Placement
}

func (c *ContextualUpdateInputItem) SetText(text string) {
	c.Text = text
}

func (c *ContextualUpdateInputItem) SetContextId(contextId string) {
	c.ContextId = contextId
}

func (c *ContextualUpdateInputItem) SetPlacement(placement *ContextualUpdatePlacement) {
	c.Placement = placement
}

func (c *ContextualUpdateInputItem) Type() string {
	return c.type_
}

func (c *ContextualUpdateInputItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContextualUpdateInputItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContextualUpdateInputItem to string"
	}
	return string(jsonData)
}

func (c *ContextualUpdateInputItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ContextualUpdateInputItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ContextualUpdateInputItem(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "contextual_update" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "contextual_update", *literals.Type)
	}
	c.type_ = "contextual_update"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContextualUpdateInputItem) MarshalJSON() ([]byte, error) {
	type embed ContextualUpdateInputItem
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "contextual_update",
	}
	return json.Marshal(marshaler)
}
