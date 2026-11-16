package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentBranchInfo struct {
	AgentId         string `json:"agent_id" xml:"agent_id" required:"true"`
	AgentName       string `json:"agent_name" xml:"agent_name" required:"true"`
	BranchId        string `json:"branch_id" xml:"branch_id" required:"true"`
	BranchName      string `json:"branch_name" xml:"branch_name" required:"true"`
	IsMain          bool   `json:"is_main" xml:"is_main" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DependentBranchInfo) GetAgentId() (value string) {
	if d == nil {
		return
	}
	return d.AgentId
}

func (d *DependentBranchInfo) GetAgentName() (value string) {
	if d == nil {
		return
	}
	return d.AgentName
}

func (d *DependentBranchInfo) GetBranchId() (value string) {
	if d == nil {
		return
	}
	return d.BranchId
}

func (d *DependentBranchInfo) GetBranchName() (value string) {
	if d == nil {
		return
	}
	return d.BranchName
}

func (d *DependentBranchInfo) GetIsMain() (value bool) {
	if d == nil {
		return
	}
	return d.IsMain
}

func (d *DependentBranchInfo) SetAgentId(agentId string) {
	d.AgentId = agentId
}

func (d *DependentBranchInfo) SetAgentName(agentName string) {
	d.AgentName = agentName
}

func (d *DependentBranchInfo) SetBranchId(branchId string) {
	d.BranchId = branchId
}

func (d *DependentBranchInfo) SetBranchName(branchName string) {
	d.BranchName = branchName
}

func (d *DependentBranchInfo) SetIsMain(isMain bool) {
	d.IsMain = isMain
}

func (d *DependentBranchInfo) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentBranchInfo) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentBranchInfo to string"
	}
	return string(jsonData)
}

func (d *DependentBranchInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentBranchInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentBranchInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentBranchInfo) MarshalJSON() ([]byte, error) {
	type embed DependentBranchInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
