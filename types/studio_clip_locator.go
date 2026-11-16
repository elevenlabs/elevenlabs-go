package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Coordinates of a clip inside a Studio project: the payload of a studio_clip reference.
type StudioClipLocator struct {
	ProjectId       string                    `json:"project_id" xml:"project_id" required:"true"`
	ChapterId       string                    `json:"chapter_id" xml:"chapter_id" required:"true"`
	ClipType        StudioClipLocatorClipType `json:"clip_type" xml:"clip_type" required:"true"`
	ClipId          string                    `json:"clip_id" xml:"clip_id" required:"true"`
	BlockId         *string                   `json:"block_id,omitempty" xml:"block_id,omitempty"`
	PreviewUrl      *string                   `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StudioClipLocator) GetProjectId() (value string) {
	if s == nil {
		return
	}
	return s.ProjectId
}

func (s *StudioClipLocator) GetChapterId() (value string) {
	if s == nil {
		return
	}
	return s.ChapterId
}

func (s *StudioClipLocator) GetClipType() (value StudioClipLocatorClipType) {
	if s == nil {
		return
	}
	return s.ClipType
}

func (s *StudioClipLocator) GetClipId() (value string) {
	if s == nil {
		return
	}
	return s.ClipId
}

func (s *StudioClipLocator) GetBlockId() (value *string) {
	if s == nil {
		return
	}
	return s.BlockId
}

func (s *StudioClipLocator) GetPreviewUrl() (value *string) {
	if s == nil {
		return
	}
	return s.PreviewUrl
}

func (s *StudioClipLocator) SetProjectId(projectId string) {
	s.ProjectId = projectId
}

func (s *StudioClipLocator) SetChapterId(chapterId string) {
	s.ChapterId = chapterId
}

func (s *StudioClipLocator) SetClipType(clipType StudioClipLocatorClipType) {
	s.ClipType = clipType
}

func (s *StudioClipLocator) SetClipId(clipId string) {
	s.ClipId = clipId
}

func (s *StudioClipLocator) SetBlockId(blockId *string) {
	s.BlockId = blockId
}

func (s *StudioClipLocator) SetPreviewUrl(previewUrl *string) {
	s.PreviewUrl = previewUrl
}

func (s *StudioClipLocator) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StudioClipLocator) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StudioClipLocator to string"
	}
	return string(jsonData)
}

func (s *StudioClipLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StudioClipLocator) MarshalJSON() ([]byte, error) {
	type embed StudioClipLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
