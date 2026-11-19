package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Metadata for async conversation delivery (Zendesk, Slack, etc.).
type AsyncConversationMetadata struct {
	DeliveryStatus                 AsyncConversationMetadataDeliveryStatus    `json:"delivery_status" xml:"delivery_status" required:"true"`
	DeliveryTimestamp              int                                        `json:"delivery_timestamp" xml:"delivery_timestamp" required:"true"`
	DeliveryError                  *string                                    `json:"delivery_error,omitempty" xml:"delivery_error,omitempty"`
	ExternalSystem                 string                                     `json:"external_system" xml:"external_system" required:"true"`
	ExternalId                     string                                     `json:"external_id" xml:"external_id" required:"true"`
	ExternalLink                   *string                                    `json:"external_link,omitempty" xml:"external_link,omitempty"`
	RetryCount                     *int                                       `json:"retry_count,omitempty" xml:"retry_count,omitempty"`
	LastRetryTimestamp             *int                                       `json:"last_retry_timestamp,omitempty" xml:"last_retry_timestamp,omitempty"`
	LastProcessedExternalMessageId *string                                    `json:"last_processed_external_message_id,omitempty" xml:"last_processed_external_message_id,omitempty"`
	IntegrationTriggerId           *string                                    `json:"integration_trigger_id,omitempty" xml:"integration_trigger_id,omitempty"`
	PlatformMetadata               *AsyncConversationMetadataPlatformMetadata `json:"platform_metadata,omitempty" xml:"platform_metadata,omitempty"`
	extraProperties                map[string]interface{}
}

func (a *AsyncConversationMetadata) GetDeliveryStatus() (value AsyncConversationMetadataDeliveryStatus) {
	if a == nil {
		return
	}
	return a.DeliveryStatus
}

func (a *AsyncConversationMetadata) GetDeliveryTimestamp() (value int) {
	if a == nil {
		return
	}
	return a.DeliveryTimestamp
}

func (a *AsyncConversationMetadata) GetDeliveryError() (value *string) {
	if a == nil {
		return
	}
	return a.DeliveryError
}

func (a *AsyncConversationMetadata) GetExternalSystem() (value string) {
	if a == nil {
		return
	}
	return a.ExternalSystem
}

func (a *AsyncConversationMetadata) GetExternalId() (value string) {
	if a == nil {
		return
	}
	return a.ExternalId
}

func (a *AsyncConversationMetadata) GetExternalLink() (value *string) {
	if a == nil {
		return
	}
	return a.ExternalLink
}

func (a *AsyncConversationMetadata) GetRetryCount() (value *int) {
	if a == nil {
		return
	}
	return a.RetryCount
}

func (a *AsyncConversationMetadata) GetLastRetryTimestamp() (value *int) {
	if a == nil {
		return
	}
	return a.LastRetryTimestamp
}

func (a *AsyncConversationMetadata) GetLastProcessedExternalMessageId() (value *string) {
	if a == nil {
		return
	}
	return a.LastProcessedExternalMessageId
}

func (a *AsyncConversationMetadata) GetIntegrationTriggerId() (value *string) {
	if a == nil {
		return
	}
	return a.IntegrationTriggerId
}

func (a *AsyncConversationMetadata) GetPlatformMetadata() (value *AsyncConversationMetadataPlatformMetadata) {
	if a == nil {
		return
	}
	return a.PlatformMetadata
}

func (a *AsyncConversationMetadata) SetDeliveryStatus(deliveryStatus AsyncConversationMetadataDeliveryStatus) {
	a.DeliveryStatus = deliveryStatus
}

func (a *AsyncConversationMetadata) SetDeliveryTimestamp(deliveryTimestamp int) {
	a.DeliveryTimestamp = deliveryTimestamp
}

func (a *AsyncConversationMetadata) SetDeliveryError(deliveryError *string) {
	a.DeliveryError = deliveryError
}

func (a *AsyncConversationMetadata) SetExternalSystem(externalSystem string) {
	a.ExternalSystem = externalSystem
}

func (a *AsyncConversationMetadata) SetExternalId(externalId string) {
	a.ExternalId = externalId
}

func (a *AsyncConversationMetadata) SetExternalLink(externalLink *string) {
	a.ExternalLink = externalLink
}

func (a *AsyncConversationMetadata) SetRetryCount(retryCount *int) {
	a.RetryCount = retryCount
}

func (a *AsyncConversationMetadata) SetLastRetryTimestamp(lastRetryTimestamp *int) {
	a.LastRetryTimestamp = lastRetryTimestamp
}

func (a *AsyncConversationMetadata) SetLastProcessedExternalMessageId(lastProcessedExternalMessageId *string) {
	a.LastProcessedExternalMessageId = lastProcessedExternalMessageId
}

func (a *AsyncConversationMetadata) SetIntegrationTriggerId(integrationTriggerId *string) {
	a.IntegrationTriggerId = integrationTriggerId
}

func (a *AsyncConversationMetadata) SetPlatformMetadata(platformMetadata *AsyncConversationMetadataPlatformMetadata) {
	a.PlatformMetadata = platformMetadata
}

func (a *AsyncConversationMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AsyncConversationMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsyncConversationMetadata to string"
	}
	return string(jsonData)
}

func (a *AsyncConversationMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AsyncConversationMetadata) MarshalJSON() ([]byte, error) {
	type embed AsyncConversationMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
