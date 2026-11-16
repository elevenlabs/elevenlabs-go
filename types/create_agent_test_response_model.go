package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentTestResponseModel struct {
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateAgentTestResponseModel) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *CreateAgentTestResponseModel) SetId(id string) {
	c.Id = id
}

func (c *CreateAgentTestResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentTestResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentTestResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentTestResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentTestResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentTestResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentTestResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentTestResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
