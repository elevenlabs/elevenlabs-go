package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The result of a bulk target edit: the updated segments and the new revision.
type DubbingBulkTargetSegmentUpdateResponse struct {
	// The edited target segments in their updated state.
	Segments []*DubbingTargetTranscriptSegment `json:"segments" xml:"segments" required:"true"`
	// The target's revision after the edits.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingBulkTargetSegmentUpdateResponse) GetSegments() (value []*DubbingTargetTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingBulkTargetSegmentUpdateResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingBulkTargetSegmentUpdateResponse) SetSegments(segments []*DubbingTargetTranscriptSegment) {
	d.Segments = segments
}

func (d *DubbingBulkTargetSegmentUpdateResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingBulkTargetSegmentUpdateResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingBulkTargetSegmentUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingBulkTargetSegmentUpdateResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingBulkTargetSegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingBulkTargetSegmentUpdateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingBulkTargetSegmentUpdateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingBulkTargetSegmentUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingBulkTargetSegmentUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
