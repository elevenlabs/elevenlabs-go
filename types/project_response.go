package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectResponse struct {
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
	TargetAudience *ProjectResponseTargetAudience `json:"target_audience,omitempty" xml:"target_audience,omitempty"`
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
	Fiction *ProjectResponseFiction `json:"fiction,omitempty" xml:"fiction,omitempty"`
	// Whether quality check is enabled for this project.
	QualityCheckOn bool `json:"quality_check_on" xml:"quality_check_on" required:"true"`
	// Whether quality check is enabled on the project when bulk converting.
	QualityCheckOnWhenBulkConvert bool `json:"quality_check_on_when_bulk_convert" xml:"quality_check_on_when_bulk_convert" required:"true"`
	// The creation meta of the project.
	CreationMeta *ProjectCreationMetaResponseModel `json:"creation_meta,omitempty" xml:"creation_meta,omitempty"`
	// The source type of the project.
	SourceType *ProjectResponseSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
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
	AspectRatio *ProjectResponseAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// Agent-related settings for the project
	AgentSettings *StudioAgentSettingsModel `json:"agent_settings,omitempty" xml:"agent_settings,omitempty"`
	// The default title voice ID.
	DefaultTitleVoiceId *string `json:"default_title_voice_id,omitempty" xml:"default_title_voice_id,omitempty"`
	// The default paragraph voice ID.
	DefaultParagraphVoiceId *string `json:"default_paragraph_voice_id,omitempty" xml:"default_paragraph_voice_id,omitempty"`
	extraProperties         map[string]interface{}
}

func (p *ProjectResponse) GetProjectId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectId
}

func (p *ProjectResponse) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProjectResponse) GetCreateDateUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreateDateUnix
}

func (p *ProjectResponse) GetCreatedByUserId() (value *string) {
	if p == nil {
		return
	}
	return p.CreatedByUserId
}

func (p *ProjectResponse) GetDefaultTitleVoiceRefId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultTitleVoiceRefId
}

func (p *ProjectResponse) GetDefaultParagraphVoiceRefId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultParagraphVoiceRefId
}

func (p *ProjectResponse) GetDefaultModelId() (value string) {
	if p == nil {
		return
	}
	return p.DefaultModelId
}

func (p *ProjectResponse) GetLastConversionDateUnix() (value *int) {
	if p == nil {
		return
	}
	return p.LastConversionDateUnix
}

func (p *ProjectResponse) GetCanBeDownloaded() (value bool) {
	if p == nil {
		return
	}
	return p.CanBeDownloaded
}

func (p *ProjectResponse) GetTitle() (value *string) {
	if p == nil {
		return
	}
	return p.Title
}

func (p *ProjectResponse) GetAuthor() (value *string) {
	if p == nil {
		return
	}
	return p.Author
}

func (p *ProjectResponse) GetDescription() (value *string) {
	if p == nil {
		return
	}
	return p.Description
}

func (p *ProjectResponse) GetGenres() (value []string) {
	if p == nil {
		return
	}
	return p.Genres
}

func (p *ProjectResponse) GetCoverImageUrl() (value *string) {
	if p == nil {
		return
	}
	return p.CoverImageUrl
}

func (p *ProjectResponse) GetTargetAudience() (value *ProjectResponseTargetAudience) {
	if p == nil {
		return
	}
	return p.TargetAudience
}

func (p *ProjectResponse) GetLanguage() (value *string) {
	if p == nil {
		return
	}
	return p.Language
}

func (p *ProjectResponse) GetContentType() (value *string) {
	if p == nil {
		return
	}
	return p.ContentType
}

func (p *ProjectResponse) GetOriginalPublicationDate() (value *string) {
	if p == nil {
		return
	}
	return p.OriginalPublicationDate
}

func (p *ProjectResponse) GetMatureContent() (value *bool) {
	if p == nil {
		return
	}
	return p.MatureContent
}

func (p *ProjectResponse) GetIsbnNumber() (value *string) {
	if p == nil {
		return
	}
	return p.IsbnNumber
}

func (p *ProjectResponse) GetVolumeNormalization() (value bool) {
	if p == nil {
		return
	}
	return p.VolumeNormalization
}

func (p *ProjectResponse) GetState() (value ProjectState) {
	if p == nil {
		return
	}
	return p.State
}

func (p *ProjectResponse) GetAccessLevel() (value ProjectResponseAccessLevel) {
	if p == nil {
		return
	}
	return p.AccessLevel
}

func (p *ProjectResponse) GetFiction() (value *ProjectResponseFiction) {
	if p == nil {
		return
	}
	return p.Fiction
}

func (p *ProjectResponse) GetQualityCheckOn() (value bool) {
	if p == nil {
		return
	}
	return p.QualityCheckOn
}

func (p *ProjectResponse) GetQualityCheckOnWhenBulkConvert() (value bool) {
	if p == nil {
		return
	}
	return p.QualityCheckOnWhenBulkConvert
}

func (p *ProjectResponse) GetCreationMeta() (value *ProjectCreationMetaResponseModel) {
	if p == nil {
		return
	}
	return p.CreationMeta
}

func (p *ProjectResponse) GetSourceType() (value *ProjectResponseSourceType) {
	if p == nil {
		return
	}
	return p.SourceType
}

func (p *ProjectResponse) GetChaptersEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.ChaptersEnabled
}

func (p *ProjectResponse) GetCaptionsEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.CaptionsEnabled
}

func (p *ProjectResponse) GetCaptionStyle() (value *CaptionStyleModel) {
	if p == nil {
		return
	}
	return p.CaptionStyle
}

func (p *ProjectResponse) GetCaptionStyleTemplateOverrides() (value map[string]*CaptionStyleModel) {
	if p == nil {
		return
	}
	return p.CaptionStyleTemplateOverrides
}

func (p *ProjectResponse) GetPublicShareId() (value *string) {
	if p == nil {
		return
	}
	return p.PublicShareId
}

func (p *ProjectResponse) GetAspectRatio() (value *ProjectResponseAspectRatio) {
	if p == nil {
		return
	}
	return p.AspectRatio
}

func (p *ProjectResponse) GetAgentSettings() (value *StudioAgentSettingsModel) {
	if p == nil {
		return
	}
	return p.AgentSettings
}

func (p *ProjectResponse) GetDefaultTitleVoiceId() (value *string) {
	if p == nil {
		return
	}
	return p.DefaultTitleVoiceId
}

func (p *ProjectResponse) GetDefaultParagraphVoiceId() (value *string) {
	if p == nil {
		return
	}
	return p.DefaultParagraphVoiceId
}

func (p *ProjectResponse) SetProjectId(projectId string) {
	p.ProjectId = projectId
}

func (p *ProjectResponse) SetName(name string) {
	p.Name = name
}

func (p *ProjectResponse) SetCreateDateUnix(createDateUnix int) {
	p.CreateDateUnix = createDateUnix
}

func (p *ProjectResponse) SetCreatedByUserId(createdByUserId *string) {
	p.CreatedByUserId = createdByUserId
}

func (p *ProjectResponse) SetDefaultTitleVoiceRefId(defaultTitleVoiceRefId string) {
	p.DefaultTitleVoiceRefId = defaultTitleVoiceRefId
}

func (p *ProjectResponse) SetDefaultParagraphVoiceRefId(defaultParagraphVoiceRefId string) {
	p.DefaultParagraphVoiceRefId = defaultParagraphVoiceRefId
}

func (p *ProjectResponse) SetDefaultModelId(defaultModelId string) {
	p.DefaultModelId = defaultModelId
}

func (p *ProjectResponse) SetLastConversionDateUnix(lastConversionDateUnix *int) {
	p.LastConversionDateUnix = lastConversionDateUnix
}

func (p *ProjectResponse) SetCanBeDownloaded(canBeDownloaded bool) {
	p.CanBeDownloaded = canBeDownloaded
}

func (p *ProjectResponse) SetTitle(title *string) {
	p.Title = title
}

func (p *ProjectResponse) SetAuthor(author *string) {
	p.Author = author
}

func (p *ProjectResponse) SetDescription(description *string) {
	p.Description = description
}

func (p *ProjectResponse) SetGenres(genres []string) {
	p.Genres = genres
}

func (p *ProjectResponse) SetCoverImageUrl(coverImageUrl *string) {
	p.CoverImageUrl = coverImageUrl
}

func (p *ProjectResponse) SetTargetAudience(targetAudience *ProjectResponseTargetAudience) {
	p.TargetAudience = targetAudience
}

func (p *ProjectResponse) SetLanguage(language *string) {
	p.Language = language
}

func (p *ProjectResponse) SetContentType(contentType *string) {
	p.ContentType = contentType
}

func (p *ProjectResponse) SetOriginalPublicationDate(originalPublicationDate *string) {
	p.OriginalPublicationDate = originalPublicationDate
}

func (p *ProjectResponse) SetMatureContent(matureContent *bool) {
	p.MatureContent = matureContent
}

func (p *ProjectResponse) SetIsbnNumber(isbnNumber *string) {
	p.IsbnNumber = isbnNumber
}

func (p *ProjectResponse) SetVolumeNormalization(volumeNormalization bool) {
	p.VolumeNormalization = volumeNormalization
}

func (p *ProjectResponse) SetState(state ProjectState) {
	p.State = state
}

func (p *ProjectResponse) SetAccessLevel(accessLevel ProjectResponseAccessLevel) {
	p.AccessLevel = accessLevel
}

func (p *ProjectResponse) SetFiction(fiction *ProjectResponseFiction) {
	p.Fiction = fiction
}

func (p *ProjectResponse) SetQualityCheckOn(qualityCheckOn bool) {
	p.QualityCheckOn = qualityCheckOn
}

func (p *ProjectResponse) SetQualityCheckOnWhenBulkConvert(qualityCheckOnWhenBulkConvert bool) {
	p.QualityCheckOnWhenBulkConvert = qualityCheckOnWhenBulkConvert
}

func (p *ProjectResponse) SetCreationMeta(creationMeta *ProjectCreationMetaResponseModel) {
	p.CreationMeta = creationMeta
}

func (p *ProjectResponse) SetSourceType(sourceType *ProjectResponseSourceType) {
	p.SourceType = sourceType
}

func (p *ProjectResponse) SetChaptersEnabled(chaptersEnabled *bool) {
	p.ChaptersEnabled = chaptersEnabled
}

func (p *ProjectResponse) SetCaptionsEnabled(captionsEnabled *bool) {
	p.CaptionsEnabled = captionsEnabled
}

func (p *ProjectResponse) SetCaptionStyle(captionStyle *CaptionStyleModel) {
	p.CaptionStyle = captionStyle
}

func (p *ProjectResponse) SetCaptionStyleTemplateOverrides(captionStyleTemplateOverrides map[string]*CaptionStyleModel) {
	p.CaptionStyleTemplateOverrides = captionStyleTemplateOverrides
}

func (p *ProjectResponse) SetPublicShareId(publicShareId *string) {
	p.PublicShareId = publicShareId
}

func (p *ProjectResponse) SetAspectRatio(aspectRatio *ProjectResponseAspectRatio) {
	p.AspectRatio = aspectRatio
}

func (p *ProjectResponse) SetAgentSettings(agentSettings *StudioAgentSettingsModel) {
	p.AgentSettings = agentSettings
}

func (p *ProjectResponse) SetDefaultTitleVoiceId(defaultTitleVoiceId *string) {
	p.DefaultTitleVoiceId = defaultTitleVoiceId
}

func (p *ProjectResponse) SetDefaultParagraphVoiceId(defaultParagraphVoiceId *string) {
	p.DefaultParagraphVoiceId = defaultParagraphVoiceId
}

func (p *ProjectResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectResponse to string"
	}
	return string(jsonData)
}

func (p *ProjectResponse) UnmarshalJSON(data []byte) error {
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

func (p *ProjectResponse) MarshalJSON() ([]byte, error) {
	type embed ProjectResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
