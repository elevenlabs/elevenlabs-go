package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAgentProcedureParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateAgentProcedureParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateAgentProcedureParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAgentProcedureParams to string"
	}
	return string(jsonData)
}

func (u *UpdateAgentProcedureParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAgentProcedureParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAgentProcedureParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateAgentProcedureParams) MarshalJSON() ([]byte, error) {
	type embed UpdateAgentProcedureParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
