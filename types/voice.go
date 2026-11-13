package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Voice struct {
	// The ID of the voice.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// The name of the voice.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// List of samples associated with the voice.
	Samples []*VoiceSample `json:"samples,omitempty" xml:"samples,omitempty"`
	// The category of the voice.
	Category *Category `json:"category,omitempty" xml:"category,omitempty"`
	// Fine-tuning information for the voice.
	FineTuning *FineTuningResponse `json:"fine_tuning,omitempty" xml:"fine_tuning,omitempty"`
	// Labels associated with the voice.
	Labels map[string]string `json:"labels,omitempty"`
	// The description of the voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The preview URL of the voice.
	PreviewUrl *string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	// The tiers the voice is available for.
	AvailableForTiers []string `json:"available_for_tiers,omitempty" xml:"available_for_tiers,omitempty"`
	// The settings of the voice.
	Settings VoiceSettings `json:"settings,omitempty" xml:"settings,omitempty"`
	// The sharing information of the voice.
	Sharing *VoiceSharingResponse `json:"sharing,omitempty" xml:"sharing,omitempty"`
	// The base model IDs for high-quality voices.
	HighQualityBaseModelIds []string `json:"high_quality_base_model_ids,omitempty" xml:"high_quality_base_model_ids,omitempty"`
	// The verified languages of the voice.
	VerifiedLanguages []*VerifiedVoiceLanguageResponseModel `json:"verified_languages,omitempty" xml:"verified_languages,omitempty"`
	// The IDs of collections this voice belongs to.
	CollectionIds []string `json:"collection_ids,omitempty" xml:"collection_ids,omitempty"`
	// The safety controls of the voice.
	SafetyControl *VoiceSafetyControl `json:"safety_control,omitempty" xml:"safety_control,omitempty"`
	// The voice verification of the voice.
	VoiceVerification *VoiceVerificationResponse `json:"voice_verification,omitempty" xml:"voice_verification,omitempty"`
	// The permission on the resource of the voice.
	PermissionOnResource *string `json:"permission_on_resource,omitempty" xml:"permission_on_resource,omitempty"`
	// Whether the voice is owned by the user.
	IsOwner *bool `json:"is_owner,omitempty" xml:"is_owner,omitempty"`
	// Whether the voice is legacy.
	IsLegacy *bool `json:"is_legacy,omitempty" xml:"is_legacy,omitempty"`
	// Whether the voice is mixed.
	IsMixed *bool `json:"is_mixed,omitempty" xml:"is_mixed,omitempty"`
	// Timestamp when the voice was marked as favorite in Unix time.
	FavoritedAtUnix *int `json:"favorited_at_unix,omitempty" xml:"favorited_at_unix,omitempty"`
	// The creation time of the voice in Unix time.
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Whether the voice is bookmarked by the current user. Only relevant for community (library-copied) voices.
	IsBookmarked *bool `json:"is_bookmarked,omitempty" xml:"is_bookmarked,omitempty"`
	// The recording quality of the voice as determined by the review pipeline.
	RecordingQuality *VoiceRecordingQuality `json:"recording_quality,omitempty" xml:"recording_quality,omitempty"`
	// The review pipeline status of the voice.
	LabellingStatus *VoiceLabellingStatus `json:"labelling_status,omitempty" xml:"labelling_status,omitempty"`
	// The reason for the recording quality assessment, as determined by the review pipeline.
	RecordingQualityReason *string `json:"recording_quality_reason,omitempty" xml:"recording_quality_reason,omitempty"`
	extraProperties        map[string]interface{}
}

func (v *Voice) GetVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *Voice) GetName() (value *string) {
	if v == nil {
		return
	}
	return v.Name
}

func (v *Voice) GetSamples() (value []*VoiceSample) {
	if v == nil {
		return
	}
	return v.Samples
}

func (v *Voice) GetCategory() (value *Category) {
	if v == nil {
		return
	}
	return v.Category
}

func (v *Voice) GetFineTuning() (value *FineTuningResponse) {
	if v == nil {
		return
	}
	return v.FineTuning
}

func (v *Voice) GetLabels() (value map[string]string) {
	if v == nil {
		return
	}
	return v.Labels
}

func (v *Voice) GetDescription() (value *string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *Voice) GetPreviewUrl() (value *string) {
	if v == nil {
		return
	}
	return v.PreviewUrl
}

func (v *Voice) GetAvailableForTiers() (value []string) {
	if v == nil {
		return
	}
	return v.AvailableForTiers
}

func (v *Voice) GetSettings() (value VoiceSettings) {
	if v == nil {
		return
	}
	return v.Settings
}

func (v *Voice) GetSharing() (value *VoiceSharingResponse) {
	if v == nil {
		return
	}
	return v.Sharing
}

func (v *Voice) GetHighQualityBaseModelIds() (value []string) {
	if v == nil {
		return
	}
	return v.HighQualityBaseModelIds
}

func (v *Voice) GetVerifiedLanguages() (value []*VerifiedVoiceLanguageResponseModel) {
	if v == nil {
		return
	}
	return v.VerifiedLanguages
}

func (v *Voice) GetCollectionIds() (value []string) {
	if v == nil {
		return
	}
	return v.CollectionIds
}

func (v *Voice) GetSafetyControl() (value *VoiceSafetyControl) {
	if v == nil {
		return
	}
	return v.SafetyControl
}

func (v *Voice) GetVoiceVerification() (value *VoiceVerificationResponse) {
	if v == nil {
		return
	}
	return v.VoiceVerification
}

func (v *Voice) GetPermissionOnResource() (value *string) {
	if v == nil {
		return
	}
	return v.PermissionOnResource
}

func (v *Voice) GetIsOwner() (value *bool) {
	if v == nil {
		return
	}
	return v.IsOwner
}

func (v *Voice) GetIsLegacy() (value *bool) {
	if v == nil {
		return
	}
	return v.IsLegacy
}

func (v *Voice) GetIsMixed() (value *bool) {
	if v == nil {
		return
	}
	return v.IsMixed
}

func (v *Voice) GetFavoritedAtUnix() (value *int) {
	if v == nil {
		return
	}
	return v.FavoritedAtUnix
}

func (v *Voice) GetCreatedAtUnix() (value *int) {
	if v == nil {
		return
	}
	return v.CreatedAtUnix
}

func (v *Voice) GetIsBookmarked() (value *bool) {
	if v == nil {
		return
	}
	return v.IsBookmarked
}

func (v *Voice) GetRecordingQuality() (value *VoiceRecordingQuality) {
	if v == nil {
		return
	}
	return v.RecordingQuality
}

func (v *Voice) GetLabellingStatus() (value *VoiceLabellingStatus) {
	if v == nil {
		return
	}
	return v.LabellingStatus
}

func (v *Voice) GetRecordingQualityReason() (value *string) {
	if v == nil {
		return
	}
	return v.RecordingQualityReason
}

func (v *Voice) SetVoiceId(voiceId string) {
	v.VoiceId = voiceId
}

func (v *Voice) SetName(name *string) {
	v.Name = name
}

func (v *Voice) SetSamples(samples []*VoiceSample) {
	v.Samples = samples
}

func (v *Voice) SetCategory(category *Category) {
	v.Category = category
}

func (v *Voice) SetFineTuning(fineTuning *FineTuningResponse) {
	v.FineTuning = fineTuning
}

func (v *Voice) SetLabels(labels map[string]string) {
	v.Labels = labels
}

func (v *Voice) SetDescription(description *string) {
	v.Description = description
}

func (v *Voice) SetPreviewUrl(previewUrl *string) {
	v.PreviewUrl = previewUrl
}

func (v *Voice) SetAvailableForTiers(availableForTiers []string) {
	v.AvailableForTiers = availableForTiers
}

func (v *Voice) SetSettings(settings VoiceSettings) {
	v.Settings = settings
}

func (v *Voice) SetSharing(sharing *VoiceSharingResponse) {
	v.Sharing = sharing
}

func (v *Voice) SetHighQualityBaseModelIds(highQualityBaseModelIds []string) {
	v.HighQualityBaseModelIds = highQualityBaseModelIds
}

func (v *Voice) SetVerifiedLanguages(verifiedLanguages []*VerifiedVoiceLanguageResponseModel) {
	v.VerifiedLanguages = verifiedLanguages
}

func (v *Voice) SetCollectionIds(collectionIds []string) {
	v.CollectionIds = collectionIds
}

func (v *Voice) SetSafetyControl(safetyControl *VoiceSafetyControl) {
	v.SafetyControl = safetyControl
}

func (v *Voice) SetVoiceVerification(voiceVerification *VoiceVerificationResponse) {
	v.VoiceVerification = voiceVerification
}

func (v *Voice) SetPermissionOnResource(permissionOnResource *string) {
	v.PermissionOnResource = permissionOnResource
}

func (v *Voice) SetIsOwner(isOwner *bool) {
	v.IsOwner = isOwner
}

func (v *Voice) SetIsLegacy(isLegacy *bool) {
	v.IsLegacy = isLegacy
}

func (v *Voice) SetIsMixed(isMixed *bool) {
	v.IsMixed = isMixed
}

func (v *Voice) SetFavoritedAtUnix(favoritedAtUnix *int) {
	v.FavoritedAtUnix = favoritedAtUnix
}

func (v *Voice) SetCreatedAtUnix(createdAtUnix *int) {
	v.CreatedAtUnix = createdAtUnix
}

func (v *Voice) SetIsBookmarked(isBookmarked *bool) {
	v.IsBookmarked = isBookmarked
}

func (v *Voice) SetRecordingQuality(recordingQuality *VoiceRecordingQuality) {
	v.RecordingQuality = recordingQuality
}

func (v *Voice) SetLabellingStatus(labellingStatus *VoiceLabellingStatus) {
	v.LabellingStatus = labellingStatus
}

func (v *Voice) SetRecordingQualityReason(recordingQualityReason *string) {
	v.RecordingQualityReason = recordingQualityReason
}

func (v *Voice) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v Voice) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: Voice to string"
	}
	return string(jsonData)
}

func (v *Voice) UnmarshalJSON(data []byte) error {
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

func (v *Voice) MarshalJSON() ([]byte, error) {
	type embed Voice
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
