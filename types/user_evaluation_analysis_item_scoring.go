package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// How this criterion is scored (binary or numeric_uniform).
type UserEvaluationAnalysisItemScoring struct {
	Mode           string
	Binary         *BinaryScoring
	NumericUniform *NumericUniformScoring
}

func (u *UserEvaluationAnalysisItemScoring) GetBinary() *BinaryScoring {
	if u == nil {
		return nil
	}
	if u.Binary == nil {
		return nil
	}
	return u.Binary
}

func (u *UserEvaluationAnalysisItemScoring) GetNumericUniform() *NumericUniformScoring {
	if u == nil {
		return nil
	}
	if u.NumericUniform == nil {
		return nil
	}
	return u.NumericUniform
}

func (u *UserEvaluationAnalysisItemScoring) GetMode() string {
	if u == nil {
		return ""
	}
	return u.Mode
}

func (u *UserEvaluationAnalysisItemScoring) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Mode string `json:"mode"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Mode = unmarshaler.Mode
	if unmarshaler.Mode == "" {
		return fmt.Errorf("%T did not include discriminant mode", u)
	}
	switch unmarshaler.Mode {
	case "binary":
		value := new(BinaryScoring)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Binary = value
	case "numeric_uniform":
		value := new(NumericUniformScoring)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.NumericUniform = value
	}
	return nil
}

func (u UserEvaluationAnalysisItemScoring) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Binary != nil {
		return internal.MarshalJSONWithDiscriminant(u.Binary, "mode", "binary")
	}
	if u.NumericUniform != nil {
		return internal.MarshalJSONWithDiscriminant(u.NumericUniform, "mode", "numeric_uniform")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UserEvaluationAnalysisItemScoringVisitor interface {
	VisitBinary(*BinaryScoring) error
	VisitNumericUniform(*NumericUniformScoring) error
}

func (u *UserEvaluationAnalysisItemScoring) Accept(visitor UserEvaluationAnalysisItemScoringVisitor) error {
	if u.Binary != nil {
		return visitor.VisitBinary(u.Binary)
	}
	if u.NumericUniform != nil {
		return visitor.VisitNumericUniform(u.NumericUniform)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UserEvaluationAnalysisItemScoring) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Binary != nil {
		fields = append(fields, "binary")
	}
	if u.NumericUniform != nil {
		fields = append(fields, "numeric_uniform")
	}
	if len(fields) == 0 {
		if u.Mode != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Mode)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Mode != "" {
		field := fields[0]
		if u.Mode != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Mode,
				u,
			)
		}
	}
	return nil
}
