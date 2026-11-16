package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Request model for creating or updating a location.
type ReceptionLocationRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string       `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	Name          StrippedStr  `json:"name" xml:"name" required:"true"`
	Address       *StrippedStr `json:"address,omitempty" xml:"address,omitempty"`
	Notes         *StrippedStr `json:"notes,omitempty" xml:"notes,omitempty"`
	// Per-location operating hours. On create, omit to have the server prefill from the most recently added active location.
	OperatingHours     *types.WorkingHours `json:"operating_hours,omitempty" xml:"operating_hours,omitempty"`
	CustomerFacingName *StrippedStr        `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
	// Maximum number of appointments that may overlap at this location at a single instant. Only enforced while the booking page setting concurrent_limit_enabled is on.
	MaxConcurrentAppointments *int `json:"max_concurrent_appointments,omitempty" xml:"max_concurrent_appointments,omitempty" min:"1" max:"999"`
	// Whether the location is active. Inactive locations take no bookings.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
}

func (r *ReceptionLocationRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionLocationRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionLocationRequest) GetAddress() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Address
}

func (r *ReceptionLocationRequest) GetNotes() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Notes
}

func (r *ReceptionLocationRequest) GetOperatingHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *ReceptionLocationRequest) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionLocationRequest) GetMaxConcurrentAppointments() (value *int) {
	if r == nil {
		return
	}
	return r.MaxConcurrentAppointments
}

func (r *ReceptionLocationRequest) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionLocationRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionLocationRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionLocationRequest) SetAddress(address *StrippedStr) {
	r.Address = address
}

func (r *ReceptionLocationRequest) SetNotes(notes *StrippedStr) {
	r.Notes = notes
}

func (r *ReceptionLocationRequest) SetOperatingHours(operatingHours *types.WorkingHours) {
	r.OperatingHours = operatingHours
}

func (r *ReceptionLocationRequest) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionLocationRequest) SetMaxConcurrentAppointments(maxConcurrentAppointments *int) {
	r.MaxConcurrentAppointments = maxConcurrentAppointments
}

func (r *ReceptionLocationRequest) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionLocationRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionLocationRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionLocationRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionLocationRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionLocationRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionLocationRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionLocationRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the location. Returned when you create one; list them with `GET /v1/reception/locations`.
	LocationId string       `json:"-" url:"-"`
	Name       *StrippedStr `json:"name,omitempty" xml:"name,omitempty"`
	Address    *StrippedStr `json:"address,omitempty" xml:"address,omitempty"`
	Notes      *StrippedStr `json:"notes,omitempty" xml:"notes,omitempty"`
	// Per-location operating hours. On create, omit to have the server prefill from the most recently added active location.
	OperatingHours     *types.WorkingHours `json:"operating_hours,omitempty" xml:"operating_hours,omitempty"`
	CustomerFacingName *StrippedStr        `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
	// Maximum number of appointments that may overlap at this location at a single instant. Only enforced while the booking page setting concurrent_limit_enabled is on.
	MaxConcurrentAppointments *int `json:"max_concurrent_appointments,omitempty" xml:"max_concurrent_appointments,omitempty"`
	// Whether the location is active. Inactive locations take no bookings.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
}

func (r *ReceptionLocationRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionLocationRequestPatch) GetLocationId() (value string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionLocationRequestPatch) GetName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionLocationRequestPatch) GetAddress() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Address
}

func (r *ReceptionLocationRequestPatch) GetNotes() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Notes
}

func (r *ReceptionLocationRequestPatch) GetOperatingHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *ReceptionLocationRequestPatch) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionLocationRequestPatch) GetMaxConcurrentAppointments() (value *int) {
	if r == nil {
		return
	}
	return r.MaxConcurrentAppointments
}

func (r *ReceptionLocationRequestPatch) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionLocationRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionLocationRequestPatch) SetLocationId(locationId string) {
	r.LocationId = locationId
}

func (r *ReceptionLocationRequestPatch) SetName(name *StrippedStr) {
	r.Name = name
}

func (r *ReceptionLocationRequestPatch) SetAddress(address *StrippedStr) {
	r.Address = address
}

func (r *ReceptionLocationRequestPatch) SetNotes(notes *StrippedStr) {
	r.Notes = notes
}

func (r *ReceptionLocationRequestPatch) SetOperatingHours(operatingHours *types.WorkingHours) {
	r.OperatingHours = operatingHours
}

func (r *ReceptionLocationRequestPatch) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionLocationRequestPatch) SetMaxConcurrentAppointments(maxConcurrentAppointments *int) {
	r.MaxConcurrentAppointments = maxConcurrentAppointments
}

func (r *ReceptionLocationRequestPatch) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionLocationRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionLocationRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionLocationRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionLocationRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionLocationRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionLocationRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listLocationsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListLocationsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListLocationsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLocationsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listLocationsRequestFieldXiWorkspaceId)
}

var (
	getLocationsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetLocationsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetLocationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLocationsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getLocationsRequestFieldXiWorkspaceId)
}

var (
	deleteLocationsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteLocationsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteLocationsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteLocationsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteLocationsRequestFieldXiWorkspaceId)
}
