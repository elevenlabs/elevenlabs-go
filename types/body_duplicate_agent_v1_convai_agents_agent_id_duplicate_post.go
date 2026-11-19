package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// A name to make the agent easier to find
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) SetName(name *string) {
	b.Name = name
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost to string"
	}
	return string(jsonData)
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost) MarshalJSON() ([]byte, error) {
	type embed BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
