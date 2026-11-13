package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BackupLlmOverride struct {
	Preference      *string `json:"preference,omitempty" xml:"preference,omitempty"`
	Order           []Llm   `json:"order" xml:"order" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BackupLlmOverride) GetPreference() (value *string) {
	if b == nil {
		return
	}
	return b.Preference
}

func (b *BackupLlmOverride) GetOrder() (value []Llm) {
	if b == nil {
		return
	}
	return b.Order
}

func (b *BackupLlmOverride) SetPreference(preference *string) {
	b.Preference = preference
}

func (b *BackupLlmOverride) SetOrder(order []Llm) {
	b.Order = order
}

func (b *BackupLlmOverride) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BackupLlmOverride) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BackupLlmOverride to string"
	}
	return string(jsonData)
}

func (b *BackupLlmOverride) UnmarshalJSON(data []byte) error {
	type unmarshaler BackupLlmOverride
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BackupLlmOverride(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BackupLlmOverride) MarshalJSON() ([]byte, error) {
	type embed BackupLlmOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
