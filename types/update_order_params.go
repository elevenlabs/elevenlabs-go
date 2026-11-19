package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateOrderParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateOrderParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateOrderParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateOrderParams to string"
	}
	return string(jsonData)
}

func (u *UpdateOrderParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateOrderParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateOrderParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateOrderParams) MarshalJSON() ([]byte, error) {
	type embed UpdateOrderParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
