package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetBookingSlugStatusParams struct {
	extraProperties map[string]interface{}
}

func (g *GetBookingSlugStatusParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetBookingSlugStatusParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetBookingSlugStatusParams to string"
	}
	return string(jsonData)
}

func (g *GetBookingSlugStatusParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetBookingSlugStatusParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetBookingSlugStatusParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetBookingSlugStatusParams) MarshalJSON() ([]byte, error) {
	type embed GetBookingSlugStatusParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
