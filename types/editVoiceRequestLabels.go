package types

import (
	"encoding/json"
	"fmt"
)

// Labels for the voice. Keys can be language, accent, gender, or age.
type EditVoiceRequestLabels struct {
	StringUnknownMap map[string]string
	String           string
	typ              string
}

func (editVoiceRequestLabels *EditVoiceRequestLabels) GetStringUnknownMap() map[string]string {
	if editVoiceRequestLabels == nil {
		return nil
	}
	return editVoiceRequestLabels.StringUnknownMap
}

func (editVoiceRequestLabels *EditVoiceRequestLabels) GetString() string {
	if editVoiceRequestLabels == nil {
		return ""
	}
	return editVoiceRequestLabels.String
}

func (editVoiceRequestLabels *EditVoiceRequestLabels) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]string
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		editVoiceRequestLabels.typ = "StringUnknownMap"
		editVoiceRequestLabels.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		editVoiceRequestLabels.typ = "String"
		editVoiceRequestLabels.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, editVoiceRequestLabels)
}

func (editVoiceRequestLabels EditVoiceRequestLabels) MarshalJSON() ([]byte, error) {
	if editVoiceRequestLabels.typ == "StringUnknownMap" || editVoiceRequestLabels.StringUnknownMap != nil {
		return json.Marshal(editVoiceRequestLabels.StringUnknownMap)
	}
	if editVoiceRequestLabels.typ == "String" || editVoiceRequestLabels.String != "" {
		return json.Marshal(editVoiceRequestLabels.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", editVoiceRequestLabels)
}

type EditVoiceRequestLabelsVisitor interface {
	VisitStringUnknownMap(map[string]string) error
	VisitString(string) error
}

func (editVoiceRequestLabels *EditVoiceRequestLabels) Accept(visitor EditVoiceRequestLabelsVisitor) error {
	if editVoiceRequestLabels.typ == "StringUnknownMap" || editVoiceRequestLabels.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(editVoiceRequestLabels.StringUnknownMap)
	}
	if editVoiceRequestLabels.typ == "String" || editVoiceRequestLabels.String != "" {
		return visitor.VisitString(editVoiceRequestLabels.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", editVoiceRequestLabels)
}
