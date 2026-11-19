package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type User struct {
	// The unique identifier of the user.
	UserId string `json:"user_id" xml:"user_id" required:"true"`
	// The unique identifier of the workspace this user's API key belongs to.
	WorkspaceId  string                `json:"workspace_id" xml:"workspace_id" required:"true"`
	Subscription *SubscriptionResponse `json:"subscription" xml:"subscription" required:"true"`
	// Whether the user is new. This field is deprecated and will be removed in the future. Use 'created_at' instead.
	IsNewUser bool `json:"is_new_user" xml:"is_new_user" required:"true"`
	// This field is deprecated and will be removed in a future major version. Instead use subscription.trust_on_invoice_creation.
	CanUseDelayedPaymentMethods bool `json:"can_use_delayed_payment_methods" xml:"can_use_delayed_payment_methods" required:"true"`
	// Whether the user's onboarding is completed.
	IsOnboardingCompleted bool `json:"is_onboarding_completed" xml:"is_onboarding_completed" required:"true"`
	// Whether the user's onboarding checklist is completed.
	IsOnboardingChecklistCompleted bool `json:"is_onboarding_checklist_completed" xml:"is_onboarding_checklist_completed" required:"true"`
	// Whether to show compliance terms (ToS, Privacy Policy, biometric consent) during onboarding. Set for users signing up from the marketing site.
	ShowComplianceTerms *bool `json:"show_compliance_terms,omitempty" xml:"show_compliance_terms,omitempty"`
	// First name of the user.
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// Whether the user's API key is hashed.
	IsApiKeyHashed *bool `json:"is_api_key_hashed,omitempty" xml:"is_api_key_hashed,omitempty"`
	// The preview of the user's API key.
	XiApiKeyPreview *string `json:"xi_api_key_preview,omitempty" xml:"xi_api_key_preview,omitempty"`
	// The referral link code of the user.
	ReferralLinkCode *string `json:"referral_link_code,omitempty" xml:"referral_link_code,omitempty"`
	// The Partnerstack partner default link of the user.
	PartnerstackPartnerDefaultLink *string `json:"partnerstack_partner_default_link,omitempty" xml:"partnerstack_partner_default_link,omitempty"`
	// The unix timestamp of the user's creation. 0 if the user was created before the unix timestamp was added.
	CreatedAt int `json:"created_at" xml:"created_at" required:"true"`
	// Seat types for workspace members.
	SeatType        SeatType `json:"seat_type" xml:"seat_type" required:"true"`
	extraProperties map[string]interface{}
}

func (u *User) GetUserId() (value string) {
	if u == nil {
		return
	}
	return u.UserId
}

func (u *User) GetWorkspaceId() (value string) {
	if u == nil {
		return
	}
	return u.WorkspaceId
}

func (u *User) GetSubscription() (value *SubscriptionResponse) {
	if u == nil {
		return
	}
	return u.Subscription
}

func (u *User) GetIsNewUser() (value bool) {
	if u == nil {
		return
	}
	return u.IsNewUser
}

func (u *User) GetCanUseDelayedPaymentMethods() (value bool) {
	if u == nil {
		return
	}
	return u.CanUseDelayedPaymentMethods
}

func (u *User) GetIsOnboardingCompleted() (value bool) {
	if u == nil {
		return
	}
	return u.IsOnboardingCompleted
}

func (u *User) GetIsOnboardingChecklistCompleted() (value bool) {
	if u == nil {
		return
	}
	return u.IsOnboardingChecklistCompleted
}

func (u *User) GetShowComplianceTerms() (value *bool) {
	if u == nil {
		return
	}
	return u.ShowComplianceTerms
}

func (u *User) GetFirstName() (value *string) {
	if u == nil {
		return
	}
	return u.FirstName
}

func (u *User) GetIsApiKeyHashed() (value *bool) {
	if u == nil {
		return
	}
	return u.IsApiKeyHashed
}

func (u *User) GetXiApiKeyPreview() (value *string) {
	if u == nil {
		return
	}
	return u.XiApiKeyPreview
}

func (u *User) GetReferralLinkCode() (value *string) {
	if u == nil {
		return
	}
	return u.ReferralLinkCode
}

func (u *User) GetPartnerstackPartnerDefaultLink() (value *string) {
	if u == nil {
		return
	}
	return u.PartnerstackPartnerDefaultLink
}

func (u *User) GetCreatedAt() (value int) {
	if u == nil {
		return
	}
	return u.CreatedAt
}

func (u *User) GetSeatType() (value SeatType) {
	if u == nil {
		return
	}
	return u.SeatType
}

func (u *User) SetUserId(userId string) {
	u.UserId = userId
}

func (u *User) SetWorkspaceId(workspaceId string) {
	u.WorkspaceId = workspaceId
}

func (u *User) SetSubscription(subscription *SubscriptionResponse) {
	u.Subscription = subscription
}

func (u *User) SetIsNewUser(isNewUser bool) {
	u.IsNewUser = isNewUser
}

func (u *User) SetCanUseDelayedPaymentMethods(canUseDelayedPaymentMethods bool) {
	u.CanUseDelayedPaymentMethods = canUseDelayedPaymentMethods
}

func (u *User) SetIsOnboardingCompleted(isOnboardingCompleted bool) {
	u.IsOnboardingCompleted = isOnboardingCompleted
}

func (u *User) SetIsOnboardingChecklistCompleted(isOnboardingChecklistCompleted bool) {
	u.IsOnboardingChecklistCompleted = isOnboardingChecklistCompleted
}

func (u *User) SetShowComplianceTerms(showComplianceTerms *bool) {
	u.ShowComplianceTerms = showComplianceTerms
}

func (u *User) SetFirstName(firstName *string) {
	u.FirstName = firstName
}

func (u *User) SetIsApiKeyHashed(isApiKeyHashed *bool) {
	u.IsApiKeyHashed = isApiKeyHashed
}

func (u *User) SetXiApiKeyPreview(xiApiKeyPreview *string) {
	u.XiApiKeyPreview = xiApiKeyPreview
}

func (u *User) SetReferralLinkCode(referralLinkCode *string) {
	u.ReferralLinkCode = referralLinkCode
}

func (u *User) SetPartnerstackPartnerDefaultLink(partnerstackPartnerDefaultLink *string) {
	u.PartnerstackPartnerDefaultLink = partnerstackPartnerDefaultLink
}

func (u *User) SetCreatedAt(createdAt int) {
	u.CreatedAt = createdAt
}

func (u *User) SetSeatType(seatType SeatType) {
	u.SeatType = seatType
}

func (u *User) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u User) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: User to string"
	}
	return string(jsonData)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type embed User
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
