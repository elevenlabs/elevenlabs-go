package types

import (
	"encoding/json"
	"fmt"
)

// Optional metadata to be included in the webhook response. This should be a JSON string representing an object with a maximum depth of 2 levels and maximum size of 16KB. Useful for tracking internal IDs, job references, or other contextual information.
type BodySpeechToTextV1SpeechToTextPostWebhookMetadata struct {
	String           string
	StringUnknownMap map[string]any
	typ              string
}

func (bodySpeechToTextV1SpeechToTextPostWebhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) GetString() string {
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata == nil {
		return ""
	}
	return bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String
}

func (bodySpeechToTextV1SpeechToTextPostWebhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) GetStringUnknownMap() map[string]any {
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata == nil {
		return nil
	}
	return bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap
}

func (bodySpeechToTextV1SpeechToTextPostWebhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ = "String"
		bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String = valueString
		return nil
	}
	var valueStringUnknownMap map[string]any
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ = "StringUnknownMap"
		bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodySpeechToTextV1SpeechToTextPostWebhookMetadata)
}

func (bodySpeechToTextV1SpeechToTextPostWebhookMetadata BodySpeechToTextV1SpeechToTextPostWebhookMetadata) MarshalJSON() ([]byte, error) {
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ == "String" || bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String != "" {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String)
	}
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ == "StringUnknownMap" || bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap != nil {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostWebhookMetadata)
}

type BodySpeechToTextV1SpeechToTextPostWebhookMetadataVisitor interface {
	VisitString(string) error
	VisitStringUnknownMap(map[string]any) error
}

func (bodySpeechToTextV1SpeechToTextPostWebhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) Accept(visitor BodySpeechToTextV1SpeechToTextPostWebhookMetadataVisitor) error {
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ == "String" || bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String != "" {
		return visitor.VisitString(bodySpeechToTextV1SpeechToTextPostWebhookMetadata.String)
	}
	if bodySpeechToTextV1SpeechToTextPostWebhookMetadata.typ == "StringUnknownMap" || bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(bodySpeechToTextV1SpeechToTextPostWebhookMetadata.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostWebhookMetadata)
}
