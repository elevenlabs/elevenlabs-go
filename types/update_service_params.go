package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Update an existing service's information.
type UpdateServiceParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateServiceParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateServiceParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateServiceParams to string"
	}
	return string(jsonData)
}

func (u *UpdateServiceParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateServiceParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateServiceParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateServiceParams) MarshalJSON() ([]byte, error) {
	type embed UpdateServiceParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
