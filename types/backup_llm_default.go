package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BackupLlmDefault struct {
	Preference      *string `json:"preference,omitempty" xml:"preference,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BackupLlmDefault) GetPreference() (value *string) {
	if b == nil {
		return
	}
	return b.Preference
}

func (b *BackupLlmDefault) SetPreference(preference *string) {
	b.Preference = preference
}

func (b *BackupLlmDefault) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BackupLlmDefault) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BackupLlmDefault to string"
	}
	return string(jsonData)
}

func (b *BackupLlmDefault) UnmarshalJSON(data []byte) error {
	type unmarshaler BackupLlmDefault
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BackupLlmDefault(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BackupLlmDefault) MarshalJSON() ([]byte, error) {
	type embed BackupLlmDefault
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
