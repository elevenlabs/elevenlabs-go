package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateLocationParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateLocationParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateLocationParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateLocationParams to string"
	}
	return string(jsonData)
}

func (u *UpdateLocationParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateLocationParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateLocationParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateLocationParams) MarshalJSON() ([]byte, error) {
	type embed UpdateLocationParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
