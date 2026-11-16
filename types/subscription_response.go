package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubscriptionResponse struct {
	// The tier of the user's subscription.
	Tier string `json:"tier" xml:"tier" required:"true"`
	// The number of characters used by the user.
	CharacterCount int `json:"character_count" xml:"character_count" required:"true"`
	// The maximum number of characters allowed in the current billing period.
	CharacterLimit int `json:"character_limit" xml:"character_limit" required:"true"`
	// Deprecated: use `max_credit_limit_extension`. Maximum number of characters that the character limit can be exceeded by. Managed by the workspace admin.
	MaxCharacterLimitExtension *int `json:"max_character_limit_extension" xml:"max_character_limit_extension" required:"true"`
	// Maximum number of credits that the credit limit can be exceeded by. Managed by the workspace admin. `"unlimited"` means no cap, `0` means usage-based billing is disabled.
	MaxCreditLimitExtension *SubscriptionResponseMaxCreditLimitExtension `json:"max_credit_limit_extension" xml:"max_credit_limit_extension" required:"true"`
	// Whether the workspace is entitled to enter overages (usage-based billing).
	CanExtendCharacterLimit bool `json:"can_extend_character_limit" xml:"can_extend_character_limit" required:"true"`
	// Deprecated: use `max_credit_limit_extension != 0`. Whether the user is allowed to extend their character limit.
	AllowedToExtendCharacterLimit bool `json:"allowed_to_extend_character_limit" xml:"allowed_to_extend_character_limit" required:"true"`
	// The Unix timestamp of the next character count reset.
	NextCharacterCountResetUnix *int `json:"next_character_count_reset_unix,omitempty" xml:"next_character_count_reset_unix,omitempty"`
	// The number of voice slots used by the user.
	VoiceSlotsUsed int `json:"voice_slots_used" xml:"voice_slots_used" required:"true"`
	// The number of professional voice slots used. For consolidated billing this is the group-wide count across all workspaces in the group; see professional_voice_slots_used_in_workspace for the current workspace only.
	ProfessionalVoiceSlotsUsed int `json:"professional_voice_slots_used" xml:"professional_voice_slots_used" required:"true"`
	// The number of professional voice slots used in the current workspace. For consolidated billing, professional_voice_slots_used counts across all workspaces in the group, while this counts only the current workspace.
	ProfessionalVoiceSlotsUsedInWorkspace int `json:"professional_voice_slots_used_in_workspace" xml:"professional_voice_slots_used_in_workspace" required:"true"`
	// The maximum number of voice slots allowed for the user.
	VoiceLimit int `json:"voice_limit" xml:"voice_limit" required:"true"`
	// The maximum number of voice add/edits allowed for the user.
	MaxVoiceAddEdits *int `json:"max_voice_add_edits,omitempty" xml:"max_voice_add_edits,omitempty"`
	// The number of voice add/edits used by the user.
	VoiceAddEditCounter int `json:"voice_add_edit_counter" xml:"voice_add_edit_counter" required:"true"`
	// The maximum number of professional voices allowed for the user.
	ProfessionalVoiceLimit int `json:"professional_voice_limit" xml:"professional_voice_limit" required:"true"`
	// Whether the user can extend their voice limit.
	CanExtendVoiceLimit bool `json:"can_extend_voice_limit" xml:"can_extend_voice_limit" required:"true"`
	// Whether the user can use instant voice cloning.
	CanUseInstantVoiceCloning bool `json:"can_use_instant_voice_cloning" xml:"can_use_instant_voice_cloning" required:"true"`
	// Whether the user can use professional voice cloning.
	CanUseProfessionalVoiceCloning bool `json:"can_use_professional_voice_cloning" xml:"can_use_professional_voice_cloning" required:"true"`
	// The currency of the user's subscription.
	Currency *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
	// Currency/amount pair.
	CurrentOverage *Price                 `json:"current_overage" xml:"current_overage" required:"true"`
	Status         SubscriptionStatusType `json:"status" xml:"status" required:"true"`
	// The billing period of the user's subscription.
	BillingPeriod *BillingPeriod `json:"billing_period,omitempty" xml:"billing_period,omitempty"`
	// The character refresh period of the user's subscription.
	CharacterRefreshPeriod *CharacterRefreshPeriod `json:"character_refresh_period,omitempty" xml:"character_refresh_period,omitempty"`
	extraProperties        map[string]interface{}
}

func (s *SubscriptionResponse) GetTier() (value string) {
	if s == nil {
		return
	}
	return s.Tier
}

func (s *SubscriptionResponse) GetCharacterCount() (value int) {
	if s == nil {
		return
	}
	return s.CharacterCount
}

func (s *SubscriptionResponse) GetCharacterLimit() (value int) {
	if s == nil {
		return
	}
	return s.CharacterLimit
}

func (s *SubscriptionResponse) GetMaxCharacterLimitExtension() (value *int) {
	if s == nil {
		return
	}
	return s.MaxCharacterLimitExtension
}

func (s *SubscriptionResponse) GetMaxCreditLimitExtension() (value *SubscriptionResponseMaxCreditLimitExtension) {
	if s == nil {
		return
	}
	return s.MaxCreditLimitExtension
}

func (s *SubscriptionResponse) GetCanExtendCharacterLimit() (value bool) {
	if s == nil {
		return
	}
	return s.CanExtendCharacterLimit
}

func (s *SubscriptionResponse) GetAllowedToExtendCharacterLimit() (value bool) {
	if s == nil {
		return
	}
	return s.AllowedToExtendCharacterLimit
}

func (s *SubscriptionResponse) GetNextCharacterCountResetUnix() (value *int) {
	if s == nil {
		return
	}
	return s.NextCharacterCountResetUnix
}

func (s *SubscriptionResponse) GetVoiceSlotsUsed() (value int) {
	if s == nil {
		return
	}
	return s.VoiceSlotsUsed
}

func (s *SubscriptionResponse) GetProfessionalVoiceSlotsUsed() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceSlotsUsed
}

func (s *SubscriptionResponse) GetProfessionalVoiceSlotsUsedInWorkspace() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceSlotsUsedInWorkspace
}

func (s *SubscriptionResponse) GetVoiceLimit() (value int) {
	if s == nil {
		return
	}
	return s.VoiceLimit
}

func (s *SubscriptionResponse) GetMaxVoiceAddEdits() (value *int) {
	if s == nil {
		return
	}
	return s.MaxVoiceAddEdits
}

func (s *SubscriptionResponse) GetVoiceAddEditCounter() (value int) {
	if s == nil {
		return
	}
	return s.VoiceAddEditCounter
}

func (s *SubscriptionResponse) GetProfessionalVoiceLimit() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceLimit
}

func (s *SubscriptionResponse) GetCanExtendVoiceLimit() (value bool) {
	if s == nil {
		return
	}
	return s.CanExtendVoiceLimit
}

func (s *SubscriptionResponse) GetCanUseInstantVoiceCloning() (value bool) {
	if s == nil {
		return
	}
	return s.CanUseInstantVoiceCloning
}

func (s *SubscriptionResponse) GetCanUseProfessionalVoiceCloning() (value bool) {
	if s == nil {
		return
	}
	return s.CanUseProfessionalVoiceCloning
}

func (s *SubscriptionResponse) GetCurrency() (value *Currency) {
	if s == nil {
		return
	}
	return s.Currency
}

func (s *SubscriptionResponse) GetCurrentOverage() (value *Price) {
	if s == nil {
		return
	}
	return s.CurrentOverage
}

func (s *SubscriptionResponse) GetStatus() (value SubscriptionStatusType) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *SubscriptionResponse) GetBillingPeriod() (value *BillingPeriod) {
	if s == nil {
		return
	}
	return s.BillingPeriod
}

func (s *SubscriptionResponse) GetCharacterRefreshPeriod() (value *CharacterRefreshPeriod) {
	if s == nil {
		return
	}
	return s.CharacterRefreshPeriod
}

func (s *SubscriptionResponse) SetTier(tier string) {
	s.Tier = tier
}

func (s *SubscriptionResponse) SetCharacterCount(characterCount int) {
	s.CharacterCount = characterCount
}

func (s *SubscriptionResponse) SetCharacterLimit(characterLimit int) {
	s.CharacterLimit = characterLimit
}

func (s *SubscriptionResponse) SetMaxCharacterLimitExtension(maxCharacterLimitExtension *int) {
	s.MaxCharacterLimitExtension = maxCharacterLimitExtension
}

func (s *SubscriptionResponse) SetMaxCreditLimitExtension(maxCreditLimitExtension *SubscriptionResponseMaxCreditLimitExtension) {
	s.MaxCreditLimitExtension = maxCreditLimitExtension
}

func (s *SubscriptionResponse) SetCanExtendCharacterLimit(canExtendCharacterLimit bool) {
	s.CanExtendCharacterLimit = canExtendCharacterLimit
}

func (s *SubscriptionResponse) SetAllowedToExtendCharacterLimit(allowedToExtendCharacterLimit bool) {
	s.AllowedToExtendCharacterLimit = allowedToExtendCharacterLimit
}

func (s *SubscriptionResponse) SetNextCharacterCountResetUnix(nextCharacterCountResetUnix *int) {
	s.NextCharacterCountResetUnix = nextCharacterCountResetUnix
}

func (s *SubscriptionResponse) SetVoiceSlotsUsed(voiceSlotsUsed int) {
	s.VoiceSlotsUsed = voiceSlotsUsed
}

func (s *SubscriptionResponse) SetProfessionalVoiceSlotsUsed(professionalVoiceSlotsUsed int) {
	s.ProfessionalVoiceSlotsUsed = professionalVoiceSlotsUsed
}

func (s *SubscriptionResponse) SetProfessionalVoiceSlotsUsedInWorkspace(professionalVoiceSlotsUsedInWorkspace int) {
	s.ProfessionalVoiceSlotsUsedInWorkspace = professionalVoiceSlotsUsedInWorkspace
}

func (s *SubscriptionResponse) SetVoiceLimit(voiceLimit int) {
	s.VoiceLimit = voiceLimit
}

func (s *SubscriptionResponse) SetMaxVoiceAddEdits(maxVoiceAddEdits *int) {
	s.MaxVoiceAddEdits = maxVoiceAddEdits
}

func (s *SubscriptionResponse) SetVoiceAddEditCounter(voiceAddEditCounter int) {
	s.VoiceAddEditCounter = voiceAddEditCounter
}

func (s *SubscriptionResponse) SetProfessionalVoiceLimit(professionalVoiceLimit int) {
	s.ProfessionalVoiceLimit = professionalVoiceLimit
}

func (s *SubscriptionResponse) SetCanExtendVoiceLimit(canExtendVoiceLimit bool) {
	s.CanExtendVoiceLimit = canExtendVoiceLimit
}

func (s *SubscriptionResponse) SetCanUseInstantVoiceCloning(canUseInstantVoiceCloning bool) {
	s.CanUseInstantVoiceCloning = canUseInstantVoiceCloning
}

func (s *SubscriptionResponse) SetCanUseProfessionalVoiceCloning(canUseProfessionalVoiceCloning bool) {
	s.CanUseProfessionalVoiceCloning = canUseProfessionalVoiceCloning
}

func (s *SubscriptionResponse) SetCurrency(currency *Currency) {
	s.Currency = currency
}

func (s *SubscriptionResponse) SetCurrentOverage(currentOverage *Price) {
	s.CurrentOverage = currentOverage
}

func (s *SubscriptionResponse) SetStatus(status SubscriptionStatusType) {
	s.Status = status
}

func (s *SubscriptionResponse) SetBillingPeriod(billingPeriod *BillingPeriod) {
	s.BillingPeriod = billingPeriod
}

func (s *SubscriptionResponse) SetCharacterRefreshPeriod(characterRefreshPeriod *CharacterRefreshPeriod) {
	s.CharacterRefreshPeriod = characterRefreshPeriod
}

func (s *SubscriptionResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubscriptionResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubscriptionResponse to string"
	}
	return string(jsonData)
}

func (s *SubscriptionResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubscriptionResponse) MarshalJSON() ([]byte, error) {
	type embed SubscriptionResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
