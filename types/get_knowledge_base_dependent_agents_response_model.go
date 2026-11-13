package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseDependentAgentsResponseModel struct {
	Agents          []*GetKnowledgeBaseDependentAgentsResponseModelAgentsItem `json:"agents" xml:"agents" required:"true"`
	Branches        []*DependentBranchInfo                                    `json:"branches,omitempty" xml:"branches,omitempty"`
	NextCursor      *string                                                   `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                                      `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) GetAgents() (value []*GetKnowledgeBaseDependentAgentsResponseModelAgentsItem) {
	if g == nil {
		return
	}
	return g.Agents
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) GetBranches() (value []*DependentBranchInfo) {
	if g == nil {
		return
	}
	return g.Branches
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) SetAgents(agents []*GetKnowledgeBaseDependentAgentsResponseModelAgentsItem) {
	g.Agents = agents
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) SetBranches(branches []*DependentBranchInfo) {
	g.Branches = branches
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseDependentAgentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseDependentAgentsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseDependentAgentsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseDependentAgentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseDependentAgentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
