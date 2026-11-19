package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectSnapshotExtendedResponseModel struct {
	// The ID of the project snapshot.
	ProjectSnapshotId string `json:"project_snapshot_id" xml:"project_snapshot_id" required:"true"`
	// The ID of the project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// The creation date of the project snapshot.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// The name of the project snapshot.
	Name string `json:"name" xml:"name" required:"true"`
	// (Deprecated)
	AudioUpload map[string]any `json:"audio_upload,omitempty"`
	// (Deprecated)
	ZipUpload           map[string]any             `json:"zip_upload,omitempty"`
	CharacterAlignments []*CharacterAlignmentModel `json:"character_alignments" xml:"character_alignments" required:"true"`
	// The total duration of the audio in seconds.
	AudioDurationSecs float64 `json:"audio_duration_secs" xml:"audio_duration_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (p *ProjectSnapshotExtendedResponseModel) GetProjectSnapshotId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectSnapshotId
}

func (p *ProjectSnapshotExtendedResponseModel) GetProjectId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectId
}

func (p *ProjectSnapshotExtendedResponseModel) GetCreatedAtUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtUnix
}

func (p *ProjectSnapshotExtendedResponseModel) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProjectSnapshotExtendedResponseModel) GetAudioUpload() (value map[string]any) {
	if p == nil {
		return
	}
	return p.AudioUpload
}

func (p *ProjectSnapshotExtendedResponseModel) GetZipUpload() (value map[string]any) {
	if p == nil {
		return
	}
	return p.ZipUpload
}

func (p *ProjectSnapshotExtendedResponseModel) GetCharacterAlignments() (value []*CharacterAlignmentModel) {
	if p == nil {
		return
	}
	return p.CharacterAlignments
}

func (p *ProjectSnapshotExtendedResponseModel) GetAudioDurationSecs() (value float64) {
	if p == nil {
		return
	}
	return p.AudioDurationSecs
}

func (p *ProjectSnapshotExtendedResponseModel) SetProjectSnapshotId(projectSnapshotId string) {
	p.ProjectSnapshotId = projectSnapshotId
}

func (p *ProjectSnapshotExtendedResponseModel) SetProjectId(projectId string) {
	p.ProjectId = projectId
}

func (p *ProjectSnapshotExtendedResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	p.CreatedAtUnix = createdAtUnix
}

func (p *ProjectSnapshotExtendedResponseModel) SetName(name string) {
	p.Name = name
}

func (p *ProjectSnapshotExtendedResponseModel) SetAudioUpload(audioUpload map[string]any) {
	p.AudioUpload = audioUpload
}

func (p *ProjectSnapshotExtendedResponseModel) SetZipUpload(zipUpload map[string]any) {
	p.ZipUpload = zipUpload
}

func (p *ProjectSnapshotExtendedResponseModel) SetCharacterAlignments(characterAlignments []*CharacterAlignmentModel) {
	p.CharacterAlignments = characterAlignments
}

func (p *ProjectSnapshotExtendedResponseModel) SetAudioDurationSecs(audioDurationSecs float64) {
	p.AudioDurationSecs = audioDurationSecs
}

func (p *ProjectSnapshotExtendedResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectSnapshotExtendedResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectSnapshotExtendedResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectSnapshotExtendedResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectSnapshotExtendedResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectSnapshotExtendedResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
