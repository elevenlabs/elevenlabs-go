package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddKnowledgeBaseResponseModel struct {
	Id   string `json:"id" xml:"id" required:"true"`
	Name string `json:"name" xml:"name" required:"true"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath      []*KnowledgeBaseFolderPathSegmentSummaryResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AddKnowledgeBaseResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AddKnowledgeBaseResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AddKnowledgeBaseResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if a == nil {
		return
	}
	return a.FolderPath
}

func (a *AddKnowledgeBaseResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AddKnowledgeBaseResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AddKnowledgeBaseResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	a.FolderPath = folderPath
}

func (a *AddKnowledgeBaseResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddKnowledgeBaseResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddKnowledgeBaseResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddKnowledgeBaseResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddKnowledgeBaseResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddKnowledgeBaseResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddKnowledgeBaseResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddKnowledgeBaseResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
