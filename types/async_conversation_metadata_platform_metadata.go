package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AsyncConversationMetadataPlatformMetadata struct {
	ConversationInitiationSource string
	PlainIntegration             *PlainPlatformMetadata
	ZendeskIntegration           *ZendeskPlatformMetadata
}

func (a *AsyncConversationMetadataPlatformMetadata) GetPlainIntegration() *PlainPlatformMetadata {
	if a == nil {
		return nil
	}
	if a.PlainIntegration == nil {
		return nil
	}
	return a.PlainIntegration
}

func (a *AsyncConversationMetadataPlatformMetadata) GetZendeskIntegration() *ZendeskPlatformMetadata {
	if a == nil {
		return nil
	}
	if a.ZendeskIntegration == nil {
		return nil
	}
	return a.ZendeskIntegration
}

func (a *AsyncConversationMetadataPlatformMetadata) GetConversationInitiationSource() string {
	if a == nil {
		return ""
	}
	return a.ConversationInitiationSource
}

func (a *AsyncConversationMetadataPlatformMetadata) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		ConversationInitiationSource string `json:"conversation_initiation_source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.ConversationInitiationSource = unmarshaler.ConversationInitiationSource
	if unmarshaler.ConversationInitiationSource == "" {
		return fmt.Errorf("%T did not include discriminant conversation_initiation_source", a)
	}
	switch unmarshaler.ConversationInitiationSource {
	case "plain_integration":
		value := new(PlainPlatformMetadata)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.PlainIntegration = value
	case "zendesk_integration":
		value := new(ZendeskPlatformMetadata)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.ZendeskIntegration = value
	}
	return nil
}

func (a AsyncConversationMetadataPlatformMetadata) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.PlainIntegration != nil {
		return internal.MarshalJSONWithDiscriminant(a.PlainIntegration, "conversation_initiation_source", "plain_integration")
	}
	if a.ZendeskIntegration != nil {
		return internal.MarshalJSONWithDiscriminant(a.ZendeskIntegration, "conversation_initiation_source", "zendesk_integration")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AsyncConversationMetadataPlatformMetadataVisitor interface {
	VisitPlainIntegration(*PlainPlatformMetadata) error
	VisitZendeskIntegration(*ZendeskPlatformMetadata) error
}

func (a *AsyncConversationMetadataPlatformMetadata) Accept(visitor AsyncConversationMetadataPlatformMetadataVisitor) error {
	if a.PlainIntegration != nil {
		return visitor.VisitPlainIntegration(a.PlainIntegration)
	}
	if a.ZendeskIntegration != nil {
		return visitor.VisitZendeskIntegration(a.ZendeskIntegration)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AsyncConversationMetadataPlatformMetadata) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.PlainIntegration != nil {
		fields = append(fields, "plain_integration")
	}
	if a.ZendeskIntegration != nil {
		fields = append(fields, "zendesk_integration")
	}
	if len(fields) == 0 {
		if a.ConversationInitiationSource != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.ConversationInitiationSource)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.ConversationInitiationSource != "" {
		field := fields[0]
		if a.ConversationInitiationSource != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.ConversationInitiationSource,
				a,
			)
		}
	}
	return nil
}
