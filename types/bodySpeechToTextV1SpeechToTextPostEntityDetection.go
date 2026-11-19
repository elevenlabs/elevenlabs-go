package types

import (
	"encoding/json"
	"fmt"
)

// Detect entities in the transcript. Can be 'all' to detect all entities, a single entity type or category string, or a list of entity types/categories. Categories include 'pii', 'phi', 'pci', 'other', 'offensive_language'. When enabled, detected entities will be returned in the 'entities' field with their text, type, and character positions.
type BodySpeechToTextV1SpeechToTextPostEntityDetection struct {
	String     string
	StringList []string
	typ        string
}

func (bodySpeechToTextV1SpeechToTextPostEntityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection) GetString() string {
	if bodySpeechToTextV1SpeechToTextPostEntityDetection == nil {
		return ""
	}
	return bodySpeechToTextV1SpeechToTextPostEntityDetection.String
}

func (bodySpeechToTextV1SpeechToTextPostEntityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection) GetStringList() []string {
	if bodySpeechToTextV1SpeechToTextPostEntityDetection == nil {
		return nil
	}
	return bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList
}

func (bodySpeechToTextV1SpeechToTextPostEntityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		bodySpeechToTextV1SpeechToTextPostEntityDetection.typ = "String"
		bodySpeechToTextV1SpeechToTextPostEntityDetection.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		bodySpeechToTextV1SpeechToTextPostEntityDetection.typ = "StringList"
		bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodySpeechToTextV1SpeechToTextPostEntityDetection)
}

func (bodySpeechToTextV1SpeechToTextPostEntityDetection BodySpeechToTextV1SpeechToTextPostEntityDetection) MarshalJSON() ([]byte, error) {
	if bodySpeechToTextV1SpeechToTextPostEntityDetection.typ == "String" || bodySpeechToTextV1SpeechToTextPostEntityDetection.String != "" {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostEntityDetection.String)
	}
	if bodySpeechToTextV1SpeechToTextPostEntityDetection.typ == "StringList" || bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList != nil {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostEntityDetection)
}

type BodySpeechToTextV1SpeechToTextPostEntityDetectionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (bodySpeechToTextV1SpeechToTextPostEntityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection) Accept(visitor BodySpeechToTextV1SpeechToTextPostEntityDetectionVisitor) error {
	if bodySpeechToTextV1SpeechToTextPostEntityDetection.typ == "String" || bodySpeechToTextV1SpeechToTextPostEntityDetection.String != "" {
		return visitor.VisitString(bodySpeechToTextV1SpeechToTextPostEntityDetection.String)
	}
	if bodySpeechToTextV1SpeechToTextPostEntityDetection.typ == "StringList" || bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList != nil {
		return visitor.VisitStringList(bodySpeechToTextV1SpeechToTextPostEntityDetection.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostEntityDetection)
}
