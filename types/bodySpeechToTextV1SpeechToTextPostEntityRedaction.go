package types

import (
	"encoding/json"
	"fmt"
)

// Redact entities from the transcript text. Accepts the same format as entity_detection: 'all', a category ('pii', 'phi'), or specific entity types. Must be a subset of entity_detection. When redaction is enabled, the entities field will not be returned.
type BodySpeechToTextV1SpeechToTextPostEntityRedaction struct {
	String     string
	StringList []string
	typ        string
}

func (bodySpeechToTextV1SpeechToTextPostEntityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction) GetString() string {
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction == nil {
		return ""
	}
	return bodySpeechToTextV1SpeechToTextPostEntityRedaction.String
}

func (bodySpeechToTextV1SpeechToTextPostEntityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction) GetStringList() []string {
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction == nil {
		return nil
	}
	return bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList
}

func (bodySpeechToTextV1SpeechToTextPostEntityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ = "String"
		bodySpeechToTextV1SpeechToTextPostEntityRedaction.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ = "StringList"
		bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodySpeechToTextV1SpeechToTextPostEntityRedaction)
}

func (bodySpeechToTextV1SpeechToTextPostEntityRedaction BodySpeechToTextV1SpeechToTextPostEntityRedaction) MarshalJSON() ([]byte, error) {
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ == "String" || bodySpeechToTextV1SpeechToTextPostEntityRedaction.String != "" {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostEntityRedaction.String)
	}
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ == "StringList" || bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList != nil {
		return json.Marshal(bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostEntityRedaction)
}

type BodySpeechToTextV1SpeechToTextPostEntityRedactionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (bodySpeechToTextV1SpeechToTextPostEntityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction) Accept(visitor BodySpeechToTextV1SpeechToTextPostEntityRedactionVisitor) error {
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ == "String" || bodySpeechToTextV1SpeechToTextPostEntityRedaction.String != "" {
		return visitor.VisitString(bodySpeechToTextV1SpeechToTextPostEntityRedaction.String)
	}
	if bodySpeechToTextV1SpeechToTextPostEntityRedaction.typ == "StringList" || bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList != nil {
		return visitor.VisitStringList(bodySpeechToTextV1SpeechToTextPostEntityRedaction.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodySpeechToTextV1SpeechToTextPostEntityRedaction)
}
