package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The owner's spoken answer to changes the assistant read back to them.
//
// Runs in the agent runtime rather than Reception: it needs the conversation's
// chat to check the owner actually answered after hearing the changes.
type ConfirmChangeParams struct {
	extraProperties map[string]interface{}
}

func (c *ConfirmChangeParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConfirmChangeParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConfirmChangeParams to string"
	}
	return string(jsonData)
}

func (c *ConfirmChangeParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ConfirmChangeParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConfirmChangeParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConfirmChangeParams) MarshalJSON() ([]byte, error) {
	type embed ConfirmChangeParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
