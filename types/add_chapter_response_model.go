package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddChapterResponseModel struct {
	Chapter         *ChapterWithContentResponseModel `json:"chapter" xml:"chapter" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AddChapterResponseModel) GetChapter() (value *ChapterWithContentResponseModel) {
	if a == nil {
		return
	}
	return a.Chapter
}

func (a *AddChapterResponseModel) SetChapter(chapter *ChapterWithContentResponseModel) {
	a.Chapter = chapter
}

func (a *AddChapterResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddChapterResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddChapterResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddChapterResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddChapterResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddChapterResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddChapterResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddChapterResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
