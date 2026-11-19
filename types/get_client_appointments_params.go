package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetClientAppointmentsParams struct {
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	extraProperties  map[string]interface{}
}

func (g *GetClientAppointmentsParams) GetIncludeCancelled() (value *bool) {
	if g == nil {
		return
	}
	return g.IncludeCancelled
}

func (g *GetClientAppointmentsParams) SetIncludeCancelled(includeCancelled *bool) {
	g.IncludeCancelled = includeCancelled
}

func (g *GetClientAppointmentsParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetClientAppointmentsParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetClientAppointmentsParams to string"
	}
	return string(jsonData)
}

func (g *GetClientAppointmentsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetClientAppointmentsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetClientAppointmentsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetClientAppointmentsParams) MarshalJSON() ([]byte, error) {
	type embed GetClientAppointmentsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
