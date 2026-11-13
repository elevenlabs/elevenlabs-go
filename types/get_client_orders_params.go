package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetClientOrdersParams struct {
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	extraProperties  map[string]interface{}
}

func (g *GetClientOrdersParams) GetIncludeCancelled() (value *bool) {
	if g == nil {
		return
	}
	return g.IncludeCancelled
}

func (g *GetClientOrdersParams) SetIncludeCancelled(includeCancelled *bool) {
	g.IncludeCancelled = includeCancelled
}

func (g *GetClientOrdersParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetClientOrdersParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetClientOrdersParams to string"
	}
	return string(jsonData)
}

func (g *GetClientOrdersParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetClientOrdersParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetClientOrdersParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetClientOrdersParams) MarshalJSON() ([]byte, error) {
	type embed GetClientOrdersParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
