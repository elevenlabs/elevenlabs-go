package tests

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type UpdateTestSuiteInvocationRequestModel struct {
	// The id of a test invocation. This is returned when tests are run.
	TestInvocationId string `json:"-" url:"-"`
	// New name for the test invocation. Set to null to restore the generated title.
	Name *string `json:"name" xml:"name" required:"true" maxLength:"256" minLength:"1"`
}

func (u *UpdateTestSuiteInvocationRequestModel) GetTestInvocationId() (value string) {
	if u == nil {
		return
	}
	return u.TestInvocationId
}

func (u *UpdateTestSuiteInvocationRequestModel) GetName() (value *string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateTestSuiteInvocationRequestModel) SetTestInvocationId(testInvocationId string) {
	u.TestInvocationId = testInvocationId
}

func (u *UpdateTestSuiteInvocationRequestModel) SetName(name *string) {
	u.Name = name
}

func (u *UpdateTestSuiteInvocationRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateTestSuiteInvocationRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateTestSuiteInvocationRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateTestSuiteInvocationRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateTestSuiteInvocationRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateTestSuiteInvocationRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

type ResubmitTestsRequestModel struct {
	// The id of a test invocation. This is returned when tests are run.
	TestInvocationId string `json:"-" url:"-"`
	// List of test run IDs to resubmit
	TestRunIds []string `json:"test_run_ids" xml:"test_run_ids" required:"true" minItems:"1" maxItems:"5000"`
	// Configuration overrides to use for testing. If not provided, the agent's default configuration will be used.
	AgentConfigOverride *types.AdhocAgentConfigOverrideForTestRequestModel `json:"agent_config_override,omitempty" xml:"agent_config_override,omitempty"`
	// Agent ID to resubmit tests for
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// ID of the branch to run the tests on. If not provided, the tests will be run on the agent's main branch.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
}

func (r *ResubmitTestsRequestModel) GetTestInvocationId() (value string) {
	if r == nil {
		return
	}
	return r.TestInvocationId
}

func (r *ResubmitTestsRequestModel) GetTestRunIds() (value []string) {
	if r == nil {
		return
	}
	return r.TestRunIds
}

func (r *ResubmitTestsRequestModel) GetAgentConfigOverride() (value *types.AdhocAgentConfigOverrideForTestRequestModel) {
	if r == nil {
		return
	}
	return r.AgentConfigOverride
}

func (r *ResubmitTestsRequestModel) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ResubmitTestsRequestModel) GetBranchId() (value *string) {
	if r == nil {
		return
	}
	return r.BranchId
}

func (r *ResubmitTestsRequestModel) SetTestInvocationId(testInvocationId string) {
	r.TestInvocationId = testInvocationId
}

func (r *ResubmitTestsRequestModel) SetTestRunIds(testRunIds []string) {
	r.TestRunIds = testRunIds
}

func (r *ResubmitTestsRequestModel) SetAgentConfigOverride(agentConfigOverride *types.AdhocAgentConfigOverrideForTestRequestModel) {
	r.AgentConfigOverride = agentConfigOverride
}

func (r *ResubmitTestsRequestModel) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ResubmitTestsRequestModel) SetBranchId(branchId *string) {
	r.BranchId = branchId
}

func (r *ResubmitTestsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ResubmitTestsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResubmitTestsRequestModel to string"
	}
	return string(jsonData)
}

func (r *ResubmitTestsRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ResubmitTestsRequestModel) MarshalJSON() ([]byte, error) {
	type embed ResubmitTestsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listInvocationsRequestFieldAgentId         = big.NewInt(1 << 0)
	listInvocationsRequestFieldBranchId        = big.NewInt(1 << 1)
	listInvocationsRequestFieldCreatedByUserId = big.NewInt(1 << 2)
	listInvocationsRequestFieldTestId          = big.NewInt(1 << 3)
	listInvocationsRequestFieldStartTime       = big.NewInt(1 << 4)
	listInvocationsRequestFieldEndTime         = big.NewInt(1 << 5)
	listInvocationsRequestFieldVersionId       = big.NewInt(1 << 6)
	listInvocationsRequestFieldPageSize        = big.NewInt(1 << 7)
	listInvocationsRequestFieldSearch          = big.NewInt(1 << 8)
	listInvocationsRequestFieldCursor          = big.NewInt(1 << 9)
)

type ListInvocationsRequest struct {
	AgentId         *string `json:"-" url:"agent_id,omitempty"`
	BranchId        *string `json:"-" url:"branch_id,omitempty"`
	CreatedByUserId *string `json:"-" url:"created_by_user_id,omitempty"`
	TestId          *string `json:"-" url:"test_id,omitempty"`
	StartTime       *int    `json:"-" url:"start_time,omitempty"`
	EndTime         *int    `json:"-" url:"end_time,omitempty"`
	VersionId       *string `json:"-" url:"version_id,omitempty"`
	PageSize        *int    `json:"-" url:"page_size,omitempty"`
	Search          *string `json:"-" url:"search,omitempty"`
	Cursor          *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListInvocationsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listInvocationsRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetBranchId(branchId *string) {
	l.BranchId = branchId
	l.require(listInvocationsRequestFieldBranchId)
}

// SetCreatedByUserId sets the CreatedByUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetCreatedByUserId(createdByUserId *string) {
	l.CreatedByUserId = createdByUserId
	l.require(listInvocationsRequestFieldCreatedByUserId)
}

// SetTestId sets the TestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetTestId(testId *string) {
	l.TestId = testId
	l.require(listInvocationsRequestFieldTestId)
}

// SetStartTime sets the StartTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetStartTime(startTime *int) {
	l.StartTime = startTime
	l.require(listInvocationsRequestFieldStartTime)
}

// SetEndTime sets the EndTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetEndTime(endTime *int) {
	l.EndTime = endTime
	l.require(listInvocationsRequestFieldEndTime)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetVersionId(versionId *string) {
	l.VersionId = versionId
	l.require(listInvocationsRequestFieldVersionId)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listInvocationsRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listInvocationsRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvocationsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listInvocationsRequestFieldCursor)
}
