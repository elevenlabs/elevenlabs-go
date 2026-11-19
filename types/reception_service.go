package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a bookable service (e.g., haircut, massage, consultation).
type ReceptionService struct {
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
	// Bookable options for this service; `variants[0]` is the default.
	Variants []*ReceptionServiceVariant `json:"variants" xml:"variants" required:"true" minItems:"1" maxItems:"20"`
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
	// Not used to limit bookings; set `max_participants` on group services instead.
	Capacity *int `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// Whether the service is offered; inactive services can't be booked.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Whether any staff member can deliver the service. When false, only the staff in `staff_ids` can. Ignored for rentals.
	AnyStaffMember *bool `json:"any_staff_member,omitempty" xml:"any_staff_member,omitempty"`
	// Staff assigned to this service; must be empty for kind='rental'.
	StaffIds []string `json:"staff_ids,omitempty" xml:"staff_ids,omitempty"`
	// Assets, such as rooms or equipment, this service can be booked on.
	AssetIds []string `json:"asset_ids,omitempty" xml:"asset_ids,omitempty"`
	// How far into the future customer-facing channels (public booking page, voice agent) may book this service. Applies to every kind.
	AdvanceBookingMinutes *int `json:"advance_booking_minutes,omitempty" xml:"advance_booking_minutes,omitempty" min:"0" max:"525600"`
	// Locations where this service is offered; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// Whether the service is listed on the public booking page.
	DisplayOnBookingPage *bool `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	// Travel time in minutes added before/after appointments for customer-location services. Blocks the staff member's calendar but is not visible to the customer. Ignored when `service_location='business'`.
	TravelMinutes *int `json:"travel_minutes,omitempty" xml:"travel_minutes,omitempty" min:"0" max:"480"`
	// Add travel time before the appointment (outbound trip).
	TravelBefore *bool `json:"travel_before,omitempty" xml:"travel_before,omitempty"`
	// Add travel time after the appointment (return trip).
	TravelAfter *bool `json:"travel_after,omitempty" xml:"travel_after,omitempty"`
	// Time-bound discounts applicable to this service or specific variants.
	Promotions []*ReceptionServicePromotion `json:"promotions,omitempty" xml:"promotions,omitempty" maxItems:"10"`
	// The service's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	// Duration of the first variant.
	DurationMinutes *int `json:"duration_minutes,omitempty" xml:"duration_minutes,omitempty"`
	// Price of the first variant.
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionService) GetPricingMode() (value *PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionService) GetIntakeQuestions() (value []*ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionService) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionService) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionService) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionService) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionService) GetVariants() (value []*ReceptionServiceVariant) {
	if r == nil {
		return
	}
	return r.Variants
}

func (r *ReceptionService) GetKind() (value *ReceptionServiceKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionService) GetBufferMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.BufferMinutes
}

func (r *ReceptionService) GetServiceLocation() (value *ServiceLocationType) {
	if r == nil {
		return
	}
	return r.ServiceLocation
}

func (r *ReceptionService) GetAddOns() (value []*ReceptionAddOn) {
	if r == nil {
		return
	}
	return r.AddOns
}

func (r *ReceptionService) GetMaxParticipants() (value *int) {
	if r == nil {
		return
	}
	return r.MaxParticipants
}

func (r *ReceptionService) GetCapacity() (value *int) {
	if r == nil {
		return
	}
	return r.Capacity
}

func (r *ReceptionService) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionService) GetAnyStaffMember() (value *bool) {
	if r == nil {
		return
	}
	return r.AnyStaffMember
}

func (r *ReceptionService) GetStaffIds() (value []string) {
	if r == nil {
		return
	}
	return r.StaffIds
}

func (r *ReceptionService) GetAssetIds() (value []string) {
	if r == nil {
		return
	}
	return r.AssetIds
}

func (r *ReceptionService) GetAdvanceBookingMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.AdvanceBookingMinutes
}

func (r *ReceptionService) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionService) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionService) GetTravelMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.TravelMinutes
}

func (r *ReceptionService) GetTravelBefore() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelBefore
}

func (r *ReceptionService) GetTravelAfter() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelAfter
}

func (r *ReceptionService) GetPromotions() (value []*ReceptionServicePromotion) {
	if r == nil {
		return
	}
	return r.Promotions
}

func (r *ReceptionService) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionService) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionService) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionService) GetDurationMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.DurationMinutes
}

func (r *ReceptionService) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionService) SetPricingMode(pricingMode *PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionService) SetIntakeQuestions(intakeQuestions []*ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionService) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionService) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionService) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionService) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionService) SetVariants(variants []*ReceptionServiceVariant) {
	r.Variants = variants
}

func (r *ReceptionService) SetKind(kind *ReceptionServiceKind) {
	r.Kind = kind
}

func (r *ReceptionService) SetBufferMinutes(bufferMinutes *int) {
	r.BufferMinutes = bufferMinutes
}

func (r *ReceptionService) SetServiceLocation(serviceLocation *ServiceLocationType) {
	r.ServiceLocation = serviceLocation
}

func (r *ReceptionService) SetAddOns(addOns []*ReceptionAddOn) {
	r.AddOns = addOns
}

func (r *ReceptionService) SetMaxParticipants(maxParticipants *int) {
	r.MaxParticipants = maxParticipants
}

func (r *ReceptionService) SetCapacity(capacity *int) {
	r.Capacity = capacity
}

func (r *ReceptionService) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionService) SetAnyStaffMember(anyStaffMember *bool) {
	r.AnyStaffMember = anyStaffMember
}

func (r *ReceptionService) SetStaffIds(staffIds []string) {
	r.StaffIds = staffIds
}

func (r *ReceptionService) SetAssetIds(assetIds []string) {
	r.AssetIds = assetIds
}

func (r *ReceptionService) SetAdvanceBookingMinutes(advanceBookingMinutes *int) {
	r.AdvanceBookingMinutes = advanceBookingMinutes
}

func (r *ReceptionService) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionService) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionService) SetTravelMinutes(travelMinutes *int) {
	r.TravelMinutes = travelMinutes
}

func (r *ReceptionService) SetTravelBefore(travelBefore *bool) {
	r.TravelBefore = travelBefore
}

func (r *ReceptionService) SetTravelAfter(travelAfter *bool) {
	r.TravelAfter = travelAfter
}

func (r *ReceptionService) SetPromotions(promotions []*ReceptionServicePromotion) {
	r.Promotions = promotions
}

func (r *ReceptionService) SetId(id string) {
	r.Id = id
}

func (r *ReceptionService) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionService) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionService) SetDurationMinutes(durationMinutes *int) {
	r.DurationMinutes = durationMinutes
}

func (r *ReceptionService) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionService) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionService) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionService to string"
	}
	return string(jsonData)
}

func (r *ReceptionService) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionService) MarshalJSON() ([]byte, error) {
	type embed ReceptionService
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
