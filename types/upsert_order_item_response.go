package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpsertOrderItemResponse struct {
	ItemId ItemId `json:"item_id" xml:"item_id" required:"true" pattern:"^proditem_[a-z0-9]{26}$"`
	// The quoted price for this item.
	Quote           *QuoteInfo `json:"quote,omitempty" xml:"quote,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpsertOrderItemResponse) GetItemId() (value ItemId) {
	if u == nil {
		return
	}
	return u.ItemId
}

func (u *UpsertOrderItemResponse) GetQuote() (value *QuoteInfo) {
	if u == nil {
		return
	}
	return u.Quote
}

func (u *UpsertOrderItemResponse) SetItemId(itemId ItemId) {
	u.ItemId = itemId
}

func (u *UpsertOrderItemResponse) SetQuote(quote *QuoteInfo) {
	u.Quote = quote
}

func (u *UpsertOrderItemResponse) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpsertOrderItemResponse) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpsertOrderItemResponse to string"
	}
	return string(jsonData)
}

func (u *UpsertOrderItemResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpsertOrderItemResponse) MarshalJSON() ([]byte, error) {
	type embed UpsertOrderItemResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
