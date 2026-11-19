package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// An MP3 or WAV file played on loop to callers waiting in the agent's concurrency wait queue. Maximum size 40 MB, maximum duration 180 seconds.
	HoldAudioFile io.Reader `json:"hold_audio_file" xml:"hold_audio_file" required:"true"`
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) GetHoldAudioFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.HoldAudioFile
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) SetHoldAudioFile(holdAudioFile io.Reader) {
	b.HoldAudioFile = holdAudioFile
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost to string"
	}
	return string(jsonData)
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost(tmp)
	return nil
}

func (b *BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost) MarshalJSON() ([]byte, error) {
	type embed BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
