package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestRunResultBucket struct {
	TestRunIds []string `json:"test_run_ids" xml:"test_run_ids" required:"true"`
	// Short one-line title for this bucket
	Title string `json:"title" xml:"title" required:"true"`
	// Short summary of why the test runs in this bucket passed or failed
	Reason          string        `json:"reason" xml:"reason" required:"true"`
	Status          TestRunStatus `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TestRunResultBucket) GetTestRunIds() (value []string) {
	if t == nil {
		return
	}
	return t.TestRunIds
}

func (t *TestRunResultBucket) GetTitle() (value string) {
	if t == nil {
		return
	}
	return t.Title
}

func (t *TestRunResultBucket) GetReason() (value string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TestRunResultBucket) GetStatus() (value TestRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TestRunResultBucket) SetTestRunIds(testRunIds []string) {
	t.TestRunIds = testRunIds
}

func (t *TestRunResultBucket) SetTitle(title string) {
	t.Title = title
}

func (t *TestRunResultBucket) SetReason(reason string) {
	t.Reason = reason
}

func (t *TestRunResultBucket) SetStatus(status TestRunStatus) {
	t.Status = status
}

func (t *TestRunResultBucket) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestRunResultBucket) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestRunResultBucket to string"
	}
	return string(jsonData)
}

func (t *TestRunResultBucket) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TestRunResultBucket
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TestRunResultBucket(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestRunResultBucket) MarshalJSON() ([]byte, error) {
	type embed TestRunResultBucket
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
