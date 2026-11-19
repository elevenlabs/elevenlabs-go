package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpsertOrderItemRequest struct {
	Item *OrderItemRequestInput `json:"item" xml:"item" required:"true"`
	// The ID of an existing item to update. Omit to create a new item.
	ItemId          *ItemId `json:"item_id,omitempty" xml:"item_id,omitempty" pattern:"^proditem_[a-z0-9]{26}$"`
	extraProperties map[string]interface{}
}

func (u *UpsertOrderItemRequest) GetItem() (value *OrderItemRequestInput) {
	if u == nil {
		return
	}
	return u.Item
}

func (u *UpsertOrderItemRequest) GetItemId() (value *ItemId) {
	if u == nil {
		return
	}
	return u.ItemId
}

func (u *UpsertOrderItemRequest) SetItem(item *OrderItemRequestInput) {
	u.Item = item
}

func (u *UpsertOrderItemRequest) SetItemId(itemId *ItemId) {
	u.ItemId = itemId
}

func (u *UpsertOrderItemRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpsertOrderItemRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpsertOrderItemRequest to string"
	}
	return string(jsonData)
}

func (u *UpsertOrderItemRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpsertOrderItemRequest) MarshalJSON() ([]byte, error) {
	type embed UpsertOrderItemRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
