package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetTestSuiteInvocationResponseModel struct {
	Id string `json:"id" xml:"id" required:"true"`
	// The user-provided name of the test invocation, if one was set.
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	AgentId         *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	BranchId        *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	VersionId       *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	RanAgainstDraft *bool   `json:"ran_against_draft,omitempty" xml:"ran_against_draft,omitempty"`
	CreatedAt       *int    `json:"created_at,omitempty" xml:"created_at,omitempty"`
	FolderId        *string `json:"folder_id,omitempty" xml:"folder_id,omitempty"`
	RepeatCount     *int    `json:"repeat_count,omitempty" xml:"repeat_count,omitempty"`
	Cancelled       *bool   `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// None when repeat_count==1 (no bucketing). Otherwise tracks bucketing lifecycle.
	BucketingStatus *BucketingStatus            `json:"bucketing_status,omitempty" xml:"bucketing_status,omitempty"`
	ResultGroups    []*TestRunResultSummary     `json:"result_groups,omitempty" xml:"result_groups,omitempty"`
	TestRuns        []*UnitTestRunResponseModel `json:"test_runs" xml:"test_runs" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetTestSuiteInvocationResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetTestSuiteInvocationResponseModel) GetName() (value *string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetTestSuiteInvocationResponseModel) GetAgentId() (value *string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetTestSuiteInvocationResponseModel) GetBranchId() (value *string) {
	if g == nil {
		return
	}
	return g.BranchId
}

func (g *GetTestSuiteInvocationResponseModel) GetVersionId() (value *string) {
	if g == nil {
		return
	}
	return g.VersionId
}

func (g *GetTestSuiteInvocationResponseModel) GetRanAgainstDraft() (value *bool) {
	if g == nil {
		return
	}
	return g.RanAgainstDraft
}

func (g *GetTestSuiteInvocationResponseModel) GetCreatedAt() (value *int) {
	if g == nil {
		return
	}
	return g.CreatedAt
}

func (g *GetTestSuiteInvocationResponseModel) GetFolderId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderId
}

func (g *GetTestSuiteInvocationResponseModel) GetRepeatCount() (value *int) {
	if g == nil {
		return
	}
	return g.RepeatCount
}

func (g *GetTestSuiteInvocationResponseModel) GetCancelled() (value *bool) {
	if g == nil {
		return
	}
	return g.Cancelled
}

func (g *GetTestSuiteInvocationResponseModel) GetBucketingStatus() (value *BucketingStatus) {
	if g == nil {
		return
	}
	return g.BucketingStatus
}

func (g *GetTestSuiteInvocationResponseModel) GetResultGroups() (value []*TestRunResultSummary) {
	if g == nil {
		return
	}
	return g.ResultGroups
}

func (g *GetTestSuiteInvocationResponseModel) GetTestRuns() (value []*UnitTestRunResponseModel) {
	if g == nil {
		return
	}
	return g.TestRuns
}

func (g *GetTestSuiteInvocationResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetTestSuiteInvocationResponseModel) SetName(name *string) {
	g.Name = name
}

func (g *GetTestSuiteInvocationResponseModel) SetAgentId(agentId *string) {
	g.AgentId = agentId
}

func (g *GetTestSuiteInvocationResponseModel) SetBranchId(branchId *string) {
	g.BranchId = branchId
}

func (g *GetTestSuiteInvocationResponseModel) SetVersionId(versionId *string) {
	g.VersionId = versionId
}

func (g *GetTestSuiteInvocationResponseModel) SetRanAgainstDraft(ranAgainstDraft *bool) {
	g.RanAgainstDraft = ranAgainstDraft
}

func (g *GetTestSuiteInvocationResponseModel) SetCreatedAt(createdAt *int) {
	g.CreatedAt = createdAt
}

func (g *GetTestSuiteInvocationResponseModel) SetFolderId(folderId *string) {
	g.FolderId = folderId
}

func (g *GetTestSuiteInvocationResponseModel) SetRepeatCount(repeatCount *int) {
	g.RepeatCount = repeatCount
}

func (g *GetTestSuiteInvocationResponseModel) SetCancelled(cancelled *bool) {
	g.Cancelled = cancelled
}

func (g *GetTestSuiteInvocationResponseModel) SetBucketingStatus(bucketingStatus *BucketingStatus) {
	g.BucketingStatus = bucketingStatus
}

func (g *GetTestSuiteInvocationResponseModel) SetResultGroups(resultGroups []*TestRunResultSummary) {
	g.ResultGroups = resultGroups
}

func (g *GetTestSuiteInvocationResponseModel) SetTestRuns(testRuns []*UnitTestRunResponseModel) {
	g.TestRuns = testRuns
}

func (g *GetTestSuiteInvocationResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetTestSuiteInvocationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTestSuiteInvocationResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetTestSuiteInvocationResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetTestSuiteInvocationResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetTestSuiteInvocationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
