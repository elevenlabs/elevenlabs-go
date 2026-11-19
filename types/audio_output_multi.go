package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioOutputMulti struct {
	Audio               string                               `json:"audio" xml:"audio" required:"true"`
	IsFinal             *bool                                `json:"is_final,omitempty" xml:"is_final,omitempty"`
	NormalizedAlignment *AudioOutputMultiNormalizedAlignment `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	Alignment           *AudioOutputMultiAlignment           `json:"alignment,omitempty" xml:"alignment,omitempty"`
	ContextId           *string                              `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties     map[string]interface{}
}

func (a *AudioOutputMulti) GetAudio() (value string) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioOutputMulti) GetIsFinal() (value *bool) {
	if a == nil {
		return
	}
	return a.IsFinal
}

func (a *AudioOutputMulti) GetNormalizedAlignment() (value *AudioOutputMultiNormalizedAlignment) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioOutputMulti) GetAlignment() (value *AudioOutputMultiAlignment) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioOutputMulti) GetContextId() (value *string) {
	if a == nil {
		return
	}
	return a.ContextId
}

func (a *AudioOutputMulti) SetAudio(audio string) {
	a.Audio = audio
}

func (a *AudioOutputMulti) SetIsFinal(isFinal *bool) {
	a.IsFinal = isFinal
}

func (a *AudioOutputMulti) SetNormalizedAlignment(normalizedAlignment *AudioOutputMultiNormalizedAlignment) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioOutputMulti) SetAlignment(alignment *AudioOutputMultiAlignment) {
	a.Alignment = alignment
}

func (a *AudioOutputMulti) SetContextId(contextId *string) {
	a.ContextId = contextId
}

func (a *AudioOutputMulti) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioOutputMulti) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioOutputMulti to string"
	}
	return string(jsonData)
}

func (a *AudioOutputMulti) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioOutputMulti
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioOutputMulti(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioOutputMulti) MarshalJSON() ([]byte, error) {
	type embed AudioOutputMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
