package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioChunk struct {
	Audio               *string                        `json:"audio,omitempty" xml:"audio,omitempty"`
	Alignment           *AudioChunkAlignment           `json:"alignment,omitempty" xml:"alignment,omitempty"`
	NormalizedAlignment *AudioChunkNormalizedAlignment `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	extraProperties     map[string]interface{}
}

func (a *AudioChunk) GetAudio() (value *string) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioChunk) GetAlignment() (value *AudioChunkAlignment) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioChunk) GetNormalizedAlignment() (value *AudioChunkNormalizedAlignment) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioChunk) SetAudio(audio *string) {
	a.Audio = audio
}

func (a *AudioChunk) SetAlignment(alignment *AudioChunkAlignment) {
	a.Alignment = alignment
}

func (a *AudioChunk) SetNormalizedAlignment(normalizedAlignment *AudioChunkNormalizedAlignment) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioChunk) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioChunk) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioChunk to string"
	}
	return string(jsonData)
}

func (a *AudioChunk) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioChunk
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioChunk(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioChunk) MarshalJSON() ([]byte, error) {
	type embed AudioChunk
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
