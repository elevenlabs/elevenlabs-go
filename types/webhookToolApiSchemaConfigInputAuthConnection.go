package types

import (
	"encoding/json"
	"fmt"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type WebhookToolApiSchemaConfigInputAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (webhookToolApiSchemaConfigInputAuthConnection *WebhookToolApiSchemaConfigInputAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if webhookToolApiSchemaConfigInputAuthConnection == nil {
		return nil
	}
	if webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator == nil {
		return nil
	}
	return webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator
}

func (webhookToolApiSchemaConfigInputAuthConnection *WebhookToolApiSchemaConfigInputAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if webhookToolApiSchemaConfigInputAuthConnection == nil {
		return nil
	}
	if webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator
}

func (webhookToolApiSchemaConfigInputAuthConnection *WebhookToolApiSchemaConfigInputAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		webhookToolApiSchemaConfigInputAuthConnection.typ = "AuthConnectionLocator"
		webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		webhookToolApiSchemaConfigInputAuthConnection.typ = "EnvironmentAuthConnectionLocator"
		webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, webhookToolApiSchemaConfigInputAuthConnection)
}

func (webhookToolApiSchemaConfigInputAuthConnection WebhookToolApiSchemaConfigInputAuthConnection) MarshalJSON() ([]byte, error) {
	if webhookToolApiSchemaConfigInputAuthConnection.typ == "AuthConnectionLocator" || webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator != nil {
		return json.Marshal(webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator)
	}
	if webhookToolApiSchemaConfigInputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", webhookToolApiSchemaConfigInputAuthConnection)
}

type WebhookToolApiSchemaConfigInputAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (webhookToolApiSchemaConfigInputAuthConnection *WebhookToolApiSchemaConfigInputAuthConnection) Accept(visitor WebhookToolApiSchemaConfigInputAuthConnectionVisitor) error {
	if webhookToolApiSchemaConfigInputAuthConnection.typ == "AuthConnectionLocator" || webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(webhookToolApiSchemaConfigInputAuthConnection.AuthConnectionLocator)
	}
	if webhookToolApiSchemaConfigInputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(webhookToolApiSchemaConfigInputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", webhookToolApiSchemaConfigInputAuthConnection)
}
