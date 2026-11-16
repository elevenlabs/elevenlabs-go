package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ModerationConfig struct {
	Sexual                *ThresholdGuardrail `json:"sexual,omitempty" xml:"sexual,omitempty"`
	Violence              *ThresholdGuardrail `json:"violence,omitempty" xml:"violence,omitempty"`
	ViolenceGraphic       *ThresholdGuardrail `json:"violence_graphic,omitempty" xml:"violence_graphic,omitempty"`
	Harassment            *ThresholdGuardrail `json:"harassment,omitempty" xml:"harassment,omitempty"`
	HarassmentThreatening *ThresholdGuardrail `json:"harassment_threatening,omitempty" xml:"harassment_threatening,omitempty"`
	Hate                  *ThresholdGuardrail `json:"hate,omitempty" xml:"hate,omitempty"`
	HateThreatening       *ThresholdGuardrail `json:"hate_threatening,omitempty" xml:"hate_threatening,omitempty"`
	SelfHarmInstructions  *ThresholdGuardrail `json:"self_harm_instructions,omitempty" xml:"self_harm_instructions,omitempty"`
	SelfHarm              *ThresholdGuardrail `json:"self_harm,omitempty" xml:"self_harm,omitempty"`
	SelfHarmIntent        *ThresholdGuardrail `json:"self_harm_intent,omitempty" xml:"self_harm_intent,omitempty"`
	SexualMinors          *ThresholdGuardrail `json:"sexual_minors,omitempty" xml:"sexual_minors,omitempty"`
	extraProperties       map[string]interface{}
}

func (m *ModerationConfig) GetSexual() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.Sexual
}

func (m *ModerationConfig) GetViolence() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.Violence
}

func (m *ModerationConfig) GetViolenceGraphic() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.ViolenceGraphic
}

func (m *ModerationConfig) GetHarassment() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.Harassment
}

func (m *ModerationConfig) GetHarassmentThreatening() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.HarassmentThreatening
}

func (m *ModerationConfig) GetHate() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.Hate
}

func (m *ModerationConfig) GetHateThreatening() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.HateThreatening
}

func (m *ModerationConfig) GetSelfHarmInstructions() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.SelfHarmInstructions
}

func (m *ModerationConfig) GetSelfHarm() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.SelfHarm
}

func (m *ModerationConfig) GetSelfHarmIntent() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.SelfHarmIntent
}

func (m *ModerationConfig) GetSexualMinors() (value *ThresholdGuardrail) {
	if m == nil {
		return
	}
	return m.SexualMinors
}

func (m *ModerationConfig) SetSexual(sexual *ThresholdGuardrail) {
	m.Sexual = sexual
}

func (m *ModerationConfig) SetViolence(violence *ThresholdGuardrail) {
	m.Violence = violence
}

func (m *ModerationConfig) SetViolenceGraphic(violenceGraphic *ThresholdGuardrail) {
	m.ViolenceGraphic = violenceGraphic
}

func (m *ModerationConfig) SetHarassment(harassment *ThresholdGuardrail) {
	m.Harassment = harassment
}

func (m *ModerationConfig) SetHarassmentThreatening(harassmentThreatening *ThresholdGuardrail) {
	m.HarassmentThreatening = harassmentThreatening
}

func (m *ModerationConfig) SetHate(hate *ThresholdGuardrail) {
	m.Hate = hate
}

func (m *ModerationConfig) SetHateThreatening(hateThreatening *ThresholdGuardrail) {
	m.HateThreatening = hateThreatening
}

func (m *ModerationConfig) SetSelfHarmInstructions(selfHarmInstructions *ThresholdGuardrail) {
	m.SelfHarmInstructions = selfHarmInstructions
}

func (m *ModerationConfig) SetSelfHarm(selfHarm *ThresholdGuardrail) {
	m.SelfHarm = selfHarm
}

func (m *ModerationConfig) SetSelfHarmIntent(selfHarmIntent *ThresholdGuardrail) {
	m.SelfHarmIntent = selfHarmIntent
}

func (m *ModerationConfig) SetSexualMinors(sexualMinors *ThresholdGuardrail) {
	m.SexualMinors = sexualMinors
}

func (m *ModerationConfig) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ModerationConfig) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ModerationConfig to string"
	}
	return string(jsonData)
}

func (m *ModerationConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ModerationConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModerationConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ModerationConfig) MarshalJSON() ([]byte, error) {
	type embed ModerationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
