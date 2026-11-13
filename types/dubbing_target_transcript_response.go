package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTargetTranscriptResponse struct {
	// BCP-47 language tag of the source transcript.
	SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
	// BCP-47 language tag this target is translated into.
	TargetLanguage string `json:"target_language" xml:"target_language" required:"true"`
	// The target segments, in playback order.
	Segments []*DubbingTargetTranscriptSegment `json:"segments" xml:"segments" required:"true"`
	// The target's revision at read time.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingTargetTranscriptResponse) GetSourceLanguage() (value *string) {
	if d == nil {
		return
	}
	return d.SourceLanguage
}

func (d *DubbingTargetTranscriptResponse) GetTargetLanguage() (value string) {
	if d == nil {
		return
	}
	return d.TargetLanguage
}

func (d *DubbingTargetTranscriptResponse) GetSegments() (value []*DubbingTargetTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingTargetTranscriptResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingTargetTranscriptResponse) SetSourceLanguage(sourceLanguage *string) {
	d.SourceLanguage = sourceLanguage
}

func (d *DubbingTargetTranscriptResponse) SetTargetLanguage(targetLanguage string) {
	d.TargetLanguage = targetLanguage
}

func (d *DubbingTargetTranscriptResponse) SetSegments(segments []*DubbingTargetTranscriptSegment) {
	d.Segments = segments
}

func (d *DubbingTargetTranscriptResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingTargetTranscriptResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTargetTranscriptResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTargetTranscriptResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingTargetTranscriptResponse) UnmarshalJSON(data []byte) error {
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

func (d *DubbingTargetTranscriptResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingTargetTranscriptResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
