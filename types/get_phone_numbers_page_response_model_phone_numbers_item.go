package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetPhoneNumbersPageResponseModelPhoneNumbersItem struct {
	Provider string
	Exotel   *GetPhoneNumberExotelResponseModel
	SipTrunk *GetPhoneNumberSipTrunkResponseModel
	Twilio   *GetPhoneNumberTwilioResponseModel
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) GetExotel() *GetPhoneNumberExotelResponseModel {
	if g == nil {
		return nil
	}
	if g.Exotel == nil {
		return nil
	}
	return g.Exotel
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) GetSipTrunk() *GetPhoneNumberSipTrunkResponseModel {
	if g == nil {
		return nil
	}
	if g.SipTrunk == nil {
		return nil
	}
	return g.SipTrunk
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) GetTwilio() *GetPhoneNumberTwilioResponseModel {
	if g == nil {
		return nil
	}
	if g.Twilio == nil {
		return nil
	}
	return g.Twilio
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) GetProvider() string {
	if g == nil {
		return ""
	}
	return g.Provider
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", g)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(GetPhoneNumberExotelResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Exotel = value
	case "sip_trunk":
		value := new(GetPhoneNumberSipTrunkResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.SipTrunk = value
	case "twilio":
		value := new(GetPhoneNumberTwilioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Twilio = value
	}
	return nil
}

func (g GetPhoneNumbersPageResponseModelPhoneNumbersItem) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(g.Exotel, "provider", "exotel")
	}
	if g.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(g.SipTrunk, "provider", "sip_trunk")
	}
	if g.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(g.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetPhoneNumbersPageResponseModelPhoneNumbersItemVisitor interface {
	VisitExotel(*GetPhoneNumberExotelResponseModel) error
	VisitSipTrunk(*GetPhoneNumberSipTrunkResponseModel) error
	VisitTwilio(*GetPhoneNumberTwilioResponseModel) error
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) Accept(visitor GetPhoneNumbersPageResponseModelPhoneNumbersItemVisitor) error {
	if g.Exotel != nil {
		return visitor.VisitExotel(g.Exotel)
	}
	if g.SipTrunk != nil {
		return visitor.VisitSipTrunk(g.SipTrunk)
	}
	if g.Twilio != nil {
		return visitor.VisitTwilio(g.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetPhoneNumbersPageResponseModelPhoneNumbersItem) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if g.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if g.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if g.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Provider)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Provider != "" {
		field := fields[0]
		if g.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Provider,
				g,
			)
		}
	}
	return nil
}
