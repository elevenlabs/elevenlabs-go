package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type MergePreviewResponseModelPhoneNumbersItem struct {
	Provider string
	Exotel   *GetPhoneNumberExotelResponseModel
	SipTrunk *GetPhoneNumberSipTrunkResponseModel
	Twilio   *GetPhoneNumberTwilioResponseModel
}

func (m *MergePreviewResponseModelPhoneNumbersItem) GetExotel() *GetPhoneNumberExotelResponseModel {
	if m == nil {
		return nil
	}
	if m.Exotel == nil {
		return nil
	}
	return m.Exotel
}

func (m *MergePreviewResponseModelPhoneNumbersItem) GetSipTrunk() *GetPhoneNumberSipTrunkResponseModel {
	if m == nil {
		return nil
	}
	if m.SipTrunk == nil {
		return nil
	}
	return m.SipTrunk
}

func (m *MergePreviewResponseModelPhoneNumbersItem) GetTwilio() *GetPhoneNumberTwilioResponseModel {
	if m == nil {
		return nil
	}
	if m.Twilio == nil {
		return nil
	}
	return m.Twilio
}

func (m *MergePreviewResponseModelPhoneNumbersItem) GetProvider() string {
	if m == nil {
		return ""
	}
	return m.Provider
}

func (m *MergePreviewResponseModelPhoneNumbersItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", m)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(GetPhoneNumberExotelResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Exotel = value
	case "sip_trunk":
		value := new(GetPhoneNumberSipTrunkResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.SipTrunk = value
	case "twilio":
		value := new(GetPhoneNumberTwilioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Twilio = value
	}
	return nil
}

func (m MergePreviewResponseModelPhoneNumbersItem) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(m.Exotel, "provider", "exotel")
	}
	if m.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(m.SipTrunk, "provider", "sip_trunk")
	}
	if m.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(m.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type MergePreviewResponseModelPhoneNumbersItemVisitor interface {
	VisitExotel(*GetPhoneNumberExotelResponseModel) error
	VisitSipTrunk(*GetPhoneNumberSipTrunkResponseModel) error
	VisitTwilio(*GetPhoneNumberTwilioResponseModel) error
}

func (m *MergePreviewResponseModelPhoneNumbersItem) Accept(visitor MergePreviewResponseModelPhoneNumbersItemVisitor) error {
	if m.Exotel != nil {
		return visitor.VisitExotel(m.Exotel)
	}
	if m.SipTrunk != nil {
		return visitor.VisitSipTrunk(m.SipTrunk)
	}
	if m.Twilio != nil {
		return visitor.VisitTwilio(m.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *MergePreviewResponseModelPhoneNumbersItem) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if m.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if m.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if m.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Provider)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Provider != "" {
		field := fields[0]
		if m.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Provider,
				m,
			)
		}
	}
	return nil
}
