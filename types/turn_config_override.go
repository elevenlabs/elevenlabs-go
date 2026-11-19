package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TurnConfigOverride struct {
	// Configuration for soft timeout functionality. Provides immediate feedback during longer LLM responses.
	SoftTimeoutConfig *SoftTimeoutConfigOverride `json:"soft_timeout_config,omitempty" xml:"soft_timeout_config,omitempty"`
	// What the agent says when it detects a noisy environment.
	NoisyEnvironmentMessage *string `json:"noisy_environment_message,omitempty" xml:"noisy_environment_message,omitempty"`
	extraProperties         map[string]interface{}
}

func (t *TurnConfigOverride) GetSoftTimeoutConfig() (value *SoftTimeoutConfigOverride) {
	if t == nil {
		return
	}
	return t.SoftTimeoutConfig
}

func (t *TurnConfigOverride) GetNoisyEnvironmentMessage() (value *string) {
	if t == nil {
		return
	}
	return t.NoisyEnvironmentMessage
}

func (t *TurnConfigOverride) SetSoftTimeoutConfig(softTimeoutConfig *SoftTimeoutConfigOverride) {
	t.SoftTimeoutConfig = softTimeoutConfig
}

func (t *TurnConfigOverride) SetNoisyEnvironmentMessage(noisyEnvironmentMessage *string) {
	t.NoisyEnvironmentMessage = noisyEnvironmentMessage
}

func (t *TurnConfigOverride) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnConfigOverride to string"
	}
	return string(jsonData)
}

func (t *TurnConfigOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TurnConfigOverride) MarshalJSON() ([]byte, error) {
	type embed TurnConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
