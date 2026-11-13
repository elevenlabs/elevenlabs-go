package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentBranchBasicInfo struct {
	Id              string `json:"id" xml:"id" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentBranchBasicInfo) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentBranchBasicInfo) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentBranchBasicInfo) SetId(id string) {
	a.Id = id
}

func (a *AgentBranchBasicInfo) SetName(name string) {
	a.Name = name
}

func (a *AgentBranchBasicInfo) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentBranchBasicInfo) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentBranchBasicInfo to string"
	}
	return string(jsonData)
}

func (a *AgentBranchBasicInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentBranchBasicInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentBranchBasicInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentBranchBasicInfo) MarshalJSON() ([]byte, error) {
	type embed AgentBranchBasicInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
