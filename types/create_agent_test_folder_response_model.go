package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentTestFolderResponseModel struct {
	Id              string `json:"id" xml:"id" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateAgentTestFolderResponseModel) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *CreateAgentTestFolderResponseModel) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateAgentTestFolderResponseModel) SetId(id string) {
	c.Id = id
}

func (c *CreateAgentTestFolderResponseModel) SetName(name string) {
	c.Name = name
}

func (c *CreateAgentTestFolderResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentTestFolderResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentTestFolderResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentTestFolderResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentTestFolderResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentTestFolderResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentTestFolderResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentTestFolderResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
