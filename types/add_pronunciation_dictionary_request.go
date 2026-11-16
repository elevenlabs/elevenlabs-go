package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddPronunciationDictionaryRequest struct {
	// The name of the pronunciation dictionary, used for identification only.
	Name string `json:"name" xml:"name" required:"true"`
	// A lexicon .pls file which we will use to initialize the project with.
	File *string `json:"file,omitempty" xml:"file,omitempty"`
	// A description of the pronunciation dictionary, used for identification only.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no access.
	WorkspaceAccess *CreateFromFilePronunciationDictionariesRequestWorkspaceAccess `json:"workspace_access,omitempty" xml:"workspace_access,omitempty"`
}

func (a *AddPronunciationDictionaryRequest) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AddPronunciationDictionaryRequest) GetFile() (value *string) {
	if a == nil {
		return
	}
	return a.File
}

func (a *AddPronunciationDictionaryRequest) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AddPronunciationDictionaryRequest) GetWorkspaceAccess() (value *CreateFromFilePronunciationDictionariesRequestWorkspaceAccess) {
	if a == nil {
		return
	}
	return a.WorkspaceAccess
}

func (a *AddPronunciationDictionaryRequest) SetName(name string) {
	a.Name = name
}

func (a *AddPronunciationDictionaryRequest) SetFile(file *string) {
	a.File = file
}

func (a *AddPronunciationDictionaryRequest) SetDescription(description *string) {
	a.Description = description
}

func (a *AddPronunciationDictionaryRequest) SetWorkspaceAccess(workspaceAccess *CreateFromFilePronunciationDictionariesRequestWorkspaceAccess) {
	a.WorkspaceAccess = workspaceAccess
}

func (a *AddPronunciationDictionaryRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddPronunciationDictionaryRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddPronunciationDictionaryRequest to string"
	}
	return string(jsonData)
}

func (a *AddPronunciationDictionaryRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

func (a *AddPronunciationDictionaryRequest) MarshalJSON() ([]byte, error) {
	type embed AddPronunciationDictionaryRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
