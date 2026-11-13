package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CloseConnection struct {
	// End the stream with an empty string.
	text            string
	extraProperties map[string]interface{}
}

func (c *CloseConnection) Text() string {
	return c.text
}

func (c *CloseConnection) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CloseConnection) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CloseConnection to string"
	}
	return string(jsonData)
}

func (c *CloseConnection) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CloseConnection
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CloseConnection(tmp)
	var literals struct {
		Text *string `json:"text"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Text != nil && *literals.Text != "" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "", *literals.Text)
	}
	c.text = ""
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "text")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CloseConnection) MarshalJSON() ([]byte, error) {
	type embed CloseConnection
	var marshaler = struct {
		embed
		Text string `json:"text"`
	}{
		embed: embed(*c),
		Text:  "",
	}
	return json.Marshal(marshaler)
}
