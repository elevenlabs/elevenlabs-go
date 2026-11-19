package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionCreateStaffRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the staff member can be booked. Inactive staff keep their booking history.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Email address of the staff member.
	Email *string              `json:"email,omitempty" xml:"email,omitempty"`
	Phone *PhoneNumberOptional `json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether this staff member currently receives SMS booking notifications. Defaults to opted out — a staff member must be explicitly opted in; flipped to False automatically when they reply STOP to the workspace number.
	SmsNotificationsOptedIn *bool `json:"sms_notifications_opted_in,omitempty" xml:"sms_notifications_opted_in,omitempty"`
	// URL of a photo of the staff member.
	PhotoUrl *string `json:"photo_url,omitempty" xml:"photo_url,omitempty" maxLength:"2083" minLength:"1"`
	// Location this staff member is assigned to; empty means available at all locations.
	LocationId *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// Custom working hours for this staff member. Null means follow the assigned location's operating_hours (requires location_id).
	WorkingHours *types.WorkingHours `json:"working_hours,omitempty" xml:"working_hours,omitempty"`
	// Whether customers can see and pick this staff member by name on the booking page. False keeps them bookable, just not shown by name.
	DisplayOnBookingPage *bool        `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	CustomerFacingName   *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
}

func (r *ReceptionCreateStaffRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateStaffRequest) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionCreateStaffRequest) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionCreateStaffRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateStaffRequest) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionCreateStaffRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionCreateStaffRequest) GetEmail() (value *string) {
	if r == nil {
		return
	}
	return r.Email
}

func (r *ReceptionCreateStaffRequest) GetPhone() (value *PhoneNumberOptional) {
	if r == nil {
		return
	}
	return r.Phone
}

func (r *ReceptionCreateStaffRequest) GetSmsNotificationsOptedIn() (value *bool) {
	if r == nil {
		return
	}
	return r.SmsNotificationsOptedIn
}

func (r *ReceptionCreateStaffRequest) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionCreateStaffRequest) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionCreateStaffRequest) GetWorkingHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.WorkingHours
}

func (r *ReceptionCreateStaffRequest) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionCreateStaffRequest) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionCreateStaffRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateStaffRequest) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionCreateStaffRequest) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionCreateStaffRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionCreateStaffRequest) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionCreateStaffRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionCreateStaffRequest) SetEmail(email *string) {
	r.Email = email
}

func (r *ReceptionCreateStaffRequest) SetPhone(phone *PhoneNumberOptional) {
	r.Phone = phone
}

func (r *ReceptionCreateStaffRequest) SetSmsNotificationsOptedIn(smsNotificationsOptedIn *bool) {
	r.SmsNotificationsOptedIn = smsNotificationsOptedIn
}

func (r *ReceptionCreateStaffRequest) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionCreateStaffRequest) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionCreateStaffRequest) SetWorkingHours(workingHours *types.WorkingHours) {
	r.WorkingHours = workingHours
}

func (r *ReceptionCreateStaffRequest) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionCreateStaffRequest) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionCreateStaffRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateStaffRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateStaffRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateStaffRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateStaffRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateStaffRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionUpdateStaffRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the staff member. Returned when you create one; list them with `GET /v1/reception/staff`.
	StaffId string `json:"-" url:"-"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty"`
	Name          *StrippedStr `json:"name,omitempty" xml:"name,omitempty"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the staff member can be booked. Inactive staff keep their booking history.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Email address of the staff member.
	Email *string              `json:"email,omitempty" xml:"email,omitempty"`
	Phone *PhoneNumberOptional `json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether this staff member currently receives SMS booking notifications. Defaults to opted out — a staff member must be explicitly opted in; flipped to False automatically when they reply STOP to the workspace number.
	SmsNotificationsOptedIn *bool `json:"sms_notifications_opted_in,omitempty" xml:"sms_notifications_opted_in,omitempty"`
	// URL of a photo of the staff member.
	PhotoUrl *string `json:"photo_url,omitempty" xml:"photo_url,omitempty" maxLength:"2083" minLength:"1"`
	// Location this staff member is assigned to; empty means available at all locations.
	LocationId *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// Custom working hours for this staff member. Null means follow the assigned location's operating_hours (requires location_id).
	WorkingHours *types.WorkingHours `json:"working_hours,omitempty" xml:"working_hours,omitempty"`
	// Whether customers can see and pick this staff member by name on the booking page. False keeps them bookable, just not shown by name.
	DisplayOnBookingPage *bool        `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	CustomerFacingName   *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
}

func (r *ReceptionUpdateStaffRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateStaffRequestPatch) GetStaffId() (value string) {
	if r == nil {
		return
	}
	return r.StaffId
}

func (r *ReceptionUpdateStaffRequestPatch) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionUpdateStaffRequestPatch) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionUpdateStaffRequestPatch) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateStaffRequestPatch) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionUpdateStaffRequestPatch) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionUpdateStaffRequestPatch) GetEmail() (value *string) {
	if r == nil {
		return
	}
	return r.Email
}

func (r *ReceptionUpdateStaffRequestPatch) GetPhone() (value *PhoneNumberOptional) {
	if r == nil {
		return
	}
	return r.Phone
}

func (r *ReceptionUpdateStaffRequestPatch) GetSmsNotificationsOptedIn() (value *bool) {
	if r == nil {
		return
	}
	return r.SmsNotificationsOptedIn
}

func (r *ReceptionUpdateStaffRequestPatch) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionUpdateStaffRequestPatch) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionUpdateStaffRequestPatch) GetWorkingHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.WorkingHours
}

func (r *ReceptionUpdateStaffRequestPatch) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionUpdateStaffRequestPatch) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionUpdateStaffRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateStaffRequestPatch) SetStaffId(staffId string) {
	r.StaffId = staffId
}

func (r *ReceptionUpdateStaffRequestPatch) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionUpdateStaffRequestPatch) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionUpdateStaffRequestPatch) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionUpdateStaffRequestPatch) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionUpdateStaffRequestPatch) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionUpdateStaffRequestPatch) SetEmail(email *string) {
	r.Email = email
}

func (r *ReceptionUpdateStaffRequestPatch) SetPhone(phone *PhoneNumberOptional) {
	r.Phone = phone
}

func (r *ReceptionUpdateStaffRequestPatch) SetSmsNotificationsOptedIn(smsNotificationsOptedIn *bool) {
	r.SmsNotificationsOptedIn = smsNotificationsOptedIn
}

func (r *ReceptionUpdateStaffRequestPatch) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionUpdateStaffRequestPatch) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionUpdateStaffRequestPatch) SetWorkingHours(workingHours *types.WorkingHours) {
	r.WorkingHours = workingHours
}

func (r *ReceptionUpdateStaffRequestPatch) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionUpdateStaffRequestPatch) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionUpdateStaffRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateStaffRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateStaffRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateStaffRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateStaffRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateStaffRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listStaffRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListStaffRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListStaffRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListStaffRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listStaffRequestFieldXiWorkspaceId)
}

var (
	getStaffRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetStaffRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetStaffRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetStaffRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getStaffRequestFieldXiWorkspaceId)
}

var (
	deleteStaffRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteStaffRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteStaffRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteStaffRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteStaffRequestFieldXiWorkspaceId)
}
