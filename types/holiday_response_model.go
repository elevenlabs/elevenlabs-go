package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A closure or time-off entry.
type HolidayResponseModel struct {
	// Display name of the closure or time off.
	Name string `json:"name" xml:"name" required:"true"`
	// Represents a calendar date as individual numeric components.
	StartDate *DateModel `json:"start_date" xml:"start_date" required:"true"`
	// Represents a calendar date as individual numeric components.
	EndDate *DateModel `json:"end_date" xml:"end_date" required:"true"`
	// Whether the entry closes the target for the whole range or replaces its working hours.
	HolidayType *HolidayType `json:"holiday_type,omitempty" xml:"holiday_type,omitempty"`
	// Replacement working hours when holiday_type is 'override'.
	OverrideHours []*TimeRangeModel `json:"override_hours,omitempty" xml:"override_hours,omitempty"`
	// What a time-off entry blocks: a staff member, an asset, or a location.
	TimeOffType TimeOffType `json:"time_off_type" xml:"time_off_type" required:"true"`
	// The staff member or asset ID; required when time_off_type is 'staff' or 'asset'.
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// For LOCATION type: the location this closure applies to.
	LocationId *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// Optional note on why the closure or time off happens, up to 500 characters.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The closure or time-off entry's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// When the entry was created, as a Unix timestamp in seconds.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// When the entry was last updated, as a Unix timestamp in seconds.
	UpdatedAtUnix   int `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	extraProperties map[string]interface{}
}

func (h *HolidayResponseModel) GetName() (value string) {
	if h == nil {
		return
	}
	return h.Name
}

func (h *HolidayResponseModel) GetStartDate() (value *DateModel) {
	if h == nil {
		return
	}
	return h.StartDate
}

func (h *HolidayResponseModel) GetEndDate() (value *DateModel) {
	if h == nil {
		return
	}
	return h.EndDate
}

func (h *HolidayResponseModel) GetHolidayType() (value *HolidayType) {
	if h == nil {
		return
	}
	return h.HolidayType
}

func (h *HolidayResponseModel) GetOverrideHours() (value []*TimeRangeModel) {
	if h == nil {
		return
	}
	return h.OverrideHours
}

func (h *HolidayResponseModel) GetTimeOffType() (value TimeOffType) {
	if h == nil {
		return
	}
	return h.TimeOffType
}

func (h *HolidayResponseModel) GetResourceId() (value *string) {
	if h == nil {
		return
	}
	return h.ResourceId
}

func (h *HolidayResponseModel) GetLocationId() (value *string) {
	if h == nil {
		return
	}
	return h.LocationId
}

func (h *HolidayResponseModel) GetReason() (value *string) {
	if h == nil {
		return
	}
	return h.Reason
}

func (h *HolidayResponseModel) GetId() (value string) {
	if h == nil {
		return
	}
	return h.Id
}

func (h *HolidayResponseModel) GetCreatedAtUnix() (value int) {
	if h == nil {
		return
	}
	return h.CreatedAtUnix
}

func (h *HolidayResponseModel) GetUpdatedAtUnix() (value int) {
	if h == nil {
		return
	}
	return h.UpdatedAtUnix
}

func (h *HolidayResponseModel) SetName(name string) {
	h.Name = name
}

func (h *HolidayResponseModel) SetStartDate(startDate *DateModel) {
	h.StartDate = startDate
}

func (h *HolidayResponseModel) SetEndDate(endDate *DateModel) {
	h.EndDate = endDate
}

func (h *HolidayResponseModel) SetHolidayType(holidayType *HolidayType) {
	h.HolidayType = holidayType
}

func (h *HolidayResponseModel) SetOverrideHours(overrideHours []*TimeRangeModel) {
	h.OverrideHours = overrideHours
}

func (h *HolidayResponseModel) SetTimeOffType(timeOffType TimeOffType) {
	h.TimeOffType = timeOffType
}

func (h *HolidayResponseModel) SetResourceId(resourceId *string) {
	h.ResourceId = resourceId
}

func (h *HolidayResponseModel) SetLocationId(locationId *string) {
	h.LocationId = locationId
}

func (h *HolidayResponseModel) SetReason(reason *string) {
	h.Reason = reason
}

func (h *HolidayResponseModel) SetId(id string) {
	h.Id = id
}

func (h *HolidayResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	h.CreatedAtUnix = createdAtUnix
}

func (h *HolidayResponseModel) SetUpdatedAtUnix(updatedAtUnix int) {
	h.UpdatedAtUnix = updatedAtUnix
}

func (h *HolidayResponseModel) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HolidayResponseModel) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HolidayResponseModel to string"
	}
	return string(jsonData)
}

func (h *HolidayResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, h); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HolidayResponseModel) MarshalJSON() ([]byte, error) {
	type embed HolidayResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
