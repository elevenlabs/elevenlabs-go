package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioOutput struct {
	Audio               string                          `json:"audio" xml:"audio" required:"true"`
	IsFinal             *bool                           `json:"is_final,omitempty" xml:"is_final,omitempty"`
	NormalizedAlignment *AudioOutputNormalizedAlignment `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	Alignment           *AudioOutputAlignment           `json:"alignment,omitempty" xml:"alignment,omitempty"`
	extraProperties     map[string]interface{}
}

func (a *AudioOutput) GetAudio() (value string) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioOutput) GetIsFinal() (value *bool) {
	if a == nil {
		return
	}
	return a.IsFinal
}

func (a *AudioOutput) GetNormalizedAlignment() (value *AudioOutputNormalizedAlignment) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioOutput) GetAlignment() (value *AudioOutputAlignment) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioOutput) SetAudio(audio string) {
	a.Audio = audio
}

func (a *AudioOutput) SetIsFinal(isFinal *bool) {
	a.IsFinal = isFinal
}

func (a *AudioOutput) SetNormalizedAlignment(normalizedAlignment *AudioOutputNormalizedAlignment) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioOutput) SetAlignment(alignment *AudioOutputAlignment) {
	a.Alignment = alignment
}

func (a *AudioOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioOutput to string"
	}
	return string(jsonData)
}

func (a *AudioOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioOutput) MarshalJSON() ([]byte, error) {
	type embed AudioOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
