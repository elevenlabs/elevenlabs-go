package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestInvocationSummaryResponseModel struct {
	// The ID of the test invocation
	Id string `json:"id" xml:"id" required:"true"`
	// The user-provided name of the test invocation, if one was set.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the agent this test invocation belongs to
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// The ID of the branch this test invocation was run on
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// The ID of the agent version this test invocation ran against. For draft or config-override runs this is the version those uncommitted changes were applied on top of. None only for runs recorded before this field existed.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// Whether the run included uncommitted changes (a saved draft or an ad-hoc config override) layered on top of version_id.
	RanAgainstDraft *bool `json:"ran_against_draft,omitempty" xml:"ran_against_draft,omitempty"`
	// Whether the test runs in this invocation did not all execute against the same version, which happens when a subset of runs was resubmitted after the original run. When true, version_id describes the most recent resubmit rather than every run.
	RunsDivergedFromVersion *bool `json:"runs_diverged_from_version,omitempty" xml:"runs_diverged_from_version,omitempty"`
	// Creation time of the test invocation in unix seconds
	CreatedAtUnixSecs int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	// Number of test runs in this invocation
	TestRunCount int `json:"test_run_count" xml:"test_run_count" required:"true"`
	// Number of test runs that passed
	PassedCount int `json:"passed_count" xml:"passed_count" required:"true"`
	// Number of test runs that failed
	FailedCount int `json:"failed_count" xml:"failed_count" required:"true"`
	// Number of test runs that are pending
	PendingCount int `json:"pending_count" xml:"pending_count" required:"true"`
	// Number of test runs that were cancelled
	CancelledCount int `json:"cancelled_count" xml:"cancelled_count" required:"true"`
	// Display title of the test invocation: the user-provided name when set, otherwise the folder name, single test name, or test count.
	Title string `json:"title" xml:"title" required:"true"`
	// The access information of the test invocation
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// Number of times each test was repeated in this invocation
	RepeatCount *int `json:"repeat_count,omitempty" xml:"repeat_count,omitempty"`
	// Whether this test invocation was cancelled
	Cancelled *bool `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// Total credits billed across test runs in this invocation. None when no run has cost data.
	CreditsUsed *int `json:"credits_used,omitempty" xml:"credits_used,omitempty"`
	// Total USD price across test runs in this invocation. None when no run has price data.
	TotalPrice      *float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TestInvocationSummaryResponseModel) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TestInvocationSummaryResponseModel) GetName() (value *string) {
	if t == nil {
		return
	}
	return t.Name
}

func (t *TestInvocationSummaryResponseModel) GetAgentId() (value *string) {
	if t == nil {
		return
	}
	return t.AgentId
}

func (t *TestInvocationSummaryResponseModel) GetBranchId() (value *string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TestInvocationSummaryResponseModel) GetVersionId() (value *string) {
	if t == nil {
		return
	}
	return t.VersionId
}

func (t *TestInvocationSummaryResponseModel) GetRanAgainstDraft() (value *bool) {
	if t == nil {
		return
	}
	return t.RanAgainstDraft
}

func (t *TestInvocationSummaryResponseModel) GetRunsDivergedFromVersion() (value *bool) {
	if t == nil {
		return
	}
	return t.RunsDivergedFromVersion
}

func (t *TestInvocationSummaryResponseModel) GetCreatedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.CreatedAtUnixSecs
}

func (t *TestInvocationSummaryResponseModel) GetTestRunCount() (value int) {
	if t == nil {
		return
	}
	return t.TestRunCount
}

func (t *TestInvocationSummaryResponseModel) GetPassedCount() (value int) {
	if t == nil {
		return
	}
	return t.PassedCount
}

func (t *TestInvocationSummaryResponseModel) GetFailedCount() (value int) {
	if t == nil {
		return
	}
	return t.FailedCount
}

func (t *TestInvocationSummaryResponseModel) GetPendingCount() (value int) {
	if t == nil {
		return
	}
	return t.PendingCount
}

func (t *TestInvocationSummaryResponseModel) GetCancelledCount() (value int) {
	if t == nil {
		return
	}
	return t.CancelledCount
}

func (t *TestInvocationSummaryResponseModel) GetTitle() (value string) {
	if t == nil {
		return
	}
	return t.Title
}

func (t *TestInvocationSummaryResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if t == nil {
		return
	}
	return t.AccessInfo
}

func (t *TestInvocationSummaryResponseModel) GetRepeatCount() (value *int) {
	if t == nil {
		return
	}
	return t.RepeatCount
}

func (t *TestInvocationSummaryResponseModel) GetCancelled() (value *bool) {
	if t == nil {
		return
	}
	return t.Cancelled
}

func (t *TestInvocationSummaryResponseModel) GetCreditsUsed() (value *int) {
	if t == nil {
		return
	}
	return t.CreditsUsed
}

func (t *TestInvocationSummaryResponseModel) GetTotalPrice() (value *float64) {
	if t == nil {
		return
	}
	return t.TotalPrice
}

func (t *TestInvocationSummaryResponseModel) SetId(id string) {
	t.Id = id
}

func (t *TestInvocationSummaryResponseModel) SetName(name *string) {
	t.Name = name
}

func (t *TestInvocationSummaryResponseModel) SetAgentId(agentId *string) {
	t.AgentId = agentId
}

func (t *TestInvocationSummaryResponseModel) SetBranchId(branchId *string) {
	t.BranchId = branchId
}

func (t *TestInvocationSummaryResponseModel) SetVersionId(versionId *string) {
	t.VersionId = versionId
}

func (t *TestInvocationSummaryResponseModel) SetRanAgainstDraft(ranAgainstDraft *bool) {
	t.RanAgainstDraft = ranAgainstDraft
}

func (t *TestInvocationSummaryResponseModel) SetRunsDivergedFromVersion(runsDivergedFromVersion *bool) {
	t.RunsDivergedFromVersion = runsDivergedFromVersion
}

func (t *TestInvocationSummaryResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	t.CreatedAtUnixSecs = createdAtUnixSecs
}

func (t *TestInvocationSummaryResponseModel) SetTestRunCount(testRunCount int) {
	t.TestRunCount = testRunCount
}

func (t *TestInvocationSummaryResponseModel) SetPassedCount(passedCount int) {
	t.PassedCount = passedCount
}

func (t *TestInvocationSummaryResponseModel) SetFailedCount(failedCount int) {
	t.FailedCount = failedCount
}

func (t *TestInvocationSummaryResponseModel) SetPendingCount(pendingCount int) {
	t.PendingCount = pendingCount
}

func (t *TestInvocationSummaryResponseModel) SetCancelledCount(cancelledCount int) {
	t.CancelledCount = cancelledCount
}

func (t *TestInvocationSummaryResponseModel) SetTitle(title string) {
	t.Title = title
}

func (t *TestInvocationSummaryResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	t.AccessInfo = accessInfo
}

func (t *TestInvocationSummaryResponseModel) SetRepeatCount(repeatCount *int) {
	t.RepeatCount = repeatCount
}

func (t *TestInvocationSummaryResponseModel) SetCancelled(cancelled *bool) {
	t.Cancelled = cancelled
}

func (t *TestInvocationSummaryResponseModel) SetCreditsUsed(creditsUsed *int) {
	t.CreditsUsed = creditsUsed
}

func (t *TestInvocationSummaryResponseModel) SetTotalPrice(totalPrice *float64) {
	t.TotalPrice = totalPrice
}

func (t *TestInvocationSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestInvocationSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestInvocationSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (t *TestInvocationSummaryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestInvocationSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed TestInvocationSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
