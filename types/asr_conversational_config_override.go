package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AsrConversationalConfigOverride struct {
	// Keywords to boost prediction probability for
	Keywords        []string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AsrConversationalConfigOverride) GetKeywords() (value []string) {
	if a == nil {
		return
	}
	return a.Keywords
}

func (a *AsrConversationalConfigOverride) SetKeywords(keywords []string) {
	a.Keywords = keywords
}

func (a *AsrConversationalConfigOverride) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AsrConversationalConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsrConversationalConfigOverride to string"
	}
	return string(jsonData)
}

func (a *AsrConversationalConfigOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AsrConversationalConfigOverride) MarshalJSON() ([]byte, error) {
	type embed AsrConversationalConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
