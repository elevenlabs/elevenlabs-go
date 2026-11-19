package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingResource struct {
	Id              string                     `json:"id" xml:"id" required:"true"`
	Version         int                        `json:"version" xml:"version" required:"true"`
	SourceLanguage  string                     `json:"source_language" xml:"source_language" required:"true"`
	TargetLanguages []string                   `json:"target_languages" xml:"target_languages" required:"true"`
	Input           *DubbingMediaReference     `json:"input" xml:"input" required:"true"`
	Background      *DubbingMediaReference     `json:"background" xml:"background" required:"true"`
	Foreground      *DubbingMediaReference     `json:"foreground" xml:"foreground" required:"true"`
	SpeakerTracks   map[string]*SpeakerTrack   `json:"speaker_tracks" required:"true"`
	SpeakerSegments map[string]*SpeakerSegment `json:"speaker_segments" required:"true"`
	Renders         map[string]*Render         `json:"renders" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingResource) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DubbingResource) GetVersion() (value int) {
	if d == nil {
		return
	}
	return d.Version
}

func (d *DubbingResource) GetSourceLanguage() (value string) {
	if d == nil {
		return
	}
	return d.SourceLanguage
}

func (d *DubbingResource) GetTargetLanguages() (value []string) {
	if d == nil {
		return
	}
	return d.TargetLanguages
}

func (d *DubbingResource) GetInput() (value *DubbingMediaReference) {
	if d == nil {
		return
	}
	return d.Input
}

func (d *DubbingResource) GetBackground() (value *DubbingMediaReference) {
	if d == nil {
		return
	}
	return d.Background
}

func (d *DubbingResource) GetForeground() (value *DubbingMediaReference) {
	if d == nil {
		return
	}
	return d.Foreground
}

func (d *DubbingResource) GetSpeakerTracks() (value map[string]*SpeakerTrack) {
	if d == nil {
		return
	}
	return d.SpeakerTracks
}

func (d *DubbingResource) GetSpeakerSegments() (value map[string]*SpeakerSegment) {
	if d == nil {
		return
	}
	return d.SpeakerSegments
}

func (d *DubbingResource) GetRenders() (value map[string]*Render) {
	if d == nil {
		return
	}
	return d.Renders
}

func (d *DubbingResource) SetId(id string) {
	d.Id = id
}

func (d *DubbingResource) SetVersion(version int) {
	d.Version = version
}

func (d *DubbingResource) SetSourceLanguage(sourceLanguage string) {
	d.SourceLanguage = sourceLanguage
}

func (d *DubbingResource) SetTargetLanguages(targetLanguages []string) {
	d.TargetLanguages = targetLanguages
}

func (d *DubbingResource) SetInput(input *DubbingMediaReference) {
	d.Input = input
}

func (d *DubbingResource) SetBackground(background *DubbingMediaReference) {
	d.Background = background
}

func (d *DubbingResource) SetForeground(foreground *DubbingMediaReference) {
	d.Foreground = foreground
}

func (d *DubbingResource) SetSpeakerTracks(speakerTracks map[string]*SpeakerTrack) {
	d.SpeakerTracks = speakerTracks
}

func (d *DubbingResource) SetSpeakerSegments(speakerSegments map[string]*SpeakerSegment) {
	d.SpeakerSegments = speakerSegments
}

func (d *DubbingResource) SetRenders(renders map[string]*Render) {
	d.Renders = renders
}

func (d *DubbingResource) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingResource) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingResource to string"
	}
	return string(jsonData)
}

func (d *DubbingResource) UnmarshalJSON(data []byte) error {
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

func (d *DubbingResource) MarshalJSON() ([]byte, error) {
	type embed DubbingResource
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
