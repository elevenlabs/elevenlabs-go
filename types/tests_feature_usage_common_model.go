package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TestsFeatureUsageCommonModel struct {
	Enabled                       *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TestsRanAfterLastModification *bool `json:"tests_ran_after_last_modification,omitempty" xml:"tests_ran_after_last_modification,omitempty"`
	TestsRanInLast7Days           *bool `json:"tests_ran_in_last_7_days,omitempty" xml:"tests_ran_in_last_7_days,omitempty"`
	extraProperties               map[string]interface{}
}

func (t *TestsFeatureUsageCommonModel) GetEnabled() (value *bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

func (t *TestsFeatureUsageCommonModel) GetTestsRanAfterLastModification() (value *bool) {
	if t == nil {
		return
	}
	return t.TestsRanAfterLastModification
}

func (t *TestsFeatureUsageCommonModel) GetTestsRanInLast7Days() (value *bool) {
	if t == nil {
		return
	}
	return t.TestsRanInLast7Days
}

func (t *TestsFeatureUsageCommonModel) SetEnabled(enabled *bool) {
	t.Enabled = enabled
}

func (t *TestsFeatureUsageCommonModel) SetTestsRanAfterLastModification(testsRanAfterLastModification *bool) {
	t.TestsRanAfterLastModification = testsRanAfterLastModification
}

func (t *TestsFeatureUsageCommonModel) SetTestsRanInLast7Days(testsRanInLast7Days *bool) {
	t.TestsRanInLast7Days = testsRanInLast7Days
}

func (t *TestsFeatureUsageCommonModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestsFeatureUsageCommonModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestsFeatureUsageCommonModel to string"
	}
	return string(jsonData)
}

func (t *TestsFeatureUsageCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler TestsFeatureUsageCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestsFeatureUsageCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestsFeatureUsageCommonModel) MarshalJSON() ([]byte, error) {
	type embed TestsFeatureUsageCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
