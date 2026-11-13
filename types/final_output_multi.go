package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FinalOutputMulti struct {
	// Indicates the generation is complete. When true, audio is null.
	IsFinal         *bool   `json:"is_final,omitempty" xml:"is_final,omitempty"`
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FinalOutputMulti) GetIsFinal() (value *bool) {
	if f == nil {
		return
	}
	return f.IsFinal
}

func (f *FinalOutputMulti) GetContextId() (value *string) {
	if f == nil {
		return
	}
	return f.ContextId
}

func (f *FinalOutputMulti) SetIsFinal(isFinal *bool) {
	f.IsFinal = isFinal
}

func (f *FinalOutputMulti) SetContextId(contextId *string) {
	f.ContextId = contextId
}

func (f *FinalOutputMulti) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FinalOutputMulti) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FinalOutputMulti to string"
	}
	return string(jsonData)
}

func (f *FinalOutputMulti) UnmarshalJSON(data []byte) error {
	type unmarshaler FinalOutputMulti
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FinalOutputMulti(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FinalOutputMulti) MarshalJSON() ([]byte, error) {
	type embed FinalOutputMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
