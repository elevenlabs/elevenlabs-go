package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FinalMulti struct {
	// Marks the end of the closing flush sequence.
	IsFinal bool `json:"is_final" xml:"is_final" required:"true"`
	// Identifies which context this response belongs to. Only set on the multi-context endpoint; omitted (None) on the single-context endpoint.
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FinalMulti) GetIsFinal() (value bool) {
	if f == nil {
		return
	}
	return f.IsFinal
}

func (f *FinalMulti) GetContextId() (value *string) {
	if f == nil {
		return
	}
	return f.ContextId
}

func (f *FinalMulti) SetIsFinal(isFinal bool) {
	f.IsFinal = isFinal
}

func (f *FinalMulti) SetContextId(contextId *string) {
	f.ContextId = contextId
}

func (f *FinalMulti) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FinalMulti) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FinalMulti to string"
	}
	return string(jsonData)
}

func (f *FinalMulti) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FinalMulti
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FinalMulti(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FinalMulti) MarshalJSON() ([]byte, error) {
	type embed FinalMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
