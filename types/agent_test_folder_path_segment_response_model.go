package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentTestFolderPathSegmentResponseModel struct {
	Id              string  `json:"id" xml:"id" required:"true"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentTestFolderPathSegmentResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentTestFolderPathSegmentResponseModel) GetName() (value *string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentTestFolderPathSegmentResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AgentTestFolderPathSegmentResponseModel) SetName(name *string) {
	a.Name = name
}

func (a *AgentTestFolderPathSegmentResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTestFolderPathSegmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTestFolderPathSegmentResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentTestFolderPathSegmentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentTestFolderPathSegmentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentTestFolderPathSegmentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTestFolderPathSegmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentTestFolderPathSegmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
