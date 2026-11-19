package types

import (
	"encoding/json"
	"fmt"
)

type McpServerConfigInputRequestHeadersValue struct {
	String                string
	ConvAiSecretLocator   *ConvAiSecretLocator
	ConvAiDynamicVariable *ConvAiDynamicVariable
	ConvAiEnvVarLocator   *ConvAiEnvVarLocator
	typ                   string
}

func (m *McpServerConfigInputRequestHeadersValue) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *McpServerConfigInputRequestHeadersValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigInputRequestHeadersValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if m == nil {
		return nil
	}
	if m.ConvAiDynamicVariable == nil {
		return nil
	}
	return m.ConvAiDynamicVariable
}

func (m *McpServerConfigInputRequestHeadersValue) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiEnvVarLocator == nil {
		return nil
	}
	return m.ConvAiEnvVarLocator
}

func (m *McpServerConfigInputRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		m.typ = "ConvAiSecretLocator"
		m.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		m.typ = "ConvAiDynamicVariable"
		m.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	var valueConvAiEnvVarLocator *ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		m.typ = "ConvAiEnvVarLocator"
		m.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpServerConfigInputRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return json.Marshal(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return json.Marshal(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return json.Marshal(m.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigInputRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (m *McpServerConfigInputRequestHeadersValue) Accept(visitor McpServerConfigInputRequestHeadersValueVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(m.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
