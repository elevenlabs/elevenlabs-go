package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentMergeProposalResponseModel struct {
	// ID of the created merge_proposal
	CreatedMergeProposalId string `json:"created_merge_proposal_id" xml:"created_merge_proposal_id" required:"true"`
	extraProperties        map[string]interface{}
}

func (c *CreateAgentMergeProposalResponseModel) GetCreatedMergeProposalId() (value string) {
	if c == nil {
		return
	}
	return c.CreatedMergeProposalId
}

func (c *CreateAgentMergeProposalResponseModel) SetCreatedMergeProposalId(createdMergeProposalId string) {
	c.CreatedMergeProposalId = createdMergeProposalId
}

func (c *CreateAgentMergeProposalResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentMergeProposalResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentMergeProposalResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentMergeProposalResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentMergeProposalResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentMergeProposalResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentMergeProposalResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentMergeProposalResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
