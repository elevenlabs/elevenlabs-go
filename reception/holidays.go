package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionCreateHolidayRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Display name of the closure or time off.
	Name string `json:"name" xml:"name" required:"true"`
	// Represents a calendar date as individual numeric components.
	StartDate *types.DateModel `json:"start_date" xml:"start_date" required:"true"`
	// Represents a calendar date as individual numeric components.
	EndDate *types.DateModel `json:"end_date" xml:"end_date" required:"true"`
	// Whether the entry closes the target for the whole range or replaces its working hours.
	HolidayType *types.HolidayType `json:"holiday_type,omitempty" xml:"holiday_type,omitempty"`
	// Replacement working hours when holiday_type is 'override'.
	OverrideHours []*types.TimeRangeModel `json:"override_hours,omitempty" xml:"override_hours,omitempty"`
	// What a time-off entry blocks: a staff member, an asset, or a location.
	TimeOffType types.TimeOffType `json:"time_off_type" xml:"time_off_type" required:"true"`
	// The staff member or asset ID; required when time_off_type is 'staff' or 'asset'.
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// For LOCATION type: the location this closure applies to.
	LocationId *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// Optional note on why the closure or time off happens, up to 500 characters.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (r *ReceptionCreateHolidayRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateHolidayRequest) GetName() (value string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateHolidayRequest) GetStartDate() (value *types.DateModel) {
	if r == nil {
		return
	}
	return r.StartDate
}

func (r *ReceptionCreateHolidayRequest) GetEndDate() (value *types.DateModel) {
	if r == nil {
		return
	}
	return r.EndDate
}

func (r *ReceptionCreateHolidayRequest) GetHolidayType() (value *types.HolidayType) {
	if r == nil {
		return
	}
	return r.HolidayType
}

func (r *ReceptionCreateHolidayRequest) GetOverrideHours() (value []*types.TimeRangeModel) {
	if r == nil {
		return
	}
	return r.OverrideHours
}

func (r *ReceptionCreateHolidayRequest) GetTimeOffType() (value types.TimeOffType) {
	if r == nil {
		return
	}
	return r.TimeOffType
}

func (r *ReceptionCreateHolidayRequest) GetResourceId() (value *string) {
	if r == nil {
		return
	}
	return r.ResourceId
}

func (r *ReceptionCreateHolidayRequest) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionCreateHolidayRequest) GetReason() (value *string) {
	if r == nil {
		return
	}
	return r.Reason
}

func (r *ReceptionCreateHolidayRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateHolidayRequest) SetName(name string) {
	r.Name = name
}

func (r *ReceptionCreateHolidayRequest) SetStartDate(startDate *types.DateModel) {
	r.StartDate = startDate
}

func (r *ReceptionCreateHolidayRequest) SetEndDate(endDate *types.DateModel) {
	r.EndDate = endDate
}

func (r *ReceptionCreateHolidayRequest) SetHolidayType(holidayType *types.HolidayType) {
	r.HolidayType = holidayType
}

func (r *ReceptionCreateHolidayRequest) SetOverrideHours(overrideHours []*types.TimeRangeModel) {
	r.OverrideHours = overrideHours
}

func (r *ReceptionCreateHolidayRequest) SetTimeOffType(timeOffType types.TimeOffType) {
	r.TimeOffType = timeOffType
}

func (r *ReceptionCreateHolidayRequest) SetResourceId(resourceId *string) {
	r.ResourceId = resourceId
}

func (r *ReceptionCreateHolidayRequest) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionCreateHolidayRequest) SetReason(reason *string) {
	r.Reason = reason
}

func (r *ReceptionCreateHolidayRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateHolidayRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateHolidayRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateHolidayRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateHolidayRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateHolidayRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type ReceptionUpdateHolidayRequestPatch struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the holiday or time-off entry. Returned when you create one; list them with `GET /v1/reception/holidays`.
	HolidayId string `json:"-" url:"-"`
	// Display name of the closure or time off.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Represents a calendar date as individual numeric components.
	StartDate *types.DateModel `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// Represents a calendar date as individual numeric components.
	EndDate *types.DateModel `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// Whether the entry closes the target for the whole range or replaces its working hours.
	HolidayType *types.HolidayType `json:"holiday_type,omitempty" xml:"holiday_type,omitempty"`
	// Replacement working hours when holiday_type is 'override'.
	OverrideHours []*types.TimeRangeModel `json:"override_hours,omitempty" xml:"override_hours,omitempty"`
	// What a time-off entry blocks: a staff member, an asset, or a location.
	TimeOffType *types.TimeOffType `json:"time_off_type,omitempty" xml:"time_off_type,omitempty"`
	// The staff member or asset ID; required when time_off_type is 'staff' or 'asset'.
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// For LOCATION type: the location this closure applies to.
	LocationId *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// Optional note on why the closure or time off happens, up to 500 characters.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (r *ReceptionUpdateHolidayRequestPatch) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateHolidayRequestPatch) GetHolidayId() (value string) {
	if r == nil {
		return
	}
	return r.HolidayId
}

func (r *ReceptionUpdateHolidayRequestPatch) GetName() (value *string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateHolidayRequestPatch) GetStartDate() (value *types.DateModel) {
	if r == nil {
		return
	}
	return r.StartDate
}

func (r *ReceptionUpdateHolidayRequestPatch) GetEndDate() (value *types.DateModel) {
	if r == nil {
		return
	}
	return r.EndDate
}

func (r *ReceptionUpdateHolidayRequestPatch) GetHolidayType() (value *types.HolidayType) {
	if r == nil {
		return
	}
	return r.HolidayType
}

func (r *ReceptionUpdateHolidayRequestPatch) GetOverrideHours() (value []*types.TimeRangeModel) {
	if r == nil {
		return
	}
	return r.OverrideHours
}

func (r *ReceptionUpdateHolidayRequestPatch) GetTimeOffType() (value *types.TimeOffType) {
	if r == nil {
		return
	}
	return r.TimeOffType
}

func (r *ReceptionUpdateHolidayRequestPatch) GetResourceId() (value *string) {
	if r == nil {
		return
	}
	return r.ResourceId
}

func (r *ReceptionUpdateHolidayRequestPatch) GetLocationId() (value *string) {
	if r == nil {
		return
	}
	return r.LocationId
}

func (r *ReceptionUpdateHolidayRequestPatch) GetReason() (value *string) {
	if r == nil {
		return
	}
	return r.Reason
}

func (r *ReceptionUpdateHolidayRequestPatch) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateHolidayRequestPatch) SetHolidayId(holidayId string) {
	r.HolidayId = holidayId
}

func (r *ReceptionUpdateHolidayRequestPatch) SetName(name *string) {
	r.Name = name
}

func (r *ReceptionUpdateHolidayRequestPatch) SetStartDate(startDate *types.DateModel) {
	r.StartDate = startDate
}

func (r *ReceptionUpdateHolidayRequestPatch) SetEndDate(endDate *types.DateModel) {
	r.EndDate = endDate
}

func (r *ReceptionUpdateHolidayRequestPatch) SetHolidayType(holidayType *types.HolidayType) {
	r.HolidayType = holidayType
}

func (r *ReceptionUpdateHolidayRequestPatch) SetOverrideHours(overrideHours []*types.TimeRangeModel) {
	r.OverrideHours = overrideHours
}

func (r *ReceptionUpdateHolidayRequestPatch) SetTimeOffType(timeOffType *types.TimeOffType) {
	r.TimeOffType = timeOffType
}

func (r *ReceptionUpdateHolidayRequestPatch) SetResourceId(resourceId *string) {
	r.ResourceId = resourceId
}

func (r *ReceptionUpdateHolidayRequestPatch) SetLocationId(locationId *string) {
	r.LocationId = locationId
}

func (r *ReceptionUpdateHolidayRequestPatch) SetReason(reason *string) {
	r.Reason = reason
}

func (r *ReceptionUpdateHolidayRequestPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateHolidayRequestPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateHolidayRequestPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateHolidayRequestPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateHolidayRequestPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateHolidayRequestPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listHolidaysRequestFieldStartUnix     = big.NewInt(1 << 0)
	listHolidaysRequestFieldEndUnix       = big.NewInt(1 << 1)
	listHolidaysRequestFieldTimeOffType   = big.NewInt(1 << 2)
	listHolidaysRequestFieldResourceId    = big.NewInt(1 << 3)
	listHolidaysRequestFieldLocationId    = big.NewInt(1 << 4)
	listHolidaysRequestFieldXiWorkspaceId = big.NewInt(1 << 5)
)

type ListHolidaysRequest struct {
	StartUnix     *int               `json:"-" url:"start_unix,omitempty"`
	EndUnix       *int               `json:"-" url:"end_unix,omitempty"`
	TimeOffType   *types.TimeOffType `json:"-" url:"time_off_type,omitempty"`
	ResourceId    *string            `json:"-" url:"resource_id,omitempty"`
	LocationId    *string            `json:"-" url:"location_id,omitempty"`
	XiWorkspaceId string             `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListHolidaysRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetStartUnix sets the StartUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetStartUnix(startUnix *int) {
	l.StartUnix = startUnix
	l.require(listHolidaysRequestFieldStartUnix)
}

// SetEndUnix sets the EndUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetEndUnix(endUnix *int) {
	l.EndUnix = endUnix
	l.require(listHolidaysRequestFieldEndUnix)
}

// SetTimeOffType sets the TimeOffType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetTimeOffType(timeOffType *types.TimeOffType) {
	l.TimeOffType = timeOffType
	l.require(listHolidaysRequestFieldTimeOffType)
}

// SetResourceId sets the ResourceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetResourceId(resourceId *string) {
	l.ResourceId = resourceId
	l.require(listHolidaysRequestFieldResourceId)
}

// SetLocationId sets the LocationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetLocationId(locationId *string) {
	l.LocationId = locationId
	l.require(listHolidaysRequestFieldLocationId)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListHolidaysRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listHolidaysRequestFieldXiWorkspaceId)
}

var (
	getHolidaysRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetHolidaysRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetHolidaysRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetHolidaysRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getHolidaysRequestFieldXiWorkspaceId)
}

var (
	deleteHolidaysRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteHolidaysRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteHolidaysRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteHolidaysRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteHolidaysRequestFieldXiWorkspaceId)
}
