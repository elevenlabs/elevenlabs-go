package projects

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost struct {
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id, version_id) encoded as a list of JSON strings for pronunciation dictionaries to be applied to the text. A list of json encoded strings is required as adding projects may occur through formData as opposed to jsonBody. To specify multiple dictionaries use multiple --form lines in your curl, such as --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"' --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []*types.PronunciationDictionaryVersionLocatorRequestModel `json:"pronunciation_dictionary_locators" xml:"pronunciation_dictionary_locators" required:"true"`
	// This will automatically mark text in this project for reconversion when the new dictionary applies or the old one no longer does.
	InvalidateAffectedText *bool `json:"invalidate_affected_text,omitempty" xml:"invalidate_affected_text,omitempty"`
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) GetPronunciationDictionaryLocators() (value []*types.PronunciationDictionaryVersionLocatorRequestModel) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) GetInvalidateAffectedText() (value *bool) {
	if b == nil {
		return
	}
	return b.InvalidateAffectedText
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*types.PronunciationDictionaryVersionLocatorRequestModel) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) SetInvalidateAffectedText(invalidateAffectedText *bool) {
	b.InvalidateAffectedText = invalidateAffectedText
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost(tmp)
	return nil
}

func (b *BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
