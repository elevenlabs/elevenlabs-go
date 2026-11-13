package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Change the seat count of an existing group session registration.
type UpdateGroupSessionSeatsParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateGroupSessionSeatsParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateGroupSessionSeatsParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateGroupSessionSeatsParams to string"
	}
	return string(jsonData)
}

func (u *UpdateGroupSessionSeatsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateGroupSessionSeatsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateGroupSessionSeatsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateGroupSessionSeatsParams) MarshalJSON() ([]byte, error) {
	type embed UpdateGroupSessionSeatsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
