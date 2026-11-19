package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentTestFolderResponseModel struct {
	Id   string `json:"id" xml:"id" required:"true"`
	Name string `json:"name" xml:"name" required:"true"`
	// The path from the root folder to the current folder.
	FolderPath []*AgentTestFolderPathSegmentResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	// The number of direct children (tests and subfolders) in this folder
	ChildrenCount   *int `json:"children_count,omitempty" xml:"children_count,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetAgentTestFolderResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetAgentTestFolderResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetAgentTestFolderResponseModel) GetFolderPath() (value []*AgentTestFolderPathSegmentResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetAgentTestFolderResponseModel) GetChildrenCount() (value *int) {
	if g == nil {
		return
	}
	return g.ChildrenCount
}

func (g *GetAgentTestFolderResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetAgentTestFolderResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetAgentTestFolderResponseModel) SetFolderPath(folderPath []*AgentTestFolderPathSegmentResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetAgentTestFolderResponseModel) SetChildrenCount(childrenCount *int) {
	g.ChildrenCount = childrenCount
}

func (g *GetAgentTestFolderResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentTestFolderResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentTestFolderResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentTestFolderResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetAgentTestFolderResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetAgentTestFolderResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAgentTestFolderResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentTestFolderResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
