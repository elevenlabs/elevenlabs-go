package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SingleTestRunRequestModel struct {
	// ID of the test to run
	TestId string `json:"test_id" xml:"test_id" required:"true"`
	// ID of the workflow node to run the test on. If not provided, the test will be run on the agent's default workflow node.
	WorkflowNodeId *string `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	// ID of the root folder to run the test on. If not provided, the test will be run on the agent's default folder.
	RootFolderId *string `json:"root_folder_id,omitempty" xml:"root_folder_id,omitempty"`
	// Name of the root folder to run the test on. If not provided, the test will be run on the agent's default folder.
	RootFolderName  *string `json:"root_folder_name,omitempty" xml:"root_folder_name,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SingleTestRunRequestModel) GetTestId() (value string) {
	if s == nil {
		return
	}
	return s.TestId
}

func (s *SingleTestRunRequestModel) GetWorkflowNodeId() (value *string) {
	if s == nil {
		return
	}
	return s.WorkflowNodeId
}

func (s *SingleTestRunRequestModel) GetRootFolderId() (value *string) {
	if s == nil {
		return
	}
	return s.RootFolderId
}

func (s *SingleTestRunRequestModel) GetRootFolderName() (value *string) {
	if s == nil {
		return
	}
	return s.RootFolderName
}

func (s *SingleTestRunRequestModel) SetTestId(testId string) {
	s.TestId = testId
}

func (s *SingleTestRunRequestModel) SetWorkflowNodeId(workflowNodeId *string) {
	s.WorkflowNodeId = workflowNodeId
}

func (s *SingleTestRunRequestModel) SetRootFolderId(rootFolderId *string) {
	s.RootFolderId = rootFolderId
}

func (s *SingleTestRunRequestModel) SetRootFolderName(rootFolderName *string) {
	s.RootFolderName = rootFolderName
}

func (s *SingleTestRunRequestModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SingleTestRunRequestModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SingleTestRunRequestModel to string"
	}
	return string(jsonData)
}

func (s *SingleTestRunRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SingleTestRunRequestModel) MarshalJSON() ([]byte, error) {
	type embed SingleTestRunRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
