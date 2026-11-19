package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionCreateAssetRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the asset can be booked. Inactive assets keep their booking history.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Custom availability hours for this asset. Null means follow the assigned location's operating_hours (requires location_id).
	AvailableHours *types.WorkingHours `json:"available_hours,omitempty" xml:"available_hours,omitempty"`
	// URL of a photo of the asset.
	PhotoUrl *string `json:"photo_url,omitempty" xml:"photo_url,omitempty" maxLength:"2083" minLength:"1"`
	// Location this asset belongs to; empty means available at all locations.
	LocationId         *string      `json:"location_id,omitempty" xml:"location_id,omitempty"`
	CustomerFacingName *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
}

func (r *ReceptionCreateAssetRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateAssetRequest) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionCreateAssetRequest) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionCreateAssetRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateAssetRequest) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionCreateAssetRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionCreateAssetRequest) GetAvailableHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.AvailableHours
}

func (r *ReceptionCreateAssetRequest) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionCreateAssetRequest) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionCreateAssetRequest) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionCreateAssetRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateAssetRequest) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionCreateAssetRequest) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionCreateAssetRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionCreateAssetRequest) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionCreateAssetRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionCreateAssetRequest) SetAvailableHours(availableHours *types.WorkingHours) {
	r.AvailableHours = availableHours
}

func (r *ReceptionCreateAssetRequest) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionCreateAssetRequest) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionCreateAssetRequest) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionCreateAssetRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateAssetRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateAssetRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateAssetRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateAssetRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateAssetRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionUpdateAssetRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the asset. Returned when you create one; list them with `GET /v1/reception/assets`.
	AssetId string `json:"-" url:"-"`
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*types.ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty"`
	Name          *StrippedStr `json:"name,omitempty" xml:"name,omitempty"`
	Description   *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the asset can be booked. Inactive assets keep their booking history.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Custom availability hours for this asset. Null means follow the assigned location's operating_hours (requires location_id).
	AvailableHours *types.WorkingHours `json:"available_hours,omitempty" xml:"available_hours,omitempty"`
	// URL of a photo of the asset.
	PhotoUrl *string `json:"photo_url,omitempty" xml:"photo_url,omitempty" maxLength:"2083" minLength:"1"`
	// Location this asset belongs to; empty means available at all locations.
	LocationId         *string      `json:"location_id,omitempty" xml:"location_id,omitempty"`
	CustomerFacingName *StrippedStr `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
}

func (r *ReceptionUpdateAssetRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateAssetRequestPatch) GetAssetId() (value string) {
	if r == nil {
		return
	}
	return r.AssetId
}

func (r *ReceptionUpdateAssetRequestPatch) GetGroups() (value []*types.ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionUpdateAssetRequestPatch) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionUpdateAssetRequestPatch) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateAssetRequestPatch) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionUpdateAssetRequestPatch) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionUpdateAssetRequestPatch) GetAvailableHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.AvailableHours
}

func (r *ReceptionUpdateAssetRequestPatch) GetPhotoUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PhotoUrl
}

func (r *ReceptionUpdateAssetRequestPatch) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionUpdateAssetRequestPatch) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionUpdateAssetRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateAssetRequestPatch) SetAssetId(assetId string) {
	r.AssetId = assetId
}

func (r *ReceptionUpdateAssetRequestPatch) SetGroups(groups []*types.ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionUpdateAssetRequestPatch) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionUpdateAssetRequestPatch) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionUpdateAssetRequestPatch) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionUpdateAssetRequestPatch) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionUpdateAssetRequestPatch) SetAvailableHours(availableHours *types.WorkingHours) {
	r.AvailableHours = availableHours
}

func (r *ReceptionUpdateAssetRequestPatch) SetPhotoUrl(photoUrl *string) {
	r.PhotoUrl = photoUrl
}

func (r *ReceptionUpdateAssetRequestPatch) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionUpdateAssetRequestPatch) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionUpdateAssetRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateAssetRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateAssetRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateAssetRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateAssetRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateAssetRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listAssetsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListAssetsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListAssetsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAssetsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listAssetsRequestFieldXiWorkspaceId)
}

var (
	getAssetsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetAssetsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAssetsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAssetsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getAssetsRequestFieldXiWorkspaceId)
}

var (
	deleteAssetsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteAssetsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteAssetsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteAssetsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteAssetsRequestFieldXiWorkspaceId)
}
