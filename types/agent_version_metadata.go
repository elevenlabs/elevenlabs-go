package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentVersionMetadata struct {
	Id                 string               `json:"id" xml:"id" required:"true"`
	AgentId            string               `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId           string               `json:"branch_id" xml:"branch_id" required:"true"`
	VersionDescription string               `json:"version_description" xml:"version_description" required:"true"`
	SeqNoInBranch      int                  `json:"seq_no_in_branch" xml:"seq_no_in_branch" required:"true"`
	TimeCommittedSecs  int                  `json:"time_committed_secs" xml:"time_committed_secs" required:"true"`
	Parents            *AgentVersionParents `json:"parents" xml:"parents" required:"true"`
	AccessInfo         *ResourceAccessInfo  `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// For a merge into the main branch, the people who published the merged changes on the source branch (access_info is whoever ran the merge). Null when not recorded.
	MergedAuthors   []*ResourceAccessInfo `json:"merged_authors,omitempty" xml:"merged_authors,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentVersionMetadata) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentVersionMetadata) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentVersionMetadata) GetBranchId() (value string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentVersionMetadata) GetVersionDescription() (value string) {
	if a == nil {
		return
	}
	return a.VersionDescription
}

func (a *AgentVersionMetadata) GetSeqNoInBranch() (value int) {
	if a == nil {
		return
	}
	return a.SeqNoInBranch
}

func (a *AgentVersionMetadata) GetTimeCommittedSecs() (value int) {
	if a == nil {
		return
	}
	return a.TimeCommittedSecs
}

func (a *AgentVersionMetadata) GetParents() (value *AgentVersionParents) {
	if a == nil {
		return
	}
	return a.Parents
}

func (a *AgentVersionMetadata) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentVersionMetadata) GetMergedAuthors() (value []*ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.MergedAuthors
}

func (a *AgentVersionMetadata) SetId(id string) {
	a.Id = id
}

func (a *AgentVersionMetadata) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentVersionMetadata) SetBranchId(branchId string) {
	a.BranchId = branchId
}

func (a *AgentVersionMetadata) SetVersionDescription(versionDescription string) {
	a.VersionDescription = versionDescription
}

func (a *AgentVersionMetadata) SetSeqNoInBranch(seqNoInBranch int) {
	a.SeqNoInBranch = seqNoInBranch
}

func (a *AgentVersionMetadata) SetTimeCommittedSecs(timeCommittedSecs int) {
	a.TimeCommittedSecs = timeCommittedSecs
}

func (a *AgentVersionMetadata) SetParents(parents *AgentVersionParents) {
	a.Parents = parents
}

func (a *AgentVersionMetadata) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentVersionMetadata) SetMergedAuthors(mergedAuthors []*ResourceAccessInfo) {
	a.MergedAuthors = mergedAuthors
}

func (a *AgentVersionMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentVersionMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentVersionMetadata to string"
	}
	return string(jsonData)
}

func (a *AgentVersionMetadata) UnmarshalJSON(data []byte) error {
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

func (a *AgentVersionMetadata) MarshalJSON() ([]byte, error) {
	type embed AgentVersionMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
