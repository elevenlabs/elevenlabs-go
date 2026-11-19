package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectExtendedResponse struct {
	// The ID of the project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// The name of the project.
	Name string `json:"name" xml:"name" required:"true"`
	// The creation date of the project.
	CreateDateUnix int `json:"create_date_unix" xml:"create_date_unix" required:"true"`
	// The user ID who created the project.
	CreatedByUserId *string `json:"created_by_user_id" xml:"created_by_user_id" required:"true"`
	// The default title project voice reference ID.
	DefaultTitleVoiceRefId string `json:"default_title_voice_ref_id" xml:"default_title_voice_ref_id" required:"true"`
	// The default paragraph project voice reference ID.
	DefaultParagraphVoiceRefId string `json:"default_paragraph_voice_ref_id" xml:"default_paragraph_voice_ref_id" required:"true"`
	// The default model ID.
	DefaultModelId string `json:"default_model_id" xml:"default_model_id" required:"true"`
	// The last conversion date of the project.
	LastConversionDateUnix *int `json:"last_conversion_date_unix,omitempty" xml:"last_conversion_date_unix,omitempty"`
	// Whether the project can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded" xml:"can_be_downloaded" required:"true"`
	// The title of the project.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The author of the project.
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// The description of the project.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// List of genres of the project.
	Genres []string `json:"genres,omitempty" xml:"genres,omitempty"`
	// The cover image URL of the project.
	CoverImageUrl *string `json:"cover_image_url,omitempty" xml:"cover_image_url,omitempty"`
	// The target audience of the project.
	TargetAudience *ProjectExtendedResponseTargetAudience `json:"target_audience,omitempty" xml:"target_audience,omitempty"`
	// Two-letter language code (ISO 639-1) of the language of the project.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The content type of the project, e.g. 'Novel' or 'Short Story'
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The original publication date of the project.
	OriginalPublicationDate *string `json:"original_publication_date,omitempty" xml:"original_publication_date,omitempty"`
	// Whether the project contains mature content.
	MatureContent *bool `json:"mature_content,omitempty" xml:"mature_content,omitempty"`
	// The ISBN number of the project.
	IsbnNumber *string `json:"isbn_number,omitempty" xml:"isbn_number,omitempty"`
	// Whether the project uses volume normalization.
	VolumeNormalization bool `json:"volume_normalization" xml:"volume_normalization" required:"true"`
	// The state of the project.
	State ProjectState `json:"state" xml:"state" required:"true"`
	// The access level of the project.
	AccessLevel ProjectResponseAccessLevel `json:"access_level" xml:"access_level" required:"true"`
	// Whether the project is fiction.
	Fiction *ProjectExtendedResponseFiction `json:"fiction,omitempty" xml:"fiction,omitempty"`
	// Whether quality check is enabled for this project.
	QualityCheckOn bool `json:"quality_check_on" xml:"quality_check_on" required:"true"`
	// Whether quality check is enabled on the project when bulk converting.
	QualityCheckOnWhenBulkConvert bool `json:"quality_check_on_when_bulk_convert" xml:"quality_check_on_when_bulk_convert" required:"true"`
	// The creation meta of the project.
	CreationMeta *ProjectCreationMetaResponseModel `json:"creation_meta,omitempty" xml:"creation_meta,omitempty"`
	// The source type of the project.
	SourceType *ProjectExtendedResponseSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// Whether chapters are enabled for the project.
	ChaptersEnabled *bool `json:"chapters_enabled,omitempty" xml:"chapters_enabled,omitempty"`
	// Whether captions are enabled for the project.
	CaptionsEnabled *bool `json:"captions_enabled,omitempty" xml:"captions_enabled,omitempty"`
	// Global styling to be applied to all captions
	CaptionStyle *CaptionStyleModel `json:"caption_style,omitempty" xml:"caption_style,omitempty"`
	// Styling changes that have been made to the provided templates
	CaptionStyleTemplateOverrides map[string]*CaptionStyleModel `json:"caption_style_template_overrides,omitempty"`
	// The public share ID of the project.
	PublicShareId *string `json:"public_share_id,omitempty" xml:"public_share_id,omitempty"`
	// The aspect ratio of the project.
	AspectRatio *ProjectExtendedResponseAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// Agent-related settings for the project
	AgentSettings *StudioAgentSettingsModel `json:"agent_settings,omitempty" xml:"agent_settings,omitempty"`
	QualityPreset QualityPresetType         `json:"quality_preset" xml:"quality_preset" required:"true"`
	// List of chapters of the project and their metadata.
	Chapters []*ChapterResponse `json:"chapters" xml:"chapters" required:"true"`
	// List of pronunciation dictionary versions of the project and their metadata.
	PronunciationDictionaryVersions []*PronunciationDictionaryVersionResponseModel `json:"pronunciation_dictionary_versions" xml:"pronunciation_dictionary_versions" required:"true"`
	// List of pronunciation dictionary locators.
	PronunciationDictionaryLocators []*PronunciationDictionaryLocatorResponseModel `json:"pronunciation_dictionary_locators" xml:"pronunciation_dictionary_locators" required:"true"`
	// Whether text normalization is applied to the project.
	ApplyTextNormalization ProjectExtendedResponseApplyTextNormalization `json:"apply_text_normalization" xml:"apply_text_normalization" required:"true"`
	// Experimental features for the project.
	Experimental map[string]any `json:"experimental,omitempty"`
	// List of uploaded assets e.g. videos, audios.
	Assets []*ProjectExtendedResponseAssetsItem `json:"assets" xml:"assets" required:"true"`
	// List of configured project voices.
	Voices []*ProjectVoiceResponseModel `json:"voices" xml:"voices" required:"true"`
	// List of voices used by the project.
	BaseVoices []*Voice `json:"base_voices,omitempty" xml:"base_voices,omitempty"`
	// The ElevenReader data if the book was published.
	PublishingRead *DirectPublishingReadResponseModel `json:"publishing_read,omitempty" xml:"publishing_read,omitempty"`
	// The default title voice ID.
	DefaultTitleVoiceId *string `json:"default_title_voice_id,omitempty" xml:"default_title_voice_id,omitempty"`
	// The default paragraph voice ID.
	DefaultParagraphVoiceId *string `json:"default_paragraph_voice_id,omitempty" xml:"default_paragraph_voice_id,omitempty"`
	extraProperties         map[string]interface{}
}

func (p *ProjectExtendedResponse) GetProjectId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectId
}

func (p *ProjectExtendedResponse) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProjectExtendedResponse) GetCreateDateUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreateDateUnix
}

func (p *ProjectExtendedResponse) GetCreatedByUserId() (value *string) {
	if p == nil {
		return
	}
	return p.CreatedByUserId
}

func (p *ProjectExtendedResponse) GetDefaultTitleVoiceRefId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultTitleVoiceRefId
}

func (p *ProjectExtendedResponse) GetDefaultParagraphVoiceRefId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultParagraphVoiceRefId
}

func (p *ProjectExtendedResponse) GetDefaultModelId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultModelId
}

func (p *ProjectExtendedResponse) GetLastConversionDateUnix() (value *int) {
	if p == nil {
		return
	}
	return p.LastConversionDateUnix
}

func (p *ProjectExtendedResponse) GetCanBeDownloaded() (value bool) {
	if p == nil {
		return
	}
	return p.CanBeDownloaded
}

func (p *ProjectExtendedResponse) GetTitle() (value *string) {
	if p == nil {
		return
	}
	return p.Title
}

func (p *ProjectExtendedResponse) GetAuthor() (value *string) {
	if p == nil {
		return
	}
	return p.Author
}

func (p *ProjectExtendedResponse) GetDescription() (value *string) {
	if p == nil {
		return
	}
	return p.Description
}

func (p *ProjectExtendedResponse) GetGenres() (value []string) {
	if p == nil {
		return
	}
	return p.Genres
}

func (p *ProjectExtendedResponse) GetCoverImageUrl() (value *string) {
	if p == nil {
		return
	}
	return p.CoverImageUrl
}

func (p *ProjectExtendedResponse) GetTargetAudience() (value *ProjectExtendedResponseTargetAudience) {
	if p == nil {
		return
	}
	return p.TargetAudience
}

func (p *ProjectExtendedResponse) GetLanguage() (value *string) {
	if p == nil {
		return
	}
	return p.Language
}

func (p *ProjectExtendedResponse) GetContentType() (value *string) {
	if p == nil {
		return
	}
	return p.ContentType
}

func (p *ProjectExtendedResponse) GetOriginalPublicationDate() (value *string) {
	if p == nil {
		return
	}
	return p.OriginalPublicationDate
}

func (p *ProjectExtendedResponse) GetMatureContent() (value *bool) {
	if p == nil {
		return
	}
	return p.MatureContent
}

func (p *ProjectExtendedResponse) GetIsbnNumber() (value *string) {
	if p == nil {
		return
	}
	return p.IsbnNumber
}

func (p *ProjectExtendedResponse) GetVolumeNormalization() (value bool) {
	if p == nil {
		return
	}
	return p.VolumeNormalization
}

func (p *ProjectExtendedResponse) GetState() (value ProjectState) {
	if p == nil {
		return
	}
	return p.State
}

func (p *ProjectExtendedResponse) GetAccessLevel() (value ProjectResponseAccessLevel) {
	if p == nil {
		return
	}
	return p.AccessLevel
}

func (p *ProjectExtendedResponse) GetFiction() (value *ProjectExtendedResponseFiction) {
	if p == nil {
		return
	}
	return p.Fiction
}

func (p *ProjectExtendedResponse) GetQualityCheckOn() (value bool) {
	if p == nil {
		return
	}
	return p.QualityCheckOn
}

func (p *ProjectExtendedResponse) GetQualityCheckOnWhenBulkConvert() (value bool) {
	if p == nil {
		return
	}
	return p.QualityCheckOnWhenBulkConvert
}

func (p *ProjectExtendedResponse) GetCreationMeta() (value *ProjectCreationMetaResponseModel) {
	if p == nil {
		return
	}
	return p.CreationMeta
}

func (p *ProjectExtendedResponse) GetSourceType() (value *ProjectExtendedResponseSourceType) {
	if p == nil {
		return
	}
	return p.SourceType
}

func (p *ProjectExtendedResponse) GetChaptersEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.ChaptersEnabled
}

func (p *ProjectExtendedResponse) GetCaptionsEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.CaptionsEnabled
}

func (p *ProjectExtendedResponse) GetCaptionStyle() (value *CaptionStyleModel) {
	if p == nil {
		return
	}
	return p.CaptionStyle
}

func (p *ProjectExtendedResponse) GetCaptionStyleTemplateOverrides() (value map[string]*CaptionStyleModel) {
	if p == nil {
		return
	}
	return p.CaptionStyleTemplateOverrides
}

func (p *ProjectExtendedResponse) GetPublicShareId() (value *string) {
	if p == nil {
		return
	}
	return p.PublicShareId
}

func (p *ProjectExtendedResponse) GetAspectRatio() (value *ProjectExtendedResponseAspectRatio) {
	if p == nil {
		return
	}
	return p.AspectRatio
}

func (p *ProjectExtendedResponse) GetAgentSettings() (value *StudioAgentSettingsModel) {
	if p == nil {
		return
	}
	return p.AgentSettings
}

func (p *ProjectExtendedResponse) GetQualityPreset() (value QualityPresetType) {
	if p == nil {
		return
	}
	return p.QualityPreset
}

func (p *ProjectExtendedResponse) GetChapters() (value []*ChapterResponse) {
	if p == nil {
		return
	}
	return p.Chapters
}

func (p *ProjectExtendedResponse) GetPronunciationDictionaryVersions() (value []*PronunciationDictionaryVersionResponseModel) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryVersions
}

func (p *ProjectExtendedResponse) GetPronunciationDictionaryLocators() (value []*PronunciationDictionaryLocatorResponseModel) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryLocators
}

func (p *ProjectExtendedResponse) GetApplyTextNormalization() (value ProjectExtendedResponseApplyTextNormalization) {
	if p == nil {
		return
	}
	return p.ApplyTextNormalization
}

func (p *ProjectExtendedResponse) GetExperimental() (value map[string]any) {
	if p == nil {
		return
	}
	return p.Experimental
}

func (p *ProjectExtendedResponse) GetAssets() (value []*ProjectExtendedResponseAssetsItem) {
	if p == nil {
		return
	}
	return p.Assets
}

func (p *ProjectExtendedResponse) GetVoices() (value []*ProjectVoiceResponseModel) {
	if p == nil {
		return
	}
	return p.Voices
}

func (p *ProjectExtendedResponse) GetBaseVoices() (value []*Voice) {
	if p == nil {
		return
	}
	return p.BaseVoices
}

func (p *ProjectExtendedResponse) GetPublishingRead() (value *DirectPublishingReadResponseModel) {
	if p == nil {
		return
	}
	return p.PublishingRead
}

func (p *ProjectExtendedResponse) GetDefaultTitleVoiceId() (value *string) {
	if p == nil {
		return
	}
	return p.DefaultTitleVoiceId
}

func (p *ProjectExtendedResponse) GetDefaultParagraphVoiceId() (value *string) {
	if p == nil {
		return
	}
	return p.DefaultParagraphVoiceId
}

func (p *ProjectExtendedResponse) SetProjectId(projectId string) {
	p.ProjectId = projectId
}

func (p *ProjectExtendedResponse) SetName(name string) {
	p.Name = name
}

func (p *ProjectExtendedResponse) SetCreateDateUnix(createDateUnix int) {
	p.CreateDateUnix = createDateUnix
}

func (p *ProjectExtendedResponse) SetCreatedByUserId(createdByUserId *string) {
	p.CreatedByUserId = createdByUserId
}

func (p *ProjectExtendedResponse) SetDefaultTitleVoiceRefId(defaultTitleVoiceRefId string) {
	p.DefaultTitleVoiceRefId = defaultTitleVoiceRefId
}

func (p *ProjectExtendedResponse) SetDefaultParagraphVoiceRefId(defaultParagraphVoiceRefId string) {
	p.DefaultParagraphVoiceRefId = defaultParagraphVoiceRefId
}

func (p *ProjectExtendedResponse) SetDefaultModelId(defaultModelId string) {
	p.DefaultModelId = defaultModelId
}

func (p *ProjectExtendedResponse) SetLastConversionDateUnix(lastConversionDateUnix *int) {
	p.LastConversionDateUnix = lastConversionDateUnix
}

func (p *ProjectExtendedResponse) SetCanBeDownloaded(canBeDownloaded bool) {
	p.CanBeDownloaded = canBeDownloaded
}

func (p *ProjectExtendedResponse) SetTitle(title *string) {
	p.Title = title
}

func (p *ProjectExtendedResponse) SetAuthor(author *string) {
	p.Author = author
}

func (p *ProjectExtendedResponse) SetDescription(description *string) {
	p.Description = description
}

func (p *ProjectExtendedResponse) SetGenres(genres []string) {
	p.Genres = genres
}

func (p *ProjectExtendedResponse) SetCoverImageUrl(coverImageUrl *string) {
	p.CoverImageUrl = coverImageUrl
}

func (p *ProjectExtendedResponse) SetTargetAudience(targetAudience *ProjectExtendedResponseTargetAudience) {
	p.TargetAudience = targetAudience
}

func (p *ProjectExtendedResponse) SetLanguage(language *string) {
	p.Language = language
}

func (p *ProjectExtendedResponse) SetContentType(contentType *string) {
	p.ContentType = contentType
}

func (p *ProjectExtendedResponse) SetOriginalPublicationDate(originalPublicationDate *string) {
	p.OriginalPublicationDate = originalPublicationDate
}

func (p *ProjectExtendedResponse) SetMatureContent(matureContent *bool) {
	p.MatureContent = matureContent
}

func (p *ProjectExtendedResponse) SetIsbnNumber(isbnNumber *string) {
	p.IsbnNumber = isbnNumber
}

func (p *ProjectExtendedResponse) SetVolumeNormalization(volumeNormalization bool) {
	p.VolumeNormalization = volumeNormalization
}

func (p *ProjectExtendedResponse) SetState(state ProjectState) {
	p.State = state
}

func (p *ProjectExtendedResponse) SetAccessLevel(accessLevel ProjectResponseAccessLevel) {
	p.AccessLevel = accessLevel
}

func (p *ProjectExtendedResponse) SetFiction(fiction *ProjectExtendedResponseFiction) {
	p.Fiction = fiction
}

func (p *ProjectExtendedResponse) SetQualityCheckOn(qualityCheckOn bool) {
	p.QualityCheckOn = qualityCheckOn
}

func (p *ProjectExtendedResponse) SetQualityCheckOnWhenBulkConvert(qualityCheckOnWhenBulkConvert bool) {
	p.QualityCheckOnWhenBulkConvert = qualityCheckOnWhenBulkConvert
}

func (p *ProjectExtendedResponse) SetCreationMeta(creationMeta *ProjectCreationMetaResponseModel) {
	p.CreationMeta = creationMeta
}

func (p *ProjectExtendedResponse) SetSourceType(sourceType *ProjectExtendedResponseSourceType) {
	p.SourceType = sourceType
}

func (p *ProjectExtendedResponse) SetChaptersEnabled(chaptersEnabled *bool) {
	p.ChaptersEnabled = chaptersEnabled
}

func (p *ProjectExtendedResponse) SetCaptionsEnabled(captionsEnabled *bool) {
	p.CaptionsEnabled = captionsEnabled
}

func (p *ProjectExtendedResponse) SetCaptionStyle(captionStyle *CaptionStyleModel) {
	p.CaptionStyle = captionStyle
}

func (p *ProjectExtendedResponse) SetCaptionStyleTemplateOverrides(captionStyleTemplateOverrides map[string]*CaptionStyleModel) {
	p.CaptionStyleTemplateOverrides = captionStyleTemplateOverrides
}

func (p *ProjectExtendedResponse) SetPublicShareId(publicShareId *string) {
	p.PublicShareId = publicShareId
}

func (p *ProjectExtendedResponse) SetAspectRatio(aspectRatio *ProjectExtendedResponseAspectRatio) {
	p.AspectRatio = aspectRatio
}

func (p *ProjectExtendedResponse) SetAgentSettings(agentSettings *StudioAgentSettingsModel) {
	p.AgentSettings = agentSettings
}

func (p *ProjectExtendedResponse) SetQualityPreset(qualityPreset QualityPresetType) {
	p.QualityPreset = qualityPreset
}

func (p *ProjectExtendedResponse) SetChapters(chapters []*ChapterResponse) {
	p.Chapters = chapters
}

func (p *ProjectExtendedResponse) SetPronunciationDictionaryVersions(pronunciationDictionaryVersions []*PronunciationDictionaryVersionResponseModel) {
	p.PronunciationDictionaryVersions = pronunciationDictionaryVersions
}

func (p *ProjectExtendedResponse) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryLocatorResponseModel) {
	p.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (p *ProjectExtendedResponse) SetApplyTextNormalization(applyTextNormalization ProjectExtendedResponseApplyTextNormalization) {
	p.ApplyTextNormalization = applyTextNormalization
}

func (p *ProjectExtendedResponse) SetExperimental(experimental map[string]any) {
	p.Experimental = experimental
}

func (p *ProjectExtendedResponse) SetAssets(assets []*ProjectExtendedResponseAssetsItem) {
	p.Assets = assets
}

func (p *ProjectExtendedResponse) SetVoices(voices []*ProjectVoiceResponseModel) {
	p.Voices = voices
}

func (p *ProjectExtendedResponse) SetBaseVoices(baseVoices []*Voice) {
	p.BaseVoices = baseVoices
}

func (p *ProjectExtendedResponse) SetPublishingRead(publishingRead *DirectPublishingReadResponseModel) {
	p.PublishingRead = publishingRead
}

func (p *ProjectExtendedResponse) SetDefaultTitleVoiceId(defaultTitleVoiceId *string) {
	p.DefaultTitleVoiceId = defaultTitleVoiceId
}

func (p *ProjectExtendedResponse) SetDefaultParagraphVoiceId(defaultParagraphVoiceId *string) {
	p.DefaultParagraphVoiceId = defaultParagraphVoiceId
}

func (p *ProjectExtendedResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectExtendedResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectExtendedResponse to string"
	}
	return string(jsonData)
}

func (p *ProjectExtendedResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectExtendedResponse) MarshalJSON() ([]byte, error) {
	type embed ProjectExtendedResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
