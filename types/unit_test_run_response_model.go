package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UnitTestRunResponseModel struct {
	TestRunId        string                                            `json:"test_run_id" xml:"test_run_id" required:"true"`
	TestInfo         *UnitTestRunResponseModelTestInfo                 `json:"test_info,omitempty" xml:"test_info,omitempty"`
	TestInvocationId string                                            `json:"test_invocation_id" xml:"test_invocation_id" required:"true"`
	AgentId          string                                            `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId         *string                                           `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	VersionId        *string                                           `json:"version_id,omitempty" xml:"version_id,omitempty"`
	RanAgainstDraft  *bool                                             `json:"ran_against_draft,omitempty" xml:"ran_against_draft,omitempty"`
	WorkflowNodeId   *string                                           `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	Status           TestRunStatus                                     `json:"status" xml:"status" required:"true"`
	AgentResponses   []*ConversationHistoryTranscriptCommonModelOutput `json:"agent_responses,omitempty" xml:"agent_responses,omitempty"`
	// Guardrails observed during this run, including blocks before an agent response.
	TriggeredGuardrails []*TriggeredGuardrailCommonModel `json:"triggered_guardrails,omitempty" xml:"triggered_guardrails,omitempty"`
	TestId              string                           `json:"test_id" xml:"test_id" required:"true"`
	TestName            *string                          `json:"test_name,omitempty" xml:"test_name,omitempty"`
	ConditionResult     *TestConditionResultCommonModel  `json:"condition_result,omitempty" xml:"condition_result,omitempty"`
	LastUpdatedAtUnix   *int                             `json:"last_updated_at_unix,omitempty" xml:"last_updated_at_unix,omitempty"`
	Metadata            *TestRunMetadata                 `json:"metadata,omitempty" xml:"metadata,omitempty"`
	RootFolderId        *string                          `json:"root_folder_id,omitempty" xml:"root_folder_id,omitempty"`
	RootFolderName      *string                          `json:"root_folder_name,omitempty" xml:"root_folder_name,omitempty"`
	Environment         *string                          `json:"environment,omitempty" xml:"environment,omitempty"`
	// Credits billed for this test run. None for runs created before cost tracking.
	CreditsUsed *int `json:"credits_used,omitempty" xml:"credits_used,omitempty"`
	// Finalized billing and provider-usage breakdown for this test run.
	Charging        *ConversationChargingCommonModel `json:"charging,omitempty" xml:"charging,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UnitTestRunResponseModel) GetTestRunId() (value string) {
	if u == nil {
		return
	}
	return u.TestRunId
}

func (u *UnitTestRunResponseModel) GetTestInfo() (value *UnitTestRunResponseModelTestInfo) {
	if u == nil {
		return
	}
	return u.TestInfo
}

func (u *UnitTestRunResponseModel) GetTestInvocationId() (value string) {
	if u == nil {
		return
	}
	return u.TestInvocationId
}

func (u *UnitTestRunResponseModel) GetAgentId() (value string) {
	if u == nil {
		return
	}
	return u.AgentId
}

func (u *UnitTestRunResponseModel) GetBranchId() (value *string) {
	if u == nil {
		return
	}
	return u.BranchId
}

func (u *UnitTestRunResponseModel) GetVersionId() (value *string) {
	if u == nil {
		return
	}
	return u.VersionId
}

func (u *UnitTestRunResponseModel) GetRanAgainstDraft() (value *bool) {
	if u == nil {
		return
	}
	return u.RanAgainstDraft
}

func (u *UnitTestRunResponseModel) GetWorkflowNodeId() (value *string) {
	if u == nil {
		return
	}
	return u.WorkflowNodeId
}

func (u *UnitTestRunResponseModel) GetStatus() (value TestRunStatus) {
	if u == nil {
		return
	}
	return u.Status
}

func (u *UnitTestRunResponseModel) GetAgentResponses() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if u == nil {
		return
	}
	return u.AgentResponses
}

func (u *UnitTestRunResponseModel) GetTriggeredGuardrails() (value []*TriggeredGuardrailCommonModel) {
	if u == nil {
		return
	}
	return u.TriggeredGuardrails
}

func (u *UnitTestRunResponseModel) GetTestId() (value string) {
	if u == nil {
		return
	}
	return u.TestId
}

func (u *UnitTestRunResponseModel) GetTestName() (value *string) {
	if u == nil {
		return
	}
	return u.TestName
}

func (u *UnitTestRunResponseModel) GetConditionResult() (value *TestConditionResultCommonModel) {
	if u == nil {
		return
	}
	return u.ConditionResult
}

func (u *UnitTestRunResponseModel) GetLastUpdatedAtUnix() (value *int) {
	if u == nil {
		return
	}
	return u.LastUpdatedAtUnix
}

func (u *UnitTestRunResponseModel) GetMetadata() (value *TestRunMetadata) {
	if u == nil {
		return
	}
	return u.Metadata
}

func (u *UnitTestRunResponseModel) GetRootFolderId() (value *string) {
	if u == nil {
		return
	}
	return u.RootFolderId
}

func (u *UnitTestRunResponseModel) GetRootFolderName() (value *string) {
	if u == nil {
		return
	}
	return u.RootFolderName
}

func (u *UnitTestRunResponseModel) GetEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.Environment
}

func (u *UnitTestRunResponseModel) GetCreditsUsed() (value *int) {
	if u == nil {
		return
	}
	return u.CreditsUsed
}

func (u *UnitTestRunResponseModel) GetCharging() (value *ConversationChargingCommonModel) {
	if u == nil {
		return
	}
	return u.Charging
}

func (u *UnitTestRunResponseModel) SetTestRunId(testRunId string) {
	u.TestRunId = testRunId
}

func (u *UnitTestRunResponseModel) SetTestInfo(testInfo *UnitTestRunResponseModelTestInfo) {
	u.TestInfo = testInfo
}

func (u *UnitTestRunResponseModel) SetTestInvocationId(testInvocationId string) {
	u.TestInvocationId = testInvocationId
}

func (u *UnitTestRunResponseModel) SetAgentId(agentId string) {
	u.AgentId = agentId
}

func (u *UnitTestRunResponseModel) SetBranchId(branchId *string) {
	u.BranchId = branchId
}

func (u *UnitTestRunResponseModel) SetVersionId(versionId *string) {
	u.VersionId = versionId
}

func (u *UnitTestRunResponseModel) SetRanAgainstDraft(ranAgainstDraft *bool) {
	u.RanAgainstDraft = ranAgainstDraft
}

func (u *UnitTestRunResponseModel) SetWorkflowNodeId(workflowNodeId *string) {
	u.WorkflowNodeId = workflowNodeId
}

func (u *UnitTestRunResponseModel) SetStatus(status TestRunStatus) {
	u.Status = status
}

func (u *UnitTestRunResponseModel) SetAgentResponses(agentResponses []*ConversationHistoryTranscriptCommonModelOutput) {
	u.AgentResponses = agentResponses
}

func (u *UnitTestRunResponseModel) SetTriggeredGuardrails(triggeredGuardrails []*TriggeredGuardrailCommonModel) {
	u.TriggeredGuardrails = triggeredGuardrails
}

func (u *UnitTestRunResponseModel) SetTestId(testId string) {
	u.TestId = testId
}

func (u *UnitTestRunResponseModel) SetTestName(testName *string) {
	u.TestName = testName
}

func (u *UnitTestRunResponseModel) SetConditionResult(conditionResult *TestConditionResultCommonModel) {
	u.ConditionResult = conditionResult
}

func (u *UnitTestRunResponseModel) SetLastUpdatedAtUnix(lastUpdatedAtUnix *int) {
	u.LastUpdatedAtUnix = lastUpdatedAtUnix
}

func (u *UnitTestRunResponseModel) SetMetadata(metadata *TestRunMetadata) {
	u.Metadata = metadata
}

func (u *UnitTestRunResponseModel) SetRootFolderId(rootFolderId *string) {
	u.RootFolderId = rootFolderId
}

func (u *UnitTestRunResponseModel) SetRootFolderName(rootFolderName *string) {
	u.RootFolderName = rootFolderName
}

func (u *UnitTestRunResponseModel) SetEnvironment(environment *string) {
	u.Environment = environment
}

func (u *UnitTestRunResponseModel) SetCreditsUsed(creditsUsed *int) {
	u.CreditsUsed = creditsUsed
}

func (u *UnitTestRunResponseModel) SetCharging(charging *ConversationChargingCommonModel) {
	u.Charging = charging
}

func (u *UnitTestRunResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestRunResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestRunResponseModel to string"
	}
	return string(jsonData)
}

func (u *UnitTestRunResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestRunResponseModel) MarshalJSON() ([]byte, error) {
	type embed UnitTestRunResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
