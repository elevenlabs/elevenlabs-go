package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FinalOutput struct {
	// Indicates the generation is complete. When true, audio is null.
	IsFinal         *bool `json:"is_final,omitempty" xml:"is_final,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FinalOutput) GetIsFinal() (value *bool) {
	if f == nil {
		return
	}
	return f.IsFinal
}

func (f *FinalOutput) SetIsFinal(isFinal *bool) {
	f.IsFinal = isFinal
}

func (f *FinalOutput) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FinalOutput) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FinalOutput to string"
	}
	return string(jsonData)
}

func (f *FinalOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler FinalOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FinalOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FinalOutput) MarshalJSON() ([]byte, error) {
	type embed FinalOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
