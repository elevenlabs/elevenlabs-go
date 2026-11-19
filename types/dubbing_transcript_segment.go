package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One segment of a source transcript.
type DubbingTranscriptSegment struct {
	// Stable identifier of the segment, used to address it in edit requests.
	Id string `json:"id" xml:"id" required:"true"`
	// The transcribed text of the segment.
	Text string `json:"text" xml:"text" required:"true"`
	// Identifier of the segment's speaker.
	SpeakerId string `json:"speaker_id" xml:"speaker_id" required:"true"`
	// Start time of the segment, in seconds.
	StartS float64 `json:"start_s" xml:"start_s" required:"true"`
	// End time of the segment, in seconds.
	EndS float64 `json:"end_s" xml:"end_s" required:"true"`
	// The caller-supplied external ID for this segment, if one was provided.
	ExternalId      *string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptSegment) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DubbingTranscriptSegment) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingTranscriptSegment) GetSpeakerId() (value string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingTranscriptSegment) GetStartS() (value float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTranscriptSegment) GetEndS() (value float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTranscriptSegment) GetExternalId() (value *string) {
	if d == nil {
		return
	}
	return d.ExternalId
}

func (d *DubbingTranscriptSegment) SetId(id string) {
	d.Id = id
}

func (d *DubbingTranscriptSegment) SetText(text string) {
	d.Text = text
}

func (d *DubbingTranscriptSegment) SetSpeakerId(speakerId string) {
	d.SpeakerId = speakerId
}

func (d *DubbingTranscriptSegment) SetStartS(startS float64) {
	d.StartS = startS
}

func (d *DubbingTranscriptSegment) SetEndS(endS float64) {
	d.EndS = endS
}

func (d *DubbingTranscriptSegment) SetExternalId(externalId *string) {
	d.ExternalId = externalId
}

func (d *DubbingTranscriptSegment) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptSegment) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptSegment to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptSegment) UnmarshalJSON(data []byte) error {
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

func (d *DubbingTranscriptSegment) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
