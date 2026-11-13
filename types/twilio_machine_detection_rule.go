package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Action to take on a live call for the given AMD verdicts.
type TwilioMachineDetectionRule struct {
	Action TwilioMachineDetectionAction `json:"action" xml:"action" required:"true"`
	// Verdicts that trigger the action.
	AnsweredBy      []TwilioAnsweredBy `json:"answered_by" xml:"answered_by" required:"true" minItems:"1"`
	extraProperties map[string]interface{}
}

func (t *TwilioMachineDetectionRule) GetAction() (value TwilioMachineDetectionAction) {
	if t == nil {
		return
	}
	return t.Action
}

func (t *TwilioMachineDetectionRule) GetAnsweredBy() (value []TwilioAnsweredBy) {
	if t == nil {
		return
	}
	return t.AnsweredBy
}

func (t *TwilioMachineDetectionRule) SetAction(action TwilioMachineDetectionAction) {
	t.Action = action
}

func (t *TwilioMachineDetectionRule) SetAnsweredBy(answeredBy []TwilioAnsweredBy) {
	t.AnsweredBy = answeredBy
}

func (t *TwilioMachineDetectionRule) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioMachineDetectionRule) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioMachineDetectionRule to string"
	}
	return string(jsonData)
}

func (t *TwilioMachineDetectionRule) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TwilioMachineDetectionRule
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TwilioMachineDetectionRule(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioMachineDetectionRule) MarshalJSON() ([]byte, error) {
	type embed TwilioMachineDetectionRule
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
