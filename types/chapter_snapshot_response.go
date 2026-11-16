package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterSnapshotResponse struct {
	// The ID of the chapter snapshot.
	ChapterSnapshotId string `json:"chapter_snapshot_id" xml:"chapter_snapshot_id" required:"true"`
	// The ID of the project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// The ID of the chapter.
	ChapterId string `json:"chapter_id" xml:"chapter_id" required:"true"`
	// The creation date of the chapter snapshot.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// The name of the chapter snapshot.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterSnapshotResponse) GetChapterSnapshotId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterSnapshotId
}

func (c *ChapterSnapshotResponse) GetProjectId() (value string) {
	if c == nil {
		return
	}
	return c.ProjectId
}

func (c *ChapterSnapshotResponse) GetChapterId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterId
}

func (c *ChapterSnapshotResponse) GetCreatedAtUnix() (value int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnix
}

func (c *ChapterSnapshotResponse) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ChapterSnapshotResponse) SetChapterSnapshotId(chapterSnapshotId string) {
	c.ChapterSnapshotId = chapterSnapshotId
}

func (c *ChapterSnapshotResponse) SetProjectId(projectId string) {
	c.ProjectId = projectId
}

func (c *ChapterSnapshotResponse) SetChapterId(chapterId string) {
	c.ChapterId = chapterId
}

func (c *ChapterSnapshotResponse) SetCreatedAtUnix(createdAtUnix int) {
	c.CreatedAtUnix = createdAtUnix
}

func (c *ChapterSnapshotResponse) SetName(name string) {
	c.Name = name
}

func (c *ChapterSnapshotResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterSnapshotResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterSnapshotResponse to string"
	}
	return string(jsonData)
}

func (c *ChapterSnapshotResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterSnapshotResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterSnapshotResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterSnapshotResponse) MarshalJSON() ([]byte, error) {
	type embed ChapterSnapshotResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
