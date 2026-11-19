package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagListingEntryResponseModel struct {
	// ID of the document or folder.
	Id string `json:"id" xml:"id" required:"true"`
	// Name of the document or folder.
	Name string `json:"name" xml:"name" required:"true"`
	// Whether the entry is a folder.
	IsFolder        bool `json:"is_folder" xml:"is_folder" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) GetIsFolder() (value bool) {
	if a == nil {
		return
	}
	return a.IsFolder
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) SetIsFolder(isFolder bool) {
	a.IsFolder = isFolder
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagListingEntryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagListingEntryResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentKnowledgeBaseRagListingEntryResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentKnowledgeBaseRagListingEntryResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentKnowledgeBaseRagListingEntryResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagListingEntryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
