package project

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Edit several source segments at once, applied atomically (all succeed or none do).
type DubbingBulkSegmentUpdateRequest struct {
	// Identifier of the dubbing project.
	ProjectId string `json:"-" url:"-"`
	// Map of segment ID to the partial update to apply to that segment. At least one entry and at most 500.
	Segments map[string]*types.DubbingSegmentUpdateRequest `json:"segments" required:"true"`
}

func (d *DubbingBulkSegmentUpdateRequest) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingBulkSegmentUpdateRequest) GetSegments() (value map[string]*types.DubbingSegmentUpdateRequest) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingBulkSegmentUpdateRequest) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingBulkSegmentUpdateRequest) SetSegments(segments map[string]*types.DubbingSegmentUpdateRequest) {
	d.Segments = segments
}

func (d *DubbingBulkSegmentUpdateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubbingBulkSegmentUpdateRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingBulkSegmentUpdateRequest to string"
	}
	return string(jsonData)
}

func (d *DubbingBulkSegmentUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingBulkSegmentUpdateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingBulkSegmentUpdateRequest(tmp)
	return nil
}

func (d *DubbingBulkSegmentUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed DubbingBulkSegmentUpdateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}

// A new source segment; its ID and word-level alignment are assigned by the server.
type DubbingSegmentCreateRequest struct {
	// Identifier of the dubbing project.
	ProjectId string `json:"-" url:"-"`
	// The text of the new segment.
	Text string `json:"text" xml:"text" required:"true"`
	// Identifier of the segment's speaker.
	SpeakerId string `json:"speaker_id" xml:"speaker_id" required:"true"`
	// Start time of the segment, in seconds.
	StartS float64 `json:"start_s" xml:"start_s" required:"true"`
	// End time of the segment, in seconds.
	EndS float64 `json:"end_s" xml:"end_s" required:"true"`
}

func (d *DubbingSegmentCreateRequest) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingSegmentCreateRequest) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingSegmentCreateRequest) GetSpeakerId() (value string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingSegmentCreateRequest) GetStartS() (value float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingSegmentCreateRequest) GetEndS() (value float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingSegmentCreateRequest) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingSegmentCreateRequest) SetText(text string) {
	d.Text = text
}

func (d *DubbingSegmentCreateRequest) SetSpeakerId(speakerId string) {
	d.SpeakerId = speakerId
}

func (d *DubbingSegmentCreateRequest) SetStartS(startS float64) {
	d.StartS = startS
}

func (d *DubbingSegmentCreateRequest) SetEndS(endS float64) {
	d.EndS = endS
}

func (d *DubbingSegmentCreateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubbingSegmentCreateRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSegmentCreateRequest to string"
	}
	return string(jsonData)
}

func (d *DubbingSegmentCreateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingSegmentCreateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingSegmentCreateRequest(tmp)
	return nil
}

func (d *DubbingSegmentCreateRequest) MarshalJSON() ([]byte, error) {
	type embed DubbingSegmentCreateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
