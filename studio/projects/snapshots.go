package projects

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost struct {
	// The ID of the Studio project snapshot.
	ProjectSnapshotId string `json:"-" url:"-"`
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// Whether to convert the audio to mpeg format.
	ConvertToMpeg *bool `json:"convert_to_mpeg,omitempty" xml:"convert_to_mpeg,omitempty"`
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetProjectSnapshotId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectSnapshotId
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetConvertToMpeg() (value *bool) {
	if b == nil {
		return
	}
	return b.ConvertToMpeg
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) SetProjectSnapshotId(projectSnapshotId string) {
	b.ProjectSnapshotId = projectSnapshotId
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) SetConvertToMpeg(convertToMpeg *bool) {
	b.ConvertToMpeg = convertToMpeg
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost to string"
	}
	return string(jsonData)
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost(tmp)
	return nil
}

func (b *BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) MarshalJSON() ([]byte, error) {
	type embed BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
