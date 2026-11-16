package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentLinkResponseModel struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The token data for the agent
	Token           *ConversationTokenResponseModel `json:"token,omitempty" xml:"token,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetAgentLinkResponseModel) GetAgentId() (value string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetAgentLinkResponseModel) GetToken() (value *ConversationTokenResponseModel) {
	if g == nil {
		return
	}
	return g.Token
}

func (g *GetAgentLinkResponseModel) SetAgentId(agentId string) {
	g.AgentId = agentId
}

func (g *GetAgentLinkResponseModel) SetToken(token *ConversationTokenResponseModel) {
	g.Token = token
}

func (g *GetAgentLinkResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentLinkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentLinkResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentLinkResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetAgentLinkResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentLinkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
