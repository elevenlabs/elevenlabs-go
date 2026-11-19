package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AsrConversationalConfigOverrideConfig struct {
	// Whether to allow overriding the keywords field.
	Keywords        *bool `json:"keywords,omitempty" xml:"keywords,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AsrConversationalConfigOverrideConfig) GetKeywords() (value *bool) {
	if a == nil {
		return
	}
	return a.Keywords
}

func (a *AsrConversationalConfigOverrideConfig) SetKeywords(keywords *bool) {
	a.Keywords = keywords
}

func (a *AsrConversationalConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AsrConversationalConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsrConversationalConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (a *AsrConversationalConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler AsrConversationalConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AsrConversationalConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AsrConversationalConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed AsrConversationalConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
