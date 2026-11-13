package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttachedTestModel struct {
	TestId          string  `json:"test_id" xml:"test_id" required:"true"`
	WorkflowNodeId  *string `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AttachedTestModel) GetTestId() (value string) {
	if a == nil {
		return
	}
	return a.TestId
}

func (a *AttachedTestModel) GetWorkflowNodeId() (value *string) {
	if a == nil {
		return
	}
	return a.WorkflowNodeId
}

func (a *AttachedTestModel) SetTestId(testId string) {
	a.TestId = testId
}

func (a *AttachedTestModel) SetWorkflowNodeId(workflowNodeId *string) {
	a.WorkflowNodeId = workflowNodeId
}

func (a *AttachedTestModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttachedTestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachedTestModel to string"
	}
	return string(jsonData)
}

func (a *AttachedTestModel) UnmarshalJSON(data []byte) error {
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

func (a *AttachedTestModel) MarshalJSON() ([]byte, error) {
	type embed AttachedTestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
