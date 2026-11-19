package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentsPageResponseModel struct {
	// A list of agents and their metadata
	Agents []*AgentSummaryResponseModel `json:"agents" xml:"agents" required:"true"`
	// The next cursor to paginate through the agents
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more agents to paginate through
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAgentsPageResponseModel) GetAgents() (value []*AgentSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.Agents
}

func (g *GetAgentsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetAgentsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAgentsPageResponseModel) SetAgents(agents []*AgentSummaryResponseModel) {
	g.Agents = agents
}

func (g *GetAgentsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetAgentsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetAgentsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetAgentsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
