package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// How to run Twilio's carrier-level answering machine detection (AMD) on a call.
type TwilioMachineDetectionConfig struct {
	// Which Twilio answering-machine-detection (AMD) mode to request for a call.
	//
	// `enable` returns a verdict as soon as Twilio can tell a human from a machine.
	// `detect_message_end` additionally waits for the greeting to finish, which is what
	// distinguishes the `machine_end_*` verdicts, at the cost of a later result.
	Mode *TwilioMachineDetectionMode `json:"mode,omitempty" xml:"mode,omitempty"`
	// Actions to take on the call per AMD verdict. Each action can appear in at most one rule, so list all of its verdicts together. Verdicts without a rule leave the call running.
	Actions         []*TwilioMachineDetectionRule `json:"actions,omitempty" xml:"actions,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TwilioMachineDetectionConfig) GetMode() (value *TwilioMachineDetectionMode) {
	if t == nil {
		return
	}
	return t.Mode
}

func (t *TwilioMachineDetectionConfig) GetActions() (value []*TwilioMachineDetectionRule) {
	if t == nil {
		return
	}
	return t.Actions
}

func (t *TwilioMachineDetectionConfig) SetMode(mode *TwilioMachineDetectionMode) {
	t.Mode = mode
}

func (t *TwilioMachineDetectionConfig) SetActions(actions []*TwilioMachineDetectionRule) {
	t.Actions = actions
}

func (t *TwilioMachineDetectionConfig) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioMachineDetectionConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioMachineDetectionConfig to string"
	}
	return string(jsonData)
}

func (t *TwilioMachineDetectionConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler TwilioMachineDetectionConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TwilioMachineDetectionConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioMachineDetectionConfig) MarshalJSON() ([]byte, error) {
	type embed TwilioMachineDetectionConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
