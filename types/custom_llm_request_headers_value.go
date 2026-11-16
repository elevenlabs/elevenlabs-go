package types

import (
	"encoding/json"
	"fmt"
)

type CustomLlmRequestHeadersValue struct {
	String                string
	ConvAiSecretLocator   *ConvAiSecretLocator
	ConvAiDynamicVariable *ConvAiDynamicVariable
	ConvAiEnvVarLocator   *ConvAiEnvVarLocator
	typ                   string
}

func (c *CustomLlmRequestHeadersValue) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *CustomLlmRequestHeadersValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiSecretLocator == nil {
		return nil
	}
	return c.ConvAiSecretLocator
}

func (c *CustomLlmRequestHeadersValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if c == nil {
		return nil
	}
	if c.ConvAiDynamicVariable == nil {
		return nil
	}
	return c.ConvAiDynamicVariable
}

func (c *CustomLlmRequestHeadersValue) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiEnvVarLocator == nil {
		return nil
	}
	return c.ConvAiEnvVarLocator
}

func (c *CustomLlmRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		c.typ = "ConvAiSecretLocator"
		c.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		c.typ = "ConvAiDynamicVariable"
		c.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	var valueConvAiEnvVarLocator *ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		c.typ = "ConvAiEnvVarLocator"
		c.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CustomLlmRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return json.Marshal(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiDynamicVariable" || c.ConvAiDynamicVariable != nil {
		return json.Marshal(c.ConvAiDynamicVariable)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return json.Marshal(c.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CustomLlmRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (c *CustomLlmRequestHeadersValue) Accept(visitor CustomLlmRequestHeadersValueVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiDynamicVariable" || c.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(c.ConvAiDynamicVariable)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(c.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
