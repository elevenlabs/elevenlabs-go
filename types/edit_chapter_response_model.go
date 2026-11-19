package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EditChapterResponseModel struct {
	Chapter         *ChapterWithContentResponseModel `json:"chapter" xml:"chapter" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EditChapterResponseModel) GetChapter() (value *ChapterWithContentResponseModel) {
	if e == nil {
		return
	}
	return e.Chapter
}

func (e *EditChapterResponseModel) SetChapter(chapter *ChapterWithContentResponseModel) {
	e.Chapter = chapter
}

func (e *EditChapterResponseModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditChapterResponseModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditChapterResponseModel to string"
	}
	return string(jsonData)
}

func (e *EditChapterResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditChapterResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditChapterResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditChapterResponseModel) MarshalJSON() ([]byte, error) {
	type embed EditChapterResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
