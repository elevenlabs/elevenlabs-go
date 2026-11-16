package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One segment of a target transcript: a source segment plus its translation.
type DubbingTargetTranscriptSegment struct {
	// Stable identifier of the segment (from the source).
	Id string `json:"id" xml:"id" required:"true"`
	// Identifier of the segment's speaker.
	SpeakerId string `json:"speaker_id" xml:"speaker_id" required:"true"`
	// Start time of the segment, in seconds.
	StartS float64 `json:"start_s" xml:"start_s" required:"true"`
	// End time of the segment, in seconds.
	EndS float64 `json:"end_s" xml:"end_s" required:"true"`
	// The source-language text of the segment.
	SourceText string `json:"source_text" xml:"source_text" required:"true"`
	// The translated text, or null if not translated yet (needs translation).
	Translation     *string `json:"translation,omitempty" xml:"translation,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingTargetTranscriptSegment) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DubbingTargetTranscriptSegment) GetSpeakerId() (value string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingTargetTranscriptSegment) GetStartS() (value float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTargetTranscriptSegment) GetEndS() (value float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTargetTranscriptSegment) GetSourceText() (value string) {
	if d == nil {
		return
	}
	return d.SourceText
}

func (d *DubbingTargetTranscriptSegment) GetTranslation() (value *string) {
	if d == nil {
		return
	}
	return d.Translation
}

func (d *DubbingTargetTranscriptSegment) SetId(id string) {
	d.Id = id
}

func (d *DubbingTargetTranscriptSegment) SetSpeakerId(speakerId string) {
	d.SpeakerId = speakerId
}

func (d *DubbingTargetTranscriptSegment) SetStartS(startS float64) {
	d.StartS = startS
}

func (d *DubbingTargetTranscriptSegment) SetEndS(endS float64) {
	d.EndS = endS
}

func (d *DubbingTargetTranscriptSegment) SetSourceText(sourceText string) {
	d.SourceText = sourceText
}

func (d *DubbingTargetTranscriptSegment) SetTranslation(translation *string) {
	d.Translation = translation
}

func (d *DubbingTargetTranscriptSegment) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTargetTranscriptSegment) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTargetTranscriptSegment to string"
	}
	return string(jsonData)
}

func (d *DubbingTargetTranscriptSegment) UnmarshalJSON(data []byte) error {
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

func (d *DubbingTargetTranscriptSegment) MarshalJSON() ([]byte, error) {
	type embed DubbingTargetTranscriptSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
