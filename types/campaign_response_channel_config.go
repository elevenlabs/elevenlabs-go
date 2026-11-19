package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CampaignResponseChannelConfig struct {
	Channel  string
	Phone    *PhoneChannelConfigOutput
	Whatsapp *WhatsAppChannelConfig
}

func (c *CampaignResponseChannelConfig) GetPhone() *PhoneChannelConfigOutput {
	if c == nil {
		return nil
	}
	if c.Phone == nil {
		return nil
	}
	return c.Phone
}

func (c *CampaignResponseChannelConfig) GetWhatsapp() *WhatsAppChannelConfig {
	if c == nil {
		return nil
	}
	if c.Whatsapp == nil {
		return nil
	}
	return c.Whatsapp
}

func (c *CampaignResponseChannelConfig) GetChannel() string {
	if c == nil {
		return ""
	}
	return c.Channel
}

func (c *CampaignResponseChannelConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Channel string `json:"channel"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Channel = unmarshaler.Channel
	if unmarshaler.Channel == "" {
		return fmt.Errorf("%T did not include discriminant channel", c)
	}
	switch unmarshaler.Channel {
	case "phone":
		value := new(PhoneChannelConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Phone = value
	case "whatsapp":
		value := new(WhatsAppChannelConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Whatsapp = value
	}
	return nil
}

func (c CampaignResponseChannelConfig) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(c.Phone, "channel", "phone")
	}
	if c.Whatsapp != nil {
		return internal.MarshalJSONWithDiscriminant(c.Whatsapp, "channel", "whatsapp")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CampaignResponseChannelConfigVisitor interface {
	VisitPhone(*PhoneChannelConfigOutput) error
	VisitWhatsapp(*WhatsAppChannelConfig) error
}

func (c *CampaignResponseChannelConfig) Accept(visitor CampaignResponseChannelConfigVisitor) error {
	if c.Phone != nil {
		return visitor.VisitPhone(c.Phone)
	}
	if c.Whatsapp != nil {
		return visitor.VisitWhatsapp(c.Whatsapp)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CampaignResponseChannelConfig) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Phone != nil {
		fields = append(fields, "phone")
	}
	if c.Whatsapp != nil {
		fields = append(fields, "whatsapp")
	}
	if len(fields) == 0 {
		if c.Channel != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Channel)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Channel != "" {
		field := fields[0]
		if c.Channel != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Channel,
				c,
			)
		}
	}
	return nil
}
