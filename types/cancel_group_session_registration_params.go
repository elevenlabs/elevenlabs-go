package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Cancel a single client's registration for a group session.
type CancelGroupSessionRegistrationParams struct {
	extraProperties map[string]interface{}
}

func (c *CancelGroupSessionRegistrationParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CancelGroupSessionRegistrationParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CancelGroupSessionRegistrationParams to string"
	}
	return string(jsonData)
}

func (c *CancelGroupSessionRegistrationParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelGroupSessionRegistrationParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelGroupSessionRegistrationParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CancelGroupSessionRegistrationParams) MarshalJSON() ([]byte, error) {
	type embed CancelGroupSessionRegistrationParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
