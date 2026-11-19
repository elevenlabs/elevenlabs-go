package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestRunMetadata struct {
	WorkspaceId     string                   `json:"workspace_id" xml:"workspace_id" required:"true"`
	TestName        string                   `json:"test_name" xml:"test_name" required:"true"`
	RanByUserEmail  string                   `json:"ran_by_user_email" xml:"ran_by_user_email" required:"true"`
	TestType        *TestRunMetadataTestType `json:"test_type,omitempty" xml:"test_type,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TestRunMetadata) GetWorkspaceId() (value string) {
	if t == nil {
		return
	}
	return t.WorkspaceId
}

func (t *TestRunMetadata) GetTestName() (value string) {
	if t == nil {
		return
	}
	return t.TestName
}

func (t *TestRunMetadata) GetRanByUserEmail() (value string) {
	if t == nil {
		return
	}
	return t.RanByUserEmail
}

func (t *TestRunMetadata) GetTestType() (value *TestRunMetadataTestType) {
	if t == nil {
		return
	}
	return t.TestType
}

func (t *TestRunMetadata) SetWorkspaceId(workspaceId string) {
	t.WorkspaceId = workspaceId
}

func (t *TestRunMetadata) SetTestName(testName string) {
	t.TestName = testName
}

func (t *TestRunMetadata) SetRanByUserEmail(ranByUserEmail string) {
	t.RanByUserEmail = ranByUserEmail
}

func (t *TestRunMetadata) SetTestType(testType *TestRunMetadataTestType) {
	t.TestType = testType
}

func (t *TestRunMetadata) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestRunMetadata) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestRunMetadata to string"
	}
	return string(jsonData)
}

func (t *TestRunMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TestRunMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TestRunMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestRunMetadata) MarshalJSON() ([]byte, error) {
	type embed TestRunMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
