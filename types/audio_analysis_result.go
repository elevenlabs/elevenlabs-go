package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioAnalysisResult struct {
	Title           string            `json:"title" xml:"title" required:"true"`
	Description     string            `json:"description" xml:"description" required:"true"`
	ContentType     *string           `json:"content_type,omitempty" xml:"content_type,omitempty"`
	OverallPacing   *string           `json:"overall_pacing,omitempty" xml:"overall_pacing,omitempty"`
	Segments        []*AudioSegment   `json:"segments,omitempty" xml:"segments,omitempty"`
	KeyMoments      []*AudioKeyMoment `json:"key_moments,omitempty" xml:"key_moments,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioAnalysisResult) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AudioAnalysisResult) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioAnalysisResult) GetContentType() (value *string) {
	if a == nil {
		return
	}
	return a.ContentType
}

func (a *AudioAnalysisResult) GetOverallPacing() (value *string) {
	if a == nil {
		return
	}
	return a.OverallPacing
}

func (a *AudioAnalysisResult) GetSegments() (value []*AudioSegment) {
	if a == nil {
		return
	}
	return a.Segments
}

func (a *AudioAnalysisResult) GetKeyMoments() (value []*AudioKeyMoment) {
	if a == nil {
		return
	}
	return a.KeyMoments
}

func (a *AudioAnalysisResult) SetTitle(title string) {
	a.Title = title
}

func (a *AudioAnalysisResult) SetDescription(description string) {
	a.Description = description
}

func (a *AudioAnalysisResult) SetContentType(contentType *string) {
	a.ContentType = contentType
}

func (a *AudioAnalysisResult) SetOverallPacing(overallPacing *string) {
	a.OverallPacing = overallPacing
}

func (a *AudioAnalysisResult) SetSegments(segments []*AudioSegment) {
	a.Segments = segments
}

func (a *AudioAnalysisResult) SetKeyMoments(keyMoments []*AudioKeyMoment) {
	a.KeyMoments = keyMoments
}

func (a *AudioAnalysisResult) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioAnalysisResult) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioAnalysisResult to string"
	}
	return string(jsonData)
}

func (a *AudioAnalysisResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioAnalysisResult) MarshalJSON() ([]byte, error) {
	type embed AudioAnalysisResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
