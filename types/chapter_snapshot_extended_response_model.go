package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterSnapshotExtendedResponseModel struct {
	// The ID of the chapter snapshot.
	ChapterSnapshotId string `json:"chapter_snapshot_id" xml:"chapter_snapshot_id" required:"true"`
	// The ID of the project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// The ID of the chapter.
	ChapterId string `json:"chapter_id" xml:"chapter_id" required:"true"`
	// The creation date of the chapter snapshot.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// The name of the chapter snapshot.
	Name                string                     `json:"name" xml:"name" required:"true"`
	CharacterAlignments []*CharacterAlignmentModel `json:"character_alignments" xml:"character_alignments" required:"true"`
	extraProperties     map[string]interface{}
}

func (c *ChapterSnapshotExtendedResponseModel) GetChapterSnapshotId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterSnapshotId
}

func (c *ChapterSnapshotExtendedResponseModel) GetProjectId() (value string) {
	if c == nil {
		return
	}
	return c.ProjectId
}

func (c *ChapterSnapshotExtendedResponseModel) GetChapterId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterId
}

func (c *ChapterSnapshotExtendedResponseModel) GetCreatedAtUnix() (value int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnix
}

func (c *ChapterSnapshotExtendedResponseModel) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ChapterSnapshotExtendedResponseModel) GetCharacterAlignments() (value []*CharacterAlignmentModel) {
	if c == nil {
		return
	}
	return c.CharacterAlignments
}

func (c *ChapterSnapshotExtendedResponseModel) SetChapterSnapshotId(chapterSnapshotId string) {
	c.ChapterSnapshotId = chapterSnapshotId
}

func (c *ChapterSnapshotExtendedResponseModel) SetProjectId(projectId string) {
	c.ProjectId = projectId
}

func (c *ChapterSnapshotExtendedResponseModel) SetChapterId(chapterId string) {
	c.ChapterId = chapterId
}

func (c *ChapterSnapshotExtendedResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	c.CreatedAtUnix = createdAtUnix
}

func (c *ChapterSnapshotExtendedResponseModel) SetName(name string) {
	c.Name = name
}

func (c *ChapterSnapshotExtendedResponseModel) SetCharacterAlignments(characterAlignments []*CharacterAlignmentModel) {
	c.CharacterAlignments = characterAlignments
}

func (c *ChapterSnapshotExtendedResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterSnapshotExtendedResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterSnapshotExtendedResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterSnapshotExtendedResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterSnapshotExtendedResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterSnapshotExtendedResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterSnapshotExtendedResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterSnapshotExtendedResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
