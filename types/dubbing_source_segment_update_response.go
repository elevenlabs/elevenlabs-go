package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The result of a source-segment add or edit: the segment and the new revision.
type DubbingSourceSegmentUpdateResponse struct {
	// One segment of a source transcript.
	Segment *DubbingTranscriptSegment `json:"segment" xml:"segment" required:"true"`
	// The project's source-transcript revision after this edit.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingSourceSegmentUpdateResponse) GetSegment() (value *DubbingTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segment
}

func (d *DubbingSourceSegmentUpdateResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingSourceSegmentUpdateResponse) SetSegment(segment *DubbingTranscriptSegment) {
	d.Segment = segment
}

func (d *DubbingSourceSegmentUpdateResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingSourceSegmentUpdateResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingSourceSegmentUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSourceSegmentUpdateResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingSourceSegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingSourceSegmentUpdateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingSourceSegmentUpdateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingSourceSegmentUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingSourceSegmentUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
