package analytics

import (
	big "math/big"
)

var (
	getLiveCountRequestFieldAgentId  = big.NewInt(1 << 0)
	getLiveCountRequestFieldAgentIds = big.NewInt(1 << 1)
)

type GetLiveCountRequest struct {
	AgentId  *string   `json:"-" url:"agent_id,omitempty"`
	AgentIds []*string `json:"-" url:"agent_ids,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetLiveCountRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLiveCountRequest) SetAgentId(agentId *string) {
	g.AgentId = agentId
	g.require(getLiveCountRequestFieldAgentId)
}

// SetAgentIds sets the AgentIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLiveCountRequest) SetAgentIds(agentIds []*string) {
	g.AgentIds = agentIds
	g.require(getLiveCountRequestFieldAgentIds)
}
