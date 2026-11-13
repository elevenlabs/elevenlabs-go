package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ImageAnalysisResult struct {
	Title            string  `json:"title" xml:"title" required:"true"`
	Description      string  `json:"description" xml:"description" required:"true"`
	ContentType      *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	MoodAndStyle     *string `json:"mood_and_style,omitempty" xml:"mood_and_style,omitempty"`
	CompositionNotes *string `json:"composition_notes,omitempty" xml:"composition_notes,omitempty"`
	// Readable text overlaid or shown in the image, if any.
	VisibleText     *string         `json:"visible_text,omitempty" xml:"visible_text,omitempty"`
	Subjects        []*ImageSubject `json:"subjects,omitempty" xml:"subjects,omitempty"`
	extraProperties map[string]interface{}
}

func (i *ImageAnalysisResult) GetTitle() (value string) {
	if i == nil {
		return
	}
	return i.Title
}

func (i *ImageAnalysisResult) GetDescription() (value string) {
	if i == nil {
		return
	}
	return i.Description
}

func (i *ImageAnalysisResult) GetContentType() (value *string) {
	if i == nil {
		return
	}
	return i.ContentType
}

func (i *ImageAnalysisResult) GetMoodAndStyle() (value *string) {
	if i == nil {
		return
	}
	return i.MoodAndStyle
}

func (i *ImageAnalysisResult) GetCompositionNotes() (value *string) {
	if i == nil {
		return
	}
	return i.CompositionNotes
}

func (i *ImageAnalysisResult) GetVisibleText() (value *string) {
	if i == nil {
		return
	}
	return i.VisibleText
}

func (i *ImageAnalysisResult) GetSubjects() (value []*ImageSubject) {
	if i == nil {
		return
	}
	return i.Subjects
}

func (i *ImageAnalysisResult) SetTitle(title string) {
	i.Title = title
}

func (i *ImageAnalysisResult) SetDescription(description string) {
	i.Description = description
}

func (i *ImageAnalysisResult) SetContentType(contentType *string) {
	i.ContentType = contentType
}

func (i *ImageAnalysisResult) SetMoodAndStyle(moodAndStyle *string) {
	i.MoodAndStyle = moodAndStyle
}

func (i *ImageAnalysisResult) SetCompositionNotes(compositionNotes *string) {
	i.CompositionNotes = compositionNotes
}

func (i *ImageAnalysisResult) SetVisibleText(visibleText *string) {
	i.VisibleText = visibleText
}

func (i *ImageAnalysisResult) SetSubjects(subjects []*ImageSubject) {
	i.Subjects = subjects
}

func (i *ImageAnalysisResult) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i ImageAnalysisResult) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: ImageAnalysisResult to string"
	}
	return string(jsonData)
}

func (i *ImageAnalysisResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *ImageAnalysisResult) MarshalJSON() ([]byte, error) {
	type embed ImageAnalysisResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
