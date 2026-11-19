package agentworkflows

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// Expr is a workflow expression, compiled to the API's expression AST. Create one with
// ExprVariable, ExprLiteral, ExprLLM, ExprAllOf, ExprAnyOf, ExprIfElse or ExprFromModel, and combine
// them with the operator methods. Use it as an edge condition with ConditionExpr.
//
// Wherever an expression is expected (the operator methods, ExprAllOf, ExprAnyOf, ExprIfElse), an
// Expr or a plain value is accepted. A plain value (a string, bool, nil or Go number) becomes a
// literal; any other type panics, since it is a programming error.
type Expr struct {
	node any
}

// LLMValueType is the type of a value the LLM extracts from the conversation.
type LLMValueType string

const (
	LLMBoolean LLMValueType = "boolean"
	LLMString  LLMValueType = "string"
	LLMInteger LLMValueType = "integer"
	LLMNumber  LLMValueType = "number"
)

// LLMValueOptions are the optional settings of ExprLLM.
type LLMValueOptions struct {
	// Type is the value's type. Empty means LLMBoolean.
	Type LLMValueType
	// Enum lists the allowed values of a string. Nil leaves it out.
	Enum []string
}

// ExprVariable is the value of a dynamic variable.
func ExprVariable(name string) Expr {
	return Expr{node: map[string]any{"type": "dynamic_variable", "name": name}}
}

// ExprLiteral is a string, number, boolean or null literal. value must be a string, bool, nil, a
// Go integer or float type (finite), or a json.Number; anything else panics.
func ExprLiteral(value any) Expr {
	switch v := value.(type) {
	case nil:
		return Expr{node: map[string]any{"type": "null_literal"}}
	case string:
		return Expr{node: map[string]any{"type": "string_literal", "value": v}}
	case bool:
		return Expr{node: map[string]any{"type": "boolean_literal", "value": v}}
	}
	if number, ok := numberText(value); ok {
		return Expr{node: map[string]any{"type": "number_literal", "value": json.Number(number)}}
	}
	panic(fmt.Sprintf("agentworkflows: ExprLiteral needs a string, bool, nil or number, not %T", value))
}

func numberText(value any) (string, bool) {
	switch v := value.(type) {
	case json.Number:
		return string(v), true
	case int:
		return strconv.FormatInt(int64(v), 10), true
	case int8:
		return strconv.FormatInt(int64(v), 10), true
	case int16:
		return strconv.FormatInt(int64(v), 10), true
	case int32:
		return strconv.FormatInt(int64(v), 10), true
	case int64:
		return strconv.FormatInt(v, 10), true
	case uint:
		return strconv.FormatUint(uint64(v), 10), true
	case uint8:
		return strconv.FormatUint(uint64(v), 10), true
	case uint16:
		return strconv.FormatUint(uint64(v), 10), true
	case uint32:
		return strconv.FormatUint(uint64(v), 10), true
	case uint64:
		return strconv.FormatUint(v, 10), true
	case float32:
		if math.IsInf(float64(v), 0) || math.IsNaN(float64(v)) {
			return "", false
		}
		return strconv.FormatFloat(float64(v), 'g', -1, 32), true
	case float64:
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return "", false
		}
		return strconv.FormatFloat(v, 'g', -1, 64), true
	}
	return "", false
}

// ExprLLM is a value the LLM extracts from the conversation, described by description.
//
//	agentworkflows.ExprLLM("How satisfied is the caller, from 1 to 10?",
//		&agentworkflows.LLMValueOptions{Type: agentworkflows.LLMNumber}).Lt(5)
func ExprLLM(description string, opts *LLMValueOptions) Expr {
	valueType := LLMBoolean
	var enum []string
	if opts != nil {
		if opts.Type != "" {
			valueType = opts.Type
		}
		enum = opts.Enum
	}
	schema := map[string]any{"type": string(valueType), "description": description}
	if enum != nil {
		values := make([]any, len(enum))
		for i, v := range enum {
			values[i] = v
		}
		schema["enum"] = values
	}
	return Expr{node: map[string]any{"type": "llm", "value_schema": schema}}
}

// ExprAllOf is true when every expression is true.
func ExprAllOf(exprs ...any) Expr {
	return Expr{node: map[string]any{"type": "and_operator", "children": children(exprs)}}
}

// ExprAnyOf is true when at least one expression is true.
func ExprAnyOf(exprs ...any) Expr {
	return Expr{node: map[string]any{"type": "or_operator", "children": children(exprs)}}
}

// ExprIfElse is whenTrue if condition is true, otherwise whenFalse.
func ExprIfElse(condition, whenTrue, whenFalse any) Expr {
	return Expr{node: map[string]any{
		"type":            "conditional_operator",
		"condition":       toAST(condition),
		"trueExpression":  toAST(whenTrue),
		"falseExpression": toAST(whenFalse),
	}}
}

// ExprFromModel wraps an existing AST node: a generated type such as *types.AstNodeInput, or the
// node in wire format (a map, or JSON text as json.RawMessage or []byte). It fails only when ast
// can't be converted to JSON.
func ExprFromModel(ast any) (Expr, error) {
	var raw []byte
	switch v := ast.(type) {
	case json.RawMessage:
		raw = v
	case []byte:
		raw = v
	}
	if raw != nil {
		node, err := wirejson.Decode(raw)
		if err != nil {
			return Expr{}, err
		}
		return Expr{node: node}, nil
	}
	node, err := wirejson.Normalize(ast)
	if err != nil {
		return Expr{}, err
	}
	return Expr{node: node}, nil
}

// Eq is true when the two values are equal.
func (e Expr) Eq(other any) Expr { return e.binary("eq_operator", other) }

// Neq is true when the two values differ.
func (e Expr) Neq(other any) Expr { return e.binary("neq_operator", other) }

// Gt is true when e is greater than other.
func (e Expr) Gt(other any) Expr { return e.binary("gt_operator", other) }

// Gte is true when e is greater than or equal to other.
func (e Expr) Gte(other any) Expr { return e.binary("gte_operator", other) }

// Lt is true when e is less than other.
func (e Expr) Lt(other any) Expr { return e.binary("lt_operator", other) }

// Lte is true when e is less than or equal to other.
func (e Expr) Lte(other any) Expr { return e.binary("lte_operator", other) }

// Add is e + other.
func (e Expr) Add(other any) Expr { return e.binary("add_operator", other) }

// Sub is e - other.
func (e Expr) Sub(other any) Expr { return e.binary("sub_operator", other) }

// Mul is e * other.
func (e Expr) Mul(other any) Expr { return e.binary("mul_operator", other) }

// Div is e / other.
func (e Expr) Div(other any) Expr { return e.binary("div_operator", other) }

// Wire returns the expression in API wire format (a fresh copy).
func (e Expr) Wire() any {
	if e.node == nil {
		return map[string]any{"type": "null_literal"}
	}
	return wirejson.Clone(e.node)
}

func (e Expr) binary(op string, other any) Expr {
	return Expr{node: map[string]any{"type": op, "left": e.Wire(), "right": toAST(other)}}
}

func children(exprs []any) []any {
	out := make([]any, len(exprs))
	for i, e := range exprs {
		out[i] = toAST(e)
	}
	return out
}

func toAST(value any) any {
	switch v := value.(type) {
	case Expr:
		return v.Wire()
	case *Expr:
		return v.Wire()
	}
	return ExprLiteral(value).Wire()
}
