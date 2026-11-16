package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Get a summary of key business analytics for a time period.
type GetAnalyticsSummaryParams struct {
	extraProperties map[string]interface{}
}

func (g *GetAnalyticsSummaryParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAnalyticsSummaryParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAnalyticsSummaryParams to string"
	}
	return string(jsonData)
}

func (g *GetAnalyticsSummaryParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAnalyticsSummaryParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAnalyticsSummaryParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAnalyticsSummaryParams) MarshalJSON() ([]byte, error) {
	type embed GetAnalyticsSummaryParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
