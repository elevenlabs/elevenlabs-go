package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentResponseModel struct {
	// ID of the created agent
	AgentId         string `json:"agent_id" xml:"agent_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateAgentResponseModel) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateAgentResponseModel) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *CreateAgentResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
