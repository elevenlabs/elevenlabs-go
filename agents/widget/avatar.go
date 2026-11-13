package widget

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// An image file to be used as the agent's avatar.
	AvatarFile io.Reader `json:"avatar_file" xml:"avatar_file" required:"true"`
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) GetAvatarFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.AvatarFile
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) SetAvatarFile(avatarFile io.Reader) {
	b.AvatarFile = avatarFile
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost to string"
	}
	return string(jsonData)
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost(tmp)
	return nil
}

func (b *BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost) MarshalJSON() ([]byte, error) {
	type embed BodyPostAgentAvatarV1ConvaiAgentsAgentIdAvatarPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
