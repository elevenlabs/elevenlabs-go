package types

import (
	"encoding/json"
	"fmt"
)

type McpToolCallMcpToolCall struct {
	McpToolCallMcpToolCallLoading          *McpToolCallMcpToolCallLoading
	McpToolCallMcpToolCallAwaitingApproval *McpToolCallMcpToolCallAwaitingApproval
	McpToolCallMcpToolCallSuccess          *McpToolCallMcpToolCallSuccess
	McpToolCallMcpToolCallFailure          *McpToolCallMcpToolCallFailure
	typ                                    string
}

func (m *McpToolCallMcpToolCall) GetMcpToolCallMcpToolCallLoading() *McpToolCallMcpToolCallLoading {
	if m == nil {
		return nil
	}
	if m.McpToolCallMcpToolCallLoading == nil {
		return nil
	}
	return m.McpToolCallMcpToolCallLoading
}

func (m *McpToolCallMcpToolCall) GetMcpToolCallMcpToolCallAwaitingApproval() *McpToolCallMcpToolCallAwaitingApproval {
	if m == nil {
		return nil
	}
	if m.McpToolCallMcpToolCallAwaitingApproval == nil {
		return nil
	}
	return m.McpToolCallMcpToolCallAwaitingApproval
}

func (m *McpToolCallMcpToolCall) GetMcpToolCallMcpToolCallSuccess() *McpToolCallMcpToolCallSuccess {
	if m == nil {
		return nil
	}
	if m.McpToolCallMcpToolCallSuccess == nil {
		return nil
	}
	return m.McpToolCallMcpToolCallSuccess
}

func (m *McpToolCallMcpToolCall) GetMcpToolCallMcpToolCallFailure() *McpToolCallMcpToolCallFailure {
	if m == nil {
		return nil
	}
	if m.McpToolCallMcpToolCallFailure == nil {
		return nil
	}
	return m.McpToolCallMcpToolCallFailure
}

func (m *McpToolCallMcpToolCall) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("state", "loading") {
		matchedLiteral = true
		var valueMcpToolCallMcpToolCallLoadingLiteral *McpToolCallMcpToolCallLoading
		if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallLoadingLiteral); err == nil {
			m.typ = "McpToolCallMcpToolCallLoading"
			m.McpToolCallMcpToolCallLoading = valueMcpToolCallMcpToolCallLoadingLiteral
			return nil
		}
	}
	if hasLiteral("state", "awaiting_approval") {
		matchedLiteral = true
		var valueMcpToolCallMcpToolCallAwaitingApprovalLiteral *McpToolCallMcpToolCallAwaitingApproval
		if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallAwaitingApprovalLiteral); err == nil {
			m.typ = "McpToolCallMcpToolCallAwaitingApproval"
			m.McpToolCallMcpToolCallAwaitingApproval = valueMcpToolCallMcpToolCallAwaitingApprovalLiteral
			return nil
		}
	}
	if hasLiteral("state", "success") {
		matchedLiteral = true
		var valueMcpToolCallMcpToolCallSuccessLiteral *McpToolCallMcpToolCallSuccess
		if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallSuccessLiteral); err == nil {
			m.typ = "McpToolCallMcpToolCallSuccess"
			m.McpToolCallMcpToolCallSuccess = valueMcpToolCallMcpToolCallSuccessLiteral
			return nil
		}
	}
	if hasLiteral("state", "failure") {
		matchedLiteral = true
		var valueMcpToolCallMcpToolCallFailureLiteral *McpToolCallMcpToolCallFailure
		if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallFailureLiteral); err == nil {
			m.typ = "McpToolCallMcpToolCallFailure"
			m.McpToolCallMcpToolCallFailure = valueMcpToolCallMcpToolCallFailureLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
	}
	var valueMcpToolCallMcpToolCallLoading *McpToolCallMcpToolCallLoading
	if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallLoading); err == nil {
		m.typ = "McpToolCallMcpToolCallLoading"
		m.McpToolCallMcpToolCallLoading = valueMcpToolCallMcpToolCallLoading
		return nil
	}
	var valueMcpToolCallMcpToolCallAwaitingApproval *McpToolCallMcpToolCallAwaitingApproval
	if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallAwaitingApproval); err == nil {
		m.typ = "McpToolCallMcpToolCallAwaitingApproval"
		m.McpToolCallMcpToolCallAwaitingApproval = valueMcpToolCallMcpToolCallAwaitingApproval
		return nil
	}
	var valueMcpToolCallMcpToolCallSuccess *McpToolCallMcpToolCallSuccess
	if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallSuccess); err == nil {
		m.typ = "McpToolCallMcpToolCallSuccess"
		m.McpToolCallMcpToolCallSuccess = valueMcpToolCallMcpToolCallSuccess
		return nil
	}
	var valueMcpToolCallMcpToolCallFailure *McpToolCallMcpToolCallFailure
	if err := json.Unmarshal(data, &valueMcpToolCallMcpToolCallFailure); err == nil {
		m.typ = "McpToolCallMcpToolCallFailure"
		m.McpToolCallMcpToolCallFailure = valueMcpToolCallMcpToolCallFailure
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpToolCallMcpToolCall) MarshalJSON() ([]byte, error) {
	if m.typ == "McpToolCallMcpToolCallLoading" || m.McpToolCallMcpToolCallLoading != nil {
		return json.Marshal(m.McpToolCallMcpToolCallLoading)
	}
	if m.typ == "McpToolCallMcpToolCallAwaitingApproval" || m.McpToolCallMcpToolCallAwaitingApproval != nil {
		return json.Marshal(m.McpToolCallMcpToolCallAwaitingApproval)
	}
	if m.typ == "McpToolCallMcpToolCallSuccess" || m.McpToolCallMcpToolCallSuccess != nil {
		return json.Marshal(m.McpToolCallMcpToolCallSuccess)
	}
	if m.typ == "McpToolCallMcpToolCallFailure" || m.McpToolCallMcpToolCallFailure != nil {
		return json.Marshal(m.McpToolCallMcpToolCallFailure)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpToolCallMcpToolCallVisitor interface {
	VisitMcpToolCallMcpToolCallLoading(*McpToolCallMcpToolCallLoading) error
	VisitMcpToolCallMcpToolCallAwaitingApproval(*McpToolCallMcpToolCallAwaitingApproval) error
	VisitMcpToolCallMcpToolCallSuccess(*McpToolCallMcpToolCallSuccess) error
	VisitMcpToolCallMcpToolCallFailure(*McpToolCallMcpToolCallFailure) error
}

func (m *McpToolCallMcpToolCall) Accept(visitor McpToolCallMcpToolCallVisitor) error {
	if m.typ == "McpToolCallMcpToolCallLoading" || m.McpToolCallMcpToolCallLoading != nil {
		return visitor.VisitMcpToolCallMcpToolCallLoading(m.McpToolCallMcpToolCallLoading)
	}
	if m.typ == "McpToolCallMcpToolCallAwaitingApproval" || m.McpToolCallMcpToolCallAwaitingApproval != nil {
		return visitor.VisitMcpToolCallMcpToolCallAwaitingApproval(m.McpToolCallMcpToolCallAwaitingApproval)
	}
	if m.typ == "McpToolCallMcpToolCallSuccess" || m.McpToolCallMcpToolCallSuccess != nil {
		return visitor.VisitMcpToolCallMcpToolCallSuccess(m.McpToolCallMcpToolCallSuccess)
	}
	if m.typ == "McpToolCallMcpToolCallFailure" || m.McpToolCallMcpToolCallFailure != nil {
		return visitor.VisitMcpToolCallMcpToolCallFailure(m.McpToolCallMcpToolCallFailure)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
