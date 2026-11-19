package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AstNodeOutput struct {
	Type                string
	AddOperator         *AstAdditionOperatorNodeOutput
	AndOperator         *AstAndOperatorNodeOutput
	BooleanLiteral      *AstBooleanNodeOutput
	ConditionalOperator *AstConditionalOperatorNodeOutput
	DivOperator         *AstDivisionOperatorNodeOutput
	DynamicVariable     *AstDynamicVariableNodeOutput
	EqOperator          *AstEqualsOperatorNodeOutput
	GtOperator          *AstGreaterThanOperatorNodeOutput
	GteOperator         *AstGreaterThanOrEqualsOperatorNodeOutput
	Llm                 *AstllmNodeOutput
	LtOperator          *AstLessThanOperatorNodeOutput
	LteOperator         *AstLessThanOrEqualsOperatorNodeOutput
	MulOperator         *AstMultiplicationOperatorNodeOutput
	NeqOperator         *AstNotEqualsOperatorNodeOutput
	NullLiteral         *AstNullNodeOutput
	NumberLiteral       *AstNumberNodeOutput
	OrOperator          *AstOrOperatorNodeOutput
	StringLiteral       *AstStringNodeOutput
	SubOperator         *AstSubtractionOperatorNodeOutput
}

func (a *AstNodeOutput) GetAddOperator() *AstAdditionOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.AddOperator == nil {
		return nil
	}
	return a.AddOperator
}

func (a *AstNodeOutput) GetAndOperator() *AstAndOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.AndOperator == nil {
		return nil
	}
	return a.AndOperator
}

func (a *AstNodeOutput) GetBooleanLiteral() *AstBooleanNodeOutput {
	if a == nil {
		return nil
	}
	if a.BooleanLiteral == nil {
		return nil
	}
	return a.BooleanLiteral
}

func (a *AstNodeOutput) GetConditionalOperator() *AstConditionalOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.ConditionalOperator == nil {
		return nil
	}
	return a.ConditionalOperator
}

func (a *AstNodeOutput) GetDivOperator() *AstDivisionOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.DivOperator == nil {
		return nil
	}
	return a.DivOperator
}

func (a *AstNodeOutput) GetDynamicVariable() *AstDynamicVariableNodeOutput {
	if a == nil {
		return nil
	}
	if a.DynamicVariable == nil {
		return nil
	}
	return a.DynamicVariable
}

func (a *AstNodeOutput) GetEqOperator() *AstEqualsOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.EqOperator == nil {
		return nil
	}
	return a.EqOperator
}

func (a *AstNodeOutput) GetGtOperator() *AstGreaterThanOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.GtOperator == nil {
		return nil
	}
	return a.GtOperator
}

func (a *AstNodeOutput) GetGteOperator() *AstGreaterThanOrEqualsOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.GteOperator == nil {
		return nil
	}
	return a.GteOperator
}

func (a *AstNodeOutput) GetLlm() *AstllmNodeOutput {
	if a == nil {
		return nil
	}
	if a.Llm == nil {
		return nil
	}
	return a.Llm
}

func (a *AstNodeOutput) GetLtOperator() *AstLessThanOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.LtOperator == nil {
		return nil
	}
	return a.LtOperator
}

func (a *AstNodeOutput) GetLteOperator() *AstLessThanOrEqualsOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.LteOperator == nil {
		return nil
	}
	return a.LteOperator
}

func (a *AstNodeOutput) GetMulOperator() *AstMultiplicationOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.MulOperator == nil {
		return nil
	}
	return a.MulOperator
}

func (a *AstNodeOutput) GetNeqOperator() *AstNotEqualsOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.NeqOperator == nil {
		return nil
	}
	return a.NeqOperator
}

func (a *AstNodeOutput) GetNullLiteral() *AstNullNodeOutput {
	if a == nil {
		return nil
	}
	if a.NullLiteral == nil {
		return nil
	}
	return a.NullLiteral
}

func (a *AstNodeOutput) GetNumberLiteral() *AstNumberNodeOutput {
	if a == nil {
		return nil
	}
	if a.NumberLiteral == nil {
		return nil
	}
	return a.NumberLiteral
}

func (a *AstNodeOutput) GetOrOperator() *AstOrOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.OrOperator == nil {
		return nil
	}
	return a.OrOperator
}

func (a *AstNodeOutput) GetStringLiteral() *AstStringNodeOutput {
	if a == nil {
		return nil
	}
	if a.StringLiteral == nil {
		return nil
	}
	return a.StringLiteral
}

func (a *AstNodeOutput) GetSubOperator() *AstSubtractionOperatorNodeOutput {
	if a == nil {
		return nil
	}
	if a.SubOperator == nil {
		return nil
	}
	return a.SubOperator
}

func (a *AstNodeOutput) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AstNodeOutput) UnmarshalJSON(data []byte) error {
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
		value := new(AstAdditionOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.AddOperator = value
	case "and_operator":
		value := new(AstAndOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.AndOperator = value
	case "boolean_literal":
		value := new(AstBooleanNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.BooleanLiteral = value
	case "conditional_operator":
		value := new(AstConditionalOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.ConditionalOperator = value
	case "div_operator":
		value := new(AstDivisionOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DivOperator = value
	case "dynamic_variable":
		value := new(AstDynamicVariableNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DynamicVariable = value
	case "eq_operator":
		value := new(AstEqualsOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.EqOperator = value
	case "gt_operator":
		value := new(AstGreaterThanOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.GtOperator = value
	case "gte_operator":
		value := new(AstGreaterThanOrEqualsOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.GteOperator = value
	case "llm":
		value := new(AstllmNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Llm = value
	case "lt_operator":
		value := new(AstLessThanOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.LtOperator = value
	case "lte_operator":
		value := new(AstLessThanOrEqualsOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.LteOperator = value
	case "mul_operator":
		value := new(AstMultiplicationOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.MulOperator = value
	case "neq_operator":
		value := new(AstNotEqualsOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NeqOperator = value
	case "null_literal":
		value := new(AstNullNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NullLiteral = value
	case "number_literal":
		value := new(AstNumberNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.NumberLiteral = value
	case "or_operator":
		value := new(AstOrOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.OrOperator = value
	case "string_literal":
		value := new(AstStringNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.StringLiteral = value
	case "sub_operator":
		value := new(AstSubtractionOperatorNodeOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.SubOperator = value
	}
	return nil
}

func (a AstNodeOutput) MarshalJSON() ([]byte, error) {
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
		return internal.MarshalJSONWithDiscriminant(a.Llm, "type", "llm")
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

type AstNodeOutputVisitor interface {
	VisitAddOperator(*AstAdditionOperatorNodeOutput) error
	VisitAndOperator(*AstAndOperatorNodeOutput) error
	VisitBooleanLiteral(*AstBooleanNodeOutput) error
	VisitConditionalOperator(*AstConditionalOperatorNodeOutput) error
	VisitDivOperator(*AstDivisionOperatorNodeOutput) error
	VisitDynamicVariable(*AstDynamicVariableNodeOutput) error
	VisitEqOperator(*AstEqualsOperatorNodeOutput) error
	VisitGtOperator(*AstGreaterThanOperatorNodeOutput) error
	VisitGteOperator(*AstGreaterThanOrEqualsOperatorNodeOutput) error
	VisitLlm(*AstllmNodeOutput) error
	VisitLtOperator(*AstLessThanOperatorNodeOutput) error
	VisitLteOperator(*AstLessThanOrEqualsOperatorNodeOutput) error
	VisitMulOperator(*AstMultiplicationOperatorNodeOutput) error
	VisitNeqOperator(*AstNotEqualsOperatorNodeOutput) error
	VisitNullLiteral(*AstNullNodeOutput) error
	VisitNumberLiteral(*AstNumberNodeOutput) error
	VisitOrOperator(*AstOrOperatorNodeOutput) error
	VisitStringLiteral(*AstStringNodeOutput) error
	VisitSubOperator(*AstSubtractionOperatorNodeOutput) error
}

func (a *AstNodeOutput) Accept(visitor AstNodeOutputVisitor) error {
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

func (a *AstNodeOutput) validate() error {
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
