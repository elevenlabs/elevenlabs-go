package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a physical resource that can be booked (rooms, equipment, vehicles).
type ReceptionAssetResource struct {
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the asset can be booked. Inactive assets keep their booking history.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Custom availability hours for this asset. Null means follow the assigned location's operating_hours (requires location_id).
	AvailableHours *WorkingHours `json:"available_hours,omitempty" xml:"available_hours,omitempty"`
	// URL of a photo of the asset.
	PhotoUrl *string `json:"photo_url,omitempty" xml:"photo_url,omitempty" maxLength:"2083" minLength:"1"`
	// Location this asset belongs to; empty means available at all locations.
	LocationId         *string      `json:"location_id,omitempty" xml:"location_id,omitempty"`
	CustomerFacingName *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
	// The asset's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix   *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionAssetResource) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionAssetResource) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionAssetResource) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionAssetResource) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionAssetResource) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionAssetResource) GetAvailableHours() (value *WorkingHours) {
	if r == nil {
		return
	}
	return r.AvailableHours
}

func (r *ReceptionAssetResource) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionAssetResource) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionAssetResource) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionAssetResource) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionAssetResource) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionAssetResource) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionAssetResource) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionAssetResource) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionAssetResource) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionAssetResource) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionAssetResource) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionAssetResource) SetAvailableHours(availableHours *WorkingHours) {
	r.AvailableHours = availableHours
}

func (r *ReceptionAssetResource) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionAssetResource) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionAssetResource) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionAssetResource) SetId(id string) {
	r.Id = id
}

func (r *ReceptionAssetResource) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionAssetResource) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionAssetResource) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionAssetResource) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionAssetResource to string"
	}
	return string(jsonData)
}

func (r *ReceptionAssetResource) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionAssetResource) MarshalJSON() ([]byte, error) {
	type embed ReceptionAssetResource
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
