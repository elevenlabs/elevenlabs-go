package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a staff member who can provide services.
type ReceptionStaffResource struct {
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
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
	WorkingHours *WorkingHours `json:"working_hours,omitempty" xml:"working_hours,omitempty"`
	// Whether customers can see and pick this staff member by name on the booking page. False keeps them bookable, just not shown by name.
	DisplayOnBookingPage *bool        `json:"display_on_booking_page,omitempty" xml:"display_on_booking_page,omitempty"`
	CustomerFacingName   *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
	// The staff member's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix   *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionStaffResource) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionStaffResource) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionStaffResource) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionStaffResource) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionStaffResource) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionStaffResource) GetEmail() (value *string) {
	if r == nil {
		return
	}
	return r.Email
}

func (r *ReceptionStaffResource) GetPhone() (value *PhoneNumberOptional) {
	if r == nil {
		return
	}
	return r.Phone
}

func (r *ReceptionStaffResource) GetSmsNotificationsOptedIn() (value *bool) {
	if r == nil {
		return
	}
	return r.SmsNotificationsOptedIn
}

func (r *ReceptionStaffResource) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionStaffResource) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionStaffResource) GetWorkingHours() (value *WorkingHours) {
	if r == nil {
		return
	}
	return r.WorkingHours
}

func (r *ReceptionStaffResource) GetDisplayOnBookingPage() (value *bool) {
	if r == nil {
		return
	}
	return r.DisplayOnBookingPage
}

func (r *ReceptionStaffResource) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionStaffResource) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionStaffResource) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionStaffResource) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionStaffResource) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionStaffResource) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionStaffResource) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionStaffResource) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionStaffResource) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionStaffResource) SetEmail(email *string) {
	r.Email = email
}

func (r *ReceptionStaffResource) SetPhone(phone *PhoneNumberOptional) {
	r.Phone = phone
}

func (r *ReceptionStaffResource) SetSmsNotificationsOptedIn(smsNotificationsOptedIn *bool) {
	r.SmsNotificationsOptedIn = smsNotificationsOptedIn
}

func (r *ReceptionStaffResource) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionStaffResource) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionStaffResource) SetWorkingHours(workingHours *WorkingHours) {
	r.WorkingHours = workingHours
}

func (r *ReceptionStaffResource) SetDisplayOnBookingPage(displayOnBookingPage *bool) {
	r.DisplayOnBookingPage = displayOnBookingPage
}

func (r *ReceptionStaffResource) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionStaffResource) SetId(id string) {
	r.Id = id
}

func (r *ReceptionStaffResource) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionStaffResource) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionStaffResource) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionStaffResource) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionStaffResource to string"
	}
	return string(jsonData)
}

func (r *ReceptionStaffResource) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionStaffResource) MarshalJSON() ([]byte, error) {
	type embed ReceptionStaffResource
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
