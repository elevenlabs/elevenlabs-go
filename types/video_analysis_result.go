package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoAnalysisResult struct {
	Title           string            `json:"title" xml:"title" required:"true"`
	Description     string            `json:"description" xml:"description" required:"true"`
	ContentType     *string           `json:"content_type,omitempty" xml:"content_type,omitempty"`
	OverallPacing   *string           `json:"overall_pacing,omitempty" xml:"overall_pacing,omitempty"`
	Subjects        []*VideoSubject   `json:"subjects,omitempty" xml:"subjects,omitempty"`
	Segments        []*VideoSegment   `json:"segments,omitempty" xml:"segments,omitempty"`
	KeyMoments      []*VideoKeyMoment `json:"key_moments,omitempty" xml:"key_moments,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VideoAnalysisResult) GetTitle() (value string) {
	if v == nil {
		return
	}
	return v.Title
}

func (v *VideoAnalysisResult) GetDescription() (value string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VideoAnalysisResult) GetContentType() (value *string) {
	if v == nil {
		return
	}
	return v.ContentType
}

func (v *VideoAnalysisResult) GetOverallPacing() (value *string) {
	if v == nil {
		return
	}
	return v.OverallPacing
}

func (v *VideoAnalysisResult) GetSubjects() (value []*VideoSubject) {
	if v == nil {
		return
	}
	return v.Subjects
}

func (v *VideoAnalysisResult) GetSegments() (value []*VideoSegment) {
	if v == nil {
		return
	}
	return v.Segments
}

func (v *VideoAnalysisResult) GetKeyMoments() (value []*VideoKeyMoment) {
	if v == nil {
		return
	}
	return v.KeyMoments
}

func (v *VideoAnalysisResult) SetTitle(title string) {
	v.Title = title
}

func (v *VideoAnalysisResult) SetDescription(description string) {
	v.Description = description
}

func (v *VideoAnalysisResult) SetContentType(contentType *string) {
	v.ContentType = contentType
}

func (v *VideoAnalysisResult) SetOverallPacing(overallPacing *string) {
	v.OverallPacing = overallPacing
}

func (v *VideoAnalysisResult) SetSubjects(subjects []*VideoSubject) {
	v.Subjects = subjects
}

func (v *VideoAnalysisResult) SetSegments(segments []*VideoSegment) {
	v.Segments = segments
}

func (v *VideoAnalysisResult) SetKeyMoments(keyMoments []*VideoKeyMoment) {
	v.KeyMoments = keyMoments
}

func (v *VideoAnalysisResult) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoAnalysisResult) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoAnalysisResult to string"
	}
	return string(jsonData)
}

func (v *VideoAnalysisResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoAnalysisResult) MarshalJSON() ([]byte, error) {
	type embed VideoAnalysisResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
