package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnalysisItemAttachmentDestination struct {
	AgentId         string `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId        string `json:"branch_id" xml:"branch_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AnalysisItemAttachmentDestination) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AnalysisItemAttachmentDestination) GetBranchId() (value string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AnalysisItemAttachmentDestination) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AnalysisItemAttachmentDestination) SetBranchId(branchId string) {
	a.BranchId = branchId
}

func (a *AnalysisItemAttachmentDestination) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemAttachmentDestination) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemAttachmentDestination to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemAttachmentDestination) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnalysisItemAttachmentDestination
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnalysisItemAttachmentDestination(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisItemAttachmentDestination) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemAttachmentDestination
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
