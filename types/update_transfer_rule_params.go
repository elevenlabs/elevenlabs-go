package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateTransferRuleParams struct {
	// Whether to offer the post_dial_digits parameter, set from the receptionists' enable_play_keypad_touch_tone_tool config. Digits saved while that is off are dropped when the receptionist is built, so the parameter is hidden rather than accepted and ignored.
	PostDialDigitsEnabled *bool `json:"post_dial_digits_enabled,omitempty" xml:"post_dial_digits_enabled,omitempty"`
	extraProperties       map[string]interface{}
}

func (u *UpdateTransferRuleParams) GetPostDialDigitsEnabled() (value *bool) {
	if u == nil {
		return
	}
	return u.PostDialDigitsEnabled
}

func (u *UpdateTransferRuleParams) SetPostDialDigitsEnabled(postDialDigitsEnabled *bool) {
	u.PostDialDigitsEnabled = postDialDigitsEnabled
}

func (u *UpdateTransferRuleParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateTransferRuleParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateTransferRuleParams to string"
	}
	return string(jsonData)
}

func (u *UpdateTransferRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateTransferRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateTransferRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateTransferRuleParams) MarshalJSON() ([]byte, error) {
	type embed UpdateTransferRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
