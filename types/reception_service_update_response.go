package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReceptionServiceUpdateResponse struct {
	// The updated service, or null when the update was held back because it reprices existing bookings.
	Service *ReceptionService `json:"service,omitempty" xml:"service,omitempty"`
	// True when the update changes the price of upcoming bookings and was not saved; send it again to the confirm-update endpoint to apply it.
	HasAffectedEvents bool `json:"has_affected_events" xml:"has_affected_events" required:"true"`
	extraProperties   map[string]interface{}
}

func (r *ReceptionServiceUpdateResponse) GetService() (value *ReceptionService) {
	if r == nil {
		return
	}
	return r.Service
}

func (r *ReceptionServiceUpdateResponse) GetHasAffectedEvents() (value bool) {
	if r == nil {
		return
	}
	return r.HasAffectedEvents
}

func (r *ReceptionServiceUpdateResponse) SetService(service *ReceptionService) {
	r.Service = service
}

func (r *ReceptionServiceUpdateResponse) SetHasAffectedEvents(hasAffectedEvents bool) {
	r.HasAffectedEvents = hasAffectedEvents
}

func (r *ReceptionServiceUpdateResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionServiceUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionServiceUpdateResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionServiceUpdateResponse) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionServiceUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionServiceUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
