package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A partial update to a source segment; omitted fields are left unchanged.
type DubbingSegmentUpdateRequest struct {
	// Identifier of the segment to edit.
	SegmentId string `json:"-" url:"-"`
	// Identifier of the dubbing project.
	ProjectId string `json:"-" url:"-"`
	// New text for the segment.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// New speaker ID for the segment.
	SpeakerId *string `json:"speaker_id,omitempty" xml:"speaker_id,omitempty"`
	// New start time, in seconds.
	StartS *float64 `json:"start_s,omitempty" xml:"start_s,omitempty"`
	// New end time, in seconds.
	EndS *float64 `json:"end_s,omitempty" xml:"end_s,omitempty"`
}

func (d *DubbingSegmentUpdateRequest) GetSegmentId() (value string) {
	if d == nil {
		return
	}
	return d.SegmentId
}

func (d *DubbingSegmentUpdateRequest) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingSegmentUpdateRequest) GetText() (value *string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingSegmentUpdateRequest) GetSpeakerId() (value *string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingSegmentUpdateRequest) GetStartS() (value *float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingSegmentUpdateRequest) GetEndS() (value *float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingSegmentUpdateRequest) SetSegmentId(segmentId string) {
	d.SegmentId = segmentId
}

func (d *DubbingSegmentUpdateRequest) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingSegmentUpdateRequest) SetText(text *string) {
	d.Text = text
}

func (d *DubbingSegmentUpdateRequest) SetSpeakerId(speakerId *string) {
	d.SpeakerId = speakerId
}

func (d *DubbingSegmentUpdateRequest) SetStartS(startS *float64) {
	d.StartS = startS
}

func (d *DubbingSegmentUpdateRequest) SetEndS(endS *float64) {
	d.EndS = endS
}

func (d *DubbingSegmentUpdateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubbingSegmentUpdateRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSegmentUpdateRequest to string"
	}
	return string(jsonData)
}

func (d *DubbingSegmentUpdateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}

func (d *DubbingSegmentUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed DubbingSegmentUpdateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
