package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The template a business-initiated conversation opened with, and how far it got.
type WhatsAppOpeningMessage struct {
	TemplateName    string  `json:"template_name" xml:"template_name" required:"true"`
	Wamid           *string `json:"wamid,omitempty" xml:"wamid,omitempty"`
	SentAtUnix      int     `json:"sent_at_unix" xml:"sent_at_unix" required:"true"`
	DeliveredAtUnix *int    `json:"delivered_at_unix,omitempty" xml:"delivered_at_unix,omitempty"`
	ReadAtUnix      *int    `json:"read_at_unix,omitempty" xml:"read_at_unix,omitempty"`
	RepliedAtUnix   *int    `json:"replied_at_unix,omitempty" xml:"replied_at_unix,omitempty"`
	FailedAtUnix    *int    `json:"failed_at_unix,omitempty" xml:"failed_at_unix,omitempty"`
	ErrorCode       *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	PricingCategory *string `json:"pricing_category,omitempty" xml:"pricing_category,omitempty"`
	Billable        *bool   `json:"billable,omitempty" xml:"billable,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppOpeningMessage) GetTemplateName() (value string) {
	if w == nil {
		return
	}
	return w.TemplateName
}

func (w *WhatsAppOpeningMessage) GetWamid() (value *string) {
	if w == nil {
		return
	}
	return w.Wamid
}

func (w *WhatsAppOpeningMessage) GetSentAtUnix() (value int) {
	if w == nil {
		return
	}
	return w.SentAtUnix
}

func (w *WhatsAppOpeningMessage) GetDeliveredAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.DeliveredAtUnix
}

func (w *WhatsAppOpeningMessage) GetReadAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.ReadAtUnix
}

func (w *WhatsAppOpeningMessage) GetRepliedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.RepliedAtUnix
}

func (w *WhatsAppOpeningMessage) GetFailedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.FailedAtUnix
}

func (w *WhatsAppOpeningMessage) GetErrorCode() (value *string) {
	if w == nil {
		return
	}
	return w.ErrorCode
}

func (w *WhatsAppOpeningMessage) GetPricingCategory() (value *string) {
	if w == nil {
		return
	}
	return w.PricingCategory
}

func (w *WhatsAppOpeningMessage) GetBillable() (value *bool) {
	if w == nil {
		return
	}
	return w.Billable
}

func (w *WhatsAppOpeningMessage) SetTemplateName(templateName string) {
	w.TemplateName = templateName
}

func (w *WhatsAppOpeningMessage) SetWamid(wamid *string) {
	w.Wamid = wamid
}

func (w *WhatsAppOpeningMessage) SetSentAtUnix(sentAtUnix int) {
	w.SentAtUnix = sentAtUnix
}

func (w *WhatsAppOpeningMessage) SetDeliveredAtUnix(deliveredAtUnix *int) {
	w.DeliveredAtUnix = deliveredAtUnix
}

func (w *WhatsAppOpeningMessage) SetReadAtUnix(readAtUnix *int) {
	w.ReadAtUnix = readAtUnix
}

func (w *WhatsAppOpeningMessage) SetRepliedAtUnix(repliedAtUnix *int) {
	w.RepliedAtUnix = repliedAtUnix
}

func (w *WhatsAppOpeningMessage) SetFailedAtUnix(failedAtUnix *int) {
	w.FailedAtUnix = failedAtUnix
}

func (w *WhatsAppOpeningMessage) SetErrorCode(errorCode *string) {
	w.ErrorCode = errorCode
}

func (w *WhatsAppOpeningMessage) SetPricingCategory(pricingCategory *string) {
	w.PricingCategory = pricingCategory
}

func (w *WhatsAppOpeningMessage) SetBillable(billable *bool) {
	w.Billable = billable
}

func (w *WhatsAppOpeningMessage) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppOpeningMessage) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppOpeningMessage to string"
	}
	return string(jsonData)
}

func (w *WhatsAppOpeningMessage) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppOpeningMessage) MarshalJSON() ([]byte, error) {
	type embed WhatsAppOpeningMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
