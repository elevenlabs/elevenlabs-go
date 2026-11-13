package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PostAgentHoldAudioResponseModel struct {
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// Custom hold audio played on loop to callers waiting in the agent's queue.
	//
	// Set by uploading a file through the agent hold-audio endpoint. Values sent in
	// agent create or update requests are ignored.
	HoldAudio       *AgentHoldAudioConfig `json:"hold_audio" xml:"hold_audio" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PostAgentHoldAudioResponseModel) GetAgentId() (value string) {
	if p == nil {
		return
	}
	return p.AgentId
}

func (p *PostAgentHoldAudioResponseModel) GetHoldAudio() (value *AgentHoldAudioConfig) {
	if p == nil {
		return
	}
	return p.HoldAudio
}

func (p *PostAgentHoldAudioResponseModel) SetAgentId(agentId string) {
	p.AgentId = agentId
}

func (p *PostAgentHoldAudioResponseModel) SetHoldAudio(holdAudio *AgentHoldAudioConfig) {
	p.HoldAudio = holdAudio
}

func (p *PostAgentHoldAudioResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PostAgentHoldAudioResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostAgentHoldAudioResponseModel to string"
	}
	return string(jsonData)
}

func (p *PostAgentHoldAudioResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PostAgentHoldAudioResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PostAgentHoldAudioResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PostAgentHoldAudioResponseModel) MarshalJSON() ([]byte, error) {
	type embed PostAgentHoldAudioResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
