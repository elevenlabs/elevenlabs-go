package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingSourceTranscriptResponse struct {
	// BCP-47 language tag of the source transcript (null if unknown).
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The source segments, in playback order.
	Segments []*DubbingTranscriptSegment `json:"segments" xml:"segments" required:"true"`
	// The project's source-transcript revision at read time.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingSourceTranscriptResponse) GetLanguage() (value *string) {
	if d == nil {
		return
	}
	return d.Language
}

func (d *DubbingSourceTranscriptResponse) GetSegments() (value []*DubbingTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingSourceTranscriptResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingSourceTranscriptResponse) SetLanguage(language *string) {
	d.Language = language
}

func (d *DubbingSourceTranscriptResponse) SetSegments(segments []*DubbingTranscriptSegment) {
	d.Segments = segments
}

func (d *DubbingSourceTranscriptResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingSourceTranscriptResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingSourceTranscriptResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSourceTranscriptResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingSourceTranscriptResponse) UnmarshalJSON(data []byte) error {
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

func (d *DubbingSourceTranscriptResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingSourceTranscriptResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
