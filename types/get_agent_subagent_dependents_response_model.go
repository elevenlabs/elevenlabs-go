package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Agents and branches that use an agent as a subagent.
type GetAgentSubagentDependentsResponseModel struct {
	Agents          []*GetAgentSubagentDependentsResponseModelAgentsItem `json:"agents" xml:"agents" required:"true"`
	Branches        []*DependentBranchInfo                               `json:"branches,omitempty" xml:"branches,omitempty"`
	NextCursor      *string                                              `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                                 `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAgentSubagentDependentsResponseModel) GetAgents() (value []*GetAgentSubagentDependentsResponseModelAgentsItem) {
	if g == nil {
		return
	}
	return g.Agents
}

func (g *GetAgentSubagentDependentsResponseModel) GetBranches() (value []*DependentBranchInfo) {
	if g == nil {
		return
	}
	return g.Branches
}

func (g *GetAgentSubagentDependentsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetAgentSubagentDependentsResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAgentSubagentDependentsResponseModel) SetAgents(agents []*GetAgentSubagentDependentsResponseModelAgentsItem) {
	g.Agents = agents
}

func (g *GetAgentSubagentDependentsResponseModel) SetBranches(branches []*DependentBranchInfo) {
	g.Branches = branches
}

func (g *GetAgentSubagentDependentsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetAgentSubagentDependentsResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetAgentSubagentDependentsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentSubagentDependentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentSubagentDependentsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentSubagentDependentsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetAgentSubagentDependentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentSubagentDependentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
