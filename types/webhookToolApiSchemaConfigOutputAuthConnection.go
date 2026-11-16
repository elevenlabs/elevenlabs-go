package types

import (
	"encoding/json"
	"fmt"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type WebhookToolApiSchemaConfigOutputAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (webhookToolApiSchemaConfigOutputAuthConnection *WebhookToolApiSchemaConfigOutputAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if webhookToolApiSchemaConfigOutputAuthConnection == nil {
		return nil
	}
	if webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator == nil {
		return nil
	}
	return webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator
}

func (webhookToolApiSchemaConfigOutputAuthConnection *WebhookToolApiSchemaConfigOutputAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if webhookToolApiSchemaConfigOutputAuthConnection == nil {
		return nil
	}
	if webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator
}

func (webhookToolApiSchemaConfigOutputAuthConnection *WebhookToolApiSchemaConfigOutputAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		webhookToolApiSchemaConfigOutputAuthConnection.typ = "AuthConnectionLocator"
		webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		webhookToolApiSchemaConfigOutputAuthConnection.typ = "EnvironmentAuthConnectionLocator"
		webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, webhookToolApiSchemaConfigOutputAuthConnection)
}

func (webhookToolApiSchemaConfigOutputAuthConnection WebhookToolApiSchemaConfigOutputAuthConnection) MarshalJSON() ([]byte, error) {
	if webhookToolApiSchemaConfigOutputAuthConnection.typ == "AuthConnectionLocator" || webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator != nil {
		return json.Marshal(webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator)
	}
	if webhookToolApiSchemaConfigOutputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", webhookToolApiSchemaConfigOutputAuthConnection)
}

type WebhookToolApiSchemaConfigOutputAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (webhookToolApiSchemaConfigOutputAuthConnection *WebhookToolApiSchemaConfigOutputAuthConnection) Accept(visitor WebhookToolApiSchemaConfigOutputAuthConnectionVisitor) error {
	if webhookToolApiSchemaConfigOutputAuthConnection.typ == "AuthConnectionLocator" || webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(webhookToolApiSchemaConfigOutputAuthConnection.AuthConnectionLocator)
	}
	if webhookToolApiSchemaConfigOutputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(webhookToolApiSchemaConfigOutputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", webhookToolApiSchemaConfigOutputAuthConnection)
}
