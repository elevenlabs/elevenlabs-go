package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestConditionResultCommonModel struct {
	Result          EvaluationSuccessResult            `json:"result" xml:"result" required:"true"`
	Rationale       *TestConditionRationaleCommonModel `json:"rationale,omitempty" xml:"rationale,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TestConditionResultCommonModel) GetResult() (value EvaluationSuccessResult) {
	if t == nil {
		return
	}
	return t.Result
}

func (t *TestConditionResultCommonModel) GetRationale() (value *TestConditionRationaleCommonModel) {
	if t == nil {
		return
	}
	return t.Rationale
}

func (t *TestConditionResultCommonModel) SetResult(result EvaluationSuccessResult) {
	t.Result = result
}

func (t *TestConditionResultCommonModel) SetRationale(rationale *TestConditionRationaleCommonModel) {
	t.Rationale = rationale
}

func (t *TestConditionResultCommonModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestConditionResultCommonModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestConditionResultCommonModel to string"
	}
	return string(jsonData)
}

func (t *TestConditionResultCommonModel) UnmarshalJSON(data []byte) error {
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

func (t *TestConditionResultCommonModel) MarshalJSON() ([]byte, error) {
	type embed TestConditionResultCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
