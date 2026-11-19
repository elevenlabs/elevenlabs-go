package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateTransferRuleParams struct {
	// Whether to offer the post_dial_digits parameter, set from the receptionists' enable_play_keypad_touch_tone_tool config. Digits saved while that is off are dropped when the receptionist is built, so the parameter is hidden rather than accepted and ignored.
	PostDialDigitsEnabled *bool `json:"post_dial_digits_enabled,omitempty" xml:"post_dial_digits_enabled,omitempty"`
	extraProperties       map[string]interface{}
}

func (c *CreateTransferRuleParams) GetPostDialDigitsEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.PostDialDigitsEnabled
}

func (c *CreateTransferRuleParams) SetPostDialDigitsEnabled(postDialDigitsEnabled *bool) {
	c.PostDialDigitsEnabled = postDialDigitsEnabled
}

func (c *CreateTransferRuleParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateTransferRuleParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTransferRuleParams to string"
	}
	return string(jsonData)
}

func (c *CreateTransferRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateTransferRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateTransferRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateTransferRuleParams) MarshalJSON() ([]byte, error) {
	type embed CreateTransferRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
