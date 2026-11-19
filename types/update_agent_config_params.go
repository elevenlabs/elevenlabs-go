package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAgentConfigParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateAgentConfigParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateAgentConfigParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAgentConfigParams to string"
	}
	return string(jsonData)
}

func (u *UpdateAgentConfigParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAgentConfigParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAgentConfigParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateAgentConfigParams) MarshalJSON() ([]byte, error) {
	type embed UpdateAgentConfigParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
