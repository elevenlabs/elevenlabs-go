package language

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Edit several target translations at once, applied atomically (all succeed or none do).
type DubbingBulkTargetSegmentUpdateRequest struct {
	// Identifier of the language target.
	LanguageId string `json:"-" url:"-"`
	// Identifier of the dubbing project.
	ProjectId string `json:"-" url:"-"`
	// Map of segment ID to the translation edit to apply to that segment. At least one entry and at most 500.
	Segments map[string]*types.DubbingTargetSegmentUpdateRequest `json:"segments" required:"true"`
}

func (d *DubbingBulkTargetSegmentUpdateRequest) GetLanguageId() (value string) {
	if d == nil {
		return
	}
	return d.LanguageId
}

func (d *DubbingBulkTargetSegmentUpdateRequest) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingBulkTargetSegmentUpdateRequest) GetSegments() (value map[string]*types.DubbingTargetSegmentUpdateRequest) {
	if d == nil {
		return
	}
	return d.Segments
}

func (d *DubbingBulkTargetSegmentUpdateRequest) SetLanguageId(languageId string) {
	d.LanguageId = languageId
}

func (d *DubbingBulkTargetSegmentUpdateRequest) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingBulkTargetSegmentUpdateRequest) SetSegments(segments map[string]*types.DubbingTargetSegmentUpdateRequest) {
	d.Segments = segments
}

func (d *DubbingBulkTargetSegmentUpdateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubbingBulkTargetSegmentUpdateRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingBulkTargetSegmentUpdateRequest to string"
	}
	return string(jsonData)
}

func (d *DubbingBulkTargetSegmentUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingBulkTargetSegmentUpdateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingBulkTargetSegmentUpdateRequest(tmp)
	return nil
}

func (d *DubbingBulkTargetSegmentUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed DubbingBulkTargetSegmentUpdateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
