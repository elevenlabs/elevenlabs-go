package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BackupLlmDisabled struct {
	Preference      *string `json:"preference,omitempty" xml:"preference,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BackupLlmDisabled) GetPreference() (value *string) {
	if b == nil {
		return
	}
	return b.Preference
}

func (b *BackupLlmDisabled) SetPreference(preference *string) {
	b.Preference = preference
}

func (b *BackupLlmDisabled) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BackupLlmDisabled) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BackupLlmDisabled to string"
	}
	return string(jsonData)
}

func (b *BackupLlmDisabled) UnmarshalJSON(data []byte) error {
	type unmarshaler BackupLlmDisabled
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BackupLlmDisabled(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BackupLlmDisabled) MarshalJSON() ([]byte, error) {
	type embed BackupLlmDisabled
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
