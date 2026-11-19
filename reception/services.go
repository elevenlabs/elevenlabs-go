package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionCreateServiceRequest struct {
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
	// Bookable options for this service. Always non-empty; `variants[0]` is the default option shown when callers don't pick a specific variant. Each entry has its own duration and (optional) price, and uniqueness is keyed by duration.
	Variants []*types.ReceptionServiceVariant `json:"variants,omitempty" xml:"variants,omitempty" minItems:"1" maxItems:"20"`
	// What a service books.
	//
	// - `classic`: a staff-delivered service (haircut, consultation, massage, ...).
	// - `rental`: an asset reservation (court, room, equipment).
	// - `group`: seats in a multi-participant session (class, workshop, tour).
	Kind *types.ReceptionServiceKind `json:"kind,omitempty" xml:"kind,omitempty"`
	// Turnover/cleanup time reserved after each booking. Applies to every service kind; default 0 means no buffer.
	BufferMinutes *int `json:"buffer_minutes,omitempty" xml:"buffer_minutes,omitempty" min:"0" max:"1440"`
	// Where a classic service is delivered.
	ServiceLocation *types.ServiceLocationType `json:"service_location,omitempty" xml:"service_location,omitempty"`
	// Optional extras that can be added to a booking of this service.
	AddOns []*types.ReceptionAddOn `json:"add_ons,omitempty" xml:"add_ons,omitempty" maxItems:"10"`
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
	Promotions []*types.ReceptionServicePromotion `json:"promotions,omitempty" xml:"promotions,omitempty" maxItems:"10"`
}

func (r *ReceptionCreateServiceRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateServiceRequest) GetPricingMode() (value *types.PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionCreateServiceRequest) GetIntakeQuestions() (value []*types.ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionCreateServiceRequest) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionCreateServiceRequest) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionCreateServiceRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateServiceRequest) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionCreateServiceRequest) GetVariants() (value []*types.ReceptionServiceVariant) {
	if r == nil {
		return
	}
	return r.Variants
}

func (r *ReceptionCreateServiceRequest) GetKind() (value *types.ReceptionServiceKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionCreateServiceRequest) GetBufferMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.BufferMinutes
}

func (r *ReceptionCreateServiceRequest) GetServiceLocation() (value *types.ServiceLocationType) {
	if r == nil {
		return
	}
	return r.ServiceLocation
}

func (r *ReceptionCreateServiceRequest) GetAddOns() (value []*types.ReceptionAddOn) {
	if r == nil {
		return
	}
	return r.AddOns
}

func (r *ReceptionCreateServiceRequest) GetMaxParticipants() (value *int) {
	if r == nil {
		return
	}
	return r.MaxParticipants
}

func (r *ReceptionCreateServiceRequest) GetCapacity() (value *int) {
	if r == nil {
		return
	}
	return r.Capacity
}

func (r *ReceptionCreateServiceRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionCreateServiceRequest) GetAnyStaffMember() (value *bool) {
	if r == nil {
		return
	}
	return r.AnyStaffMember
}

func (r *ReceptionCreateServiceRequest) GetStaffIds() (value []string) {
	if r == nil {
		return
	}
	return r.StaffIds
}

func (r *ReceptionCreateServiceRequest) GetAssetIds() (value []string) {
	if r == nil {
		return
	}
	return r.AssetIds
}

func (r *ReceptionCreateServiceRequest) GetAdvanceBookingMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.AdvanceBookingMinutes
}

func (r *ReceptionCreateServiceRequest) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionCreateServiceRequest) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionCreateServiceRequest) GetTravelMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.TravelMinutes
}

func (r *ReceptionCreateServiceRequest) GetTravelBefore() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelBefore
}

func (r *ReceptionCreateServiceRequest) GetTravelAfter() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelAfter
}

func (r *ReceptionCreateServiceRequest) GetPromotions() (value []*types.ReceptionServicePromotion) {
	if r == nil {
		return
	}
	return r.Promotions
}

func (r *ReceptionCreateServiceRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateServiceRequest) SetPricingMode(pricingMode *types.PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionCreateServiceRequest) SetIntakeQuestions(intakeQuestions []*types.ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionCreateServiceRequest) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionCreateServiceRequest) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionCreateServiceRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionCreateServiceRequest) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionCreateServiceRequest) SetVariants(variants []*types.ReceptionServiceVariant) {
	r.Variants = variants
}

func (r *ReceptionCreateServiceRequest) SetKind(kind *types.ReceptionServiceKind) {
	r.Kind = kind
}

func (r *ReceptionCreateServiceRequest) SetBufferMinutes(bufferMinutes *int) {
	r.BufferMinutes = bufferMinutes
}

func (r *ReceptionCreateServiceRequest) SetServiceLocation(serviceLocation *types.ServiceLocationType) {
	r.ServiceLocation = serviceLocation
}

func (r *ReceptionCreateServiceRequest) SetAddOns(addOns []*types.ReceptionAddOn) {
	r.AddOns = addOns
}

func (r *ReceptionCreateServiceRequest) SetMaxParticipants(maxParticipants *int) {
	r.MaxParticipants = maxParticipants
}

func (r *ReceptionCreateServiceRequest) SetCapacity(capacity *int) {
	r.Capacity = capacity
}

func (r *ReceptionCreateServiceRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionCreateServiceRequest) SetAnyStaffMember(anyStaffMember *bool) {
	r.AnyStaffMember = anyStaffMember
}

func (r *ReceptionCreateServiceRequest) SetStaffIds(staffIds []string) {
	r.StaffIds = staffIds
}

func (r *ReceptionCreateServiceRequest) SetAssetIds(assetIds []string) {
	r.AssetIds = assetIds
}

func (r *ReceptionCreateServiceRequest) SetAdvanceBookingMinutes(advanceBookingMinutes *int) {
	r.AdvanceBookingMinutes = advanceBookingMinutes
}

func (r *ReceptionCreateServiceRequest) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionCreateServiceRequest) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionCreateServiceRequest) SetTravelMinutes(travelMinutes *int) {
	r.TravelMinutes = travelMinutes
}

func (r *ReceptionCreateServiceRequest) SetTravelBefore(travelBefore *bool) {
	r.TravelBefore = travelBefore
}

func (r *ReceptionCreateServiceRequest) SetTravelAfter(travelAfter *bool) {
	r.TravelAfter = travelAfter
}

func (r *ReceptionCreateServiceRequest) SetPromotions(promotions []*types.ReceptionServicePromotion) {
	r.Promotions = promotions
}

func (r *ReceptionCreateServiceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateServiceRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateServiceRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateServiceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateServiceRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateServiceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionUpdateServiceRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the service. Returned when you create one; list them with `GET /v1/reception/services`.
	ServiceId string `json:"-" url:"-"`
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
	// Bookable options for this service. Always non-empty; `variants[0]` is the default option shown when callers don't pick a specific variant. Each entry has its own duration and (optional) price, and uniqueness is keyed by duration.
	Variants []*types.ReceptionServiceVariant `json:"variants,omitempty" xml:"variants,omitempty"`
	// What a service books.
	//
	// - `classic`: a staff-delivered service (haircut, consultation, massage, ...).
	// - `rental`: an asset reservation (court, room, equipment).
	// - `group`: seats in a multi-participant session (class, workshop, tour).
	Kind *types.ReceptionServiceKind `json:"kind,omitempty" xml:"kind,omitempty"`
	// Turnover/cleanup time reserved after each booking. Applies to every service kind; default 0 means no buffer.
	BufferMinutes *int `json:"buffer_minutes,omitempty" xml:"buffer_minutes,omitempty"`
	// Where a classic service is delivered.
	ServiceLocation *types.ServiceLocationType `json:"service_location,omitempty" xml:"service_location,omitempty"`
	// Optional extras that can be added to a booking of this service.
	AddOns []*types.ReceptionAddOn `json:"add_ons,omitempty" xml:"add_ons,omitempty"`
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
	AdvanceBookingMinutes *int `json:"advance_booking_minutes,omitempty" xml:"advance_booking_minutes,omitempty"`
	// Locations where this service is offered; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// Whether the service is listed on the public booking page.
	DisplayOnBookingPage *bool `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	// Travel time in minutes added before/after appointments for customer-location services. Blocks the staff member's calendar but is not visible to the customer. Ignored when `service_location='business'`.
	TravelMinutes *int `json:"travel_minutes,omitempty" xml:"travel_minutes,omitempty"`
	// Add travel time before the appointment (outbound trip).
	TravelBefore *bool `json:"travel_before,omitempty" xml:"travel_before,omitempty"`
	// Add travel time after the appointment (return trip).
	TravelAfter *bool `json:"travel_after,omitempty" xml:"travel_after,omitempty"`
	// Time-bound discounts applicable to this service or specific variants.
	Promotions []*types.ReceptionServicePromotion `json:"promotions,omitempty" xml:"promotions,omitempty"`
	// Sets the first variant's duration. Ignored when `variants` is sent.
	DurationMinutes *int `json:"duration_minutes,omitempty" xml:"duration_minutes,omitempty" max:"43200"`
	// Sets the first variant's price. Ignored when `variants` is sent.
	Price *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
}

func (r *ReceptionUpdateServiceRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateServiceRequestPatch) GetServiceId() (value string) {
	if r == nil {
		return
	}
	return r.ServiceId
}

func (r *ReceptionUpdateServiceRequestPatch) GetPricingMode() (value *types.PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionUpdateServiceRequestPatch) GetIntakeQuestions() (value []*types.ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionUpdateServiceRequestPatch) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionUpdateServiceRequestPatch) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionUpdateServiceRequestPatch) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateServiceRequestPatch) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionUpdateServiceRequestPatch) GetVariants() (value []*types.ReceptionServiceVariant) {
	if r == nil {
		return
	}
	return r.Variants
}

func (r *ReceptionUpdateServiceRequestPatch) GetKind() (value *types.ReceptionServiceKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionUpdateServiceRequestPatch) GetBufferMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.BufferMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) GetServiceLocation() (value *types.ServiceLocationType) {
	if r == nil {
		return
	}
	return r.ServiceLocation
}

func (r *ReceptionUpdateServiceRequestPatch) GetAddOns() (value []*types.ReceptionAddOn) {
	if r == nil {
		return
	}
	return r.AddOns
}

func (r *ReceptionUpdateServiceRequestPatch) GetMaxParticipants() (value *int) {
	if r == nil {
		return
	}
	return r.MaxParticipants
}

func (r *ReceptionUpdateServiceRequestPatch) GetCapacity() (value *int) {
	if r == nil {
		return
	}
	return r.Capacity
}

func (r *ReceptionUpdateServiceRequestPatch) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionUpdateServiceRequestPatch) GetAnyStaffMember() (value *bool) {
	if r == nil {
		return
	}
	return r.AnyStaffMember
}

func (r *ReceptionUpdateServiceRequestPatch) GetStaffIds() (value []string) {
	if r == nil {
		return
	}
	return r.StaffIds
}

func (r *ReceptionUpdateServiceRequestPatch) GetAssetIds() (value []string) {
	if r == nil {
		return
	}
	return r.AssetIds
}

func (r *ReceptionUpdateServiceRequestPatch) GetAdvanceBookingMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.AdvanceBookingMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionUpdateServiceRequestPatch) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionUpdateServiceRequestPatch) GetTravelMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.TravelMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) GetTravelBefore() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelBefore
}

func (r *ReceptionUpdateServiceRequestPatch) GetTravelAfter() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelAfter
}

func (r *ReceptionUpdateServiceRequestPatch) GetPromotions() (value []*types.ReceptionServicePromotion) {
	if r == nil {
		return
	}
	return r.Promotions
}

func (r *ReceptionUpdateServiceRequestPatch) GetDurationMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.DurationMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionUpdateServiceRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateServiceRequestPatch) SetServiceId(serviceId string) {
	r.ServiceId = serviceId
}

func (r *ReceptionUpdateServiceRequestPatch) SetPricingMode(pricingMode *types.PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionUpdateServiceRequestPatch) SetIntakeQuestions(intakeQuestions []*types.ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionUpdateServiceRequestPatch) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionUpdateServiceRequestPatch) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionUpdateServiceRequestPatch) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionUpdateServiceRequestPatch) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionUpdateServiceRequestPatch) SetVariants(variants []*types.ReceptionServiceVariant) {
	r.Variants = variants
}

func (r *ReceptionUpdateServiceRequestPatch) SetKind(kind *types.ReceptionServiceKind) {
	r.Kind = kind
}

func (r *ReceptionUpdateServiceRequestPatch) SetBufferMinutes(bufferMinutes *int) {
	r.BufferMinutes = bufferMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) SetServiceLocation(serviceLocation *types.ServiceLocationType) {
	r.ServiceLocation = serviceLocation
}

func (r *ReceptionUpdateServiceRequestPatch) SetAddOns(addOns []*types.ReceptionAddOn) {
	r.AddOns = addOns
}

func (r *ReceptionUpdateServiceRequestPatch) SetMaxParticipants(maxParticipants *int) {
	r.MaxParticipants = maxParticipants
}

func (r *ReceptionUpdateServiceRequestPatch) SetCapacity(capacity *int) {
	r.Capacity = capacity
}

func (r *ReceptionUpdateServiceRequestPatch) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionUpdateServiceRequestPatch) SetAnyStaffMember(anyStaffMember *bool) {
	r.AnyStaffMember = anyStaffMember
}

func (r *ReceptionUpdateServiceRequestPatch) SetStaffIds(staffIds []string) {
	r.StaffIds = staffIds
}

func (r *ReceptionUpdateServiceRequestPatch) SetAssetIds(assetIds []string) {
	r.AssetIds = assetIds
}

func (r *ReceptionUpdateServiceRequestPatch) SetAdvanceBookingMinutes(advanceBookingMinutes *int) {
	r.AdvanceBookingMinutes = advanceBookingMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionUpdateServiceRequestPatch) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionUpdateServiceRequestPatch) SetTravelMinutes(travelMinutes *int) {
	r.TravelMinutes = travelMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) SetTravelBefore(travelBefore *bool) {
	r.TravelBefore = travelBefore
}

func (r *ReceptionUpdateServiceRequestPatch) SetTravelAfter(travelAfter *bool) {
	r.TravelAfter = travelAfter
}

func (r *ReceptionUpdateServiceRequestPatch) SetPromotions(promotions []*types.ReceptionServicePromotion) {
	r.Promotions = promotions
}

func (r *ReceptionUpdateServiceRequestPatch) SetDurationMinutes(durationMinutes *int) {
	r.DurationMinutes = durationMinutes
}

func (r *ReceptionUpdateServiceRequestPatch) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionUpdateServiceRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateServiceRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateServiceRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateServiceRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateServiceRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateServiceRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionConfirmServiceUpdateRequest struct {
	// Also update the price on the service's future bookings.
	Reprice *bool `json:"-" url:"reprice,omitempty" explode:"true" serializationStyle:"form" queryParam:"reprice"`
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the service. Returned when you create one; list them with `GET /v1/reception/services`.
	ServiceId string `json:"-" url:"-"`
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
	// Bookable options for this service. Always non-empty; `variants[0]` is the default option shown when callers don't pick a specific variant. Each entry has its own duration and (optional) price, and uniqueness is keyed by duration.
	Variants []*types.ReceptionServiceVariant `json:"variants,omitempty" xml:"variants,omitempty"`
	// What a service books.
	//
	// - `classic`: a staff-delivered service (haircut, consultation, massage, ...).
	// - `rental`: an asset reservation (court, room, equipment).
	// - `group`: seats in a multi-participant session (class, workshop, tour).
	Kind *types.ReceptionServiceKind `json:"kind,omitempty" xml:"kind,omitempty"`
	// Turnover/cleanup time reserved after each booking. Applies to every service kind; default 0 means no buffer.
	BufferMinutes *int `json:"buffer_minutes,omitempty" xml:"buffer_minutes,omitempty"`
	// Where a classic service is delivered.
	ServiceLocation *types.ServiceLocationType `json:"service_location,omitempty" xml:"service_location,omitempty"`
	// Optional extras that can be added to a booking of this service.
	AddOns []*types.ReceptionAddOn `json:"add_ons,omitempty" xml:"add_ons,omitempty"`
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
	AdvanceBookingMinutes *int `json:"advance_booking_minutes,omitempty" xml:"advance_booking_minutes,omitempty"`
	// Locations where this service is offered; empty list means all locations.
	LocationIds []string `json:"location_ids,omitempty" xml:"location_ids,omitempty"`
	// Whether the service is listed on the public booking page.
	DisplayOnBookingPage *bool `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	// Travel time in minutes added before/after appointments for customer-location services. Blocks the staff member's calendar but is not visible to the customer. Ignored when `service_location='business'`.
	TravelMinutes *int `json:"travel_minutes,omitempty" xml:"travel_minutes,omitempty"`
	// Add travel time before the appointment (outbound trip).
	TravelBefore *bool `json:"travel_before,omitempty" xml:"travel_before,omitempty"`
	// Add travel time after the appointment (return trip).
	TravelAfter *bool `json:"travel_after,omitempty" xml:"travel_after,omitempty"`
	// Time-bound discounts applicable to this service or specific variants.
	Promotions []*types.ReceptionServicePromotion `json:"promotions,omitempty" xml:"promotions,omitempty"`
	// Sets the first variant's duration. Ignored when `variants` is sent.
	DurationMinutes *int `json:"duration_minutes,omitempty" xml:"duration_minutes,omitempty" max:"43200"`
	// Sets the first variant's price. Ignored when `variants` is sent.
	Price *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
}

func (r *ReceptionConfirmServiceUpdateRequest) GetReprice() (value *bool) {
	if r == nil {
		return
	}
	return r.Reprice
}

func (r *ReceptionConfirmServiceUpdateRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionConfirmServiceUpdateRequest) GetServiceId() (value string) {
	if r == nil {
		return
	}
	return r.ServiceId
}

func (r *ReceptionConfirmServiceUpdateRequest) GetPricingMode() (value *types.PricingMode) {
	if r == nil {
		return
	}
	return r.PricingMode
}

func (r *ReceptionConfirmServiceUpdateRequest) GetIntakeQuestions() (value []*types.ReceptionIntakeQuestion) {
	if r == nil {
		return
	}
	return r.IntakeQuestions
}

func (r *ReceptionConfirmServiceUpdateRequest) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionConfirmServiceUpdateRequest) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionConfirmServiceUpdateRequest) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionConfirmServiceUpdateRequest) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionConfirmServiceUpdateRequest) GetVariants() (value []*types.ReceptionServiceVariant) {
	if r == nil {
		return
	}
	return r.Variants
}

func (r *ReceptionConfirmServiceUpdateRequest) GetKind() (value *types.ReceptionServiceKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionConfirmServiceUpdateRequest) GetBufferMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.BufferMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) GetServiceLocation() (value *types.ServiceLocationType) {
	if r == nil {
		return
	}
	return r.ServiceLocation
}

func (r *ReceptionConfirmServiceUpdateRequest) GetAddOns() (value []*types.ReceptionAddOn) {
	if r == nil {
		return
	}
	return r.AddOns
}

func (r *ReceptionConfirmServiceUpdateRequest) GetMaxParticipants() (value *int) {
	if r == nil {
		return
	}
	return r.MaxParticipants
}

func (r *ReceptionConfirmServiceUpdateRequest) GetCapacity() (value *int) {
	if r == nil {
		return
	}
	return r.Capacity
}

func (r *ReceptionConfirmServiceUpdateRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionConfirmServiceUpdateRequest) GetAnyStaffMember() (value *bool) {
	if r == nil {
		return
	}
	return r.AnyStaffMember
}

func (r *ReceptionConfirmServiceUpdateRequest) GetStaffIds() (value []string) {
	if r == nil {
		return
	}
	return r.StaffIds
}

func (r *ReceptionConfirmServiceUpdateRequest) GetAssetIds() (value []string) {
	if r == nil {
		return
	}
	return r.AssetIds
}

func (r *ReceptionConfirmServiceUpdateRequest) GetAdvanceBookingMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.AdvanceBookingMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) GetLocationIds() (value []string) {
	if r == nil {
		return
	}
	return r.LocationIds
}

func (r *ReceptionConfirmServiceUpdateRequest) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionConfirmServiceUpdateRequest) GetTravelMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.TravelMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) GetTravelBefore() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelBefore
}

func (r *ReceptionConfirmServiceUpdateRequest) GetTravelAfter() (value *bool) {
	if r == nil {
		return
	}
	return r.TravelAfter
}

func (r *ReceptionConfirmServiceUpdateRequest) GetPromotions() (value []*types.ReceptionServicePromotion) {
	if r == nil {
		return
	}
	return r.Promotions
}

func (r *ReceptionConfirmServiceUpdateRequest) GetDurationMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.DurationMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionConfirmServiceUpdateRequest) SetReprice(reprice *bool) {
	r.Reprice = reprice
}

func (r *ReceptionConfirmServiceUpdateRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionConfirmServiceUpdateRequest) SetServiceId(serviceId string) {
	r.ServiceId = serviceId
}

func (r *ReceptionConfirmServiceUpdateRequest) SetPricingMode(pricingMode *types.PricingMode) {
	r.PricingMode = pricingMode
}

func (r *ReceptionConfirmServiceUpdateRequest) SetIntakeQuestions(intakeQuestions []*types.ReceptionIntakeQuestion) {
	r.IntakeQuestions = intakeQuestions
}

func (r *ReceptionConfirmServiceUpdateRequest) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionConfirmServiceUpdateRequest) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionConfirmServiceUpdateRequest) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionConfirmServiceUpdateRequest) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionConfirmServiceUpdateRequest) SetVariants(variants []*types.ReceptionServiceVariant) {
	r.Variants = variants
}

func (r *ReceptionConfirmServiceUpdateRequest) SetKind(kind *types.ReceptionServiceKind) {
	r.Kind = kind
}

func (r *ReceptionConfirmServiceUpdateRequest) SetBufferMinutes(bufferMinutes *int) {
	r.BufferMinutes = bufferMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) SetServiceLocation(serviceLocation *types.ServiceLocationType) {
	r.ServiceLocation = serviceLocation
}

func (r *ReceptionConfirmServiceUpdateRequest) SetAddOns(addOns []*types.ReceptionAddOn) {
	r.AddOns = addOns
}

func (r *ReceptionConfirmServiceUpdateRequest) SetMaxParticipants(maxParticipants *int) {
	r.MaxParticipants = maxParticipants
}

func (r *ReceptionConfirmServiceUpdateRequest) SetCapacity(capacity *int) {
	r.Capacity = capacity
}

func (r *ReceptionConfirmServiceUpdateRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionConfirmServiceUpdateRequest) SetAnyStaffMember(anyStaffMember *bool) {
	r.AnyStaffMember = anyStaffMember
}

func (r *ReceptionConfirmServiceUpdateRequest) SetStaffIds(staffIds []string) {
	r.StaffIds = staffIds
}

func (r *ReceptionConfirmServiceUpdateRequest) SetAssetIds(assetIds []string) {
	r.AssetIds = assetIds
}

func (r *ReceptionConfirmServiceUpdateRequest) SetAdvanceBookingMinutes(advanceBookingMinutes *int) {
	r.AdvanceBookingMinutes = advanceBookingMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) SetLocationIds(locationIds []string) {
	r.LocationIds = locationIds
}

func (r *ReceptionConfirmServiceUpdateRequest) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionConfirmServiceUpdateRequest) SetTravelMinutes(travelMinutes *int) {
	r.TravelMinutes = travelMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) SetTravelBefore(travelBefore *bool) {
	r.TravelBefore = travelBefore
}

func (r *ReceptionConfirmServiceUpdateRequest) SetTravelAfter(travelAfter *bool) {
	r.TravelAfter = travelAfter
}

func (r *ReceptionConfirmServiceUpdateRequest) SetPromotions(promotions []*types.ReceptionServicePromotion) {
	r.Promotions = promotions
}

func (r *ReceptionConfirmServiceUpdateRequest) SetDurationMinutes(durationMinutes *int) {
	r.DurationMinutes = durationMinutes
}

func (r *ReceptionConfirmServiceUpdateRequest) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionConfirmServiceUpdateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionConfirmServiceUpdateRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionConfirmServiceUpdateRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionConfirmServiceUpdateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionConfirmServiceUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionConfirmServiceUpdateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listServicesRequestFieldStaffId       = big.NewInt(1 << 0)
	listServicesRequestFieldAssetId       = big.NewInt(1 << 1)
	listServicesRequestFieldIsActive      = big.NewInt(1 << 2)
	listServicesRequestFieldKind          = big.NewInt(1 << 3)
	listServicesRequestFieldXiWorkspaceId = big.NewInt(1 << 4)
)

type ListServicesRequest struct {
	StaffId       *string                     `json:"-" url:"staff_id,omitempty"`
	AssetId       *string                     `json:"-" url:"asset_id,omitempty"`
	IsActive      *bool                       `json:"-" url:"is_active,omitempty"`
	Kind          *types.ReceptionServiceKind `json:"-" url:"kind,omitempty"`
	XiWorkspaceId string                      `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListServicesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetStaffId sets the StaffId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListServicesRequest) SetStaffId(staffId *string) {
	l.StaffId = staffId
	l.require(listServicesRequestFieldStaffId)
}

// SetAssetId sets the AssetId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListServicesRequest) SetAssetId(assetId *string) {
	l.AssetId = assetId
	l.require(listServicesRequestFieldAssetId)
}

// SetIsActive sets the IsActive field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListServicesRequest) SetIsActive(isActive *bool) {
	l.IsActive = isActive
	l.require(listServicesRequestFieldIsActive)
}

// SetKind sets the Kind field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListServicesRequest) SetKind(kind *types.ReceptionServiceKind) {
	l.Kind = kind
	l.require(listServicesRequestFieldKind)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListServicesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listServicesRequestFieldXiWorkspaceId)
}

var (
	getServicesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetServicesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetServicesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetServicesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getServicesRequestFieldXiWorkspaceId)
}

var (
	deleteServicesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteServicesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteServicesRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteServicesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteServicesRequestFieldXiWorkspaceId)
}
