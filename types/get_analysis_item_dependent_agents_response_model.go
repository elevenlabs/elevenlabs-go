package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAnalysisItemDependentAgentsResponseModel struct {
	Agents          []*GetAnalysisItemDependentAgentsResponseModelAgentsItem `json:"agents" xml:"agents" required:"true"`
	Branches        []*DependentBranchInfo                                   `json:"branches,omitempty" xml:"branches,omitempty"`
	NextCursor      *string                                                  `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                                     `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAnalysisItemDependentAgentsResponseModel) GetAgents() (value []*GetAnalysisItemDependentAgentsResponseModelAgentsItem) {
	if g == nil {
		return
	}
	return g.Agents
}

func (g *GetAnalysisItemDependentAgentsResponseModel) GetBranches() (value []*DependentBranchInfo) {
	if g == nil {
		return
	}
	return g.Branches
}

func (g *GetAnalysisItemDependentAgentsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetAnalysisItemDependentAgentsResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAnalysisItemDependentAgentsResponseModel) SetAgents(agents []*GetAnalysisItemDependentAgentsResponseModelAgentsItem) {
	g.Agents = agents
}

func (g *GetAnalysisItemDependentAgentsResponseModel) SetBranches(branches []*DependentBranchInfo) {
	g.Branches = branches
}

func (g *GetAnalysisItemDependentAgentsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetAnalysisItemDependentAgentsResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetAnalysisItemDependentAgentsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAnalysisItemDependentAgentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAnalysisItemDependentAgentsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAnalysisItemDependentAgentsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetAnalysisItemDependentAgentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAnalysisItemDependentAgentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
