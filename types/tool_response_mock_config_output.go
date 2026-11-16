package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolResponseMockConfigOutput struct {
	// If the list is empty, the mock will always activate.
	ParameterConditions []*UnitTestToolCallParameter `json:"parameter_conditions,omitempty" xml:"parameter_conditions,omitempty"`
	// The return value the LLM sees when this mock is active.
	MockResult string `json:"mock_result" xml:"mock_result" required:"true"`
	// If true, the mock result is surfaced to the LLM as a tool error rather than a successful result.
	IsError         *bool `json:"is_error,omitempty" xml:"is_error,omitempty"`
	extraProperties map[string]interface{}
}

func (t *ToolResponseMockConfigOutput) GetParameterConditions() (value []*UnitTestToolCallParameter) {
	if t == nil {
		return
	}
	return t.ParameterConditions
}

func (t *ToolResponseMockConfigOutput) GetMockResult() (value string) {
	if t == nil {
		return
	}
	return t.MockResult
}

func (t *ToolResponseMockConfigOutput) GetIsError() (value *bool) {
	if t == nil {
		return
	}
	return t.IsError
}

func (t *ToolResponseMockConfigOutput) SetParameterConditions(parameterConditions []*UnitTestToolCallParameter) {
	t.ParameterConditions = parameterConditions
}

func (t *ToolResponseMockConfigOutput) SetMockResult(mockResult string) {
	t.MockResult = mockResult
}

func (t *ToolResponseMockConfigOutput) SetIsError(isError *bool) {
	t.IsError = isError
}

func (t *ToolResponseMockConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolResponseMockConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolResponseMockConfigOutput to string"
	}
	return string(jsonData)
}

func (t *ToolResponseMockConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias ToolResponseMockConfigOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = ToolResponseMockConfigOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolResponseMockConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ToolResponseMockConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
