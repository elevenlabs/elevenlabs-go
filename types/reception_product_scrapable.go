package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Scrape-reachable subset of a product (retail/catalog item).
type ReceptionProductScrapable struct {
	// Whether a catalog item's price and duration are known up front.
	//
	// - `fixed`: the catalog price and duration apply (the default).
	// - `quote`: neither is known until the owner has seen the job. The
	// receptionist collects intake answers and files a quote request instead
	// of booking a slot or taking an order.
	PricingMode *PricingMode `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Price of the product; null for no listed price.
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
	extraProperties map[string]interface{}
}

func (r *ReceptionProductScrapable) GetPricingMode() (value *PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionProductScrapable) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionProductScrapable) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionProductScrapable) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionProductScrapable) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionProductScrapable) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionProductScrapable) SetPricingMode(pricingMode *PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionProductScrapable) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionProductScrapable) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionProductScrapable) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionProductScrapable) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionProductScrapable) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionProductScrapable) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionProductScrapable) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionProductScrapable to string"
	}
	return string(jsonData)
}

func (r *ReceptionProductScrapable) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionProductScrapable) MarshalJSON() ([]byte, error) {
	type embed ReceptionProductScrapable
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
