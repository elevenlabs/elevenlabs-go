package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Update an existing client's information.
type UpdateClientParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateClientParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateClientParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateClientParams to string"
	}
	return string(jsonData)
}

func (u *UpdateClientParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateClientParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateClientParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateClientParams) MarshalJSON() ([]byte, error) {
	type embed UpdateClientParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
