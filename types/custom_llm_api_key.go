package types

import (
	"encoding/json"
	"fmt"
)

// The API key for authentication. Either a workspace secret reference {'secret_id': '...'} or an environment variable reference {'env_var_label': '...'}.
type CustomLlmApiKey struct {
	ConvAiSecretLocator *ConvAiSecretLocator
	ConvAiEnvVarLocator *ConvAiEnvVarLocator
	typ                 string
}

func (c *CustomLlmApiKey) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiSecretLocator == nil {
		return nil
	}
	return c.ConvAiSecretLocator
}

func (c *CustomLlmApiKey) GetConvAiEnvVarLocator() *ConvAiEnvVarLocator {
	if c == nil {
		return nil
	}
	if c.ConvAiEnvVarLocator == nil {
		return nil
	}
	return c.ConvAiEnvVarLocator
}

func (c *CustomLlmApiKey) UnmarshalJSON(data []byte) error {
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

func (c CustomLlmApiKey) MarshalJSON() ([]byte, error) {
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return json.Marshal(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return json.Marshal(c.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CustomLlmApiKeyVisitor interface {
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiEnvVarLocator(*ConvAiEnvVarLocator) error
}

func (c *CustomLlmApiKey) Accept(visitor CustomLlmApiKeyVisitor) error {
	if c.typ == "ConvAiSecretLocator" || c.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(c.ConvAiSecretLocator)
	}
	if c.typ == "ConvAiEnvVarLocator" || c.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(c.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
