package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RunAgentTestsRequestModel struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// Optional name for the test invocation.
	Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"256" minLength:"1"`
	// List of tests to run on the agent
	Tests []*SingleTestRunRequestModel `json:"tests" xml:"tests" required:"true" minItems:"1" maxItems:"5000"`
	// Configuration overrides to use for testing. If not provided, the agent's default configuration will be used.
	AgentConfigOverride *AdhocAgentConfigOverrideForTestRequestModel `json:"agent_config_override,omitempty" xml:"agent_config_override,omitempty"`
	// ID of the branch to run the tests on. If not provided, the tests will be run on the agent's main branch.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// Number of times to run each test. When greater than 1, results are grouped and summarized.
	RepeatCount *int `json:"repeat_count,omitempty" xml:"repeat_count,omitempty" min:"1" max:"50"`
}

func (r *RunAgentTestsRequestModel) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *RunAgentTestsRequestModel) GetName() (value *string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *RunAgentTestsRequestModel) GetTests() (value []*SingleTestRunRequestModel) {
	if r == nil {
		return
	}
	return r.Tests
}

func (r *RunAgentTestsRequestModel) GetAgentConfigOverride() (value *AdhocAgentConfigOverrideForTestRequestModel) {
	if r == nil {
		return
	}
	return r.AgentConfigOverride
}

func (r *RunAgentTestsRequestModel) GetBranchId() (value *string) {
	if r == nil {
		return
	}
	return r.BranchId
}

func (r *RunAgentTestsRequestModel) GetRepeatCount() (value *int) {
	if r == nil {
		return
	}
	return r.RepeatCount
}

func (r *RunAgentTestsRequestModel) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *RunAgentTestsRequestModel) SetName(name *string) {
	r.Name = name
}

func (r *RunAgentTestsRequestModel) SetTests(tests []*SingleTestRunRequestModel) {
	r.Tests = tests
}

func (r *RunAgentTestsRequestModel) SetAgentConfigOverride(agentConfigOverride *AdhocAgentConfigOverrideForTestRequestModel) {
	r.AgentConfigOverride = agentConfigOverride
}

func (r *RunAgentTestsRequestModel) SetBranchId(branchId *string) {
	r.BranchId = branchId
}

func (r *RunAgentTestsRequestModel) SetRepeatCount(repeatCount *int) {
	r.RepeatCount = repeatCount
}

func (r *RunAgentTestsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r RunAgentTestsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RunAgentTestsRequestModel to string"
	}
	return string(jsonData)
}

func (r *RunAgentTestsRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *RunAgentTestsRequestModel) MarshalJSON() ([]byte, error) {
	type embed RunAgentTestsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
