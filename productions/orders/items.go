package orders

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost struct {
	// The ID of the order.
	OrderId OrderId                       `json:"-" url:"-"`
	Request *types.UpsertOrderItemRequest `json:"request" xml:"request" required:"true"`
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) GetOrderId() (value OrderId) {
	if b == nil {
		return
	}
	return b.OrderId
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) GetRequest() (value *types.UpsertOrderItemRequest) {
	if b == nil {
		return
	}
	return b.Request
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) SetOrderId(orderId OrderId) {
	b.OrderId = orderId
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) SetRequest(request *types.UpsertOrderItemRequest) {
	b.Request = request
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost(tmp)
	return nil
}

func (b *BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
