package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LibraryVoiceResponse struct {
	// The public owner id of the voice.
	PublicOwnerId string `json:"public_owner_id" xml:"public_owner_id" required:"true"`
	// The id of the voice.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// The date the voice was added to the library in Unix time.
	DateUnix int `json:"date_unix" xml:"date_unix" required:"true"`
	// The name of the voice.
	Name string `json:"name" xml:"name" required:"true"`
	// The accent of the voice.
	Accent string `json:"accent" xml:"accent" required:"true"`
	// The gender of the voice.
	Gender string `json:"gender" xml:"gender" required:"true"`
	// The age of the voice.
	Age string `json:"age" xml:"age" required:"true"`
	// The descriptive of the voice.
	Descriptive string `json:"descriptive" xml:"descriptive" required:"true"`
	// The use case of the voice.
	UseCase string `json:"use_case" xml:"use_case" required:"true"`
	// The category of the voice.
	Category LibraryVoiceResponseCategory `json:"category" xml:"category" required:"true"`
	// The language of the voice.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The locale of the voice.
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The description of the voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The preview URL of the voice.
	PreviewUrl *string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	// The usage character count of the voice in the last year.
	UsageCharacterCount1Y int `json:"usage_character_count_1y" xml:"usage_character_count_1y" required:"true"`
	// The usage character count of the voice in the last 7 days.
	UsageCharacterCount7D int `json:"usage_character_count_7d" xml:"usage_character_count_7d" required:"true"`
	// The play API usage character count of the voice in the last year.
	PlayApiUsageCharacterCount1Y int `json:"play_api_usage_character_count_1y" xml:"play_api_usage_character_count_1y" required:"true"`
	// The number of times the voice has been cloned.
	ClonedByCount int `json:"cloned_by_count" xml:"cloned_by_count" required:"true"`
	// The rate multiplier of the voice.
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The rate of the voice in USD per 1000 credits. null if default
	FiatRate *float64 `json:"fiat_rate,omitempty" xml:"fiat_rate,omitempty"`
	// Whether free users are allowed to use the voice.
	FreeUsersAllowed bool `json:"free_users_allowed" xml:"free_users_allowed" required:"true"`
	// Whether live moderation is enabled for the voice.
	LiveModerationEnabled bool `json:"live_moderation_enabled" xml:"live_moderation_enabled" required:"true"`
	// Whether the voice is featured.
	Featured bool `json:"featured" xml:"featured" required:"true"`
	// The verified languages of the voice.
	VerifiedLanguages []*VerifiedVoiceLanguageResponseModel `json:"verified_languages,omitempty" xml:"verified_languages,omitempty"`
	// The notice period of the voice.
	NoticePeriod *int `json:"notice_period,omitempty" xml:"notice_period,omitempty"`
	// The Instagram username of the voice.
	InstagramUsername *string `json:"instagram_username,omitempty" xml:"instagram_username,omitempty"`
	// The Twitter username of the voice.
	TwitterUsername *string `json:"twitter_username,omitempty" xml:"twitter_username,omitempty"`
	// The YouTube username of the voice.
	YoutubeUsername *string `json:"youtube_username,omitempty" xml:"youtube_username,omitempty"`
	// The TikTok username of the voice.
	TiktokUsername *string `json:"tiktok_username,omitempty" xml:"tiktok_username,omitempty"`
	// The image URL of the voice.
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// Whether the voice was added by the user.
	IsAddedByUser *bool `json:"is_added_by_user,omitempty" xml:"is_added_by_user,omitempty"`
	// Whether the voice is bookmarked by the current user. Only relevant when is_added_by_user is True.
	IsBookmarked    *bool `json:"is_bookmarked,omitempty" xml:"is_bookmarked,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LibraryVoiceResponse) GetPublicOwnerId() (value string) {
	if l == nil {
		return
	}
	return l.PublicOwnerId
}

func (l *LibraryVoiceResponse) GetVoiceId() (value string) {
	if l == nil {
		return
	}
	return l.VoiceId
}

func (l *LibraryVoiceResponse) GetDateUnix() (value int) {
	if l == nil {
		return
	}
	return l.DateUnix
}

func (l *LibraryVoiceResponse) GetName() (value string) {
	if l == nil {
		return
	}
	return l.Name
}

func (l *LibraryVoiceResponse) GetAccent() (value string) {
	if l == nil {
		return
	}
	return l.Accent
}

func (l *LibraryVoiceResponse) GetGender() (value string) {
	if l == nil {
		return
	}
	return l.Gender
}

func (l *LibraryVoiceResponse) GetAge() (value string) {
	if l == nil {
		return
	}
	return l.Age
}

func (l *LibraryVoiceResponse) GetDescriptive() (value string) {
	if l == nil {
		return
	}
	return l.Descriptive
}

func (l *LibraryVoiceResponse) GetUseCase() (value string) {
	if l == nil {
		return
	}
	return l.UseCase
}

func (l *LibraryVoiceResponse) GetCategory() (value LibraryVoiceResponseCategory) {
	if l == nil {
		return
	}
	return l.Category
}

func (l *LibraryVoiceResponse) GetLanguage() (value *string) {
	if l == nil {
		return
	}
	return l.Language
}

func (l *LibraryVoiceResponse) GetLocale() (value *string) {
	if l == nil {
		return
	}
	return l.Locale
}

func (l *LibraryVoiceResponse) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LibraryVoiceResponse) GetPreviewUrl() (value *string) {
	if l == nil {
		return
	}
	return l.PreviewUrl
}

func (l *LibraryVoiceResponse) GetUsageCharacterCount1Y() (value int) {
	if l == nil {
		return
	}
	return l.UsageCharacterCount1Y
}

func (l *LibraryVoiceResponse) GetUsageCharacterCount7D() (value int) {
	if l == nil {
		return
	}
	return l.UsageCharacterCount7D
}

func (l *LibraryVoiceResponse) GetPlayApiUsageCharacterCount1Y() (value int) {
	if l == nil {
		return
	}
	return l.PlayApiUsageCharacterCount1Y
}

func (l *LibraryVoiceResponse) GetClonedByCount() (value int) {
	if l == nil {
		return
	}
	return l.ClonedByCount
}

func (l *LibraryVoiceResponse) GetRate() (value *float64) {
	if l == nil {
		return
	}
	return l.Rate
}

func (l *LibraryVoiceResponse) GetFiatRate() (value *float64) {
	if l == nil {
		return
	}
	return l.FiatRate
}

func (l *LibraryVoiceResponse) GetFreeUsersAllowed() (value bool) {
	if l == nil {
		return
	}
	return l.FreeUsersAllowed
}

func (l *LibraryVoiceResponse) GetLiveModerationEnabled() (value bool) {
	if l == nil {
		return
	}
	return l.LiveModerationEnabled
}

func (l *LibraryVoiceResponse) GetFeatured() (value bool) {
	if l == nil {
		return
	}
	return l.Featured
}

func (l *LibraryVoiceResponse) GetVerifiedLanguages() (value []*VerifiedVoiceLanguageResponseModel) {
	if l == nil {
		return
	}
	return l.VerifiedLanguages
}

func (l *LibraryVoiceResponse) GetNoticePeriod() (value *int) {
	if l == nil {
		return
	}
	return l.NoticePeriod
}

func (l *LibraryVoiceResponse) GetInstagramUsername() (value *string) {
	if l == nil {
		return
	}
	return l.InstagramUsername
}

func (l *LibraryVoiceResponse) GetTwitterUsername() (value *string) {
	if l == nil {
		return
	}
	return l.TwitterUsername
}

func (l *LibraryVoiceResponse) GetYoutubeUsername() (value *string) {
	if l == nil {
		return
	}
	return l.YoutubeUsername
}

func (l *LibraryVoiceResponse) GetTiktokUsername() (value *string) {
	if l == nil {
		return
	}
	return l.TiktokUsername
}

func (l *LibraryVoiceResponse) GetImageUrl() (value *string) {
	if l == nil {
		return
	}
	return l.ImageUrl
}

func (l *LibraryVoiceResponse) GetIsAddedByUser() (value *bool) {
	if l == nil {
		return
	}
	return l.IsAddedByUser
}

func (l *LibraryVoiceResponse) GetIsBookmarked() (value *bool) {
	if l == nil {
		return
	}
	return l.IsBookmarked
}

func (l *LibraryVoiceResponse) SetPublicOwnerId(publicOwnerId string) {
	l.PublicOwnerId = publicOwnerId
}

func (l *LibraryVoiceResponse) SetVoiceId(voiceId string) {
	l.VoiceId = voiceId
}

func (l *LibraryVoiceResponse) SetDateUnix(dateUnix int) {
	l.DateUnix = dateUnix
}

func (l *LibraryVoiceResponse) SetName(name string) {
	l.Name = name
}

func (l *LibraryVoiceResponse) SetAccent(accent string) {
	l.Accent = accent
}

func (l *LibraryVoiceResponse) SetGender(gender string) {
	l.Gender = gender
}

func (l *LibraryVoiceResponse) SetAge(age string) {
	l.Age = age
}

func (l *LibraryVoiceResponse) SetDescriptive(descriptive string) {
	l.Descriptive = descriptive
}

func (l *LibraryVoiceResponse) SetUseCase(useCase string) {
	l.UseCase = useCase
}

func (l *LibraryVoiceResponse) SetCategory(category LibraryVoiceResponseCategory) {
	l.Category = category
}

func (l *LibraryVoiceResponse) SetLanguage(language *string) {
	l.Language = language
}

func (l *LibraryVoiceResponse) SetLocale(locale *string) {
	l.Locale = locale
}

func (l *LibraryVoiceResponse) SetDescription(description *string) {
	l.Description = description
}

func (l *LibraryVoiceResponse) SetPreviewUrl(previewUrl *string) {
	l.PreviewUrl = previewUrl
}

func (l *LibraryVoiceResponse) SetUsageCharacterCount1Y(usageCharacterCount1Y int) {
	l.UsageCharacterCount1Y = usageCharacterCount1Y
}

func (l *LibraryVoiceResponse) SetUsageCharacterCount7D(usageCharacterCount7D int) {
	l.UsageCharacterCount7D = usageCharacterCount7D
}

func (l *LibraryVoiceResponse) SetPlayApiUsageCharacterCount1Y(playApiUsageCharacterCount1Y int) {
	l.PlayApiUsageCharacterCount1Y = playApiUsageCharacterCount1Y
}

func (l *LibraryVoiceResponse) SetClonedByCount(clonedByCount int) {
	l.ClonedByCount = clonedByCount
}

func (l *LibraryVoiceResponse) SetRate(rate *float64) {
	l.Rate = rate
}

func (l *LibraryVoiceResponse) SetFiatRate(fiatRate *float64) {
	l.FiatRate = fiatRate
}

func (l *LibraryVoiceResponse) SetFreeUsersAllowed(freeUsersAllowed bool) {
	l.FreeUsersAllowed = freeUsersAllowed
}

func (l *LibraryVoiceResponse) SetLiveModerationEnabled(liveModerationEnabled bool) {
	l.LiveModerationEnabled = liveModerationEnabled
}

func (l *LibraryVoiceResponse) SetFeatured(featured bool) {
	l.Featured = featured
}

func (l *LibraryVoiceResponse) SetVerifiedLanguages(verifiedLanguages []*VerifiedVoiceLanguageResponseModel) {
	l.VerifiedLanguages = verifiedLanguages
}

func (l *LibraryVoiceResponse) SetNoticePeriod(noticePeriod *int) {
	l.NoticePeriod = noticePeriod
}

func (l *LibraryVoiceResponse) SetInstagramUsername(instagramUsername *string) {
	l.InstagramUsername = instagramUsername
}

func (l *LibraryVoiceResponse) SetTwitterUsername(twitterUsername *string) {
	l.TwitterUsername = twitterUsername
}

func (l *LibraryVoiceResponse) SetYoutubeUsername(youtubeUsername *string) {
	l.YoutubeUsername = youtubeUsername
}

func (l *LibraryVoiceResponse) SetTiktokUsername(tiktokUsername *string) {
	l.TiktokUsername = tiktokUsername
}

func (l *LibraryVoiceResponse) SetImageUrl(imageUrl *string) {
	l.ImageUrl = imageUrl
}

func (l *LibraryVoiceResponse) SetIsAddedByUser(isAddedByUser *bool) {
	l.IsAddedByUser = isAddedByUser
}

func (l *LibraryVoiceResponse) SetIsBookmarked(isBookmarked *bool) {
	l.IsBookmarked = isBookmarked
}

func (l *LibraryVoiceResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LibraryVoiceResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LibraryVoiceResponse to string"
	}
	return string(jsonData)
}

func (l *LibraryVoiceResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LibraryVoiceResponse) MarshalJSON() ([]byte, error) {
	type embed LibraryVoiceResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
