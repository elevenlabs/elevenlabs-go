package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetScheduleParams struct {
	IncludeLocationFilter *bool `json:"include_location_filter,omitempty" xml:"include_location_filter,omitempty"`
	extraProperties       map[string]interface{}
}

func (g *GetScheduleParams) GetIncludeLocationFilter() (value *bool) {
	if g == nil {
		return
	}
	return g.IncludeLocationFilter
}

func (g *GetScheduleParams) SetIncludeLocationFilter(includeLocationFilter *bool) {
	g.IncludeLocationFilter = includeLocationFilter
}

func (g *GetScheduleParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetScheduleParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetScheduleParams to string"
	}
	return string(jsonData)
}

func (g *GetScheduleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetScheduleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetScheduleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetScheduleParams) MarshalJSON() ([]byte, error) {
	type embed GetScheduleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
