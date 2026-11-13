package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LeadView struct {
	LeadId string `json:"lead_id" xml:"lead_id" required:"true"`
	// The `id` given to this recipient when the campaign was created.
	RecipientId *string `json:"recipient_id" xml:"recipient_id" required:"true"`
	// The number the lead's next call dials. It changes when the lead moves on to another listed number or the contact asks to be called on a different one.
	PhoneNumber string `json:"phone_number" xml:"phone_number" required:"true"`
	// The lead's listed numbers, in calling order.
	PhoneNumbers []string `json:"phone_numbers,omitempty" xml:"phone_numbers,omitempty"`
	// The position in `phone_numbers` of the number being called.
	PhoneNumbersIndex int `json:"phone_numbers_index" xml:"phone_numbers_index" required:"true"`
	// The contact's identity: the `user_id` submitted in `conversation_initiation_client_data`, else the lead's phone number.
	UserId string `json:"user_id" xml:"user_id" required:"true"`
	// Calls placed, counted against the policy's `max_total_calls`.
	AttemptCount int `json:"attempt_count" xml:"attempt_count" required:"true"`
	// Failed calls to the current number; resets on a new number.
	FailureCount int `json:"failure_count" xml:"failure_count" required:"true"`
	// Callbacks the contact asked for.
	CallbackCount int `json:"callback_count" xml:"callback_count" required:"true"`
	// Times the contact asked to be called on another number.
	RedirectCount int `json:"redirect_count" xml:"redirect_count" required:"true"`
	// Set when the contact asked not to be called, the lead was removed, or the campaign stopped. The lead isn't called again.
	DoNotCall bool `json:"do_not_call" xml:"do_not_call" required:"true"`
	// Why the lead won't be called again; unset while it may be.
	TerminalDisposition *TerminalDisposition `json:"terminal_disposition" xml:"terminal_disposition" required:"true"`
	// When the lead joined the campaign; unset on older leads.
	CreatedAtUnix *int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	UpdatedAtUnix int  `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	// The dynamic variables submitted for the lead, with secrets redacted.
	DynamicVariables map[string]any            `json:"dynamic_variables" required:"true"`
	Overrides        *FlattenedConfigOverrides `json:"overrides" xml:"overrides" required:"true"`
	// The lead's latest message delivery; set on message campaigns.
	Delivery        *LeadDeliveryView `json:"delivery,omitempty" xml:"delivery,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LeadView) GetLeadId() (value string) {
	if l == nil {
		return
	}
	return l.LeadId
}

func (l *LeadView) GetRecipientId() (value *string) {
	if l == nil {
		return
	}
	return l.RecipientId
}

func (l *LeadView) GetPhoneNumber() (value string) {
	if l == nil {
		return
	}
	return l.PhoneNumber
}

func (l *LeadView) GetPhoneNumbers() (value []string) {
	if l == nil {
		return
	}
	return l.PhoneNumbers
}

func (l *LeadView) GetPhoneNumbersIndex() (value int) {
	if l == nil {
		return
	}
	return l.PhoneNumbersIndex
}

func (l *LeadView) GetUserId() (value string) {
	if l == nil {
		return
	}
	return l.UserId
}

func (l *LeadView) GetAttemptCount() (value int) {
	if l == nil {
		return
	}
	return l.AttemptCount
}

func (l *LeadView) GetFailureCount() (value int) {
	if l == nil {
		return
	}
	return l.FailureCount
}

func (l *LeadView) GetCallbackCount() (value int) {
	if l == nil {
		return
	}
	return l.CallbackCount
}

func (l *LeadView) GetRedirectCount() (value int) {
	if l == nil {
		return
	}
	return l.RedirectCount
}

func (l *LeadView) GetDoNotCall() (value bool) {
	if l == nil {
		return
	}
	return l.DoNotCall
}

func (l *LeadView) GetTerminalDisposition() (value *TerminalDisposition) {
	if l == nil {
		return
	}
	return l.TerminalDisposition
}

func (l *LeadView) GetCreatedAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.CreatedAtUnix
}

func (l *LeadView) GetUpdatedAtUnix() (value int) {
	if l == nil {
		return
	}
	return l.UpdatedAtUnix
}

func (l *LeadView) GetDynamicVariables() (value map[string]any) {
	if l == nil {
		return
	}
	return l.DynamicVariables
}

func (l *LeadView) GetOverrides() (value *FlattenedConfigOverrides) {
	if l == nil {
		return
	}
	return l.Overrides
}

func (l *LeadView) GetDelivery() (value *LeadDeliveryView) {
	if l == nil {
		return
	}
	return l.Delivery
}

func (l *LeadView) SetLeadId(leadId string) {
	l.LeadId = leadId
}

func (l *LeadView) SetRecipientId(recipientId *string) {
	l.RecipientId = recipientId
}

func (l *LeadView) SetPhoneNumber(phoneNumber string) {
	l.PhoneNumber = phoneNumber
}

func (l *LeadView) SetPhoneNumbers(phoneNumbers []string) {
	l.PhoneNumbers = phoneNumbers
}

func (l *LeadView) SetPhoneNumbersIndex(phoneNumbersIndex int) {
	l.PhoneNumbersIndex = phoneNumbersIndex
}

func (l *LeadView) SetUserId(userId string) {
	l.UserId = userId
}

func (l *LeadView) SetAttemptCount(attemptCount int) {
	l.AttemptCount = attemptCount
}

func (l *LeadView) SetFailureCount(failureCount int) {
	l.FailureCount = failureCount
}

func (l *LeadView) SetCallbackCount(callbackCount int) {
	l.CallbackCount = callbackCount
}

func (l *LeadView) SetRedirectCount(redirectCount int) {
	l.RedirectCount = redirectCount
}

func (l *LeadView) SetDoNotCall(doNotCall bool) {
	l.DoNotCall = doNotCall
}

func (l *LeadView) SetTerminalDisposition(terminalDisposition *TerminalDisposition) {
	l.TerminalDisposition = terminalDisposition
}

func (l *LeadView) SetCreatedAtUnix(createdAtUnix *int) {
	l.CreatedAtUnix = createdAtUnix
}

func (l *LeadView) SetUpdatedAtUnix(updatedAtUnix int) {
	l.UpdatedAtUnix = updatedAtUnix
}

func (l *LeadView) SetDynamicVariables(dynamicVariables map[string]any) {
	l.DynamicVariables = dynamicVariables
}

func (l *LeadView) SetOverrides(overrides *FlattenedConfigOverrides) {
	l.Overrides = overrides
}

func (l *LeadView) SetDelivery(delivery *LeadDeliveryView) {
	l.Delivery = delivery
}

func (l *LeadView) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LeadView) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LeadView to string"
	}
	return string(jsonData)
}

func (l *LeadView) UnmarshalJSON(data []byte) error {
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

func (l *LeadView) MarshalJSON() ([]byte, error) {
	type embed LeadView
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
