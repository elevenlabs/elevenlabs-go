package studio

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyCreateStudioProjectV1StudioProjectsPost struct {
	// The name of the Studio project, used for identification only.
	Name string `json:"name" xml:"name" required:"true"`
	// The voice_id that corresponds to the default voice used for new titles.
	DefaultTitleVoiceId *string `json:"default_title_voice_id,omitempty" xml:"default_title_voice_id,omitempty"`
	// The voice_id that corresponds to the default voice used for new paragraphs.
	DefaultParagraphVoiceId *string `json:"default_paragraph_voice_id,omitempty" xml:"default_paragraph_voice_id,omitempty"`
	// The ID of the model to be used for this Studio project, you can query GET /v1/models to list all available models.
	DefaultModelId *string `json:"default_model_id,omitempty" xml:"default_model_id,omitempty"`
	// An optional URL from which we will extract content to initialize the Studio project. If this is set, 'from_url' and 'from_content' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	FromUrl *string `json:"from_url,omitempty" xml:"from_url,omitempty"`
	// An optional .epub, .pdf, .txt or similar file can be provided. If provided, we will initialize the Studio project with its content. If this is set, 'from_url' and 'from_content' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	FromDocument *string `json:"from_document,omitempty" xml:"from_document,omitempty"`
	//
	// An optional content to initialize the Studio project with. If this is set, 'from_url' and 'from_document' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	//
	// Example:
	// [{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]
	//
	FromContentJson *string `json:"from_content_json,omitempty" xml:"from_content_json,omitempty"`
	// Output quality of the generated audio. Must be one of:
	// 'standard' - standard output format, 128kbps with 44.1kHz sample rate.
	// 'high' - high quality output format, 192kbps with 44.1kHz sample rate and major improvements on our side.
	// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and highest improvements on our side.
	// 'ultra_lossless' - ultra quality output format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side in a fully lossless format.
	// If not provided, defaults to the highest quality preset available on your subscription tier.
	//
	QualityPreset *types.QualityPresetType `json:"quality_preset,omitempty" xml:"quality_preset,omitempty"`
	// An optional name of the author of the Studio project, this will be added as metadata to the mp3 file on Studio project or chapter download.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// An optional name of the author of the Studio project, this will be added as metadata to the mp3 file on Studio project or chapter download.
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// An optional description of the Studio project.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// An optional list of genres associated with the Studio project.
	Genres []string `json:"genres,omitempty" xml:"genres,omitempty"`
	// An optional target audience of the Studio project.
	TargetAudience *CreateProjectsRequestTargetAudience `json:"target_audience,omitempty" xml:"target_audience,omitempty"`
	// An optional language of the Studio project. Two-letter language code (ISO 639-1).
	Language *string `json:"language,omitempty" xml:"language,omitempty" maxLength:"2" minLength:"2"`
	// An optional content type of the Studio project.
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// An optional original publication date of the Studio project, in the format YYYY-MM-DD or YYYY.
	OriginalPublicationDate *string `json:"original_publication_date,omitempty" xml:"original_publication_date,omitempty" pattern:"^\\d{4}-\\d{2}-\\d{2}$|^\\d{4}$"`
	// An optional specification of whether this Studio project contains mature content.
	MatureContent *bool `json:"mature_content,omitempty" xml:"mature_content,omitempty"`
	// An optional ISBN number of the Studio project you want to create, this will be added as metadata to the mp3 file on Studio project or chapter download.
	IsbnNumber *string `json:"isbn_number,omitempty" xml:"isbn_number,omitempty"`
	// [Deprecated] When the Studio project is downloaded, should the returned audio have postprocessing in order to make it compliant with audiobook normalized volume requirements
	AcxVolumeNormalization *bool `json:"acx_volume_normalization,omitempty" xml:"acx_volume_normalization,omitempty"`
	// When the Studio project is downloaded, should the returned audio have postprocessing in order to make it compliant with audiobook normalized volume requirements
	VolumeNormalization *bool `json:"volume_normalization,omitempty" xml:"volume_normalization,omitempty"`
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id, version_id) encoded as a list of JSON strings for pronunciation dictionaries to be applied to the text. A list of json encoded strings is required as adding projects may occur through formData as opposed to jsonBody. To specify multiple dictionaries use multiple --form lines in your curl, such as --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"' --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []string `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	//
	// A url that will be called by our service when the Studio project is converted. Request will contain a json blob containing the status of the conversion
	// Messages:
	// 1. When project was converted successfully:
	// {
	// type: "project_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// conversion_status: "success",
	// project_snapshot_id: "22m00Tcm4TlvDq8ikMAT",
	// error_details: None,
	// }
	// }
	// 2. When project conversion failed:
	// {
	// type: "project_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// conversion_status: "error",
	// project_snapshot_id: None,
	// error_details: "Error details if conversion failed"
	// }
	// }
	//
	// 3. When chapter was converted successfully:
	// {
	// type: "chapter_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// chapter_id: "22m00Tcm4TlvDq8ikMAT",
	// conversion_status: "success",
	// chapter_snapshot_id: "23m00Tcm4TlvDq8ikMAV",
	// error_details: None,
	// }
	// }
	// 4. When chapter conversion failed:
	// {
	// type: "chapter_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// chapter_id: "22m00Tcm4TlvDq8ikMAT",
	// conversion_status: "error",
	// chapter_snapshot_id: None,
	// error_details: "Error details if conversion failed"
	// }
	// }
	//
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	// An optional specification of whether the content of this Studio project is fiction.
	Fiction *CreateProjectsRequestFiction `json:"fiction,omitempty" xml:"fiction,omitempty"`
	//
	// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
	// When set to 'auto', the system will automatically decide whether to apply text normalization
	// (e.g., spelling out numbers). With 'on', text normalization will always be applied, while
	// with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
	//
	ApplyTextNormalization *CreateProjectsRequestApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// Whether to auto convert the Studio project to audio or not.
	AutoConvert *bool `json:"auto_convert,omitempty" xml:"auto_convert,omitempty"`
	// [Alpha Feature] Whether automatically assign voices to phrases in the create Project.
	AutoAssignVoices *bool `json:"auto_assign_voices,omitempty" xml:"auto_assign_voices,omitempty"`
	// The type of Studio project to create.
	SourceType *CreateProjectsRequestSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// Optional voice settings overrides for the project, encoded as a list of JSON strings.
	//
	// Example:
	// ["{\"voice_id\": \"21m00Tcm4TlvDq8ikWAM\", \"stability\": 0.7, \"similarity_boost\": 0.8, \"style\": 0.5, \"speed\": 1.0, \"use_speaker_boost\": true}"]
	//
	VoiceSettings []string `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	// If true, creates a corresponding read for direct publishing in draft state
	CreatePublishingRead *bool `json:"create_publishing_read,omitempty" xml:"create_publishing_read,omitempty"`
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetDefaultTitleVoiceId() (value *string) {
	if b == nil {
		return
	}
	return b.DefaultTitleVoiceId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetDefaultParagraphVoiceId() (value *string) {
	if b == nil {
		return
	}
	return b.DefaultParagraphVoiceId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetDefaultModelId() (value *string) {
	if b == nil {
		return
	}
	return b.DefaultModelId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetFromUrl() (value *string) {
	if b == nil {
		return
	}
	return b.FromUrl
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetFromDocument() (value *string) {
	if b == nil {
		return
	}
	return b.FromDocument
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetFromContentJson() (value *string) {
	if b == nil {
		return
	}
	return b.FromContentJson
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetQualityPreset() (value *types.QualityPresetType) {
	if b == nil {
		return
	}
	return b.QualityPreset
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetAuthor() (value *string) {
	if b == nil {
		return
	}
	return b.Author
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetGenres() (value []string) {
	if b == nil {
		return
	}
	return b.Genres
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetTargetAudience() (value *CreateProjectsRequestTargetAudience) {
	if b == nil {
		return
	}
	return b.TargetAudience
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.Language
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetContentType() (value *string) {
	if b == nil {
		return
	}
	return b.ContentType
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetOriginalPublicationDate() (value *string) {
	if b == nil {
		return
	}
	return b.OriginalPublicationDate
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetMatureContent() (value *bool) {
	if b == nil {
		return
	}
	return b.MatureContent
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetIsbnNumber() (value *string) {
	if b == nil {
		return
	}
	return b.IsbnNumber
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetAcxVolumeNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.AcxVolumeNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetVolumeNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.VolumeNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetPronunciationDictionaryLocators() (value []string) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetCallbackUrl() (value *string) {
	if b == nil {
		return
	}
	return b.CallbackUrl
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetFiction() (value *CreateProjectsRequestFiction) {
	if b == nil {
		return
	}
	return b.Fiction
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetApplyTextNormalization() (value *CreateProjectsRequestApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetAutoConvert() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoConvert
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetAutoAssignVoices() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoAssignVoices
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetSourceType() (value *CreateProjectsRequestSourceType) {
	if b == nil {
		return
	}
	return b.SourceType
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetVoiceSettings() (value []string) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetCreatePublishingRead() (value *bool) {
	if b == nil {
		return
	}
	return b.CreatePublishingRead
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetDefaultTitleVoiceId(defaultTitleVoiceId *string) {
	b.DefaultTitleVoiceId = defaultTitleVoiceId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetDefaultParagraphVoiceId(defaultParagraphVoiceId *string) {
	b.DefaultParagraphVoiceId = defaultParagraphVoiceId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetDefaultModelId(defaultModelId *string) {
	b.DefaultModelId = defaultModelId
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetFromUrl(fromUrl *string) {
	b.FromUrl = fromUrl
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetFromDocument(fromDocument *string) {
	b.FromDocument = fromDocument
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetFromContentJson(fromContentJson *string) {
	b.FromContentJson = fromContentJson
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetQualityPreset(qualityPreset *types.QualityPresetType) {
	b.QualityPreset = qualityPreset
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetAuthor(author *string) {
	b.Author = author
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetGenres(genres []string) {
	b.Genres = genres
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetTargetAudience(targetAudience *CreateProjectsRequestTargetAudience) {
	b.TargetAudience = targetAudience
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetLanguage(language *string) {
	b.Language = language
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetContentType(contentType *string) {
	b.ContentType = contentType
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetOriginalPublicationDate(originalPublicationDate *string) {
	b.OriginalPublicationDate = originalPublicationDate
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetMatureContent(matureContent *bool) {
	b.MatureContent = matureContent
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetIsbnNumber(isbnNumber *string) {
	b.IsbnNumber = isbnNumber
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetAcxVolumeNormalization(acxVolumeNormalization *bool) {
	b.AcxVolumeNormalization = acxVolumeNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetVolumeNormalization(volumeNormalization *bool) {
	b.VolumeNormalization = volumeNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []string) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetCallbackUrl(callbackUrl *string) {
	b.CallbackUrl = callbackUrl
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetFiction(fiction *CreateProjectsRequestFiction) {
	b.Fiction = fiction
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetApplyTextNormalization(applyTextNormalization *CreateProjectsRequestApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetAutoConvert(autoConvert *bool) {
	b.AutoConvert = autoConvert
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetAutoAssignVoices(autoAssignVoices *bool) {
	b.AutoAssignVoices = autoAssignVoices
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetSourceType(sourceType *CreateProjectsRequestSourceType) {
	b.SourceType = sourceType
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetVoiceSettings(voiceSettings []string) {
	b.VoiceSettings = voiceSettings
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) SetCreatePublishingRead(createPublishingRead *bool) {
	b.CreatePublishingRead = createPublishingRead
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateStudioProjectV1StudioProjectsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateStudioProjectV1StudioProjectsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateStudioProjectV1StudioProjectsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateStudioProjectV1StudioProjectsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type CreateProjectsRequestTargetAudience string

const (
	CreateProjectsRequestTargetAudienceChildren   CreateProjectsRequestTargetAudience = "children"
	CreateProjectsRequestTargetAudienceYoungAdult CreateProjectsRequestTargetAudience = "young adult"
	CreateProjectsRequestTargetAudienceAdult      CreateProjectsRequestTargetAudience = "adult"
	CreateProjectsRequestTargetAudienceAllAges    CreateProjectsRequestTargetAudience = "all ages"
)

func NewCreateProjectsRequestTargetAudienceFromString(s string) (CreateProjectsRequestTargetAudience, error) {
	switch s {
	case "children":
		return CreateProjectsRequestTargetAudienceChildren, nil
	case "young adult":
		return CreateProjectsRequestTargetAudienceYoungAdult, nil
	case "adult":
		return CreateProjectsRequestTargetAudienceAdult, nil
	case "all ages":
		return CreateProjectsRequestTargetAudienceAllAges, nil
	}
	var t CreateProjectsRequestTargetAudience
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateProjectsRequestTargetAudience) Ptr() *CreateProjectsRequestTargetAudience {
	return &e
}

type CreateProjectsRequestFiction string

const (
	CreateProjectsRequestFictionFiction    CreateProjectsRequestFiction = "fiction"
	CreateProjectsRequestFictionNonFiction CreateProjectsRequestFiction = "non-fiction"
)

func NewCreateProjectsRequestFictionFromString(s string) (CreateProjectsRequestFiction, error) {
	switch s {
	case "fiction":
		return CreateProjectsRequestFictionFiction, nil
	case "non-fiction":
		return CreateProjectsRequestFictionNonFiction, nil
	}
	var t CreateProjectsRequestFiction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateProjectsRequestFiction) Ptr() *CreateProjectsRequestFiction {
	return &e
}

type CreateProjectsRequestApplyTextNormalization string

const (
	CreateProjectsRequestApplyTextNormalizationAuto         CreateProjectsRequestApplyTextNormalization = "auto"
	CreateProjectsRequestApplyTextNormalizationOn           CreateProjectsRequestApplyTextNormalization = "on"
	CreateProjectsRequestApplyTextNormalizationOff          CreateProjectsRequestApplyTextNormalization = "off"
	CreateProjectsRequestApplyTextNormalizationApplyEnglish CreateProjectsRequestApplyTextNormalization = "apply_english"
)

func NewCreateProjectsRequestApplyTextNormalizationFromString(s string) (CreateProjectsRequestApplyTextNormalization, error) {
	switch s {
	case "auto":
		return CreateProjectsRequestApplyTextNormalizationAuto, nil
	case "on":
		return CreateProjectsRequestApplyTextNormalizationOn, nil
	case "off":
		return CreateProjectsRequestApplyTextNormalizationOff, nil
	case "apply_english":
		return CreateProjectsRequestApplyTextNormalizationApplyEnglish, nil
	}
	var t CreateProjectsRequestApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateProjectsRequestApplyTextNormalization) Ptr() *CreateProjectsRequestApplyTextNormalization {
	return &e
}

type CreateProjectsRequestSourceType string

const (
	CreateProjectsRequestSourceTypeBlank      CreateProjectsRequestSourceType = "blank"
	CreateProjectsRequestSourceTypeBook       CreateProjectsRequestSourceType = "book"
	CreateProjectsRequestSourceTypeArticle    CreateProjectsRequestSourceType = "article"
	CreateProjectsRequestSourceTypeGenfm      CreateProjectsRequestSourceType = "genfm"
	CreateProjectsRequestSourceTypeVideo      CreateProjectsRequestSourceType = "video"
	CreateProjectsRequestSourceTypeScreenplay CreateProjectsRequestSourceType = "screenplay"
)

func NewCreateProjectsRequestSourceTypeFromString(s string) (CreateProjectsRequestSourceType, error) {
	switch s {
	case "blank":
		return CreateProjectsRequestSourceTypeBlank, nil
	case "book":
		return CreateProjectsRequestSourceTypeBook, nil
	case "article":
		return CreateProjectsRequestSourceTypeArticle, nil
	case "genfm":
		return CreateProjectsRequestSourceTypeGenfm, nil
	case "video":
		return CreateProjectsRequestSourceTypeVideo, nil
	case "screenplay":
		return CreateProjectsRequestSourceTypeScreenplay, nil
	}
	var t CreateProjectsRequestSourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateProjectsRequestSourceType) Ptr() *CreateProjectsRequestSourceType {
	return &e
}

type BodyUpdateStudioProjectV1StudioProjectsProjectIdPost struct {
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// The name of the Studio project, used for identification only.
	Name string `json:"name" xml:"name" required:"true"`
	// The voice_id that corresponds to the default voice used for new titles.
	DefaultTitleVoiceId string `json:"default_title_voice_id" xml:"default_title_voice_id" required:"true"`
	// The voice_id that corresponds to the default voice used for new paragraphs.
	DefaultParagraphVoiceId string `json:"default_paragraph_voice_id" xml:"default_paragraph_voice_id" required:"true"`
	// An optional name of the author of the Studio project, this will be added as metadata to the mp3 file on Studio project or chapter download.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// An optional name of the author of the Studio project, this will be added as metadata to the mp3 file on Studio project or chapter download.
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// An optional ISBN number of the Studio project you want to create, this will be added as metadata to the mp3 file on Studio project or chapter download.
	IsbnNumber *string `json:"isbn_number,omitempty" xml:"isbn_number,omitempty"`
	// When the Studio project is downloaded, should the returned audio have postprocessing in order to make it compliant with audiobook normalized volume requirements
	VolumeNormalization *bool `json:"volume_normalization,omitempty" xml:"volume_normalization,omitempty"`
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetDefaultTitleVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.DefaultTitleVoiceId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetDefaultParagraphVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.DefaultParagraphVoiceId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetAuthor() (value *string) {
	if b == nil {
		return
	}
	return b.Author
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetIsbnNumber() (value *string) {
	if b == nil {
		return
	}
	return b.IsbnNumber
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetVolumeNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.VolumeNormalization
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetName(name string) {
	b.Name = name
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetDefaultTitleVoiceId(defaultTitleVoiceId string) {
	b.DefaultTitleVoiceId = defaultTitleVoiceId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetDefaultParagraphVoiceId(defaultParagraphVoiceId string) {
	b.DefaultParagraphVoiceId = defaultParagraphVoiceId
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetAuthor(author *string) {
	b.Author = author
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetIsbnNumber(isbnNumber *string) {
	b.IsbnNumber = isbnNumber
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) SetVolumeNormalization(volumeNormalization *bool) {
	b.VolumeNormalization = volumeNormalization
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateStudioProjectV1StudioProjectsProjectIdPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateStudioProjectV1StudioProjectsProjectIdPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateStudioProjectV1StudioProjectsProjectIdPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	getProjectsRequestFieldShareId = big.NewInt(1 << 0)
)

type GetProjectsRequest struct {
	ShareId *string `json:"-" url:"share_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetProjectsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetShareId sets the ShareId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetProjectsRequest) SetShareId(shareId *string) {
	g.ShareId = shareId
	g.require(getProjectsRequestFieldShareId)
}
