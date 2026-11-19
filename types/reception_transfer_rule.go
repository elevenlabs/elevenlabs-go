package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A configured phone number that the agent can transfer calls to.
type ReceptionTransferRule struct {
	Label       StrippedStr `json:"label" xml:"label" required:"true"`
	PhoneNumber PhoneNumber `json:"phone_number" xml:"phone_number" required:"true"`
	Condition   StrippedStr `json:"condition" xml:"condition" required:"true"`
	// DTMF digits automatically played after the transfer destination picks up (e.g. '1' to answer a 'press 1 to accept' screen, or 'www1234' to key an extension after a pause). Valid characters are 0-9, *, #, w (0.5s pause) and W (1s pause). Twilio transfers only.
	PostDialDigits *string `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty" maxLength:"50" minLength:"1" pattern:"^[0-9*#wW]*$"`
	// ID of the transfer rule.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix   *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionTransferRule) GetLabel() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Label
}

func (r *ReceptionTransferRule) GetPhoneNumber() (value PhoneNumber) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

func (r *ReceptionTransferRule) GetCondition() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Condition
}

func (r *ReceptionTransferRule) GetPostDialDigits() (value *string) {
	if r == nil {
		return
	}
	return r.PostDialDigits
}

func (r *ReceptionTransferRule) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionTransferRule) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionTransferRule) SetLabel(label StrippedStr) {
	r.Label = label
}

func (r *ReceptionTransferRule) SetPhoneNumber(phoneNumber PhoneNumber) {
	r.PhoneNumber = phoneNumber
}

func (r *ReceptionTransferRule) SetCondition(condition StrippedStr) {
	r.Condition = condition
}

func (r *ReceptionTransferRule) SetPostDialDigits(postDialDigits *string) {
	r.PostDialDigits = postDialDigits
}

func (r *ReceptionTransferRule) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionTransferRule) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionTransferRule) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionTransferRule) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionTransferRule to string"
	}
	return string(jsonData)
}

func (r *ReceptionTransferRule) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionTransferRule) MarshalJSON() ([]byte, error) {
	type embed ReceptionTransferRule
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
