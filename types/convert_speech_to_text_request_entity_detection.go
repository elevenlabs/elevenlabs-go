package types

import (
	"encoding/json"
	"fmt"
)

// Detect entities in the transcript. Can be 'all' to detect all entities, a single entity type or category string, or a list of entity types/categories. Categories include 'pii', 'phi', 'pci', 'other', 'offensive_language'. When enabled, detected entities will be returned in the 'entities' field with their text, type, and character positions.
type ConvertSpeechToTextRequestEntityDetection struct {
	String     string
	StringList []string
	typ        string
}

func (c *ConvertSpeechToTextRequestEntityDetection) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ConvertSpeechToTextRequestEntityDetection) GetStringList() []string {
	if c == nil {
		return nil
	}
	return c.StringList
}

func (c *ConvertSpeechToTextRequestEntityDetection) UnmarshalJSON(data []byte) error {
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

func (c ConvertSpeechToTextRequestEntityDetection) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "StringList" || c.StringList != nil {
		return json.Marshal(c.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConvertSpeechToTextRequestEntityDetectionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (c *ConvertSpeechToTextRequestEntityDetection) Accept(visitor ConvertSpeechToTextRequestEntityDetectionVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "StringList" || c.StringList != nil {
		return visitor.VisitStringList(c.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
