package types

import (
	"encoding/json"
	"fmt"
)

// Optional metadata to be included in the webhook response. This should be a JSON string representing an object with a maximum depth of 2 levels and maximum size of 16KB. Useful for tracking internal IDs, job references, or other contextual information.
type ConvertSpeechToTextRequestWebhookMetadata struct {
	String           string
	StringUnknownMap map[string]any
	typ              string
}

func (c *ConvertSpeechToTextRequestWebhookMetadata) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ConvertSpeechToTextRequestWebhookMetadata) GetStringUnknownMap() map[string]any {
	if c == nil {
		return nil
	}
	return c.StringUnknownMap
}

func (c *ConvertSpeechToTextRequestWebhookMetadata) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	var valueStringUnknownMap map[string]any
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		c.typ = "StringUnknownMap"
		c.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConvertSpeechToTextRequestWebhookMetadata) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "StringUnknownMap" || c.StringUnknownMap != nil {
		return json.Marshal(c.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConvertSpeechToTextRequestWebhookMetadataVisitor interface {
	VisitString(string) error
	VisitStringUnknownMap(map[string]any) error
}

func (c *ConvertSpeechToTextRequestWebhookMetadata) Accept(visitor ConvertSpeechToTextRequestWebhookMetadataVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "StringUnknownMap" || c.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(c.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
