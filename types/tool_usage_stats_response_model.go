package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolUsageStatsResponseModel struct {
	// The total number of calls to the tool
	TotalCalls      *int    `json:"total_calls,omitempty" xml:"total_calls,omitempty"`
	AvgLatencySecs  float64 `json:"avg_latency_secs" xml:"avg_latency_secs" required:"true"`
	extraProperties map[string]interface{}
}

func (t *ToolUsageStatsResponseModel) GetTotalCalls() (value *int) {
	if t == nil {
		return
	}
	return t.TotalCalls
}

func (t *ToolUsageStatsResponseModel) GetAvgLatencySecs() (value float64) {
	if t == nil {
		return
	}
	return t.AvgLatencySecs
}

func (t *ToolUsageStatsResponseModel) SetTotalCalls(totalCalls *int) {
	t.TotalCalls = totalCalls
}

func (t *ToolUsageStatsResponseModel) SetAvgLatencySecs(avgLatencySecs float64) {
	t.AvgLatencySecs = avgLatencySecs
}

func (t *ToolUsageStatsResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolUsageStatsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolUsageStatsResponseModel to string"
	}
	return string(jsonData)
}

func (t *ToolUsageStatsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias ToolUsageStatsResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = ToolUsageStatsResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolUsageStatsResponseModel) MarshalJSON() ([]byte, error) {
	type embed ToolUsageStatsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
