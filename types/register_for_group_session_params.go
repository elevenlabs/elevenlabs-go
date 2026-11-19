package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Register a client for a scheduled group session.
type RegisterForGroupSessionParams struct {
	extraProperties map[string]interface{}
}

func (r *RegisterForGroupSessionParams) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RegisterForGroupSessionParams) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegisterForGroupSessionParams to string"
	}
	return string(jsonData)
}

func (r *RegisterForGroupSessionParams) UnmarshalJSON(data []byte) error {
	type unmarshaler RegisterForGroupSessionParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RegisterForGroupSessionParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "smb_tool_type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RegisterForGroupSessionParams) MarshalJSON() ([]byte, error) {
	type embed RegisterForGroupSessionParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
