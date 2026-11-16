package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferRuleInputProviderConfigsItem struct {
	Type     string
	Exotel   *ExotelTransferInput
	SipTrunk *SipTransferInput
	Twilio   *TwilioTransferInput
}

func (t *TransferRuleInputProviderConfigsItem) GetExotel() *ExotelTransferInput {
	if t == nil {
		return nil
	}
	if t.Exotel == nil {
		return nil
	}
	return t.Exotel
}

func (t *TransferRuleInputProviderConfigsItem) GetSipTrunk() *SipTransferInput {
	if t == nil {
		return nil
	}
	if t.SipTrunk == nil {
		return nil
	}
	return t.SipTrunk
}

func (t *TransferRuleInputProviderConfigsItem) GetTwilio() *TwilioTransferInput {
	if t == nil {
		return nil
	}
	if t.Twilio == nil {
		return nil
	}
	return t.Twilio
}

func (t *TransferRuleInputProviderConfigsItem) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TransferRuleInputProviderConfigsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", t)
	}
	switch unmarshaler.Type {
	case "exotel":
		value := new(ExotelTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Exotel = value
	case "sip_trunk":
		value := new(SipTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.SipTrunk = value
	case "twilio":
		value := new(TwilioTransferInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Twilio = value
	}
	return nil
}

func (t TransferRuleInputProviderConfigsItem) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(t.Exotel, "type", "exotel")
	}
	if t.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(t.SipTrunk, "type", "sip_trunk")
	}
	if t.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(t.Twilio, "type", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TransferRuleInputProviderConfigsItemVisitor interface {
	VisitExotel(*ExotelTransferInput) error
	VisitSipTrunk(*SipTransferInput) error
	VisitTwilio(*TwilioTransferInput) error
}

func (t *TransferRuleInputProviderConfigsItem) Accept(visitor TransferRuleInputProviderConfigsItemVisitor) error {
	if t.Exotel != nil {
		return visitor.VisitExotel(t.Exotel)
	}
	if t.SipTrunk != nil {
		return visitor.VisitSipTrunk(t.SipTrunk)
	}
	if t.Twilio != nil {
		return visitor.VisitTwilio(t.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TransferRuleInputProviderConfigsItem) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if t.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if t.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if t.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Type)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Type != "" {
		field := fields[0]
		if t.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Type,
				t,
			)
		}
	}
	return nil
}
