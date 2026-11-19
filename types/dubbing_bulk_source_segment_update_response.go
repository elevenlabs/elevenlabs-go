package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The result of a bulk source edit: the updated segments and the new revision.
type DubbingBulkSourceSegmentUpdateResponse struct {
	// The edited segments in their updated state.
	Segments []*DubbingTranscriptSegment `json:"segments" xml:"segments" required:"true"`
	// The project's source-transcript revision after the edits.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingBulkSourceSegmentUpdateResponse) GetSegments() (value []*DubbingTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingBulkSourceSegmentUpdateResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingBulkSourceSegmentUpdateResponse) SetSegments(segments []*DubbingTranscriptSegment) {
	d.Segments = segments
}

func (d *DubbingBulkSourceSegmentUpdateResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingBulkSourceSegmentUpdateResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingBulkSourceSegmentUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingBulkSourceSegmentUpdateResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingBulkSourceSegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingBulkSourceSegmentUpdateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingBulkSourceSegmentUpdateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingBulkSourceSegmentUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingBulkSourceSegmentUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
