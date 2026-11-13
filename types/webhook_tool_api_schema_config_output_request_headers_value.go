package types

import (
	"encoding/json"
	"fmt"
)

type WebhookToolApiSchemaConfigOutputRequestHeadersValue struct {
	String                string
	ConvAiSecretLocator   *ConvAiSecretLocator
	ConvAiDynamicVariable *ConvAiDynamicVariable
	ConvAiEnvVarLocator   *ConvAiEnvVarLocator
	typ                   string
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) GetString() string {
	if w == nil {
		return ""
	}
	return w.String
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if w == nil {
		return nil
	}
	if w.ConvAiSecretLocator == nil {
		return nil
	}
	return w.ConvAiSecretLocator
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if w == nil {
		return nil
	}
	if w.ConvAiDynamicVariable == nil {
		return nil
	}
	return w.ConvAiDynamicVariable
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if w == nil {
		return nil
	}
	if w.ConvAiEnvVarLocator == nil {
		return nil
	}
	return w.ConvAiEnvVarLocator
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		w.typ = "String"
		w.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		w.typ = "ConvAiSecretLocator"
		w.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		w.typ = "ConvAiDynamicVariable"
		w.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	var valueConvAiEnvVarLocator *ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		w.typ = "ConvAiEnvVarLocator"
		w.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
}

func (w WebhookToolApiSchemaConfigOutputRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if w.typ == "String" || w.String != "" {
		return json.Marshal(w.String)
	}
	if w.typ == "ConvAiSecretLocator" || w.ConvAiSecretLocator != nil {
		return json.Marshal(w.ConvAiSecretLocator)
	}
	if w.typ == "ConvAiDynamicVariable" || w.ConvAiDynamicVariable != nil {
		return json.Marshal(w.ConvAiDynamicVariable)
	}
	if w.typ == "ConvAiEnvVarLocator" || w.ConvAiEnvVarLocator != nil {
		return json.Marshal(w.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WebhookToolApiSchemaConfigOutputRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (w *WebhookToolApiSchemaConfigOutputRequestHeadersValue) Accept(visitor WebhookToolApiSchemaConfigOutputRequestHeadersValueVisitor) error {
	if w.typ == "String" || w.String != "" {
		return visitor.VisitString(w.String)
	}
	if w.typ == "ConvAiSecretLocator" || w.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(w.ConvAiSecretLocator)
	}
	if w.typ == "ConvAiDynamicVariable" || w.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(w.ConvAiDynamicVariable)
	}
	if w.typ == "ConvAiEnvVarLocator" || w.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(w.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}
