package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetBookingPageSettingsParams struct {
	extraProperties map[string]interface{}
}

func (g *GetBookingPageSettingsParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetBookingPageSettingsParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetBookingPageSettingsParams to string"
	}
	return string(jsonData)
}

func (g *GetBookingPageSettingsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetBookingPageSettingsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetBookingPageSettingsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetBookingPageSettingsParams) MarshalJSON() ([]byte, error) {
	type embed GetBookingPageSettingsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
