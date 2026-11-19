package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferRuleOutputProviderConfigsItem struct {
	Type     string
	Exotel   *ExotelTransferOutput
	SipTrunk *SipTransferOutput
	Twilio   *TwilioTransferOutput
}

func (t *TransferRuleOutputProviderConfigsItem) GetExotel() *ExotelTransferOutput {
	if t == nil {
		return nil
	}
	if t.Exotel == nil {
		return nil
	}
	return t.Exotel
}

func (t *TransferRuleOutputProviderConfigsItem) GetSipTrunk() *SipTransferOutput {
	if t == nil {
		return nil
	}
	if t.SipTrunk == nil {
		return nil
	}
	return t.SipTrunk
}

func (t *TransferRuleOutputProviderConfigsItem) GetTwilio() *TwilioTransferOutput {
	if t == nil {
		return nil
	}
	if t.Twilio == nil {
		return nil
	}
	return t.Twilio
}

func (t *TransferRuleOutputProviderConfigsItem) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TransferRuleOutputProviderConfigsItem) UnmarshalJSON(data []byte) error {
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
		value := new(ExotelTransferOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Exotel = value
	case "sip_trunk":
		value := new(SipTransferOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.SipTrunk = value
	case "twilio":
		value := new(TwilioTransferOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Twilio = value
	}
	return nil
}

func (t TransferRuleOutputProviderConfigsItem) MarshalJSON() ([]byte, error) {
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

type TransferRuleOutputProviderConfigsItemVisitor interface {
	VisitExotel(*ExotelTransferOutput) error
	VisitSipTrunk(*SipTransferOutput) error
	VisitTwilio(*TwilioTransferOutput) error
}

func (t *TransferRuleOutputProviderConfigsItem) Accept(visitor TransferRuleOutputProviderConfigsItemVisitor) error {
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

func (t *TransferRuleOutputProviderConfigsItem) validate() error {
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
