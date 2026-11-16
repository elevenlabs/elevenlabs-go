package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RegexParameterEvaluationStrategy struct {
	type_ string
	// A regex pattern to match the agent's response against.
	Pattern         string `json:"pattern" xml:"pattern" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RegexParameterEvaluationStrategy) GetPattern() (value string) {
	if r == nil {
		return
	}
	return r.Pattern
}

func (r *RegexParameterEvaluationStrategy) SetPattern(pattern string) {
	r.Pattern = pattern
}

func (r *RegexParameterEvaluationStrategy) Type() string {
	return r.type_
}

func (r *RegexParameterEvaluationStrategy) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RegexParameterEvaluationStrategy) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegexParameterEvaluationStrategy to string"
	}
	return string(jsonData)
}

func (r *RegexParameterEvaluationStrategy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RegexParameterEvaluationStrategy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RegexParameterEvaluationStrategy(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "regex" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", r, "regex", *literals.Type)
	}
	r.type_ = "regex"
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RegexParameterEvaluationStrategy) MarshalJSON() ([]byte, error) {
	type embed RegexParameterEvaluationStrategy
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*r),
		Type:  "regex",
	}
	return json.Marshal(marshaler)
}
