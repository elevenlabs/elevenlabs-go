package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TurnConfigOverrideConfig struct {
	SoftTimeoutConfig *SoftTimeoutConfigOverrideConfig `json:"soft_timeout_config,omitempty" xml:"soft_timeout_config,omitempty"`
	// Whether to allow overriding the noisy_environment_message field.
	NoisyEnvironmentMessage *bool `json:"noisy_environment_message,omitempty" xml:"noisy_environment_message,omitempty"`
	extraProperties         map[string]interface{}
}

func (t *TurnConfigOverrideConfig) GetSoftTimeoutConfig() (value *SoftTimeoutConfigOverrideConfig) {
	if t == nil {
		return
	}
	return t.SoftTimeoutConfig
}

func (t *TurnConfigOverrideConfig) GetNoisyEnvironmentMessage() (value *bool) {
	if t == nil {
		return
	}
	return t.NoisyEnvironmentMessage
}

func (t *TurnConfigOverrideConfig) SetSoftTimeoutConfig(softTimeoutConfig *SoftTimeoutConfigOverrideConfig) {
	t.SoftTimeoutConfig = softTimeoutConfig
}

func (t *TurnConfigOverrideConfig) SetNoisyEnvironmentMessage(noisyEnvironmentMessage *bool) {
	t.NoisyEnvironmentMessage = noisyEnvironmentMessage
}

func (t *TurnConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (t *TurnConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler TurnConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TurnConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TurnConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed TurnConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
