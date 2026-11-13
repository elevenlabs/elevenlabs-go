package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UnitTestGuardrailTriggerEvaluationModel struct {
	// Match any guardrail triggered by the generated turn.
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UnitTestGuardrailTriggerEvaluationModel) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UnitTestGuardrailTriggerEvaluationModel) SetType(type_ *string) {
	u.Type = type_
}

func (u *UnitTestGuardrailTriggerEvaluationModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestGuardrailTriggerEvaluationModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestGuardrailTriggerEvaluationModel to string"
	}
	return string(jsonData)
}

func (u *UnitTestGuardrailTriggerEvaluationModel) UnmarshalJSON(data []byte) error {
	type unmarshaler UnitTestGuardrailTriggerEvaluationModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnitTestGuardrailTriggerEvaluationModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestGuardrailTriggerEvaluationModel) MarshalJSON() ([]byte, error) {
	type embed UnitTestGuardrailTriggerEvaluationModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
