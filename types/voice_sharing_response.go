package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSharingResponse struct {
	// The status of the voice sharing.
	Status *VoiceSharingState `json:"status,omitempty" xml:"status,omitempty"`
	// The sample ID of the history item.
	HistoryItemSampleId *string `json:"history_item_sample_id,omitempty" xml:"history_item_sample_id,omitempty"`
	// The Unix time, in seconds, at which the voice was shared.
	DateUnix *int `json:"date_unix,omitempty" xml:"date_unix,omitempty"`
	// A list of whitelisted emails.
	WhitelistedEmails []string `json:"whitelisted_emails,omitempty" xml:"whitelisted_emails,omitempty"`
	// The ID of the public owner.
	PublicOwnerId *string `json:"public_owner_id,omitempty" xml:"public_owner_id,omitempty"`
	// The ID of the original voice.
	OriginalVoiceId *string `json:"original_voice_id,omitempty" xml:"original_voice_id,omitempty"`
	// Whether financial rewards are enabled.
	FinancialRewardsEnabled *bool `json:"financial_rewards_enabled,omitempty" xml:"financial_rewards_enabled,omitempty"`
	// Whether free users are allowed.
	FreeUsersAllowed *bool `json:"free_users_allowed,omitempty" xml:"free_users_allowed,omitempty"`
	// Whether live moderation is enabled.
	LiveModerationEnabled *bool `json:"live_moderation_enabled,omitempty" xml:"live_moderation_enabled,omitempty"`
	// The rate of the voice sharing.
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The rate of the voice sharing in USD per 1000 credits.
	FiatRate *float64 `json:"fiat_rate,omitempty" xml:"fiat_rate,omitempty"`
	// The notice period of the voice sharing.
	NoticePeriod *int `json:"notice_period,omitempty" xml:"notice_period,omitempty"`
	// The Unix time, in seconds, at which users who added the voice from the Voice Library lose access to it.
	DisableAtUnix *int `json:"disable_at_unix,omitempty" xml:"disable_at_unix,omitempty"`
	// Whether voice mixing is allowed.
	VoiceMixingAllowed *bool `json:"voice_mixing_allowed,omitempty" xml:"voice_mixing_allowed,omitempty"`
	// Whether the voice is featured.
	Featured *bool `json:"featured,omitempty" xml:"featured,omitempty"`
	// The category of the voice.
	Category *VoiceSharingResponseCategory `json:"category,omitempty" xml:"category,omitempty"`
	// Whether the reader app is enabled.
	ReaderAppEnabled *bool `json:"reader_app_enabled,omitempty" xml:"reader_app_enabled,omitempty"`
	// The image URL of the voice.
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// The ban reason of the voice.
	BanReason *string `json:"ban_reason,omitempty" xml:"ban_reason,omitempty"`
	// The number of likes on the voice.
	LikedByCount *int `json:"liked_by_count,omitempty" xml:"liked_by_count,omitempty"`
	// The number of clones on the voice.
	ClonedByCount *int `json:"cloned_by_count,omitempty" xml:"cloned_by_count,omitempty"`
	// The name of the voice.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The description of the voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The labels of the voice.
	Labels map[string]string `json:"labels,omitempty"`
	// The review status of the voice.
	ReviewStatus *ReviewStatus `json:"review_status,omitempty" xml:"review_status,omitempty"`
	// The review message of the voice.
	ReviewMessage *string `json:"review_message,omitempty" xml:"review_message,omitempty"`
	// Whether the voice is enabled in the library.
	EnabledInLibrary *bool `json:"enabled_in_library,omitempty" xml:"enabled_in_library,omitempty"`
	// The Instagram username of the voice.
	InstagramUsername *string `json:"instagram_username,omitempty" xml:"instagram_username,omitempty"`
	// The Twitter/X username of the voice.
	TwitterUsername *string `json:"twitter_username,omitempty" xml:"twitter_username,omitempty"`
	// The YouTube username of the voice.
	YoutubeUsername *string `json:"youtube_username,omitempty" xml:"youtube_username,omitempty"`
	// The TikTok username of the voice.
	TiktokUsername *string `json:"tiktok_username,omitempty" xml:"tiktok_username,omitempty"`
	// The moderation check of the voice.
	ModerationCheck *VoiceSharingModerationCheckResponseModel `json:"moderation_check,omitempty" xml:"moderation_check,omitempty"`
	// The reader restricted on of the voice.
	ReaderRestrictedOn []*ReaderResourceResponseModel `json:"reader_restricted_on,omitempty" xml:"reader_restricted_on,omitempty"`
	extraProperties    map[string]interface{}
}

func (v *VoiceSharingResponse) GetStatus() (value *VoiceSharingState) {
	if v == nil {
		return
	}
	return v.Status
}

func (v *VoiceSharingResponse) GetHistoryItemSampleId() (value *string) {
	if v == nil {
		return
	}
	return v.HistoryItemSampleId
}

func (v *VoiceSharingResponse) GetDateUnix() (value *int) {
	if v == nil {
		return
	}
	return v.DateUnix
}

func (v *VoiceSharingResponse) GetWhitelistedEmails() (value []string) {
	if v == nil {
		return
	}
	return v.WhitelistedEmails
}

func (v *VoiceSharingResponse) GetPublicOwnerId() (value *string) {
	if v == nil {
		return
	}
	return v.PublicOwnerId
}

func (v *VoiceSharingResponse) GetOriginalVoiceId() (value *string) {
	if v == nil {
		return
	}
	return v.OriginalVoiceId
}

func (v *VoiceSharingResponse) GetFinancialRewardsEnabled() (value *bool) {
	if v == nil {
		return
	}
	return v.FinancialRewardsEnabled
}

func (v *VoiceSharingResponse) GetFreeUsersAllowed() (value *bool) {
	if v == nil {
		return
	}
	return v.FreeUsersAllowed
}

func (v *VoiceSharingResponse) GetLiveModerationEnabled() (value *bool) {
	if v == nil {
		return
	}
	return v.LiveModerationEnabled
}

func (v *VoiceSharingResponse) GetRate() (value *float64) {
	if v == nil {
		return
	}
	return v.Rate
}

func (v *VoiceSharingResponse) GetFiatRate() (value *float64) {
	if v == nil {
		return
	}
	return v.FiatRate
}

func (v *VoiceSharingResponse) GetNoticePeriod() (value *int) {
	if v == nil {
		return
	}
	return v.NoticePeriod
}

func (v *VoiceSharingResponse) GetDisableAtUnix() (value *int) {
	if v == nil {
		return
	}
	return v.DisableAtUnix
}

func (v *VoiceSharingResponse) GetVoiceMixingAllowed() (value *bool) {
	if v == nil {
		return
	}
	return v.VoiceMixingAllowed
}

func (v *VoiceSharingResponse) GetFeatured() (value *bool) {
	if v == nil {
		return
	}
	return v.Featured
}

func (v *VoiceSharingResponse) GetCategory() (value *VoiceSharingResponseCategory) {
	if v == nil {
		return
	}
	return v.Category
}

func (v *VoiceSharingResponse) GetReaderAppEnabled() (value *bool) {
	if v == nil {
		return
	}
	return v.ReaderAppEnabled
}

func (v *VoiceSharingResponse) GetImageUrl() (value *string) {
	if v == nil {
		return
	}
	return v.ImageUrl
}

func (v *VoiceSharingResponse) GetBanReason() (value *string) {
	if v == nil {
		return
	}
	return v.BanReason
}

func (v *VoiceSharingResponse) GetLikedByCount() (value *int) {
	if v == nil {
		return
	}
	return v.LikedByCount
}

func (v *VoiceSharingResponse) GetClonedByCount() (value *int) {
	if v == nil {
		return
	}
	return v.ClonedByCount
}

func (v *VoiceSharingResponse) GetName() (value *string) {
	if v == nil {
		return
	}
	return v.Name
}

func (v *VoiceSharingResponse) GetDescription() (value *string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VoiceSharingResponse) GetLabels() (value map[string]string) {
	if v == nil {
		return
	}
	return v.Labels
}

func (v *VoiceSharingResponse) GetReviewStatus() (value *ReviewStatus) {
	if v == nil {
		return
	}
	return v.ReviewStatus
}

func (v *VoiceSharingResponse) GetReviewMessage() (value *string) {
	if v == nil {
		return
	}
	return v.ReviewMessage
}

func (v *VoiceSharingResponse) GetEnabledInLibrary() (value *bool) {
	if v == nil {
		return
	}
	return v.EnabledInLibrary
}

func (v *VoiceSharingResponse) GetInstagramUsername() (value *string) {
	if v == nil {
		return
	}
	return v.InstagramUsername
}

func (v *VoiceSharingResponse) GetTwitterUsername() (value *string) {
	if v == nil {
		return
	}
	return v.TwitterUsername
}

func (v *VoiceSharingResponse) GetYoutubeUsername() (value *string) {
	if v == nil {
		return
	}
	return v.YoutubeUsername
}

func (v *VoiceSharingResponse) GetTiktokUsername() (value *string) {
	if v == nil {
		return
	}
	return v.TiktokUsername
}

func (v *VoiceSharingResponse) GetModerationCheck() (value *VoiceSharingModerationCheckResponseModel) {
	if v == nil {
		return
	}
	return v.ModerationCheck
}

func (v *VoiceSharingResponse) GetReaderRestrictedOn() (value []*ReaderResourceResponseModel) {
	if v == nil {
		return
	}
	return v.ReaderRestrictedOn
}

func (v *VoiceSharingResponse) SetStatus(status *VoiceSharingState) {
	v.Status = status
}

func (v *VoiceSharingResponse) SetHistoryItemSampleId(historyItemSampleId *string) {
	v.HistoryItemSampleId = historyItemSampleId
}

func (v *VoiceSharingResponse) SetDateUnix(dateUnix *int) {
	v.DateUnix = dateUnix
}

func (v *VoiceSharingResponse) SetWhitelistedEmails(whitelistedEmails []string) {
	v.WhitelistedEmails = whitelistedEmails
}

func (v *VoiceSharingResponse) SetPublicOwnerId(publicOwnerId *string) {
	v.PublicOwnerId = publicOwnerId
}

func (v *VoiceSharingResponse) SetOriginalVoiceId(originalVoiceId *string) {
	v.OriginalVoiceId = originalVoiceId
}

func (v *VoiceSharingResponse) SetFinancialRewardsEnabled(financialRewardsEnabled *bool) {
	v.FinancialRewardsEnabled = financialRewardsEnabled
}

func (v *VoiceSharingResponse) SetFreeUsersAllowed(freeUsersAllowed *bool) {
	v.FreeUsersAllowed = freeUsersAllowed
}

func (v *VoiceSharingResponse) SetLiveModerationEnabled(liveModerationEnabled *bool) {
	v.LiveModerationEnabled = liveModerationEnabled
}

func (v *VoiceSharingResponse) SetRate(rate *float64) {
	v.Rate = rate
}

func (v *VoiceSharingResponse) SetFiatRate(fiatRate *float64) {
	v.FiatRate = fiatRate
}

func (v *VoiceSharingResponse) SetNoticePeriod(noticePeriod *int) {
	v.NoticePeriod = noticePeriod
}

func (v *VoiceSharingResponse) SetDisableAtUnix(disableAtUnix *int) {
	v.DisableAtUnix = disableAtUnix
}

func (v *VoiceSharingResponse) SetVoiceMixingAllowed(voiceMixingAllowed *bool) {
	v.VoiceMixingAllowed = voiceMixingAllowed
}

func (v *VoiceSharingResponse) SetFeatured(featured *bool) {
	v.Featured = featured
}

func (v *VoiceSharingResponse) SetCategory(category *VoiceSharingResponseCategory) {
	v.Category = category
}

func (v *VoiceSharingResponse) SetReaderAppEnabled(readerAppEnabled *bool) {
	v.ReaderAppEnabled = readerAppEnabled
}

func (v *VoiceSharingResponse) SetImageUrl(imageUrl *string) {
	v.ImageUrl = imageUrl
}

func (v *VoiceSharingResponse) SetBanReason(banReason *string) {
	v.BanReason = banReason
}

func (v *VoiceSharingResponse) SetLikedByCount(likedByCount *int) {
	v.LikedByCount = likedByCount
}

func (v *VoiceSharingResponse) SetClonedByCount(clonedByCount *int) {
	v.ClonedByCount = clonedByCount
}

func (v *VoiceSharingResponse) SetName(name *string) {
	v.Name = name
}

func (v *VoiceSharingResponse) SetDescription(description *string) {
	v.Description = description
}

func (v *VoiceSharingResponse) SetLabels(labels map[string]string) {
	v.Labels = labels
}

func (v *VoiceSharingResponse) SetReviewStatus(reviewStatus *ReviewStatus) {
	v.ReviewStatus = reviewStatus
}

func (v *VoiceSharingResponse) SetReviewMessage(reviewMessage *string) {
	v.ReviewMessage = reviewMessage
}

func (v *VoiceSharingResponse) SetEnabledInLibrary(enabledInLibrary *bool) {
	v.EnabledInLibrary = enabledInLibrary
}

func (v *VoiceSharingResponse) SetInstagramUsername(instagramUsername *string) {
	v.InstagramUsername = instagramUsername
}

func (v *VoiceSharingResponse) SetTwitterUsername(twitterUsername *string) {
	v.TwitterUsername = twitterUsername
}

func (v *VoiceSharingResponse) SetYoutubeUsername(youtubeUsername *string) {
	v.YoutubeUsername = youtubeUsername
}

func (v *VoiceSharingResponse) SetTiktokUsername(tiktokUsername *string) {
	v.TiktokUsername = tiktokUsername
}

func (v *VoiceSharingResponse) SetModerationCheck(moderationCheck *VoiceSharingModerationCheckResponseModel) {
	v.ModerationCheck = moderationCheck
}

func (v *VoiceSharingResponse) SetReaderRestrictedOn(readerRestrictedOn []*ReaderResourceResponseModel) {
	v.ReaderRestrictedOn = readerRestrictedOn
}

func (v *VoiceSharingResponse) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSharingResponse) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSharingResponse to string"
	}
	return string(jsonData)
}

func (v *VoiceSharingResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceSharingResponse) MarshalJSON() ([]byte, error) {
	type embed VoiceSharingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
