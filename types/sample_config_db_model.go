package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SampleConfigDbModel struct {
	IsSample        *bool                          `json:"is_sample,omitempty" xml:"is_sample,omitempty"`
	ParentId        *string                        `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	ParentType      *SampleConfigDbModelParentType `json:"parent_type,omitempty" xml:"parent_type,omitempty"`
	ChapterIds      []string                       `json:"chapter_ids,omitempty" xml:"chapter_ids,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SampleConfigDbModel) GetIsSample() (value *bool) {
	if s == nil {
		return
	}
	return s.IsSample
}

func (s *SampleConfigDbModel) GetParentId() (value *string) {
	if s == nil {
		return
	}
	return s.ParentId
}

func (s *SampleConfigDbModel) GetParentType() (value *SampleConfigDbModelParentType) {
	if s == nil {
		return
	}
	return s.ParentType
}

func (s *SampleConfigDbModel) GetChapterIds() (value []string) {
	if s == nil {
		return
	}
	return s.ChapterIds
}

func (s *SampleConfigDbModel) SetIsSample(isSample *bool) {
	s.IsSample = isSample
}

func (s *SampleConfigDbModel) SetParentId(parentId *string) {
	s.ParentId = parentId
}

func (s *SampleConfigDbModel) SetParentType(parentType *SampleConfigDbModelParentType) {
	s.ParentType = parentType
}

func (s *SampleConfigDbModel) SetChapterIds(chapterIds []string) {
	s.ChapterIds = chapterIds
}

func (s *SampleConfigDbModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SampleConfigDbModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SampleConfigDbModel to string"
	}
	return string(jsonData)
}

func (s *SampleConfigDbModel) UnmarshalJSON(data []byte) error {
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

func (s *SampleConfigDbModel) MarshalJSON() ([]byte, error) {
	type embed SampleConfigDbModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
