package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioChunkMulti struct {
	Audio               *string                             `json:"audio,omitempty" xml:"audio,omitempty"`
	Alignment           *AudioChunkMultiAlignment           `json:"alignment,omitempty" xml:"alignment,omitempty"`
	NormalizedAlignment *AudioChunkMultiNormalizedAlignment `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	// Identifies which context this response belongs to. Only set on the multi-context endpoint; omitted (None) on the single-context endpoint.
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioChunkMulti) GetAudio() (value *string) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioChunkMulti) GetAlignment() (value *AudioChunkMultiAlignment) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioChunkMulti) GetNormalizedAlignment() (value *AudioChunkMultiNormalizedAlignment) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioChunkMulti) GetContextId() (value *string) {
	if a == nil {
		return
	}
	return a.ContextId
}

func (a *AudioChunkMulti) SetAudio(audio *string) {
	a.Audio = audio
}

func (a *AudioChunkMulti) SetAlignment(alignment *AudioChunkMultiAlignment) {
	a.Alignment = alignment
}

func (a *AudioChunkMulti) SetNormalizedAlignment(normalizedAlignment *AudioChunkMultiNormalizedAlignment) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioChunkMulti) SetContextId(contextId *string) {
	a.ContextId = contextId
}

func (a *AudioChunkMulti) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioChunkMulti) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioChunkMulti to string"
	}
	return string(jsonData)
}

func (a *AudioChunkMulti) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioChunkMulti
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioChunkMulti(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioChunkMulti) MarshalJSON() ([]byte, error) {
	type embed AudioChunkMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
