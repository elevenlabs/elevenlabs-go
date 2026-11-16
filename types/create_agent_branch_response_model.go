package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentBranchResponseModel struct {
	// ID of the created branch
	CreatedBranchId string `json:"created_branch_id" xml:"created_branch_id" required:"true"`
	// ID of the first version on the created branch
	CreatedVersionId string `json:"created_version_id" xml:"created_version_id" required:"true"`
	extraProperties  map[string]interface{}
}

func (c *CreateAgentBranchResponseModel) GetCreatedBranchId() (value string) {
	if c == nil {
		return
	}
	return c.CreatedBranchId
}

func (c *CreateAgentBranchResponseModel) GetCreatedVersionId() (value string) {
	if c == nil {
		return
	}
	return c.CreatedVersionId
}

func (c *CreateAgentBranchResponseModel) SetCreatedBranchId(createdBranchId string) {
	c.CreatedBranchId = createdBranchId
}

func (c *CreateAgentBranchResponseModel) SetCreatedVersionId(createdVersionId string) {
	c.CreatedVersionId = createdVersionId
}

func (c *CreateAgentBranchResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentBranchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentBranchResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentBranchResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentBranchResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentBranchResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentBranchResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentBranchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
