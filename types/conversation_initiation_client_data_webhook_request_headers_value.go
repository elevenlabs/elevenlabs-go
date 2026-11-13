package types

import (
	"encoding/json"
	"fmt"
)

type ConversationInitiationClientDataWebhookRequestHeadersValue struct {
	String              string
	ConvAiSecretLocator *ConvAiSecretLocator
	ConvAiEnvVarLocator *ConvAiEnvVarLocator
	typ                 string
}

func (c *ConversationInitiationClientDataWebhookRequestHeadersValue) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ConversationInitiationClientDataWebhookRequestHeadersValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiSecretLocator == nil {
		return nil
	}
	return c.ConvAiSecretLocator
}

func (c *ConversationInitiationClientDataWebhookRequestHeadersValue) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiEnvVarLocator == nil {
		return nil
	}
	return c.ConvAiEnvVarLocator
}

func (c *ConversationInitiationClientDataWebhookRequestHeadersValue) UnmarshalJSON(data []byte) error {
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
	var valueConvAiEnvVarLocator *ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		c.typ = "ConvAiEnvVarLocator"
		c.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConversationInitiationClientDataWebhookRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return json.Marshal(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return json.Marshal(c.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConversationInitiationClientDataWebhookRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (c *ConversationInitiationClientDataWebhookRequestHeadersValue) Accept(visitor ConversationInitiationClientDataWebhookRequestHeadersValueVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(c.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
