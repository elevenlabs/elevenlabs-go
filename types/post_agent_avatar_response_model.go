package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PostAgentAvatarResponseModel struct {
	AgentId         string  `json:"agent_id" xml:"agent_id" required:"true"`
	AvatarUrl       *string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PostAgentAvatarResponseModel) GetAgentId() (value string) {
	if p == nil {
		return
	}
	return p.AgentId
}

func (p *PostAgentAvatarResponseModel) GetAvatarUrl() (value *string) {
	if p == nil {
		return
	}
	return p.AvatarUrl
}

func (p *PostAgentAvatarResponseModel) SetAgentId(agentId string) {
	p.AgentId = agentId
}

func (p *PostAgentAvatarResponseModel) SetAvatarUrl(avatarUrl *string) {
	p.AvatarUrl = avatarUrl
}

func (p *PostAgentAvatarResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PostAgentAvatarResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostAgentAvatarResponseModel to string"
	}
	return string(jsonData)
}

func (p *PostAgentAvatarResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PostAgentAvatarResponseModel) MarshalJSON() ([]byte, error) {
	type embed PostAgentAvatarResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
