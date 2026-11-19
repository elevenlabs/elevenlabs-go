package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A partial edit to a target segment. An omitted field is left unchanged; a provided `null`
// clears it (see each field for what clearing means).
type DubbingTargetSegmentUpdateRequest struct {
	// Identifier of the segment to edit.
	SegmentId string `json:"-" url:"-"`
	// Identifier of the language target.
	LanguageId string `json:"-" url:"-"`
	// Identifier of the dubbing project.
	ProjectId string `json:"-" url:"-"`
	// New translated text, or null to mark the segment for re-translation.
	Translation *string `json:"translation,omitempty" xml:"translation,omitempty"`
}

func (d *DubbingTargetSegmentUpdateRequest) GetSegmentId() (value string) {
	if d == nil {
		return
	}
	return d.SegmentId
}

func (d *DubbingTargetSegmentUpdateRequest) GetLanguageId() (value string) {
	if d == nil {
		return
	}
	return d.LanguageId
}

func (d *DubbingTargetSegmentUpdateRequest) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingTargetSegmentUpdateRequest) GetTranslation() (value *string) {
	if d == nil {
		return
	}
	return d.Translation
}

func (d *DubbingTargetSegmentUpdateRequest) SetSegmentId(segmentId string) {
	d.SegmentId = segmentId
}

func (d *DubbingTargetSegmentUpdateRequest) SetLanguageId(languageId string) {
	d.LanguageId = languageId
}

func (d *DubbingTargetSegmentUpdateRequest) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingTargetSegmentUpdateRequest) SetTranslation(translation *string) {
	d.Translation = translation
}

func (d *DubbingTargetSegmentUpdateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubbingTargetSegmentUpdateRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTargetSegmentUpdateRequest to string"
	}
	return string(jsonData)
}

func (d *DubbingTargetSegmentUpdateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}

func (d *DubbingTargetSegmentUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed DubbingTargetSegmentUpdateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
