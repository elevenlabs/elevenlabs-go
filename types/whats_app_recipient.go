package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A WhatsApp user by phone number and/or BSUID; the phone wins where one is needed.
type WhatsAppRecipient struct {
	Phone           *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Bsuid           *string `json:"bsuid,omitempty" xml:"bsuid,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppRecipient) GetPhone() (value *string) {
	if w == nil {
		return
	}
	return w.Phone
}

func (w *WhatsAppRecipient) GetBsuid() (value *string) {
	if w == nil {
		return
	}
	return w.Bsuid
}

func (w *WhatsAppRecipient) SetPhone(phone *string) {
	w.Phone = phone
}

func (w *WhatsAppRecipient) SetBsuid(bsuid *string) {
	w.Bsuid = bsuid
}

func (w *WhatsAppRecipient) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppRecipient) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppRecipient to string"
	}
	return string(jsonData)
}

func (w *WhatsAppRecipient) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppRecipient) MarshalJSON() ([]byte, error) {
	type embed WhatsAppRecipient
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
