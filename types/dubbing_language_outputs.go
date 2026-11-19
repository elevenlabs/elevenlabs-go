package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Signed, time-limited download URLs for a language target's outputs.
type DubbingLanguageOutputs struct {
	// Signed URL for the dubbed lossless audio track, in FLAC. The link expires one hour after it is issued; re-read the language target for a fresh one.
	LosslessAudio *string `json:"lossless_audio,omitempty" xml:"lossless_audio,omitempty"`
	// Signed URL for the translated lines as WebVTT subtitles, timed to the source speech. Null when none are available, such as for languages dubbed before subtitles were produced. The link expires one hour after it is issued.
	Subtitles       *string `json:"subtitles,omitempty" xml:"subtitles,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingLanguageOutputs) GetLosslessAudio() (value *string) {
	if d == nil {
		return
	}
	return d.LosslessAudio
}

func (d *DubbingLanguageOutputs) GetSubtitles() (value *string) {
	if d == nil {
		return
	}
	return d.Subtitles
}

func (d *DubbingLanguageOutputs) SetLosslessAudio(losslessAudio *string) {
	d.LosslessAudio = losslessAudio
}

func (d *DubbingLanguageOutputs) SetSubtitles(subtitles *string) {
	d.Subtitles = subtitles
}

func (d *DubbingLanguageOutputs) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingLanguageOutputs) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingLanguageOutputs to string"
	}
	return string(jsonData)
}

func (d *DubbingLanguageOutputs) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingLanguageOutputs) MarshalJSON() ([]byte, error) {
	type embed DubbingLanguageOutputs
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
