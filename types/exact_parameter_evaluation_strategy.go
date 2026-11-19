package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExactParameterEvaluationStrategy struct {
	type_ string
	// The exact string value that the parameter must match.
	ExpectedValue   string `json:"expected_value" xml:"expected_value" required:"true"`
	extraProperties map[string]interface{}
}

func (e *ExactParameterEvaluationStrategy) GetExpectedValue() (value string) {
	if e == nil {
		return
	}
	return e.ExpectedValue
}

func (e *ExactParameterEvaluationStrategy) SetExpectedValue(expectedValue string) {
	e.ExpectedValue = expectedValue
}

func (e *ExactParameterEvaluationStrategy) Type() string {
	return e.type_
}

func (e *ExactParameterEvaluationStrategy) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExactParameterEvaluationStrategy) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExactParameterEvaluationStrategy to string"
	}
	return string(jsonData)
}

func (e *ExactParameterEvaluationStrategy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExactParameterEvaluationStrategy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExactParameterEvaluationStrategy(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "exact" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "exact", *literals.Type)
	}
	e.type_ = "exact"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExactParameterEvaluationStrategy) MarshalJSON() ([]byte, error) {
	type embed ExactParameterEvaluationStrategy
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*e),
		Type:  "exact",
	}
	return json.Marshal(marshaler)
}
