package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Look up a client by their exact phone number.
type GetClientByPhoneParams struct {
	extraProperties map[string]interface{}
}

func (g *GetClientByPhoneParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetClientByPhoneParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetClientByPhoneParams to string"
	}
	return string(jsonData)
}

func (g *GetClientByPhoneParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetClientByPhoneParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetClientByPhoneParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetClientByPhoneParams) MarshalJSON() ([]byte, error) {
	type embed GetClientByPhoneParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
