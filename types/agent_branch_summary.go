package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentBranchSummary struct {
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
	// ID of the parent branch
	ParentBranchId *string `json:"parent_branch_id,omitempty" xml:"parent_branch_id,omitempty"`
	// Whether a draft exists for the branch
	DraftExists *bool `json:"draft_exists,omitempty" xml:"draft_exists,omitempty"`
	// Unix seconds when the caller's draft on this branch was first created, or null when they have no draft. A draft created before last_committed_at was written against a config the branch has since moved past, so it may not reflect the current one.
	DraftCreatedAt *int `json:"draft_created_at,omitempty" xml:"draft_created_at,omitempty"`
	// Whether the caller's draft on this branch was created before the branch's last commit, meaning it was written against a config the branch has since moved past and may not reflect the current one.
	DraftIsBehindTip *bool `json:"draft_is_behind_tip,omitempty" xml:"draft_is_behind_tip,omitempty"`
	// Number of calls in the last 7 days
	Calls7D *int `json:"calls_7d,omitempty" xml:"calls_7d,omitempty"`
	// Number of commits on this branch not yet on main, relative to their common ancestor. Null if it could not be computed (e.g. no common ancestor, or the branch history exceeds the comparison budget).
	CommitsAhead *int `json:"commits_ahead,omitempty" xml:"commits_ahead,omitempty"`
	// Number of commits on main not yet incorporated into this branch, relative to their common ancestor. Null if it could not be computed (e.g. no common ancestor, or the branch history exceeds the comparison budget).
	CommitsBehind *int `json:"commits_behind,omitempty" xml:"commits_behind,omitempty"`
	// ID of the branch this branch's tip version was merged into, if any
	MergedIntoBranchId *string `json:"merged_into_branch_id,omitempty" xml:"merged_into_branch_id,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *AgentBranchSummary) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentBranchSummary) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentBranchSummary) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentBranchSummary) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AgentBranchSummary) GetCreatedAt() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAt
}

func (a *AgentBranchSummary) GetLastCommittedAt() (value int) {
	if a == nil {
		return
	}
	return a.LastCommittedAt
}

func (a *AgentBranchSummary) GetIsArchived() (value bool) {
	if a == nil {
		return
	}
	return a.IsArchived
}

func (a *AgentBranchSummary) GetProtectionStatus() (value *BranchProtectionStatus) {
	if a == nil {
		return
	}
	return a.ProtectionStatus
}

func (a *AgentBranchSummary) GetAutoRebase() (value *bool) {
	if a == nil {
		return
	}
	return a.AutoRebase
}

func (a *AgentBranchSummary) GetSyncStatus() (value *BranchSyncStatus) {
	if a == nil {
		return
	}
	return a.SyncStatus
}

func (a *AgentBranchSummary) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentBranchSummary) GetCurrentLivePercentage() (value *float64) {
	if a == nil {
		return
	}
	return a.CurrentLivePercentage
}

func (a *AgentBranchSummary) GetParentBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.ParentBranchId
}

func (a *AgentBranchSummary) GetDraftExists() (value *bool) {
	if a == nil {
		return
	}
	return a.DraftExists
}

func (a *AgentBranchSummary) GetDraftCreatedAt() (value *int) {
	if a == nil {
		return
	}
	return a.DraftCreatedAt
}

func (a *AgentBranchSummary) GetDraftIsBehindTip() (value *bool) {
	if a == nil {
		return
	}
	return a.DraftIsBehindTip
}

func (a *AgentBranchSummary) GetCalls7D() (value *int) {
	if a == nil {
		return
	}
	return a.Calls7D
}

func (a *AgentBranchSummary) GetCommitsAhead() (value *int) {
	if a == nil {
		return
	}
	return a.CommitsAhead
}

func (a *AgentBranchSummary) GetCommitsBehind() (value *int) {
	if a == nil {
		return
	}
	return a.CommitsBehind
}

func (a *AgentBranchSummary) GetMergedIntoBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.MergedIntoBranchId
}

func (a *AgentBranchSummary) SetId(id string) {
	a.Id = id
}

func (a *AgentBranchSummary) SetName(name string) {
	a.Name = name
}

func (a *AgentBranchSummary) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentBranchSummary) SetDescription(description string) {
	a.Description = description
}

func (a *AgentBranchSummary) SetCreatedAt(createdAt int) {
	a.CreatedAt = createdAt
}

func (a *AgentBranchSummary) SetLastCommittedAt(lastCommittedAt int) {
	a.LastCommittedAt = lastCommittedAt
}

func (a *AgentBranchSummary) SetIsArchived(isArchived bool) {
	a.IsArchived = isArchived
}

func (a *AgentBranchSummary) SetProtectionStatus(protectionStatus *BranchProtectionStatus) {
	a.ProtectionStatus = protectionStatus
}

func (a *AgentBranchSummary) SetAutoRebase(autoRebase *bool) {
	a.AutoRebase = autoRebase
}

func (a *AgentBranchSummary) SetSyncStatus(syncStatus *BranchSyncStatus) {
	a.SyncStatus = syncStatus
}

func (a *AgentBranchSummary) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentBranchSummary) SetCurrentLivePercentage(currentLivePercentage *float64) {
	a.CurrentLivePercentage = currentLivePercentage
}

func (a *AgentBranchSummary) SetParentBranchId(parentBranchId *string) {
	a.ParentBranchId = parentBranchId
}

func (a *AgentBranchSummary) SetDraftExists(draftExists *bool) {
	a.DraftExists = draftExists
}

func (a *AgentBranchSummary) SetDraftCreatedAt(draftCreatedAt *int) {
	a.DraftCreatedAt = draftCreatedAt
}

func (a *AgentBranchSummary) SetDraftIsBehindTip(draftIsBehindTip *bool) {
	a.DraftIsBehindTip = draftIsBehindTip
}

func (a *AgentBranchSummary) SetCalls7D(calls7D *int) {
	a.Calls7D = calls7D
}

func (a *AgentBranchSummary) SetCommitsAhead(commitsAhead *int) {
	a.CommitsAhead = commitsAhead
}

func (a *AgentBranchSummary) SetCommitsBehind(commitsBehind *int) {
	a.CommitsBehind = commitsBehind
}

func (a *AgentBranchSummary) SetMergedIntoBranchId(mergedIntoBranchId *string) {
	a.MergedIntoBranchId = mergedIntoBranchId
}

func (a *AgentBranchSummary) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentBranchSummary) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentBranchSummary to string"
	}
	return string(jsonData)
}

func (a *AgentBranchSummary) UnmarshalJSON(data []byte) error {
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

func (a *AgentBranchSummary) MarshalJSON() ([]byte, error) {
	type embed AgentBranchSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
