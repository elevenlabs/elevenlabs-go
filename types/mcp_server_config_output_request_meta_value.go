package types

import (
	"encoding/json"
	"fmt"
)

type McpServerConfigOutputRequestMetaValue struct {
	String                string
	Integer               int
	Double                float64
	Boolean               bool
	ConvAiSecretLocator   *ConvAiSecretLocator
	ConvAiDynamicVariable *ConvAiDynamicVariable
	ConvAiEnvVarLocator   *ConvAiEnvVarLocator
	typ                   string
}

func (m *McpServerConfigOutputRequestMetaValue) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *McpServerConfigOutputRequestMetaValue) GetInteger() int {
	if m == nil {
		return 0
	}
	return m.Integer
}

func (m *McpServerConfigOutputRequestMetaValue) GetDouble() float64 {
	if m == nil {
		return 0
	}
	return m.Double
}

func (m *McpServerConfigOutputRequestMetaValue) GetBoolean() bool {
	if m == nil {
		return false
	}
	return m.Boolean
}

func (m *McpServerConfigOutputRequestMetaValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigOutputRequestMetaValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if m == nil {
		return nil
	}
	if m.ConvAiDynamicVariable == nil {
		return nil
	}
	return m.ConvAiDynamicVariable
}

func (m *McpServerConfigOutputRequestMetaValue) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiEnvVarLocator == nil {
		return nil
	}
	return m.ConvAiEnvVarLocator
}

func (m *McpServerConfigOutputRequestMetaValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		m.typ = "Integer"
		m.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		m.typ = "Double"
		m.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		m.typ = "Boolean"
		m.Boolean = valueBoolean
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

func (m McpServerConfigOutputRequestMetaValue) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return json.Marshal(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return json.Marshal(m.Double)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return json.Marshal(m.Boolean)
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

type McpServerConfigOutputRequestMetaValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (m *McpServerConfigOutputRequestMetaValue) Accept(visitor McpServerConfigOutputRequestMetaValueVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return visitor.VisitInteger(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return visitor.VisitDouble(m.Double)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return visitor.VisitBoolean(m.Boolean)
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
