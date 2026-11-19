package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AstNodeInput struct {
	Type                string
	AddOperator         *AstAdditionOperatorNodeInput
	AndOperator         *AstAndOperatorNodeInput
	BooleanLiteral      *AstBooleanNodeInput
	ConditionalOperator *AstConditionalOperatorNodeInput
	DivOperator         *AstDivisionOperatorNodeInput
	DynamicVariable     *AstDynamicVariableNodeInput
	EqOperator          *AstEqualsOperatorNodeInput
	GtOperator          *AstGreaterThanOperatorNodeInput
	GteOperator         *AstGreaterThanOrEqualsOperatorNodeInput
	Llm                 *AstllmNodeInput
	LtOperator          *AstLessThanOperatorNodeInput
	LteOperator         *AstLessThanOrEqualsOperatorNodeInput
	MulOperator         *AstMultiplicationOperatorNodeInput
	NeqOperator         *AstNotEqualsOperatorNodeInput
	NullLiteral         *AstNullNodeInput
	NumberLiteral       *AstNumberNodeInput
	OrOperator          *AstOrOperatorNodeInput
	StringLiteral       *AstStringNodeInput
	SubOperator         *AstSubtractionOperatorNodeInput
}

func (a *AstNodeInput) GetAddOperator() *AstAdditionOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.AddOperator == nil {
		return nil
	}
	return a.AddOperator
}

func (a *AstNodeInput) GetAndOperator() *AstAndOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.AndOperator == nil {
		return nil
	}
	return a.AndOperator
}

func (a *AstNodeInput) GetBooleanLiteral() *AstBooleanNodeInput {
	if a == nil {
		return nil
	}
	if a.BooleanLiteral == nil {
		return nil
	}
	return a.BooleanLiteral
}

func (a *AstNodeInput) GetConditionalOperator() *AstConditionalOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.ConditionalOperator == nil {
		return nil
	}
	return a.ConditionalOperator
}

func (a *AstNodeInput) GetDivOperator() *AstDivisionOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.DivOperator == nil {
		return nil
	}
	return a.DivOperator
}

func (a *AstNodeInput) GetDynamicVariable() *AstDynamicVariableNodeInput {
	if a == nil {
		return nil
	}
	if a.DynamicVariable == nil {
		return nil
	}
	return a.DynamicVariable
}

func (a *AstNodeInput) GetEqOperator() *AstEqualsOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.EqOperator == nil {
		return nil
	}
	return a.EqOperator
}

func (a *AstNodeInput) GetGtOperator() *AstGreaterThanOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.GtOperator == nil {
		return nil
	}
	return a.GtOperator
}

func (a *AstNodeInput) GetGteOperator() *AstGreaterThanOrEqualsOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.GteOperator == nil {
		return nil
	}
	return a.GteOperator
}

func (a *AstNodeInput) GetLlm() *AstllmNodeInput {
	if a == nil {
		return nil
	}
	if a.Llm == nil {
		return nil
	}
	return a.Llm
}

func (a *AstNodeInput) GetLtOperator() *AstLessThanOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.LtOperator == nil {
		return nil
	}
	return a.LtOperator
}

func (a *AstNodeInput) GetLteOperator() *AstLessThanOrEqualsOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.LteOperator == nil {
		return nil
	}
	return a.LteOperator
}

func (a *AstNodeInput) GetMulOperator() *AstMultiplicationOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.MulOperator == nil {
		return nil
	}
	return a.MulOperator
}

func (a *AstNodeInput) GetNeqOperator() *AstNotEqualsOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.NeqOperator == nil {
		return nil
	}
	return a.NeqOperator
}

func (a *AstNodeInput) GetNullLiteral() *AstNullNodeInput {
	if a == nil {
		return nil
	}
	if a.NullLiteral == nil {
		return nil
	}
	return a.NullLiteral
}

func (a *AstNodeInput) GetNumberLiteral() *AstNumberNodeInput {
	if a == nil {
		return nil
	}
	if a.NumberLiteral == nil {
		return nil
	}
	return a.NumberLiteral
}

func (a *AstNodeInput) GetOrOperator() *AstOrOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.OrOperator == nil {
		return nil
	}
	return a.OrOperator
}

func (a *AstNodeInput) GetStringLiteral() *AstStringNodeInput {
	if a == nil {
		return nil
	}
	if a.StringLiteral == nil {
		return nil
	}
	return a.StringLiteral
}

func (a *AstNodeInput) GetSubOperator() *AstSubtractionOperatorNodeInput {
	if a == nil {
		return nil
	}
	if a.SubOperator == nil {
		return nil
	}
	return a.SubOperator
}

func (a *AstNodeInput) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AstNodeInput) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", a)
	}
	switch unmarshaler.Type {
	case "add_operator":
		value := new(AstAdditionOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.AddOperator = value
	case "and_operator":
		value := new(AstAndOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.AndOperator = value
	case "boolean_literal":
		value := new(AstBooleanNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.BooleanLiteral = value
	case "conditional_operator":
		value := new(AstConditionalOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.ConditionalOperator = value
	case "div_operator":
		value := new(AstDivisionOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DivOperator = value
	case "dynamic_variable":
		value := new(AstDynamicVariableNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DynamicVariable = value
	case "eq_operator":
		value := new(AstEqualsOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.EqOperator = value
	case "gt_operator":
		value := new(AstGreaterThanOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.GtOperator = value
	case "gte_operator":
		value := new(AstGreaterThanOrEqualsOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.GteOperator = value
	case "llm":
		value := new(AstllmNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Llm = value
	case "lt_operator":
		value := new(AstLessThanOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.LtOperator = value
	case "lte_operator":
		value := new(AstLessThanOrEqualsOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.LteOperator = value
	case "mul_operator":
		value := new(AstMultiplicationOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.MulOperator = value
	case "neq_operator":
		value := new(AstNotEqualsOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NeqOperator = value
	case "null_literal":
		value := new(AstNullNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NullLiteral = value
	case "number_literal":
		value := new(AstNumberNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NumberLiteral = value
	case "or_operator":
		value := new(AstOrOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.OrOperator = value
	case "string_literal":
		value := new(AstStringNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.StringLiteral = value
	case "sub_operator":
		value := new(AstSubtractionOperatorNodeInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.SubOperator = value
	}
	return nil
}

func (a AstNodeInput) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.AddOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.AddOperator, "type", "add_operator")
	}
	if a.AndOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.AndOperator, "type", "and_operator")
	}
	if a.BooleanLiteral != nil {
		return internal.MarshalJSONWithDiscriminant(a.BooleanLiteral, "type", "boolean_literal")
	}
	if a.ConditionalOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.ConditionalOperator, "type", "conditional_operator")
	}
	if a.DivOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.DivOperator, "type", "div_operator")
	}
	if a.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(a.DynamicVariable, "type", "dynamic_variable")
	}
	if a.EqOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.EqOperator, "type", "eq_operator")
	}
	if a.GtOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.GtOperator, "type", "gt_operator")
	}
	if a.GteOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.GteOperator, "type", "gte_operator")
	}
	if a.Llm != nil {
		return internal.MarshalJSONWithExtraProperty(a.Llm, "type", "llm")
	}
	if a.LtOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.LtOperator, "type", "lt_operator")
	}
	if a.LteOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.LteOperator, "type", "lte_operator")
	}
	if a.MulOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.MulOperator, "type", "mul_operator")
	}
	if a.NeqOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.NeqOperator, "type", "neq_operator")
	}
	if a.NullLiteral != nil {
		return internal.MarshalJSONWithDiscriminant(a.NullLiteral, "type", "null_literal")
	}
	if a.NumberLiteral != nil {
		return internal.MarshalJSONWithDiscriminant(a.NumberLiteral, "type", "number_literal")
	}
	if a.OrOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.OrOperator, "type", "or_operator")
	}
	if a.StringLiteral != nil {
		return internal.MarshalJSONWithDiscriminant(a.StringLiteral, "type", "string_literal")
	}
	if a.SubOperator != nil {
		return internal.MarshalJSONWithDiscriminant(a.SubOperator, "type", "sub_operator")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AstNodeInputVisitor interface {
	VisitAddOperator(*AstAdditionOperatorNodeInput) error
	VisitAndOperator(*AstAndOperatorNodeInput) error
	VisitBooleanLiteral(*AstBooleanNodeInput) error
	VisitConditionalOperator(*AstConditionalOperatorNodeInput) error
	VisitDivOperator(*AstDivisionOperatorNodeInput) error
	VisitDynamicVariable(*AstDynamicVariableNodeInput) error
	VisitEqOperator(*AstEqualsOperatorNodeInput) error
	VisitGtOperator(*AstGreaterThanOperatorNodeInput) error
	VisitGteOperator(*AstGreaterThanOrEqualsOperatorNodeInput) error
	VisitLlm(*AstllmNodeInput) error
	VisitLtOperator(*AstLessThanOperatorNodeInput) error
	VisitLteOperator(*AstLessThanOrEqualsOperatorNodeInput) error
	VisitMulOperator(*AstMultiplicationOperatorNodeInput) error
	VisitNeqOperator(*AstNotEqualsOperatorNodeInput) error
	VisitNullLiteral(*AstNullNodeInput) error
	VisitNumberLiteral(*AstNumberNodeInput) error
	VisitOrOperator(*AstOrOperatorNodeInput) error
	VisitStringLiteral(*AstStringNodeInput) error
	VisitSubOperator(*AstSubtractionOperatorNodeInput) error
}

func (a *AstNodeInput) Accept(visitor AstNodeInputVisitor) error {
	if a.AddOperator != nil {
		return visitor.VisitAddOperator(a.AddOperator)
	}
	if a.AndOperator != nil {
		return visitor.VisitAndOperator(a.AndOperator)
	}
	if a.BooleanLiteral != nil {
		return visitor.VisitBooleanLiteral(a.BooleanLiteral)
	}
	if a.ConditionalOperator != nil {
		return visitor.VisitConditionalOperator(a.ConditionalOperator)
	}
	if a.DivOperator != nil {
		return visitor.VisitDivOperator(a.DivOperator)
	}
	if a.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(a.DynamicVariable)
	}
	if a.EqOperator != nil {
		return visitor.VisitEqOperator(a.EqOperator)
	}
	if a.GtOperator != nil {
		return visitor.VisitGtOperator(a.GtOperator)
	}
	if a.GteOperator != nil {
		return visitor.VisitGteOperator(a.GteOperator)
	}
	if a.Llm != nil {
		return visitor.VisitLlm(a.Llm)
	}
	if a.LtOperator != nil {
		return visitor.VisitLtOperator(a.LtOperator)
	}
	if a.LteOperator != nil {
		return visitor.VisitLteOperator(a.LteOperator)
	}
	if a.MulOperator != nil {
		return visitor.VisitMulOperator(a.MulOperator)
	}
	if a.NeqOperator != nil {
		return visitor.VisitNeqOperator(a.NeqOperator)
	}
	if a.NullLiteral != nil {
		return visitor.VisitNullLiteral(a.NullLiteral)
	}
	if a.NumberLiteral != nil {
		return visitor.VisitNumberLiteral(a.NumberLiteral)
	}
	if a.OrOperator != nil {
		return visitor.VisitOrOperator(a.OrOperator)
	}
	if a.StringLiteral != nil {
		return visitor.VisitStringLiteral(a.StringLiteral)
	}
	if a.SubOperator != nil {
		return visitor.VisitSubOperator(a.SubOperator)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AstNodeInput) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.AddOperator != nil {
		fields = append(fields, "add_operator")
	}
	if a.AndOperator != nil {
		fields = append(fields, "and_operator")
	}
	if a.BooleanLiteral != nil {
		fields = append(fields, "boolean_literal")
	}
	if a.ConditionalOperator != nil {
		fields = append(fields, "conditional_operator")
	}
	if a.DivOperator != nil {
		fields = append(fields, "div_operator")
	}
	if a.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if a.EqOperator != nil {
		fields = append(fields, "eq_operator")
	}
	if a.GtOperator != nil {
		fields = append(fields, "gt_operator")
	}
	if a.GteOperator != nil {
		fields = append(fields, "gte_operator")
	}
	if a.Llm != nil {
		fields = append(fields, "llm")
	}
	if a.LtOperator != nil {
		fields = append(fields, "lt_operator")
	}
	if a.LteOperator != nil {
		fields = append(fields, "lte_operator")
	}
	if a.MulOperator != nil {
		fields = append(fields, "mul_operator")
	}
	if a.NeqOperator != nil {
		fields = append(fields, "neq_operator")
	}
	if a.NullLiteral != nil {
		fields = append(fields, "null_literal")
	}
	if a.NumberLiteral != nil {
		fields = append(fields, "number_literal")
	}
	if a.OrOperator != nil {
		fields = append(fields, "or_operator")
	}
	if a.StringLiteral != nil {
		fields = append(fields, "string_literal")
	}
	if a.SubOperator != nil {
		fields = append(fields, "sub_operator")
	}
	if len(fields) == 0 {
		if a.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Type)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Type != "" {
		field := fields[0]
		if a.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Type,
				a,
			)
		}
	}
	return nil
}
