package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListOrdersResponse struct {
	// The list of orders matching the query.
	Orders          []*OrderSummary `json:"orders" xml:"orders" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListOrdersResponse) GetOrders() (value []*OrderSummary) {
	if l == nil {
		return
	}
	return l.Orders
}

func (l *ListOrdersResponse) SetOrders(orders []*OrderSummary) {
	l.Orders = orders
}

func (l *ListOrdersResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListOrdersResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListOrdersResponse to string"
	}
	return string(jsonData)
}

func (l *ListOrdersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListOrdersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListOrdersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListOrdersResponse) MarshalJSON() ([]byte, error) {
	type embed ListOrdersResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
