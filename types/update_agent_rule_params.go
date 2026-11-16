package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAgentRuleParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateAgentRuleParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateAgentRuleParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAgentRuleParams to string"
	}
	return string(jsonData)
}

func (u *UpdateAgentRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAgentRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAgentRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateAgentRuleParams) MarshalJSON() ([]byte, error) {
	type embed UpdateAgentRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
