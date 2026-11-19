package voices

import (
	"encoding/json"
	"fmt"
)

// Labels for the voice. Keys can be language, accent, gender, or age.
type AddVoiceRequestLabels struct {
	StringUnknownMap map[string]string
	String           string
	typ              string
}

func (addVoiceRequestLabels *AddVoiceRequestLabels) GetStringUnknownMap() map[string]string {
	if addVoiceRequestLabels == nil {
		return nil
	}
	return addVoiceRequestLabels.StringUnknownMap
}

func (addVoiceRequestLabels *AddVoiceRequestLabels) GetString() string {
	if addVoiceRequestLabels == nil {
		return ""
	}
	return addVoiceRequestLabels.String
}

func (addVoiceRequestLabels *AddVoiceRequestLabels) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]string
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		addVoiceRequestLabels.typ = "StringUnknownMap"
		addVoiceRequestLabels.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		addVoiceRequestLabels.typ = "String"
		addVoiceRequestLabels.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, addVoiceRequestLabels)
}

func (addVoiceRequestLabels AddVoiceRequestLabels) MarshalJSON() ([]byte, error) {
	if addVoiceRequestLabels.typ == "StringUnknownMap" || addVoiceRequestLabels.StringUnknownMap != nil {
		return json.Marshal(addVoiceRequestLabels.StringUnknownMap)
	}
	if addVoiceRequestLabels.typ == "String" || addVoiceRequestLabels.String != "" {
		return json.Marshal(addVoiceRequestLabels.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", addVoiceRequestLabels)
}

type AddVoiceRequestLabelsVisitor interface {
	VisitStringUnknownMap(map[string]string) error
	VisitString(string) error
}

func (addVoiceRequestLabels *AddVoiceRequestLabels) Accept(visitor AddVoiceRequestLabelsVisitor) error {
	if addVoiceRequestLabels.typ == "StringUnknownMap" || addVoiceRequestLabels.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(addVoiceRequestLabels.StringUnknownMap)
	}
	if addVoiceRequestLabels.typ == "String" || addVoiceRequestLabels.String != "" {
		return visitor.VisitString(addVoiceRequestLabels.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", addVoiceRequestLabels)
}
