package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionCreateProductRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Whether a catalog item's price and duration are known up front.
	//
	// - `fixed`: the catalog price and duration apply (the default).
	// - `quote`: neither is known until the owner has seen the job. The
	// receptionist collects intake answers and files a quote request instead
	// of booking a slot or taking an order.
	PricingMode *types.PricingMode `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
	// Questions the agent asks before booking or ordering this item. Required for a useful `pricing_mode='quote'` item -- without them a quote request arrives with nothing to price -- but also valid on fixed-price items that need extra detail (current hair color before a dye job, allergies before a treatment).
	IntakeQuestions []*types.ReceptionIntakeQuestion `json:"intake_questions,omitempty" xml:"intake_questions,omitempty" maxItems:"10"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Price of the product; null for no listed price.
	Price *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
	// False hides the product without deleting it.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// False keeps the product listed (callers can ask about it) but not orderable.
	Orderable *bool `json:"orderable,omitempty" xml:"orderable,omitempty"`
	// Fulfillment options customers may choose when ordering this product.
	AllowedFulfillmentTypes []types.OrderFulfillmentType `json:"allowed_fulfillment_types,omitempty" xml:"allowed_fulfillment_types,omitempty"`
	// Locations where this product is sold; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// Whether this product is counted. A counted product starts with nothing on hand at every location.
	TrackQuantity *bool `json:"track_quantity,omitempty" xml:"track_quantity,omitempty"`
}

func (r *ReceptionCreateProductRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateProductRequest) GetPricingMode() (value *types.PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionCreateProductRequest) GetIntakeQuestions() (value []*types.ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionCreateProductRequest) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionCreateProductRequest) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionCreateProductRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateProductRequest) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionCreateProductRequest) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionCreateProductRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionCreateProductRequest) GetOrderable() (value *bool) {
	if r == nil {
		return
	}
	return r.Orderable
}

func (r *ReceptionCreateProductRequest) GetAllowedFulfillmentTypes() (value []types.OrderFulfillmentType) {
	if r == nil {
		return
	}
	return r.AllowedFulfillmentTypes
}

func (r *ReceptionCreateProductRequest) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionCreateProductRequest) GetTrackQuantity() (value *bool) {
	if r == nil {
		return
	}
	return r.TrackQuantity
}

func (r *ReceptionCreateProductRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateProductRequest) SetPricingMode(pricingMode *types.PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionCreateProductRequest) SetIntakeQuestions(intakeQuestions []*types.ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionCreateProductRequest) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionCreateProductRequest) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionCreateProductRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionCreateProductRequest) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionCreateProductRequest) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionCreateProductRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionCreateProductRequest) SetOrderable(orderable *bool) {
	r.Orderable = orderable
}

func (r *ReceptionCreateProductRequest) SetAllowedFulfillmentTypes(allowedFulfillmentTypes []types.OrderFulfillmentType) {
	r.AllowedFulfillmentTypes = allowedFulfillmentTypes
}

func (r *ReceptionCreateProductRequest) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionCreateProductRequest) SetTrackQuantity(trackQuantity *bool) {
	r.TrackQuantity = trackQuantity
}

func (r *ReceptionCreateProductRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateProductRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateProductRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateProductRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateProductRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateProductRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionUpdateProductRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the product. Returned when you create one; list them with `GET /v1/reception/products`.
	ProductId string `json:"-" url:"-"`
	// Whether a catalog item's price and duration are known up front.
	//
	// - `fixed`: the catalog price and duration apply (the default).
	// - `quote`: neither is known until the owner has seen the job. The
	// receptionist collects intake answers and files a quote request instead
	// of booking a slot or taking an order.
	PricingMode *types.PricingMode `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
	// Questions the agent asks before booking or ordering this item. Required for a useful `pricing_mode='quote'` item -- without them a quote request arrives with nothing to price -- but also valid on fixed-price items that need extra detail (current hair color before a dye job, allergies before a treatment).
	IntakeQuestions []*types.ReceptionIntakeQuestion `json:"intake_questions,omitempty" xml:"intake_questions,omitempty"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty"`
	Name          *StrippedStr `json:"name,omitempty" xml:"name,omitempty"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Price of the product; null for no listed price.
	Price *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
	// False hides the product without deleting it.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// False keeps the product listed (callers can ask about it) but not orderable.
	Orderable *bool `json:"orderable,omitempty" xml:"orderable,omitempty"`
	// Fulfillment options customers may choose when ordering this product.
	AllowedFulfillmentTypes []types.OrderFulfillmentType `json:"allowed_fulfillment_types,omitempty" xml:"allowed_fulfillment_types,omitempty"`
	// Locations where this product is sold; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// Whether this product is counted. Switching it only changes the mode: turning it on keeps whatever is already on hand, turning it off drops the counts. Use the stock endpoints to change the amounts themselves.
	TrackQuantity *bool `json:"track_quantity,omitempty" xml:"track_quantity,omitempty"`
}

func (r *ReceptionUpdateProductRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateProductRequestPatch) GetProductId() (value string) {
	if r == nil {
		return
	}
	return r.ProductId
}

func (r *ReceptionUpdateProductRequestPatch) GetPricingMode() (value *types.PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionUpdateProductRequestPatch) GetIntakeQuestions() (value []*types.ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionUpdateProductRequestPatch) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionUpdateProductRequestPatch) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionUpdateProductRequestPatch) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateProductRequestPatch) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionUpdateProductRequestPatch) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionUpdateProductRequestPatch) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionUpdateProductRequestPatch) GetOrderable() (value *bool) {
	if r == nil {
		return
	}
	return r.Orderable
}

func (r *ReceptionUpdateProductRequestPatch) GetAllowedFulfillmentTypes() (value []types.OrderFulfillmentType) {
	if r == nil {
		return
	}
	return r.AllowedFulfillmentTypes
}

func (r *ReceptionUpdateProductRequestPatch) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionUpdateProductRequestPatch) GetTrackQuantity() (value *bool) {
	if r == nil {
		return
	}
	return r.TrackQuantity
}

func (r *ReceptionUpdateProductRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateProductRequestPatch) SetProductId(productId string) {
	r.ProductId = productId
}

func (r *ReceptionUpdateProductRequestPatch) SetPricingMode(pricingMode *types.PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionUpdateProductRequestPatch) SetIntakeQuestions(intakeQuestions []*types.ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionUpdateProductRequestPatch) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionUpdateProductRequestPatch) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionUpdateProductRequestPatch) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionUpdateProductRequestPatch) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionUpdateProductRequestPatch) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionUpdateProductRequestPatch) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionUpdateProductRequestPatch) SetOrderable(orderable *bool) {
	r.Orderable = orderable
}

func (r *ReceptionUpdateProductRequestPatch) SetAllowedFulfillmentTypes(allowedFulfillmentTypes []types.OrderFulfillmentType) {
	r.AllowedFulfillmentTypes = allowedFulfillmentTypes
}

func (r *ReceptionUpdateProductRequestPatch) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionUpdateProductRequestPatch) SetTrackQuantity(trackQuantity *bool) {
	r.TrackQuantity = trackQuantity
}

func (r *ReceptionUpdateProductRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateProductRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateProductRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateProductRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateProductRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateProductRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listProductsRequestFieldGroup         = big.NewInt(1 << 0)
	listProductsRequestFieldIsActive      = big.NewInt(1 << 1)
	listProductsRequestFieldXiWorkspaceId = big.NewInt(1 << 2)
)

type ListProductsRequest struct {
	Group         *string `json:"-" url:"group,omitempty"`
	IsActive      *bool   `json:"-" url:"is_active,omitempty"`
	XiWorkspaceId string  `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListProductsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetGroup sets the Group field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProductsRequest) SetGroup(group *string) {
	l.Group = group
	l.require(listProductsRequestFieldGroup)
}

// SetIsActive sets the IsActive field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProductsRequest) SetIsActive(isActive *bool) {
	l.IsActive = isActive
	l.require(listProductsRequestFieldIsActive)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProductsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listProductsRequestFieldXiWorkspaceId)
}

var (
	getProductsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetProductsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetProductsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetProductsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getProductsRequestFieldXiWorkspaceId)
}

var (
	deleteProductsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteProductsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteProductsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteProductsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteProductsRequestFieldXiWorkspaceId)
}
