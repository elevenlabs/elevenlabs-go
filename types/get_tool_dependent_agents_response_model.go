package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetToolDependentAgentsResponseModel struct {
	Agents          []*GetToolDependentAgentsResponseModelAgentsItem `json:"agents" xml:"agents" required:"true"`
	Branches        []*DependentBranchInfo                           `json:"branches,omitempty" xml:"branches,omitempty"`
	NextCursor      *string                                          `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                             `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetToolDependentAgentsResponseModel) GetAgents() (value []*GetToolDependentAgentsResponseModelAgentsItem) {
	if g == nil {
		return
	}
	return g.Agents
}

func (g *GetToolDependentAgentsResponseModel) GetBranches() (value []*DependentBranchInfo) {
	if g == nil {
		return
	}
	return g.Branches
}

func (g *GetToolDependentAgentsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetToolDependentAgentsResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetToolDependentAgentsResponseModel) SetAgents(agents []*GetToolDependentAgentsResponseModelAgentsItem) {
	g.Agents = agents
}

func (g *GetToolDependentAgentsResponseModel) SetBranches(branches []*DependentBranchInfo) {
	g.Branches = branches
}

func (g *GetToolDependentAgentsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetToolDependentAgentsResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetToolDependentAgentsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetToolDependentAgentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetToolDependentAgentsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetToolDependentAgentsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetToolDependentAgentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetToolDependentAgentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
