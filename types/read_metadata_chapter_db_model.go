package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReadMetadataChapterDbModel struct {
	ChapterName        string   `json:"chapter_name" xml:"chapter_name" required:"true"`
	WordCount          int      `json:"word_count" xml:"word_count" required:"true"`
	CharCount          int      `json:"char_count" xml:"char_count" required:"true"`
	StartingCharOffset int      `json:"starting_char_offset" xml:"starting_char_offset" required:"true"`
	HasParsedHtml      *bool    `json:"has_parsed_html,omitempty" xml:"has_parsed_html,omitempty"`
	HasSummary         *bool    `json:"has_summary,omitempty" xml:"has_summary,omitempty"`
	DurationSeconds    *float64 `json:"duration_seconds,omitempty" xml:"duration_seconds,omitempty"`
	FileNumber         *string  `json:"file_number,omitempty" xml:"file_number,omitempty"`
	IsFallbackName     *bool    `json:"is_fallback_name,omitempty" xml:"is_fallback_name,omitempty"`
	ChapterId          *string  `json:"chapter_id,omitempty" xml:"chapter_id,omitempty"`
	extraProperties    map[string]interface{}
}

func (r *ReadMetadataChapterDbModel) GetChapterName() (value string) {
	if r == nil {
		return
	}
	return r.ChapterName
}

func (r *ReadMetadataChapterDbModel) GetWordCount() (value int) {
	if r == nil {
		return
	}
	return r.WordCount
}

func (r *ReadMetadataChapterDbModel) GetCharCount() (value int) {
	if r == nil {
		return
	}
	return r.CharCount
}

func (r *ReadMetadataChapterDbModel) GetStartingCharOffset() (value int) {
	if r == nil {
		return
	}
	return r.StartingCharOffset
}

func (r *ReadMetadataChapterDbModel) GetHasParsedHtml() (value *bool) {
	if r == nil {
		return
	}
	return r.HasParsedHtml
}

func (r *ReadMetadataChapterDbModel) GetHasSummary() (value *bool) {
	if r == nil {
		return
	}
	return r.HasSummary
}

func (r *ReadMetadataChapterDbModel) GetDurationSeconds() (value *float64) {
	if r == nil {
		return
	}
	return r.DurationSeconds
}

func (r *ReadMetadataChapterDbModel) GetFileNumber() (value *string) {
	if r == nil {
		return
	}
	return r.FileNumber
}

func (r *ReadMetadataChapterDbModel) GetIsFallbackName() (value *bool) {
	if r == nil {
		return
	}
	return r.IsFallbackName
}

func (r *ReadMetadataChapterDbModel) GetChapterId() (value *string) {
	if r == nil {
		return
	}
	return r.ChapterId
}

func (r *ReadMetadataChapterDbModel) SetChapterName(chapterName string) {
	r.ChapterName = chapterName
}

func (r *ReadMetadataChapterDbModel) SetWordCount(wordCount int) {
	r.WordCount = wordCount
}

func (r *ReadMetadataChapterDbModel) SetCharCount(charCount int) {
	r.CharCount = charCount
}

func (r *ReadMetadataChapterDbModel) SetStartingCharOffset(startingCharOffset int) {
	r.StartingCharOffset = startingCharOffset
}

func (r *ReadMetadataChapterDbModel) SetHasParsedHtml(hasParsedHtml *bool) {
	r.HasParsedHtml = hasParsedHtml
}

func (r *ReadMetadataChapterDbModel) SetHasSummary(hasSummary *bool) {
	r.HasSummary = hasSummary
}

func (r *ReadMetadataChapterDbModel) SetDurationSeconds(durationSeconds *float64) {
	r.DurationSeconds = durationSeconds
}

func (r *ReadMetadataChapterDbModel) SetFileNumber(fileNumber *string) {
	r.FileNumber = fileNumber
}

func (r *ReadMetadataChapterDbModel) SetIsFallbackName(isFallbackName *bool) {
	r.IsFallbackName = isFallbackName
}

func (r *ReadMetadataChapterDbModel) SetChapterId(chapterId *string) {
	r.ChapterId = chapterId
}

func (r *ReadMetadataChapterDbModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReadMetadataChapterDbModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReadMetadataChapterDbModel to string"
	}
	return string(jsonData)
}

func (r *ReadMetadataChapterDbModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReadMetadataChapterDbModel) MarshalJSON() ([]byte, error) {
	type embed ReadMetadataChapterDbModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
