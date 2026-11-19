package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a product, menu item, or catalog item.
//
// Used for restaurants (menu), e-commerce (products), salons (retail items), etc.
type ReceptionProduct struct {
	// Whether a catalog item's price and duration are known up front.
	//
	// - `fixed`: the catalog price and duration apply (the default).
	// - `quote`: neither is known until the owner has seen the job. The
	// receptionist collects intake answers and files a quote request instead
	// of booking a slot or taking an order.
	PricingMode *PricingMode `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
	// Questions the agent asks before booking or ordering this item. Required for a useful `pricing_mode='quote'` item -- without them a quote request arrives with nothing to price -- but also valid on fixed-price items that need extra detail (current hair color before a dye job, allergies before a treatment).
	IntakeQuestions []*ReceptionIntakeQuestion `json:"intake_questions,omitempty" xml:"intake_questions,omitempty" maxItems:"10"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
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
	AllowedFulfillmentTypes []OrderFulfillmentType `json:"allowed_fulfillment_types,omitempty" xml:"allowed_fulfillment_types,omitempty"`
	// Locations where this product is sold; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// On-hand quantity by location. None means unlimited. A missing location key means zero.
	Stock map[string]*int `json:"stock,omitempty"`
	// The product's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix   *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionProduct) GetPricingMode() (value *PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionProduct) GetIntakeQuestions() (value []*ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionProduct) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionProduct) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionProduct) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionProduct) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionProduct) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionProduct) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionProduct) GetOrderable() (value *bool) {
	if r == nil {
		return
	}
	return r.Orderable
}

func (r *ReceptionProduct) GetAllowedFulfillmentTypes() (value []OrderFulfillmentType) {
	if r == nil {
		return
	}
	return r.AllowedFulfillmentTypes
}

func (r *ReceptionProduct) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionProduct) GetStock() (value map[string]*int) {
	if r == nil {
		return
	}
	return r.Stock
}

func (r *ReceptionProduct) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionProduct) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionProduct) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionProduct) SetPricingMode(pricingMode *PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionProduct) SetIntakeQuestions(intakeQuestions []*ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionProduct) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionProduct) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionProduct) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionProduct) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionProduct) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionProduct) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionProduct) SetOrderable(orderable *bool) {
	r.Orderable = orderable
}

func (r *ReceptionProduct) SetAllowedFulfillmentTypes(allowedFulfillmentTypes []OrderFulfillmentType) {
	r.AllowedFulfillmentTypes = allowedFulfillmentTypes
}

func (r *ReceptionProduct) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionProduct) SetStock(stock map[string]*int) {
	r.Stock = stock
}

func (r *ReceptionProduct) SetId(id string) {
	r.Id = id
}

func (r *ReceptionProduct) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionProduct) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionProduct) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionProduct) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionProduct to string"
	}
	return string(jsonData)
}

func (r *ReceptionProduct) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionProduct) MarshalJSON() ([]byte, error) {
	type embed ReceptionProduct
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
