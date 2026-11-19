package types

import (
	"encoding/json"
	"fmt"
)

// Detect entities on committed transcripts. Can be 'all', a single entity type or category, or a list of types/categories ('pii', 'phi', 'pci', 'other', 'offensive_language'). When enabled, detected entities are delivered in a separate 'committed_transcript_entities' event with their text, type, and character positions.
type V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection struct {
	String     string
	StringList []string
	typ        string
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection) GetString() string {
	if v == nil {
		return ""
	}
	return v.String
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection) GetStringList() []string {
	if v == nil {
		return nil
	}
	return v.StringList
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		v.typ = "String"
		v.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		v.typ = "StringList"
		v.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection) MarshalJSON() ([]byte, error) {
	if v.typ == "String" || v.String != "" {
		return json.Marshal(v.String)
	}
	if v.typ == "StringList" || v.StringList != nil {
		return json.Marshal(v.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V1SpeechToTextRealtimeSpeechToTextRequestEntityDetectionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection) Accept(visitor V1SpeechToTextRealtimeSpeechToTextRequestEntityDetectionVisitor) error {
	if v.typ == "String" || v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.typ == "StringList" || v.StringList != nil {
		return visitor.VisitStringList(v.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}
