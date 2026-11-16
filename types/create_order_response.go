package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateOrderResponse struct {
	OrderId OrderId `json:"order_id" xml:"order_id" required:"true" pattern:"^prodorder_[a-z0-9]{26}$"`
	// Whether this is a sandbox order.
	Sandbox         *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateOrderResponse) GetOrderId() (value OrderId) {
	if c == nil {
		return
	}
	return c.OrderId
}

func (c *CreateOrderResponse) GetSandbox() (value *bool) {
	if c == nil {
		return
	}
	return c.Sandbox
}

func (c *CreateOrderResponse) SetOrderId(orderId OrderId) {
	c.OrderId = orderId
}

func (c *CreateOrderResponse) SetSandbox(sandbox *bool) {
	c.Sandbox = sandbox
}

func (c *CreateOrderResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateOrderResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateOrderResponse to string"
	}
	return string(jsonData)
}

func (c *CreateOrderResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateOrderResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateOrderResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateOrderResponse) MarshalJSON() ([]byte, error) {
	type embed CreateOrderResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
