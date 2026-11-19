package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnalysisItemAttachmentResultData struct {
	AgentId         string                       `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId        string                       `json:"branch_id" xml:"branch_id" required:"true"`
	Action          AnalysisItemAttachmentAction `json:"action" xml:"action" required:"true"`
	Changed         bool                         `json:"changed" xml:"changed" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AnalysisItemAttachmentResultData) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AnalysisItemAttachmentResultData) GetBranchId() (value string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AnalysisItemAttachmentResultData) GetAction() (value AnalysisItemAttachmentAction) {
	if a == nil {
		return
	}
	return a.Action
}

func (a *AnalysisItemAttachmentResultData) GetChanged() (value bool) {
	if a == nil {
		return
	}
	return a.Changed
}

func (a *AnalysisItemAttachmentResultData) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AnalysisItemAttachmentResultData) SetBranchId(branchId string) {
	a.BranchId = branchId
}

func (a *AnalysisItemAttachmentResultData) SetAction(action AnalysisItemAttachmentAction) {
	a.Action = action
}

func (a *AnalysisItemAttachmentResultData) SetChanged(changed bool) {
	a.Changed = changed
}

func (a *AnalysisItemAttachmentResultData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemAttachmentResultData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemAttachmentResultData to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemAttachmentResultData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnalysisItemAttachmentResultData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnalysisItemAttachmentResultData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisItemAttachmentResultData) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemAttachmentResultData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
