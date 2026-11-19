package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Scrape-reachable subset of a service (classic, rental or group).
//
// The extraction pipeline uses one LLM prompt per kind (so the model sees
// kind-specific guidance), but every code path produces this single shape:
// onboarding, imports and the unified “smb_services“ collection all
// discriminate by the “kind“ field.
type ReceptionServiceScrapable struct {
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
	// Bookable options for this service. Always non-empty; `variants[0]` is the default option shown when callers don't pick a specific variant. Each entry has its own duration and (optional) price, and uniqueness is keyed by duration.
	Variants []*ReceptionServiceVariant `json:"variants,omitempty" xml:"variants,omitempty" minItems:"1" maxItems:"20"`
	// What a service books.
	//
	// - `classic`: a staff-delivered service (haircut, consultation, massage, ...).
	// - `rental`: an asset reservation (court, room, equipment).
	// - `group`: seats in a multi-participant session (class, workshop, tour).
	Kind *ReceptionServiceKind `json:"kind,omitempty" xml:"kind,omitempty"`
	// Turnover/cleanup time reserved after each booking. Applies to every service kind; default 0 means no buffer.
	BufferMinutes *int `json:"buffer_minutes,omitempty" xml:"buffer_minutes,omitempty" min:"0" max:"1440"`
	// Where a classic service is delivered.
	ServiceLocation *ServiceLocationType `json:"service_location,omitempty" xml:"service_location,omitempty"`
	// Optional extras that can be added to a booking of this service.
	AddOns []*ReceptionAddOn `json:"add_ons,omitempty" xml:"add_ons,omitempty" maxItems:"10"`
	// Capacity for `kind='group'`: how many seats a single session can sell across all registrations. Must be null for classic and rental services (those are always exclusive to one client per booking).
	MaxParticipants *int `json:"max_participants,omitempty" xml:"max_participants,omitempty" min:"2" max:"200"`
	extraProperties map[string]interface{}
}

func (r *ReceptionServiceScrapable) GetPricingMode() (value *PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionServiceScrapable) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionServiceScrapable) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionServiceScrapable) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionServiceScrapable) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionServiceScrapable) GetVariants() (value []*ReceptionServiceVariant) {
	if r == nil {
		return
	}
	return r.Variants
}

func (r *ReceptionServiceScrapable) GetKind() (value *ReceptionServiceKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionServiceScrapable) GetBufferMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.BufferMinutes
}

func (r *ReceptionServiceScrapable) GetServiceLocation() (value *ServiceLocationType) {
	if r == nil {
		return
	}
	return r.ServiceLocation
}

func (r *ReceptionServiceScrapable) GetAddOns() (value []*ReceptionAddOn) {
	if r == nil {
		return
	}
	return r.AddOns
}

func (r *ReceptionServiceScrapable) GetMaxParticipants() (value *int) {
	if r == nil {
		return
	}
	return r.MaxParticipants
}

func (r *ReceptionServiceScrapable) SetPricingMode(pricingMode *PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionServiceScrapable) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionServiceScrapable) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionServiceScrapable) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionServiceScrapable) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionServiceScrapable) SetVariants(variants []*ReceptionServiceVariant) {
	r.Variants = variants
}

func (r *ReceptionServiceScrapable) SetKind(kind *ReceptionServiceKind) {
	r.Kind = kind
}

func (r *ReceptionServiceScrapable) SetBufferMinutes(bufferMinutes *int) {
	r.BufferMinutes = bufferMinutes
}

func (r *ReceptionServiceScrapable) SetServiceLocation(serviceLocation *ServiceLocationType) {
	r.ServiceLocation = serviceLocation
}

func (r *ReceptionServiceScrapable) SetAddOns(addOns []*ReceptionAddOn) {
	r.AddOns = addOns
}

func (r *ReceptionServiceScrapable) SetMaxParticipants(maxParticipants *int) {
	r.MaxParticipants = maxParticipants
}

func (r *ReceptionServiceScrapable) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionServiceScrapable) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionServiceScrapable to string"
	}
	return string(jsonData)
}

func (r *ReceptionServiceScrapable) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionServiceScrapable) MarshalJSON() ([]byte, error) {
	type embed ReceptionServiceScrapable
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
