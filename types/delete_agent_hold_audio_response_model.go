package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DeleteAgentHoldAudioResponseModel struct {
	AgentId         string `json:"agent_id" xml:"agent_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DeleteAgentHoldAudioResponseModel) GetAgentId() (value string) {
	if d == nil {
		return
	}
	return d.AgentId
}

func (d *DeleteAgentHoldAudioResponseModel) SetAgentId(agentId string) {
	d.AgentId = agentId
}

func (d *DeleteAgentHoldAudioResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteAgentHoldAudioResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteAgentHoldAudioResponseModel to string"
	}
	return string(jsonData)
}

func (d *DeleteAgentHoldAudioResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DeleteAgentHoldAudioResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DeleteAgentHoldAudioResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteAgentHoldAudioResponseModel) MarshalJSON() ([]byte, error) {
	type embed DeleteAgentHoldAudioResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
