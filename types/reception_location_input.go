package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A single business location stored inline in agent info.
type ReceptionLocationInput struct {
	Name               StrippedStr   `json:"name" xml:"name" required:"true"`
	Address            *StrippedStr  `json:"address,omitempty" xml:"address,omitempty"`
	Notes              *StrippedStr  `json:"notes,omitempty" xml:"notes,omitempty"`
	OperatingHours     *WorkingHours `json:"operating_hours,omitempty" xml:"operating_hours,omitempty"`
	CustomerFacingName *StrippedStr  `json:"customer_facing_name,omitempty" xml:"customer_facing_name,omitempty"`
	// Maximum number of appointments that may overlap at this location at a single instant. Only enforced while the booking page setting concurrent_limit_enabled is on.
	MaxConcurrentAppointments *int   `json:"max_concurrent_appointments,omitempty" xml:"max_concurrent_appointments,omitempty" min:"1" max:"999"`
	Id                        string `json:"id" xml:"id" required:"true"`
	// Unix timestamp when this location was deactivated. None when active.
	DeactivatedAtUnix *int `json:"deactivated_at_unix,omitempty" xml:"deactivated_at_unix,omitempty"`
	// Why this location was deactivated. None when active.
	DeactivationReason *ReceptionLocationDeactivationReason `json:"deactivation_reason,omitempty" xml:"deactivation_reason,omitempty"`
	extraProperties    map[string]interface{}
}

func (r *ReceptionLocationInput) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionLocationInput) GetAddress() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Address
}

func (r *ReceptionLocationInput) GetNotes() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Notes
}

func (r *ReceptionLocationInput) GetOperatingHours() (value *WorkingHours) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *ReceptionLocationInput) GetCustomerFacingName() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.CustomerFacingName
}

func (r *ReceptionLocationInput) GetMaxConcurrentAppointments() (value *int) {
	if r == nil {
		return
	}
	return r.MaxConcurrentAppointments
}

func (r *ReceptionLocationInput) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionLocationInput) GetDeactivatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.DeactivatedAtUnix
}

func (r *ReceptionLocationInput) GetDeactivationReason() (value *ReceptionLocationDeactivationReason) {
	if r == nil {
		return
	}
	return r.DeactivationReason
}

func (r *ReceptionLocationInput) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionLocationInput) SetAddress(address *StrippedStr) {
	r.Address = address
}

func (r *ReceptionLocationInput) SetNotes(notes *StrippedStr) {
	r.Notes = notes
}

func (r *ReceptionLocationInput) SetOperatingHours(operatingHours *WorkingHours) {
	r.OperatingHours = operatingHours
}

func (r *ReceptionLocationInput) SetCustomerFacingName(customerFacingName *StrippedStr) {
	r.CustomerFacingName = customerFacingName
}

func (r *ReceptionLocationInput) SetMaxConcurrentAppointments(maxConcurrentAppointments *int) {
	r.MaxConcurrentAppointments = maxConcurrentAppointments
}

func (r *ReceptionLocationInput) SetId(id string) {
	r.Id = id
}

func (r *ReceptionLocationInput) SetDeactivatedAtUnix(deactivatedAtUnix *int) {
	r.DeactivatedAtUnix = deactivatedAtUnix
}

func (r *ReceptionLocationInput) SetDeactivationReason(deactivationReason *ReceptionLocationDeactivationReason) {
	r.DeactivationReason = deactivationReason
}

func (r *ReceptionLocationInput) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionLocationInput) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionLocationInput to string"
	}
	return string(jsonData)
}

func (r *ReceptionLocationInput) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionLocationInput) MarshalJSON() ([]byte, error) {
	type embed ReceptionLocationInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
