package types

import (
	"encoding/json"
	"fmt"
)

type ApiIntegrationWebhookOverridesRequestHeadersValue struct {
	String                string
	ConvAiDynamicVariable *ConvAiDynamicVariable
	typ                   string
}

func (a *ApiIntegrationWebhookOverridesRequestHeadersValue) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApiIntegrationWebhookOverridesRequestHeadersValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if a == nil {
		return nil
	}
	if a.ConvAiDynamicVariable == nil {
		return nil
	}
	return a.ConvAiDynamicVariable
}

func (a *ApiIntegrationWebhookOverridesRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueConvAiDynamicVariable *ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		a.typ = "ConvAiDynamicVariable"
		a.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApiIntegrationWebhookOverridesRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "ConvAiDynamicVariable" || a.ConvAiDynamicVariable != nil {
		return json.Marshal(a.ConvAiDynamicVariable)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiIntegrationWebhookOverridesRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
}

func (a *ApiIntegrationWebhookOverridesRequestHeadersValue) Accept(visitor ApiIntegrationWebhookOverridesRequestHeadersValueVisitor) error {
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "ConvAiDynamicVariable" || a.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(a.ConvAiDynamicVariable)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}
