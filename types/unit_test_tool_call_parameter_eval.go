package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UnitTestToolCallParameterEval struct {
	Type     string
	Anything *MatchAnythingParameterEvaluationStrategy
	Exact    *ExactParameterEvaluationStrategy
	Llm      *LlmParameterEvaluationStrategy
	Regex    *RegexParameterEvaluationStrategy
}

func (u *UnitTestToolCallParameterEval) GetAnything() *MatchAnythingParameterEvaluationStrategy {
	if u == nil {
		return nil
	}
	if u.Anything == nil {
		return nil
	}
	return u.Anything
}

func (u *UnitTestToolCallParameterEval) GetExact() *ExactParameterEvaluationStrategy {
	if u == nil {
		return nil
	}
	if u.Exact == nil {
		return nil
	}
	return u.Exact
}

func (u *UnitTestToolCallParameterEval) GetLlm() *LlmParameterEvaluationStrategy {
	if u == nil {
		return nil
	}
	if u.Llm == nil {
		return nil
	}
	return u.Llm
}

func (u *UnitTestToolCallParameterEval) GetRegex() *RegexParameterEvaluationStrategy {
	if u == nil {
		return nil
	}
	if u.Regex == nil {
		return nil
	}
	return u.Regex
}

func (u *UnitTestToolCallParameterEval) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UnitTestToolCallParameterEval) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "anything":
		value := new(MatchAnythingParameterEvaluationStrategy)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Anything = value
	case "exact":
		value := new(ExactParameterEvaluationStrategy)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Exact = value
	case "llm":
		value := new(LlmParameterEvaluationStrategy)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Llm = value
	case "regex":
		value := new(RegexParameterEvaluationStrategy)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Regex = value
	}
	return nil
}

func (u UnitTestToolCallParameterEval) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Anything != nil {
		return internal.MarshalJSONWithDiscriminant(u.Anything, "type", "anything")
	}
	if u.Exact != nil {
		return internal.MarshalJSONWithDiscriminant(u.Exact, "type", "exact")
	}
	if u.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(u.Llm, "type", "llm")
	}
	if u.Regex != nil {
		return internal.MarshalJSONWithDiscriminant(u.Regex, "type", "regex")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UnitTestToolCallParameterEvalVisitor interface {
	VisitAnything(*MatchAnythingParameterEvaluationStrategy) error
	VisitExact(*ExactParameterEvaluationStrategy) error
	VisitLlm(*LlmParameterEvaluationStrategy) error
	VisitRegex(*RegexParameterEvaluationStrategy) error
}

func (u *UnitTestToolCallParameterEval) Accept(visitor UnitTestToolCallParameterEvalVisitor) error {
	if u.Anything != nil {
		return visitor.VisitAnything(u.Anything)
	}
	if u.Exact != nil {
		return visitor.VisitExact(u.Exact)
	}
	if u.Llm != nil {
		return visitor.VisitLlm(u.Llm)
	}
	if u.Regex != nil {
		return visitor.VisitRegex(u.Regex)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UnitTestToolCallParameterEval) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Anything != nil {
		fields = append(fields, "anything")
	}
	if u.Exact != nil {
		fields = append(fields, "exact")
	}
	if u.Llm != nil {
		fields = append(fields, "llm")
	}
	if u.Regex != nil {
		fields = append(fields, "regex")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}
