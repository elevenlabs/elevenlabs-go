package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentVersionParents struct {
	InBranchParentId     *string `json:"in_branch_parent_id,omitempty" xml:"in_branch_parent_id,omitempty"`
	OutOfBranchParentId  *string `json:"out_of_branch_parent_id,omitempty" xml:"out_of_branch_parent_id,omitempty"`
	MergedIntoBranchId   *string `json:"merged_into_branch_id,omitempty" xml:"merged_into_branch_id,omitempty"`
	MergedFromBranchId   *string `json:"merged_from_branch_id,omitempty" xml:"merged_from_branch_id,omitempty"`
	MergedFromVersionId  *string `json:"merged_from_version_id,omitempty" xml:"merged_from_version_id,omitempty"`
	RebasedFromVersionId *string `json:"rebased_from_version_id,omitempty" xml:"rebased_from_version_id,omitempty"`
	extraProperties      map[string]interface{}
}

func (a *AgentVersionParents) GetInBranchParentId() (value *string) {
	if a == nil {
		return
	}
	return a.InBranchParentId
}

func (a *AgentVersionParents) GetOutOfBranchParentId() (value *string) {
	if a == nil {
		return
	}
	return a.OutOfBranchParentId
}

func (a *AgentVersionParents) GetMergedIntoBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.MergedIntoBranchId
}

func (a *AgentVersionParents) GetMergedFromBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.MergedFromBranchId
}

func (a *AgentVersionParents) GetMergedFromVersionId() (value *string) {
	if a == nil {
		return
	}
	return a.MergedFromVersionId
}

func (a *AgentVersionParents) GetRebasedFromVersionId() (value *string) {
	if a == nil {
		return
	}
	return a.RebasedFromVersionId
}

func (a *AgentVersionParents) SetInBranchParentId(inBranchParentId *string) {
	a.InBranchParentId = inBranchParentId
}

func (a *AgentVersionParents) SetOutOfBranchParentId(outOfBranchParentId *string) {
	a.OutOfBranchParentId = outOfBranchParentId
}

func (a *AgentVersionParents) SetMergedIntoBranchId(mergedIntoBranchId *string) {
	a.MergedIntoBranchId = mergedIntoBranchId
}

func (a *AgentVersionParents) SetMergedFromBranchId(mergedFromBranchId *string) {
	a.MergedFromBranchId = mergedFromBranchId
}

func (a *AgentVersionParents) SetMergedFromVersionId(mergedFromVersionId *string) {
	a.MergedFromVersionId = mergedFromVersionId
}

func (a *AgentVersionParents) SetRebasedFromVersionId(rebasedFromVersionId *string) {
	a.RebasedFromVersionId = rebasedFromVersionId
}

func (a *AgentVersionParents) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentVersionParents) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentVersionParents to string"
	}
	return string(jsonData)
}

func (a *AgentVersionParents) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentVersionParents) MarshalJSON() ([]byte, error) {
	type embed AgentVersionParents
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
