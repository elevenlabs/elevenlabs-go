package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DirectPublishingReadResponseModel struct {
	ReadId                       string                                           `json:"read_id" xml:"read_id" required:"true"`
	CreatedAtUnix                int                                              `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	UpdatedAtUnix                int                                              `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	WordCount                    int                                              `json:"word_count" xml:"word_count" required:"true"`
	CharCount                    int                                              `json:"char_count" xml:"char_count" required:"true"`
	Chapters                     []*ReadMetadataChapterDbModel                    `json:"chapters" xml:"chapters" required:"true"`
	Title                        *string                                          `json:"title,omitempty" xml:"title,omitempty"`
	Author                       *string                                          `json:"author,omitempty" xml:"author,omitempty"`
	Description                  *string                                          `json:"description,omitempty" xml:"description,omitempty"`
	ArticleImageUrl              *string                                          `json:"article_image_url,omitempty" xml:"article_image_url,omitempty"`
	Language                     *string                                          `json:"language,omitempty" xml:"language,omitempty"`
	Locale                       *string                                          `json:"locale,omitempty" xml:"locale,omitempty"`
	DisplayMode                  *DirectPublishingReadResponseModelDisplayMode    `json:"display_mode,omitempty" xml:"display_mode,omitempty"`
	Genre                        []DirectPublishingReadResponseModelGenreItem     `json:"genre,omitempty" xml:"genre,omitempty"`
	Fiction                      *string                                          `json:"fiction,omitempty" xml:"fiction,omitempty"`
	ContentType                  *string                                          `json:"content_type,omitempty" xml:"content_type,omitempty"`
	OriginalFileType             *string                                          `json:"original_file_type,omitempty" xml:"original_file_type,omitempty"`
	TargetAudience               *DirectPublishingReadResponseModelTargetAudience `json:"target_audience,omitempty" xml:"target_audience,omitempty"`
	MatureContent                *bool                                            `json:"mature_content,omitempty" xml:"mature_content,omitempty"`
	SafesearchAdult              *bool                                            `json:"safesearch_adult,omitempty" xml:"safesearch_adult,omitempty"`
	Origin                       *string                                          `json:"origin,omitempty" xml:"origin,omitempty"`
	PublicationDate              *string                                          `json:"publication_date,omitempty" xml:"publication_date,omitempty"`
	Isbn                         *string                                          `json:"isbn,omitempty" xml:"isbn,omitempty"`
	Ean                          *string                                          `json:"ean,omitempty" xml:"ean,omitempty"`
	LegalTerms                   *ReadLegalTerms                                  `json:"legal_terms,omitempty" xml:"legal_terms,omitempty"`
	ContentGuidelinesTerms       *ReadLegalTerms                                  `json:"content_guidelines_terms,omitempty" xml:"content_guidelines_terms,omitempty"`
	LastUpdatedFromProjectUnix   *int                                             `json:"last_updated_from_project_unix,omitempty" xml:"last_updated_from_project_unix,omitempty"`
	PublishingProjectId          *string                                          `json:"publishing_project_id,omitempty" xml:"publishing_project_id,omitempty"`
	PublishingState              *string                                          `json:"publishing_state,omitempty" xml:"publishing_state,omitempty"`
	PublisherProfileId           *string                                          `json:"publisher_profile_id,omitempty" xml:"publisher_profile_id,omitempty"`
	QualityScore                 *int                                             `json:"quality_score,omitempty" xml:"quality_score,omitempty"`
	Publisher                    *string                                          `json:"publisher,omitempty" xml:"publisher,omitempty"`
	Copyright                    *string                                          `json:"copyright,omitempty" xml:"copyright,omitempty"`
	Subtitle                     *string                                          `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	DistributionTerritories      []string                                         `json:"distribution_territories,omitempty" xml:"distribution_territories,omitempty"`
	Edition                      *string                                          `json:"edition,omitempty" xml:"edition,omitempty"`
	Contributors                 []*Contributor                                   `json:"contributors,omitempty" xml:"contributors,omitempty"`
	PayoutType                   *DirectPublishingReadResponseModelPayoutType     `json:"payout_type,omitempty" xml:"payout_type,omitempty"`
	ListPrice                    *float64                                         `json:"list_price,omitempty" xml:"list_price,omitempty"`
	Currency                     *string                                          `json:"currency,omitempty" xml:"currency,omitempty"`
	OriginalAudioProjectExportId *string                                          `json:"original_audio_project_export_id,omitempty" xml:"original_audio_project_export_id,omitempty"`
	OriginalAudioDocumentId      *string                                          `json:"original_audio_document_id,omitempty" xml:"original_audio_document_id,omitempty"`
	SeriesId                     *string                                          `json:"series_id,omitempty" xml:"series_id,omitempty"`
	Volume                       *int                                             `json:"volume,omitempty" xml:"volume,omitempty"`
	PublishedAtUnix              *int                                             `json:"published_at_unix,omitempty" xml:"published_at_unix,omitempty"`
	ReadSlug                     *string                                          `json:"read_slug,omitempty" xml:"read_slug,omitempty"`
	PreviewAudioObject           *PreviewAudioDbModel                             `json:"preview_audio_object,omitempty" xml:"preview_audio_object,omitempty"`
	SampleConfig                 *SampleConfigDbModel                             `json:"sample_config,omitempty" xml:"sample_config,omitempty"`
	Review                       *ReviewResponseModel                             `json:"review,omitempty" xml:"review,omitempty"`
	VoiceId                      *string                                          `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	CanUseAssistant              *bool                                            `json:"can_use_assistant,omitempty" xml:"can_use_assistant,omitempty"`
	IsVoiceChangerOn             *bool                                            `json:"is_voice_changer_on,omitempty" xml:"is_voice_changer_on,omitempty"`
	RestrictedToUserEmailDomains []string                                         `json:"restricted_to_user_email_domains,omitempty" xml:"restricted_to_user_email_domains,omitempty"`
	extraProperties              map[string]interface{}
}

func (d *DirectPublishingReadResponseModel) GetReadId() (value string) {
	if d == nil {
		return
	}
	return d.ReadId
}

func (d *DirectPublishingReadResponseModel) GetCreatedAtUnix() (value int) {
	if d == nil {
		return
	}
	return d.CreatedAtUnix
}

func (d *DirectPublishingReadResponseModel) GetUpdatedAtUnix() (value int) {
	if d == nil {
		return
	}
	return d.UpdatedAtUnix
}

func (d *DirectPublishingReadResponseModel) GetWordCount() (value int) {
	if d == nil {
		return
	}
	return d.WordCount
}

func (d *DirectPublishingReadResponseModel) GetCharCount() (value int) {
	if d == nil {
		return
	}
	return d.CharCount
}

func (d *DirectPublishingReadResponseModel) GetChapters() (value []*ReadMetadataChapterDbModel) {
	if d == nil {
		return
	}
	return d.Chapters
}

func (d *DirectPublishingReadResponseModel) GetTitle() (value *string) {
	if d == nil {
		return
	}
	return d.Title
}

func (d *DirectPublishingReadResponseModel) GetAuthor() (value *string) {
	if d == nil {
		return
	}
	return d.Author
}

func (d *DirectPublishingReadResponseModel) GetDescription() (value *string) {
	if d == nil {
		return
	}
	return d.Description
}

func (d *DirectPublishingReadResponseModel) GetArticleImageUrl() (value *string) {
	if d == nil {
		return
	}
	return d.ArticleImageUrl
}

func (d *DirectPublishingReadResponseModel) GetLanguage() (value *string) {
	if d == nil {
		return
	}
	return d.Language
}

func (d *DirectPublishingReadResponseModel) GetLocale() (value *string) {
	if d == nil {
		return
	}
	return d.Locale
}

func (d *DirectPublishingReadResponseModel) GetDisplayMode() (value *DirectPublishingReadResponseModelDisplayMode) {
	if d == nil {
		return
	}
	return d.DisplayMode
}

func (d *DirectPublishingReadResponseModel) GetGenre() (value []DirectPublishingReadResponseModelGenreItem) {
	if d == nil {
		return
	}
	return d.Genre
}

func (d *DirectPublishingReadResponseModel) GetFiction() (value *string) {
	if d == nil {
		return
	}
	return d.Fiction
}

func (d *DirectPublishingReadResponseModel) GetContentType() (value *string) {
	if d == nil {
		return
	}
	return d.ContentType
}

func (d *DirectPublishingReadResponseModel) GetOriginalFileType() (value *string) {
	if d == nil {
		return
	}
	return d.OriginalFileType
}

func (d *DirectPublishingReadResponseModel) GetTargetAudience() (value *DirectPublishingReadResponseModelTargetAudience) {
	if d == nil {
		return
	}
	return d.TargetAudience
}

func (d *DirectPublishingReadResponseModel) GetMatureContent() (value *bool) {
	if d == nil {
		return
	}
	return d.MatureContent
}

func (d *DirectPublishingReadResponseModel) GetSafesearchAdult() (value *bool) {
	if d == nil {
		return
	}
	return d.SafesearchAdult
}

func (d *DirectPublishingReadResponseModel) GetOrigin() (value *string) {
	if d == nil {
		return
	}
	return d.Origin
}

func (d *DirectPublishingReadResponseModel) GetPublicationDate() (value *string) {
	if d == nil {
		return
	}
	return d.PublicationDate
}

func (d *DirectPublishingReadResponseModel) GetIsbn() (value *string) {
	if d == nil {
		return
	}
	return d.Isbn
}

func (d *DirectPublishingReadResponseModel) GetEan() (value *string) {
	if d == nil {
		return
	}
	return d.Ean
}

func (d *DirectPublishingReadResponseModel) GetLegalTerms() (value *ReadLegalTerms) {
	if d == nil {
		return
	}
	return d.LegalTerms
}

func (d *DirectPublishingReadResponseModel) GetContentGuidelinesTerms() (value *ReadLegalTerms) {
	if d == nil {
		return
	}
	return d.ContentGuidelinesTerms
}

func (d *DirectPublishingReadResponseModel) GetLastUpdatedFromProjectUnix() (value *int) {
	if d == nil {
		return
	}
	return d.LastUpdatedFromProjectUnix
}

func (d *DirectPublishingReadResponseModel) GetPublishingProjectId() (value *string) {
	if d == nil {
		return
	}
	return d.PublishingProjectId
}

func (d *DirectPublishingReadResponseModel) GetPublishingState() (value *string) {
	if d == nil {
		return
	}
	return d.PublishingState
}

func (d *DirectPublishingReadResponseModel) GetPublisherProfileId() (value *string) {
	if d == nil {
		return
	}
	return d.PublisherProfileId
}

func (d *DirectPublishingReadResponseModel) GetQualityScore() (value *int) {
	if d == nil {
		return
	}
	return d.QualityScore
}

func (d *DirectPublishingReadResponseModel) GetPublisher() (value *string) {
	if d == nil {
		return
	}
	return d.Publisher
}

func (d *DirectPublishingReadResponseModel) GetCopyright() (value *string) {
	if d == nil {
		return
	}
	return d.Copyright
}

func (d *DirectPublishingReadResponseModel) GetSubtitle() (value *string) {
	if d == nil {
		return
	}
	return d.Subtitle
}

func (d *DirectPublishingReadResponseModel) GetDistributionTerritories() (value []string) {
	if d == nil {
		return
	}
	return d.DistributionTerritories
}

func (d *DirectPublishingReadResponseModel) GetEdition() (value *string) {
	if d == nil {
		return
	}
	return d.Edition
}

func (d *DirectPublishingReadResponseModel) GetContributors() (value []*Contributor) {
	if d == nil {
		return
	}
	return d.Contributors
}

func (d *DirectPublishingReadResponseModel) GetPayoutType() (value *DirectPublishingReadResponseModelPayoutType) {
	if d == nil {
		return
	}
	return d.PayoutType
}

func (d *DirectPublishingReadResponseModel) GetListPrice() (value *float64) {
	if d == nil {
		return
	}
	return d.ListPrice
}

func (d *DirectPublishingReadResponseModel) GetCurrency() (value *string) {
	if d == nil {
		return
	}
	return d.Currency
}

func (d *DirectPublishingReadResponseModel) GetOriginalAudioProjectExportId() (value *string) {
	if d == nil {
		return
	}
	return d.OriginalAudioProjectExportId
}

func (d *DirectPublishingReadResponseModel) GetOriginalAudioDocumentId() (value *string) {
	if d == nil {
		return
	}
	return d.OriginalAudioDocumentId
}

func (d *DirectPublishingReadResponseModel) GetSeriesId() (value *string) {
	if d == nil {
		return
	}
	return d.SeriesId
}

func (d *DirectPublishingReadResponseModel) GetVolume() (value *int) {
	if d == nil {
		return
	}
	return d.Volume
}

func (d *DirectPublishingReadResponseModel) GetPublishedAtUnix() (value *int) {
	if d == nil {
		return
	}
	return d.PublishedAtUnix
}

func (d *DirectPublishingReadResponseModel) GetReadSlug() (value *string) {
	if d == nil {
		return
	}
	return d.ReadSlug
}

func (d *DirectPublishingReadResponseModel) GetPreviewAudioObject() (value *PreviewAudioDbModel) {
	if d == nil {
		return
	}
	return d.PreviewAudioObject
}

func (d *DirectPublishingReadResponseModel) GetSampleConfig() (value *SampleConfigDbModel) {
	if d == nil {
		return
	}
	return d.SampleConfig
}

func (d *DirectPublishingReadResponseModel) GetReview() (value *ReviewResponseModel) {
	if d == nil {
		return
	}
	return d.Review
}

func (d *DirectPublishingReadResponseModel) GetVoiceId() (value *string) {
	if d == nil {
		return
	}
	return d.VoiceId
}

func (d *DirectPublishingReadResponseModel) GetCanUseAssistant() (value *bool) {
	if d == nil {
		return
	}
	return d.CanUseAssistant
}

func (d *DirectPublishingReadResponseModel) GetIsVoiceChangerOn() (value *bool) {
	if d == nil {
		return
	}
	return d.IsVoiceChangerOn
}

func (d *DirectPublishingReadResponseModel) GetRestrictedToUserEmailDomains() (value []string) {
	if d == nil {
		return
	}
	return d.RestrictedToUserEmailDomains
}

func (d *DirectPublishingReadResponseModel) SetReadId(readId string) {
	d.ReadId = readId
}

func (d *DirectPublishingReadResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	d.CreatedAtUnix = createdAtUnix
}

func (d *DirectPublishingReadResponseModel) SetUpdatedAtUnix(updatedAtUnix int) {
	d.UpdatedAtUnix = updatedAtUnix
}

func (d *DirectPublishingReadResponseModel) SetWordCount(wordCount int) {
	d.WordCount = wordCount
}

func (d *DirectPublishingReadResponseModel) SetCharCount(charCount int) {
	d.CharCount = charCount
}

func (d *DirectPublishingReadResponseModel) SetChapters(chapters []*ReadMetadataChapterDbModel) {
	d.Chapters = chapters
}

func (d *DirectPublishingReadResponseModel) SetTitle(title *string) {
	d.Title = title
}

func (d *DirectPublishingReadResponseModel) SetAuthor(author *string) {
	d.Author = author
}

func (d *DirectPublishingReadResponseModel) SetDescription(description *string) {
	d.Description = description
}

func (d *DirectPublishingReadResponseModel) SetArticleImageUrl(articleImageUrl *string) {
	d.ArticleImageUrl = articleImageUrl
}

func (d *DirectPublishingReadResponseModel) SetLanguage(language *string) {
	d.Language = language
}

func (d *DirectPublishingReadResponseModel) SetLocale(locale *string) {
	d.Locale = locale
}

func (d *DirectPublishingReadResponseModel) SetDisplayMode(displayMode *DirectPublishingReadResponseModelDisplayMode) {
	d.DisplayMode = displayMode
}

func (d *DirectPublishingReadResponseModel) SetGenre(genre []DirectPublishingReadResponseModelGenreItem) {
	d.Genre = genre
}

func (d *DirectPublishingReadResponseModel) SetFiction(fiction *string) {
	d.Fiction = fiction
}

func (d *DirectPublishingReadResponseModel) SetContentType(contentType *string) {
	d.ContentType = contentType
}

func (d *DirectPublishingReadResponseModel) SetOriginalFileType(originalFileType *string) {
	d.OriginalFileType = originalFileType
}

func (d *DirectPublishingReadResponseModel) SetTargetAudience(targetAudience *DirectPublishingReadResponseModelTargetAudience) {
	d.TargetAudience = targetAudience
}

func (d *DirectPublishingReadResponseModel) SetMatureContent(matureContent *bool) {
	d.MatureContent = matureContent
}

func (d *DirectPublishingReadResponseModel) SetSafesearchAdult(safesearchAdult *bool) {
	d.SafesearchAdult = safesearchAdult
}

func (d *DirectPublishingReadResponseModel) SetOrigin(origin *string) {
	d.Origin = origin
}

func (d *DirectPublishingReadResponseModel) SetPublicationDate(publicationDate *string) {
	d.PublicationDate = publicationDate
}

func (d *DirectPublishingReadResponseModel) SetIsbn(isbn *string) {
	d.Isbn = isbn
}

func (d *DirectPublishingReadResponseModel) SetEan(ean *string) {
	d.Ean = ean
}

func (d *DirectPublishingReadResponseModel) SetLegalTerms(legalTerms *ReadLegalTerms) {
	d.LegalTerms = legalTerms
}

func (d *DirectPublishingReadResponseModel) SetContentGuidelinesTerms(contentGuidelinesTerms *ReadLegalTerms) {
	d.ContentGuidelinesTerms = contentGuidelinesTerms
}

func (d *DirectPublishingReadResponseModel) SetLastUpdatedFromProjectUnix(lastUpdatedFromProjectUnix *int) {
	d.LastUpdatedFromProjectUnix = lastUpdatedFromProjectUnix
}

func (d *DirectPublishingReadResponseModel) SetPublishingProjectId(publishingProjectId *string) {
	d.PublishingProjectId = publishingProjectId
}

func (d *DirectPublishingReadResponseModel) SetPublishingState(publishingState *string) {
	d.PublishingState = publishingState
}

func (d *DirectPublishingReadResponseModel) SetPublisherProfileId(publisherProfileId *string) {
	d.PublisherProfileId = publisherProfileId
}

func (d *DirectPublishingReadResponseModel) SetQualityScore(qualityScore *int) {
	d.QualityScore = qualityScore
}

func (d *DirectPublishingReadResponseModel) SetPublisher(publisher *string) {
	d.Publisher = publisher
}

func (d *DirectPublishingReadResponseModel) SetCopyright(copyright *string) {
	d.Copyright = copyright
}

func (d *DirectPublishingReadResponseModel) SetSubtitle(subtitle *string) {
	d.Subtitle = subtitle
}

func (d *DirectPublishingReadResponseModel) SetDistributionTerritories(distributionTerritories []string) {
	d.DistributionTerritories = distributionTerritories
}

func (d *DirectPublishingReadResponseModel) SetEdition(edition *string) {
	d.Edition = edition
}

func (d *DirectPublishingReadResponseModel) SetContributors(contributors []*Contributor) {
	d.Contributors = contributors
}

func (d *DirectPublishingReadResponseModel) SetPayoutType(payoutType *DirectPublishingReadResponseModelPayoutType) {
	d.PayoutType = payoutType
}

func (d *DirectPublishingReadResponseModel) SetListPrice(listPrice *float64) {
	d.ListPrice = listPrice
}

func (d *DirectPublishingReadResponseModel) SetCurrency(currency *string) {
	d.Currency = currency
}

func (d *DirectPublishingReadResponseModel) SetOriginalAudioProjectExportId(originalAudioProjectExportId *string) {
	d.OriginalAudioProjectExportId = originalAudioProjectExportId
}

func (d *DirectPublishingReadResponseModel) SetOriginalAudioDocumentId(originalAudioDocumentId *string) {
	d.OriginalAudioDocumentId = originalAudioDocumentId
}

func (d *DirectPublishingReadResponseModel) SetSeriesId(seriesId *string) {
	d.SeriesId = seriesId
}

func (d *DirectPublishingReadResponseModel) SetVolume(volume *int) {
	d.Volume = volume
}

func (d *DirectPublishingReadResponseModel) SetPublishedAtUnix(publishedAtUnix *int) {
	d.PublishedAtUnix = publishedAtUnix
}

func (d *DirectPublishingReadResponseModel) SetReadSlug(readSlug *string) {
	d.ReadSlug = readSlug
}

func (d *DirectPublishingReadResponseModel) SetPreviewAudioObject(previewAudioObject *PreviewAudioDbModel) {
	d.PreviewAudioObject = previewAudioObject
}

func (d *DirectPublishingReadResponseModel) SetSampleConfig(sampleConfig *SampleConfigDbModel) {
	d.SampleConfig = sampleConfig
}

func (d *DirectPublishingReadResponseModel) SetReview(review *ReviewResponseModel) {
	d.Review = review
}

func (d *DirectPublishingReadResponseModel) SetVoiceId(voiceId *string) {
	d.VoiceId = voiceId
}

func (d *DirectPublishingReadResponseModel) SetCanUseAssistant(canUseAssistant *bool) {
	d.CanUseAssistant = canUseAssistant
}

func (d *DirectPublishingReadResponseModel) SetIsVoiceChangerOn(isVoiceChangerOn *bool) {
	d.IsVoiceChangerOn = isVoiceChangerOn
}

func (d *DirectPublishingReadResponseModel) SetRestrictedToUserEmailDomains(restrictedToUserEmailDomains []string) {
	d.RestrictedToUserEmailDomains = restrictedToUserEmailDomains
}

func (d *DirectPublishingReadResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DirectPublishingReadResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DirectPublishingReadResponseModel to string"
	}
	return string(jsonData)
}

func (d *DirectPublishingReadResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DirectPublishingReadResponseModel) MarshalJSON() ([]byte, error) {
	type embed DirectPublishingReadResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
