package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowPhoneNumberNodeModelInputProviderConfigsItem struct {
	Type     string
	Exotel   *ExotelTransferInput
	SipTrunk *SipTransferInput
	Twilio   *TwilioTransferInput
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) GetExotel() *ExotelTransferInput {
	if w == nil {
		return nil
	}
	if w.Exotel == nil {
		return nil
	}
	return w.Exotel
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) GetSipTrunk() *SipTransferInput {
	if w == nil {
		return nil
	}
	if w.SipTrunk == nil {
		return nil
	}
	return w.SipTrunk
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) GetTwilio() *TwilioTransferInput {
	if w == nil {
		return nil
	}
	if w.Twilio == nil {
		return nil
	}
	return w.Twilio
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "exotel":
		value := new(ExotelTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Exotel = value
	case "sip_trunk":
		value := new(SipTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.SipTrunk = value
	case "twilio":
		value := new(TwilioTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Twilio = value
	}
	return nil
}

func (w WorkflowPhoneNumberNodeModelInputProviderConfigsItem) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(w.Exotel, "type", "exotel")
	}
	if w.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(w.SipTrunk, "type", "sip_trunk")
	}
	if w.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(w.Twilio, "type", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowPhoneNumberNodeModelInputProviderConfigsItemVisitor interface {
	VisitExotel(*ExotelTransferInput) error
	VisitSipTrunk(*SipTransferInput) error
	VisitTwilio(*TwilioTransferInput) error
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) Accept(visitor WorkflowPhoneNumberNodeModelInputProviderConfigsItemVisitor) error {
	if w.Exotel != nil {
		return visitor.VisitExotel(w.Exotel)
	}
	if w.SipTrunk != nil {
		return visitor.VisitSipTrunk(w.SipTrunk)
	}
	if w.Twilio != nil {
		return visitor.VisitTwilio(w.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowPhoneNumberNodeModelInputProviderConfigsItem) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if w.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if w.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}
