package types

import (
	"encoding/json"
	"fmt"
)

// Redact entities from committed transcripts. Accepts the same format as entity_detection: 'all', a category ('pii', 'phi', 'pci', 'other', 'offensive_language'), or specific entity types. The redacted text is delivered in a separate 'redacted_transcript' event after the committed transcript. If entity detection fails, an 'error' event is sent and the connection is closed. Cannot be combined with transcript_edit or infer_speaker_names.
type V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction struct {
	String     string
	StringList []string
	typ        string
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction) GetString() string {
	if v == nil {
		return ""
	}
	return v.String
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction) GetStringList() []string {
	if v == nil {
		return nil
	}
	return v.StringList
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction) UnmarshalJSON(data []byte) error {
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

func (v V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction) MarshalJSON() ([]byte, error) {
	if v.typ == "String" || v.String != "" {
		return json.Marshal(v.String)
	}
	if v.typ == "StringList" || v.StringList != nil {
		return json.Marshal(v.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V1SpeechToTextRealtimeSpeechToTextRequestEntityRedactionVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (v *V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction) Accept(visitor V1SpeechToTextRealtimeSpeechToTextRequestEntityRedactionVisitor) error {
	if v.typ == "String" || v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.typ == "StringList" || v.StringList != nil {
		return visitor.VisitStringList(v.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}
