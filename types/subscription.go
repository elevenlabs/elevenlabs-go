package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Subscription struct {
	// The tier of the user's subscription.
	Tier string `json:"tier" xml:"tier" required:"true"`
	// The number of characters used by the user.
	CharacterCount int `json:"character_count" xml:"character_count" required:"true"`
	// The maximum number of characters allowed in the current billing period.
	CharacterLimit int `json:"character_limit" xml:"character_limit" required:"true"`
	// Deprecated: use `max_credit_limit_extension`. Maximum number of characters that the character limit can be exceeded by. Managed by the workspace admin.
	MaxCharacterLimitExtension *int `json:"max_character_limit_extension" xml:"max_character_limit_extension" required:"true"`
	// Maximum number of credits that the credit limit can be exceeded by. Managed by the workspace admin. `"unlimited"` means no cap, `0` means usage-based billing is disabled.
	MaxCreditLimitExtension *SubscriptionMaxCreditLimitExtension `json:"max_credit_limit_extension" xml:"max_credit_limit_extension" required:"true"`
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
	// The next invoice for the user.
	NextInvoice *InvoiceResponse `json:"next_invoice,omitempty" xml:"next_invoice,omitempty"`
	// The open invoices for the user.
	OpenInvoices []*InvoiceResponse `json:"open_invoices" xml:"open_invoices" required:"true"`
	// Whether the user has open invoices.
	HasOpenInvoices bool `json:"has_open_invoices" xml:"has_open_invoices" required:"true"`
	// Whether cancelling the subscription now ends it immediately, because an unpaid invoice is being collected, rather than at the end of the billing period.
	CancelsImmediately *bool `json:"cancels_immediately,omitempty" xml:"cancels_immediately,omitempty"`
	// The pending change for the user.
	PendingChange *SubscriptionPendingChange `json:"pending_change,omitempty" xml:"pending_change,omitempty"`
	// True if any workspace owned by this user's auth account has redeemed the starter first-month discount coupon.
	HasUsedStarterCouponOnAccount *bool `json:"has_used_starter_coupon_on_account,omitempty" xml:"has_used_starter_coupon_on_account,omitempty"`
	// True if any workspace owned by this user's auth account has redeemed the creator first-month discount coupon.
	HasUsedCreatorCouponOnAccount *bool `json:"has_used_creator_coupon_on_account,omitempty" xml:"has_used_creator_coupon_on_account,omitempty"`
	// True if this user's auth account qualifies for the Starter first-month promo, using the same check as checkout. The starter_discount feature flag still decides whether the promo is offered.
	IsEligibleForStarterPromo *bool `json:"is_eligible_for_starter_promo,omitempty" xml:"is_eligible_for_starter_promo,omitempty"`
	extraProperties           map[string]interface{}
}

func (s *Subscription) GetTier() (value string) {
	if s == nil {
		return
	}
	return s.Tier
}

func (s *Subscription) GetCharacterCount() (value int) {
	if s == nil {
		return
	}
	return s.CharacterCount
}

func (s *Subscription) GetCharacterLimit() (value int) {
	if s == nil {
		return
	}
	return s.CharacterLimit
}

func (s *Subscription) GetMaxCharacterLimitExtension() (value *int) {
	if s == nil {
		return
	}
	return s.MaxCharacterLimitExtension
}

func (s *Subscription) GetMaxCreditLimitExtension() (value *SubscriptionMaxCreditLimitExtension) {
	if s == nil {
		return
	}
	return s.MaxCreditLimitExtension
}

func (s *Subscription) GetCanExtendCharacterLimit() (value bool) {
	if s == nil {
		return
	}
	return s.CanExtendCharacterLimit
}

func (s *Subscription) GetAllowedToExtendCharacterLimit() (value bool) {
	if s == nil {
		return
	}
	return s.AllowedToExtendCharacterLimit
}

func (s *Subscription) GetNextCharacterCountResetUnix() (value *int) {
	if s == nil {
		return
	}
	return s.NextCharacterCountResetUnix
}

func (s *Subscription) GetVoiceSlotsUsed() (value int) {
	if s == nil {
		return
	}
	return s.VoiceSlotsUsed
}

func (s *Subscription) GetProfessionalVoiceSlotsUsed() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceSlotsUsed
}

func (s *Subscription) GetProfessionalVoiceSlotsUsedInWorkspace() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceSlotsUsedInWorkspace
}

func (s *Subscription) GetVoiceLimit() (value int) {
	if s == nil {
		return
	}
	return s.VoiceLimit
}

func (s *Subscription) GetMaxVoiceAddEdits() (value *int) {
	if s == nil {
		return
	}
	return s.MaxVoiceAddEdits
}

func (s *Subscription) GetVoiceAddEditCounter() (value int) {
	if s == nil {
		return
	}
	return s.VoiceAddEditCounter
}

func (s *Subscription) GetProfessionalVoiceLimit() (value int) {
	if s == nil {
		return
	}
	return s.ProfessionalVoiceLimit
}

func (s *Subscription) GetCanExtendVoiceLimit() (value bool) {
	if s == nil {
		return
	}
	return s.CanExtendVoiceLimit
}

func (s *Subscription) GetCanUseInstantVoiceCloning() (value bool) {
	if s == nil {
		return
	}
	return s.CanUseInstantVoiceCloning
}

func (s *Subscription) GetCanUseProfessionalVoiceCloning() (value bool) {
	if s == nil {
		return
	}
	return s.CanUseProfessionalVoiceCloning
}

func (s *Subscription) GetCurrency() (value *Currency) {
	if s == nil {
		return
	}
	return s.Currency
}

func (s *Subscription) GetCurrentOverage() (value *Price) {
	if s == nil {
		return
	}
	return s.CurrentOverage
}

func (s *Subscription) GetStatus() (value SubscriptionStatusType) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *Subscription) GetBillingPeriod() (value *BillingPeriod) {
	if s == nil {
		return
	}
	return s.BillingPeriod
}

func (s *Subscription) GetCharacterRefreshPeriod() (value *CharacterRefreshPeriod) {
	if s == nil {
		return
	}
	return s.CharacterRefreshPeriod
}

func (s *Subscription) GetNextInvoice() (value *InvoiceResponse) {
	if s == nil {
		return
	}
	return s.NextInvoice
}

func (s *Subscription) GetOpenInvoices() (value []*InvoiceResponse) {
	if s == nil {
		return
	}
	return s.OpenInvoices
}

func (s *Subscription) GetHasOpenInvoices() (value bool) {
	if s == nil {
		return
	}
	return s.HasOpenInvoices
}

func (s *Subscription) GetCancelsImmediately() (value *bool) {
	if s == nil {
		return
	}
	return s.CancelsImmediately
}

func (s *Subscription) GetPendingChange() (value *SubscriptionPendingChange) {
	if s == nil {
		return
	}
	return s.PendingChange
}

func (s *Subscription) GetHasUsedStarterCouponOnAccount() (value *bool) {
	if s == nil {
		return
	}
	return s.HasUsedStarterCouponOnAccount
}

func (s *Subscription) GetHasUsedCreatorCouponOnAccount() (value *bool) {
	if s == nil {
		return
	}
	return s.HasUsedCreatorCouponOnAccount
}

func (s *Subscription) GetIsEligibleForStarterPromo() (value *bool) {
	if s == nil {
		return
	}
	return s.IsEligibleForStarterPromo
}

func (s *Subscription) SetTier(tier string) {
	s.Tier = tier
}

func (s *Subscription) SetCharacterCount(characterCount int) {
	s.CharacterCount = characterCount
}

func (s *Subscription) SetCharacterLimit(characterLimit int) {
	s.CharacterLimit = characterLimit
}

func (s *Subscription) SetMaxCharacterLimitExtension(maxCharacterLimitExtension *int) {
	s.MaxCharacterLimitExtension = maxCharacterLimitExtension
}

func (s *Subscription) SetMaxCreditLimitExtension(maxCreditLimitExtension *SubscriptionMaxCreditLimitExtension) {
	s.MaxCreditLimitExtension = maxCreditLimitExtension
}

func (s *Subscription) SetCanExtendCharacterLimit(canExtendCharacterLimit bool) {
	s.CanExtendCharacterLimit = canExtendCharacterLimit
}

func (s *Subscription) SetAllowedToExtendCharacterLimit(allowedToExtendCharacterLimit bool) {
	s.AllowedToExtendCharacterLimit = allowedToExtendCharacterLimit
}

func (s *Subscription) SetNextCharacterCountResetUnix(nextCharacterCountResetUnix *int) {
	s.NextCharacterCountResetUnix = nextCharacterCountResetUnix
}

func (s *Subscription) SetVoiceSlotsUsed(voiceSlotsUsed int) {
	s.VoiceSlotsUsed = voiceSlotsUsed
}

func (s *Subscription) SetProfessionalVoiceSlotsUsed(professionalVoiceSlotsUsed int) {
	s.ProfessionalVoiceSlotsUsed = professionalVoiceSlotsUsed
}

func (s *Subscription) SetProfessionalVoiceSlotsUsedInWorkspace(professionalVoiceSlotsUsedInWorkspace int) {
	s.ProfessionalVoiceSlotsUsedInWorkspace = professionalVoiceSlotsUsedInWorkspace
}

func (s *Subscription) SetVoiceLimit(voiceLimit int) {
	s.VoiceLimit = voiceLimit
}

func (s *Subscription) SetMaxVoiceAddEdits(maxVoiceAddEdits *int) {
	s.MaxVoiceAddEdits = maxVoiceAddEdits
}

func (s *Subscription) SetVoiceAddEditCounter(voiceAddEditCounter int) {
	s.VoiceAddEditCounter = voiceAddEditCounter
}

func (s *Subscription) SetProfessionalVoiceLimit(professionalVoiceLimit int) {
	s.ProfessionalVoiceLimit = professionalVoiceLimit
}

func (s *Subscription) SetCanExtendVoiceLimit(canExtendVoiceLimit bool) {
	s.CanExtendVoiceLimit = canExtendVoiceLimit
}

func (s *Subscription) SetCanUseInstantVoiceCloning(canUseInstantVoiceCloning bool) {
	s.CanUseInstantVoiceCloning = canUseInstantVoiceCloning
}

func (s *Subscription) SetCanUseProfessionalVoiceCloning(canUseProfessionalVoiceCloning bool) {
	s.CanUseProfessionalVoiceCloning = canUseProfessionalVoiceCloning
}

func (s *Subscription) SetCurrency(currency *Currency) {
	s.Currency = currency
}

func (s *Subscription) SetCurrentOverage(currentOverage *Price) {
	s.CurrentOverage = currentOverage
}

func (s *Subscription) SetStatus(status SubscriptionStatusType) {
	s.Status = status
}

func (s *Subscription) SetBillingPeriod(billingPeriod *BillingPeriod) {
	s.BillingPeriod = billingPeriod
}

func (s *Subscription) SetCharacterRefreshPeriod(characterRefreshPeriod *CharacterRefreshPeriod) {
	s.CharacterRefreshPeriod = characterRefreshPeriod
}

func (s *Subscription) SetNextInvoice(nextInvoice *InvoiceResponse) {
	s.NextInvoice = nextInvoice
}

func (s *Subscription) SetOpenInvoices(openInvoices []*InvoiceResponse) {
	s.OpenInvoices = openInvoices
}

func (s *Subscription) SetHasOpenInvoices(hasOpenInvoices bool) {
	s.HasOpenInvoices = hasOpenInvoices
}

func (s *Subscription) SetCancelsImmediately(cancelsImmediately *bool) {
	s.CancelsImmediately = cancelsImmediately
}

func (s *Subscription) SetPendingChange(pendingChange *SubscriptionPendingChange) {
	s.PendingChange = pendingChange
}

func (s *Subscription) SetHasUsedStarterCouponOnAccount(hasUsedStarterCouponOnAccount *bool) {
	s.HasUsedStarterCouponOnAccount = hasUsedStarterCouponOnAccount
}

func (s *Subscription) SetHasUsedCreatorCouponOnAccount(hasUsedCreatorCouponOnAccount *bool) {
	s.HasUsedCreatorCouponOnAccount = hasUsedCreatorCouponOnAccount
}

func (s *Subscription) SetIsEligibleForStarterPromo(isEligibleForStarterPromo *bool) {
	s.IsEligibleForStarterPromo = isEligibleForStarterPromo
}

func (s *Subscription) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s Subscription) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Subscription to string"
	}
	return string(jsonData)
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
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

func (s *Subscription) MarshalJSON() ([]byte, error) {
	type embed Subscription
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
