package products

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReceptionUpdateProductStockRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the product. Returned when you create one; list them with `GET /v1/reception/products`.
	ProductId string `json:"-" url:"-"`
	// Location whose count is set.
	LocationId string `json:"location_id" xml:"location_id" required:"true"`
	// The new on-hand quantity.
	OnHand int `json:"on_hand" xml:"on_hand" required:"true" min:"0"`
	// On-hand value the caller last saw. The write is rejected with a 409 if it no longer matches, so a sale made while the caller was editing is never overwritten.
	ExpectedOnHand int `json:"expected_on_hand" xml:"expected_on_hand" required:"true" min:"0"`
}

func (r *ReceptionUpdateProductStockRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateProductStockRequest) GetProductId() (value string) {
	if r == nil {
		return
	}
	return r.ProductId
}

func (r *ReceptionUpdateProductStockRequest) GetLocationId() (value string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionUpdateProductStockRequest) GetOnHand() (value int) {
	if r == nil {
		return
	}
	return r.OnHand
}

func (r *ReceptionUpdateProductStockRequest) GetExpectedOnHand() (value int) {
	if r == nil {
		return
	}
	return r.ExpectedOnHand
}

func (r *ReceptionUpdateProductStockRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateProductStockRequest) SetProductId(productId string) {
	r.ProductId = productId
}

func (r *ReceptionUpdateProductStockRequest) SetLocationId(locationId string) {
	r.LocationId = locationId
}

func (r *ReceptionUpdateProductStockRequest) SetOnHand(onHand int) {
	r.OnHand = onHand
}

func (r *ReceptionUpdateProductStockRequest) SetExpectedOnHand(expectedOnHand int) {
	r.ExpectedOnHand = expectedOnHand
}

func (r *ReceptionUpdateProductStockRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateProductStockRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateProductStockRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateProductStockRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionUpdateProductStockRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionUpdateProductStockRequest(tmp)
	return nil
}

func (r *ReceptionUpdateProductStockRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateProductStockRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionAdjustProductStockRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the product. Returned when you create one; list them with `GET /v1/reception/products`.
	ProductId string `json:"-" url:"-"`
	// Location whose count changes.
	LocationId string `json:"location_id" xml:"location_id" required:"true"`
	// Signed change to apply to on-hand. Prefer this over an absolute write whenever the intent is 'a few more/fewer'.
	Delta int `json:"delta" xml:"delta" required:"true"`
}

func (r *ReceptionAdjustProductStockRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionAdjustProductStockRequest) GetProductId() (value string) {
	if r == nil {
		return
	}
	return r.ProductId
}

func (r *ReceptionAdjustProductStockRequest) GetLocationId() (value string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionAdjustProductStockRequest) GetDelta() (value int) {
	if r == nil {
		return
	}
	return r.Delta
}

func (r *ReceptionAdjustProductStockRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionAdjustProductStockRequest) SetProductId(productId string) {
	r.ProductId = productId
}

func (r *ReceptionAdjustProductStockRequest) SetLocationId(locationId string) {
	r.LocationId = locationId
}

func (r *ReceptionAdjustProductStockRequest) SetDelta(delta int) {
	r.Delta = delta
}

func (r *ReceptionAdjustProductStockRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionAdjustProductStockRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionAdjustProductStockRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionAdjustProductStockRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionAdjustProductStockRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionAdjustProductStockRequest(tmp)
	return nil
}

func (r *ReceptionAdjustProductStockRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionAdjustProductStockRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
