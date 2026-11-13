package types

import (
	"encoding/json"
	"fmt"
)

// Redact entities from the transcript text. Accepts the same format as entity_detection: 'all', a category ('pii', 'phi'), or specific entity types. Must be a subset of entity_detection. When redaction is enabled, the entities field will not be returned.
type ConvertSpeechToTextRequestEntityRedaction struct {
	String     string
	StringList []string
	typ        string
}

func (c *ConvertSpeechToTextRequestEntityRedaction) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ConvertSpeechToTextRequestEntityRedaction) GetStringList() []string {
	if c == nil {
		return nil
	}
	return c.StringList
}

func (c *ConvertSpeechToTextRequestEntityRedaction) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		c.typ = "StringList"
		c.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConvertSpeechToTextRequestEntityRedaction) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "StringList" || c.StringList != nil {
		return json.Marshal(c.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConvertSpeechToTextRequestEntityRedactionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (c *ConvertSpeechToTextRequestEntityRedaction) Accept(visitor ConvertSpeechToTextRequestEntityRedactionVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "StringList" || c.StringList != nil {
		return visitor.VisitStringList(c.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
