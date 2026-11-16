package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectSnapshotResponse struct {
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
	ZipUpload       map[string]any `json:"zip_upload,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProjectSnapshotResponse) GetProjectSnapshotId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectSnapshotId
}

func (p *ProjectSnapshotResponse) GetProjectId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectId
}

func (p *ProjectSnapshotResponse) GetCreatedAtUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtUnix
}

func (p *ProjectSnapshotResponse) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProjectSnapshotResponse) GetAudioUpload() (value map[string]any) {
	if p == nil {
		return
	}
	return p.AudioUpload
}

func (p *ProjectSnapshotResponse) GetZipUpload() (value map[string]any) {
	if p == nil {
		return
	}
	return p.ZipUpload
}

func (p *ProjectSnapshotResponse) SetProjectSnapshotId(projectSnapshotId string) {
	p.ProjectSnapshotId = projectSnapshotId
}

func (p *ProjectSnapshotResponse) SetProjectId(projectId string) {
	p.ProjectId = projectId
}

func (p *ProjectSnapshotResponse) SetCreatedAtUnix(createdAtUnix int) {
	p.CreatedAtUnix = createdAtUnix
}

func (p *ProjectSnapshotResponse) SetName(name string) {
	p.Name = name
}

func (p *ProjectSnapshotResponse) SetAudioUpload(audioUpload map[string]any) {
	p.AudioUpload = audioUpload
}

func (p *ProjectSnapshotResponse) SetZipUpload(zipUpload map[string]any) {
	p.ZipUpload = zipUpload
}

func (p *ProjectSnapshotResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectSnapshotResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectSnapshotResponse to string"
	}
	return string(jsonData)
}

func (p *ProjectSnapshotResponse) UnmarshalJSON(data []byte) error {
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

func (p *ProjectSnapshotResponse) MarshalJSON() ([]byte, error) {
	type embed ProjectSnapshotResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
