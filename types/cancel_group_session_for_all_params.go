package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Cancel an entire group session and notify every registered participant.
// Destructive -- prefer “cancel_group_session_registration“ for cancelling
// a single attendee.
type CancelGroupSessionForAllParams struct {
	extraProperties map[string]interface{}
}

func (c *CancelGroupSessionForAllParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CancelGroupSessionForAllParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CancelGroupSessionForAllParams to string"
	}
	return string(jsonData)
}

func (c *CancelGroupSessionForAllParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelGroupSessionForAllParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelGroupSessionForAllParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CancelGroupSessionForAllParams) MarshalJSON() ([]byte, error) {
	type embed CancelGroupSessionForAllParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
