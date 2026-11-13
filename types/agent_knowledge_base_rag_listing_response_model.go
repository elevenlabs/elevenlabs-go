package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagListingResponseModel struct {
	// Entries the agent would be shown, folders first and then by name. Large folders are cut to the number of entries the agent is shown.
	Entries []*AgentKnowledgeBaseRagListingEntryResponseModel `json:"entries" xml:"entries" required:"true"`
	// Number of entries in the listed folder, including any not returned.
	TotalCount      int `json:"total_count" xml:"total_count" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentKnowledgeBaseRagListingResponseModel) GetEntries() (value []*AgentKnowledgeBaseRagListingEntryResponseModel) {
	if a == nil {
		return
	}
	return a.Entries
}

func (a *AgentKnowledgeBaseRagListingResponseModel) GetTotalCount() (value int) {
	if a == nil {
		return
	}
	return a.TotalCount
}

func (a *AgentKnowledgeBaseRagListingResponseModel) SetEntries(entries []*AgentKnowledgeBaseRagListingEntryResponseModel) {
	a.Entries = entries
}

func (a *AgentKnowledgeBaseRagListingResponseModel) SetTotalCount(totalCount int) {
	a.TotalCount = totalCount
}

func (a *AgentKnowledgeBaseRagListingResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagListingResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagListingResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagListingResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentKnowledgeBaseRagListingResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentKnowledgeBaseRagListingResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentKnowledgeBaseRagListingResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagListingResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
