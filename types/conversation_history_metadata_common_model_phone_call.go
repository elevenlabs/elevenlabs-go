package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryMetadataCommonModelPhoneCall struct {
	Type        string
	Exotel      *ConversationHistoryExotelPhoneCallModel
	SipTrunking *ConversationHistorySipTrunkingPhoneCallModel
	Twilio      *ConversationHistoryTwilioPhoneCallModel
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) GetExotel() *ConversationHistoryExotelPhoneCallModel {
	if c == nil {
		return nil
	}
	if c.Exotel == nil {
		return nil
	}
	return c.Exotel
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) GetSipTrunking() *ConversationHistorySipTrunkingPhoneCallModel {
	if c == nil {
		return nil
	}
	if c.SipTrunking == nil {
		return nil
	}
	return c.SipTrunking
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) GetTwilio() *ConversationHistoryTwilioPhoneCallModel {
	if c == nil {
		return nil
	}
	if c.Twilio == nil {
		return nil
	}
	return c.Twilio
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "exotel":
		value := new(ConversationHistoryExotelPhoneCallModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Exotel = value
	case "sip_trunking":
		value := new(ConversationHistorySipTrunkingPhoneCallModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.SipTrunking = value
	case "twilio":
		value := new(ConversationHistoryTwilioPhoneCallModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Twilio = value
	}
	return nil
}

func (c ConversationHistoryMetadataCommonModelPhoneCall) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(c.Exotel, "type", "exotel")
	}
	if c.SipTrunking != nil {
		return internal.MarshalJSONWithDiscriminant(c.SipTrunking, "type", "sip_trunking")
	}
	if c.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(c.Twilio, "type", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ConversationHistoryMetadataCommonModelPhoneCallVisitor interface {
	VisitExotel(*ConversationHistoryExotelPhoneCallModel) error
	VisitSipTrunking(*ConversationHistorySipTrunkingPhoneCallModel) error
	VisitTwilio(*ConversationHistoryTwilioPhoneCallModel) error
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) Accept(visitor ConversationHistoryMetadataCommonModelPhoneCallVisitor) error {
	if c.Exotel != nil {
		return visitor.VisitExotel(c.Exotel)
	}
	if c.SipTrunking != nil {
		return visitor.VisitSipTrunking(c.SipTrunking)
	}
	if c.Twilio != nil {
		return visitor.VisitTwilio(c.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ConversationHistoryMetadataCommonModelPhoneCall) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if c.SipTrunking != nil {
		fields = append(fields, "sip_trunking")
	}
	if c.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}
