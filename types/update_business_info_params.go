package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateBusinessInfoParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateBusinessInfoParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateBusinessInfoParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateBusinessInfoParams to string"
	}
	return string(jsonData)
}

func (u *UpdateBusinessInfoParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateBusinessInfoParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateBusinessInfoParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateBusinessInfoParams) MarshalJSON() ([]byte, error) {
	type embed UpdateBusinessInfoParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
