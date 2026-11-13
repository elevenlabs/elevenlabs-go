package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmSchemaOverride struct {
	// Prompt override for the LLM. If not provided, the original schema description is used.
	Prompt          *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmSchemaOverride) GetPrompt() (value *string) {
	if l == nil {
		return
	}
	return l.Prompt
}

func (l *LlmSchemaOverride) SetPrompt(prompt *string) {
	l.Prompt = prompt
}

func (l *LlmSchemaOverride) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmSchemaOverride) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmSchemaOverride to string"
	}
	return string(jsonData)
}

func (l *LlmSchemaOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "source")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmSchemaOverride) MarshalJSON() ([]byte, error) {
	type embed LlmSchemaOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
