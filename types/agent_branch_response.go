package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentBranchResponse struct {
	Id               string                  `json:"id" xml:"id" required:"true"`
	Name             string                  `json:"name" xml:"name" required:"true"`
	AgentId          string                  `json:"agent_id" xml:"agent_id" required:"true"`
	Description      string                  `json:"description" xml:"description" required:"true"`
	CreatedAt        int                     `json:"created_at" xml:"created_at" required:"true"`
	LastCommittedAt  int                     `json:"last_committed_at" xml:"last_committed_at" required:"true"`
	IsArchived       bool                    `json:"is_archived" xml:"is_archived" required:"true"`
	ProtectionStatus *BranchProtectionStatus `json:"protection_status,omitempty" xml:"protection_status,omitempty"`
	AutoRebase       *bool                   `json:"auto_rebase,omitempty" xml:"auto_rebase,omitempty"`
	// Set while the branch is kept in sync with another branch (Main today), pulling in its new versions automatically. Null when it is not synced.
	SyncStatus *BranchSyncStatus `json:"sync_status,omitempty" xml:"sync_status,omitempty"`
	// Access information for the branch
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// Percentage of traffic live on the branch
	CurrentLivePercentage *float64 `json:"current_live_percentage,omitempty" xml:"current_live_percentage,omitempty"`
	// Parent branch of the branch
	ParentBranch *AgentBranchBasicInfo `json:"parent_branch,omitempty" xml:"parent_branch,omitempty"`
	// Most recent versions on the branch
	MostRecentVersions []*AgentVersionMetadata `json:"most_recent_versions,omitempty" xml:"most_recent_versions,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *AgentBranchResponse) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentBranchResponse) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentBranchResponse) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentBranchResponse) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AgentBranchResponse) GetCreatedAt() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAt
}

func (a *AgentBranchResponse) GetLastCommittedAt() (value int) {
	if a == nil {
		return
	}
	return a.LastCommittedAt
}

func (a *AgentBranchResponse) GetIsArchived() (value bool) {
	if a == nil {
		return
	}
	return a.IsArchived
}

func (a *AgentBranchResponse) GetProtectionStatus() (value *BranchProtectionStatus) {
	if a == nil {
		return
	}
	return a.ProtectionStatus
}

func (a *AgentBranchResponse) GetAutoRebase() (value *bool) {
	if a == nil {
		return
	}
	return a.AutoRebase
}

func (a *AgentBranchResponse) GetSyncStatus() (value *BranchSyncStatus) {
	if a == nil {
		return
	}
	return a.SyncStatus
}

func (a *AgentBranchResponse) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentBranchResponse) GetCurrentLivePercentage() (value *float64) {
	if a == nil {
		return
	}
	return a.CurrentLivePercentage
}

func (a *AgentBranchResponse) GetParentBranch() (value *AgentBranchBasicInfo) {
	if a == nil {
		return
	}
	return a.ParentBranch
}

func (a *AgentBranchResponse) GetMostRecentVersions() (value []*AgentVersionMetadata) {
	if a == nil {
		return
	}
	return a.MostRecentVersions
}

func (a *AgentBranchResponse) SetId(id string) {
	a.Id = id
}

func (a *AgentBranchResponse) SetName(name string) {
	a.Name = name
}

func (a *AgentBranchResponse) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentBranchResponse) SetDescription(description string) {
	a.Description = description
}

func (a *AgentBranchResponse) SetCreatedAt(createdAt int) {
	a.CreatedAt = createdAt
}

func (a *AgentBranchResponse) SetLastCommittedAt(lastCommittedAt int) {
	a.LastCommittedAt = lastCommittedAt
}

func (a *AgentBranchResponse) SetIsArchived(isArchived bool) {
	a.IsArchived = isArchived
}

func (a *AgentBranchResponse) SetProtectionStatus(protectionStatus *BranchProtectionStatus) {
	a.ProtectionStatus = protectionStatus
}

func (a *AgentBranchResponse) SetAutoRebase(autoRebase *bool) {
	a.AutoRebase = autoRebase
}

func (a *AgentBranchResponse) SetSyncStatus(syncStatus *BranchSyncStatus) {
	a.SyncStatus = syncStatus
}

func (a *AgentBranchResponse) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentBranchResponse) SetCurrentLivePercentage(currentLivePercentage *float64) {
	a.CurrentLivePercentage = currentLivePercentage
}

func (a *AgentBranchResponse) SetParentBranch(parentBranch *AgentBranchBasicInfo) {
	a.ParentBranch = parentBranch
}

func (a *AgentBranchResponse) SetMostRecentVersions(mostRecentVersions []*AgentVersionMetadata) {
	a.MostRecentVersions = mostRecentVersions
}

func (a *AgentBranchResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentBranchResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentBranchResponse to string"
	}
	return string(jsonData)
}

func (a *AgentBranchResponse) UnmarshalJSON(data []byte) error {
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

func (a *AgentBranchResponse) MarshalJSON() ([]byte, error) {
	type embed AgentBranchResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
