package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestRunResultSummary struct {
	TestId          string                 `json:"test_id" xml:"test_id" required:"true"`
	TestName        string                 `json:"test_name" xml:"test_name" required:"true"`
	WorkflowNodeId  *string                `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	Buckets         []*TestRunResultBucket `json:"buckets" xml:"buckets" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TestRunResultSummary) GetTestId() (value string) {
	if t == nil {
		return
	}
	return t.TestId
}

func (t *TestRunResultSummary) GetTestName() (value string) {
	if t == nil {
		return
	}
	return t.TestName
}

func (t *TestRunResultSummary) GetWorkflowNodeId() (value *string) {
	if t == nil {
		return
	}
	return t.WorkflowNodeId
}

func (t *TestRunResultSummary) GetBuckets() (value []*TestRunResultBucket) {
	if t == nil {
		return
	}
	return t.Buckets
}

func (t *TestRunResultSummary) SetTestId(testId string) {
	t.TestId = testId
}

func (t *TestRunResultSummary) SetTestName(testName string) {
	t.TestName = testName
}

func (t *TestRunResultSummary) SetWorkflowNodeId(workflowNodeId *string) {
	t.WorkflowNodeId = workflowNodeId
}

func (t *TestRunResultSummary) SetBuckets(buckets []*TestRunResultBucket) {
	t.Buckets = buckets
}

func (t *TestRunResultSummary) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestRunResultSummary) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestRunResultSummary to string"
	}
	return string(jsonData)
}

func (t *TestRunResultSummary) UnmarshalJSON(data []byte) error {
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

func (t *TestRunResultSummary) MarshalJSON() ([]byte, error) {
	type embed TestRunResultSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
