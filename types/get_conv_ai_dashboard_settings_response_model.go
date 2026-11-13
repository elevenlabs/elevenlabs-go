package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetConvAiDashboardSettingsResponseModel struct {
	Charts          []*GetConvAiDashboardSettingsResponseModelChartsItem `json:"charts,omitempty" xml:"charts,omitempty" maxItems:"4"`
	extraProperties map[string]interface{}
}

func (g *GetConvAiDashboardSettingsResponseModel) GetCharts() (value []*GetConvAiDashboardSettingsResponseModelChartsItem) {
	if g == nil {
		return
	}
	return g.Charts
}

func (g *GetConvAiDashboardSettingsResponseModel) SetCharts(charts []*GetConvAiDashboardSettingsResponseModelChartsItem) {
	g.Charts = charts
}

func (g *GetConvAiDashboardSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConvAiDashboardSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConvAiDashboardSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConvAiDashboardSettingsResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler GetConvAiDashboardSettingsResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetConvAiDashboardSettingsResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetConvAiDashboardSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConvAiDashboardSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
