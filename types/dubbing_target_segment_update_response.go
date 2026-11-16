package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The result of a target-translation edit: the updated segment and the new revision.
type DubbingTargetSegmentUpdateResponse struct {
	// One segment of a target transcript: a source segment plus its translation.
	Segment *DubbingTargetTranscriptSegment `json:"segment" xml:"segment" required:"true"`
	// The target's revision after this edit.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingTargetSegmentUpdateResponse) GetSegment() (value *DubbingTargetTranscriptSegment) {
	if d == nil {
		return
	}
	return d.Segment
}

func (d *DubbingTargetSegmentUpdateResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingTargetSegmentUpdateResponse) SetSegment(segment *DubbingTargetTranscriptSegment) {
	d.Segment = segment
}

func (d *DubbingTargetSegmentUpdateResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingTargetSegmentUpdateResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTargetSegmentUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTargetSegmentUpdateResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingTargetSegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTargetSegmentUpdateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTargetSegmentUpdateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTargetSegmentUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingTargetSegmentUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
