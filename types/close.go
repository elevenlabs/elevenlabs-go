package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Close struct {
	type_           string
	extraProperties map[string]interface{}
}

func (c *Close) Type() string {
	return c.type_
}

func (c *Close) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c Close) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Close to string"
	}
	return string(jsonData)
}

func (c *Close) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias Close
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = Close(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "close" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "close", *literals.Type)
	}
	c.type_ = "close"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *Close) MarshalJSON() ([]byte, error) {
	type embed Close
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "close",
	}
	return json.Marshal(marshaler)
}
