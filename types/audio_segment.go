package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioSegment struct {
	StartMs         int     `json:"start_ms" xml:"start_ms" required:"true"`
	EndMs           int     `json:"end_ms" xml:"end_ms" required:"true"`
	Description     string  `json:"description" xml:"description" required:"true"`
	SegmentType     *string `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	HasSpeech       *bool   `json:"has_speech,omitempty" xml:"has_speech,omitempty"`
	HasMusic        *bool   `json:"has_music,omitempty" xml:"has_music,omitempty"`
	Pacing          *string `json:"pacing,omitempty" xml:"pacing,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioSegment) GetStartMs() (value int) {
	if a == nil {
		return
	}
	return a.StartMs
}

func (a *AudioSegment) GetEndMs() (value int) {
	if a == nil {
		return
	}
	return a.EndMs
}

func (a *AudioSegment) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioSegment) GetSegmentType() (value *string) {
	if a == nil {
		return
	}
	return a.SegmentType
}

func (a *AudioSegment) GetHasSpeech() (value *bool) {
	if a == nil {
		return
	}
	return a.HasSpeech
}

func (a *AudioSegment) GetHasMusic() (value *bool) {
	if a == nil {
		return
	}
	return a.HasMusic
}

func (a *AudioSegment) GetPacing() (value *string) {
	if a == nil {
		return
	}
	return a.Pacing
}

func (a *AudioSegment) SetStartMs(startMs int) {
	a.StartMs = startMs
}

func (a *AudioSegment) SetEndMs(endMs int) {
	a.EndMs = endMs
}

func (a *AudioSegment) SetDescription(description string) {
	a.Description = description
}

func (a *AudioSegment) SetSegmentType(segmentType *string) {
	a.SegmentType = segmentType
}

func (a *AudioSegment) SetHasSpeech(hasSpeech *bool) {
	a.HasSpeech = hasSpeech
}

func (a *AudioSegment) SetHasMusic(hasMusic *bool) {
	a.HasMusic = hasMusic
}

func (a *AudioSegment) SetPacing(pacing *string) {
	a.Pacing = pacing
}

func (a *AudioSegment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioSegment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioSegment to string"
	}
	return string(jsonData)
}

func (a *AudioSegment) UnmarshalJSON(data []byte) error {
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

func (a *AudioSegment) MarshalJSON() ([]byte, error) {
	type embed AudioSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
