package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OrderItemInfo struct {
	ItemId ItemId                  `json:"item_id" xml:"item_id" required:"true" pattern:"^proditem_[a-z0-9]{26}$"`
	Item   *OrderItemRequestOutput `json:"item" xml:"item" required:"true"`
	// The quoted price for this item.
	Quote           *QuoteInfo `json:"quote,omitempty" xml:"quote,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OrderItemInfo) GetItemId() (value ItemId) {
	if o == nil {
		return
	}
	return o.ItemId
}

func (o *OrderItemInfo) GetItem() (value *OrderItemRequestOutput) {
	if o == nil {
		return
	}
	return o.Item
}

func (o *OrderItemInfo) GetQuote() (value *QuoteInfo) {
	if o == nil {
		return
	}
	return o.Quote
}

func (o *OrderItemInfo) SetItemId(itemId ItemId) {
	o.ItemId = itemId
}

func (o *OrderItemInfo) SetItem(item *OrderItemRequestOutput) {
	o.Item = item
}

func (o *OrderItemInfo) SetQuote(quote *QuoteInfo) {
	o.Quote = quote
}

func (o *OrderItemInfo) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrderItemInfo) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrderItemInfo to string"
	}
	return string(jsonData)
}

func (o *OrderItemInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrderItemInfo) MarshalJSON() ([]byte, error) {
	type embed OrderItemInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
