package agentworkflows

import "github.com/elevenlabs/elevenlabs-go/internal/wirejson"

// Condition says when an edge may be followed. Create one with ConditionLLM, ConditionExpr,
// ConditionAlways, ConditionSuccess or ConditionFailure. The zero value is not a valid condition:
// passing it to When or as a Back option records an invalid_option error.
type Condition struct {
	wire map[string]any
}

// ConditionOptions are the optional settings of ConditionLLM and ConditionExpr.
type ConditionOptions struct {
	// Label is shown on the edge in the workflow editor.
	Label string
}

func newCondition(wire map[string]any, label string) Condition {
	if label != "" {
		wire["label"] = label
	}
	return Condition{wire: wire}
}

// ConditionLLM is followed when the LLM judges condition to be true.
func ConditionLLM(condition string, opts *ConditionOptions) Condition {
	return newCondition(map[string]any{"type": "llm", "condition": condition}, conditionLabel(opts))
}

// ConditionExpr is followed when expression evaluates to true.
func ConditionExpr(expression Expr, opts *ConditionOptions) Condition {
	return newCondition(map[string]any{"type": "expression", "expression": expression.Wire()}, conditionLabel(opts))
}

// ConditionAlways is always followed. A verb called without a routing method uses it.
func ConditionAlways() Condition {
	return newCondition(map[string]any{"type": "unconditional"}, "")
}

// ConditionSuccess is followed when every tool in a tool node succeeded.
func ConditionSuccess() Condition {
	return newCondition(map[string]any{"type": "result", "successful": true}, "")
}

// ConditionFailure is followed when any tool in a tool node failed.
func ConditionFailure() Condition {
	return newCondition(map[string]any{"type": "result", "successful": false}, "")
}

func conditionLabel(opts *ConditionOptions) string {
	if opts == nil {
		return ""
	}
	return opts.Label
}

// Wire returns the condition in API wire format (a fresh copy).
func (c Condition) Wire() map[string]any {
	if c.wire == nil {
		return map[string]any{"type": "unconditional"}
	}
	return wirejson.Clone(c.wire).(map[string]any)
}
