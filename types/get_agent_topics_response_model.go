package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentTopicsResponseModel struct {
	Topics              []*AgentTopicResponseModel `json:"topics" xml:"topics" required:"true"`
	WindowStartUnixSecs int                        `json:"window_start_unix_secs" xml:"window_start_unix_secs" required:"true"`
	WindowEndUnixSecs   int                        `json:"window_end_unix_secs" xml:"window_end_unix_secs" required:"true"`
	// Number of daily topic-discovery runs the returned metrics were summed over.
	AggregatedRunCount *int    `json:"aggregated_run_count,omitempty" xml:"aggregated_run_count,omitempty"`
	HasMore            *bool   `json:"has_more,omitempty" xml:"has_more,omitempty"`
	NextCursor         *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties    map[string]interface{}
}

func (g *GetAgentTopicsResponseModel) GetTopics() (value []*AgentTopicResponseModel) {
	if g == nil {
		return
	}
	return g.Topics
}

func (g *GetAgentTopicsResponseModel) GetWindowStartUnixSecs() (value int) {
	if g == nil {
		return
	}
	return g.WindowStartUnixSecs
}

func (g *GetAgentTopicsResponseModel) GetWindowEndUnixSecs() (value int) {
	if g == nil {
		return
	}
	return g.WindowEndUnixSecs
}

func (g *GetAgentTopicsResponseModel) GetAggregatedRunCount() (value *int) {
	if g == nil {
		return
	}
	return g.AggregatedRunCount
}

func (g *GetAgentTopicsResponseModel) GetHasMore() (value *bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAgentTopicsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetAgentTopicsResponseModel) SetTopics(topics []*AgentTopicResponseModel) {
	g.Topics = topics
}

func (g *GetAgentTopicsResponseModel) SetWindowStartUnixSecs(windowStartUnixSecs int) {
	g.WindowStartUnixSecs = windowStartUnixSecs
}

func (g *GetAgentTopicsResponseModel) SetWindowEndUnixSecs(windowEndUnixSecs int) {
	g.WindowEndUnixSecs = windowEndUnixSecs
}

func (g *GetAgentTopicsResponseModel) SetAggregatedRunCount(aggregatedRunCount *int) {
	g.AggregatedRunCount = aggregatedRunCount
}

func (g *GetAgentTopicsResponseModel) SetHasMore(hasMore *bool) {
	g.HasMore = hasMore
}

func (g *GetAgentTopicsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetAgentTopicsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentTopicsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentTopicsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentTopicsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAgentTopicsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentTopicsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
