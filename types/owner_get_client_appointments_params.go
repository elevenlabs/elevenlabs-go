package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Appointments with the booking channel, internal notes and intake answers.
type OwnerGetClientAppointmentsParams struct {
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	extraProperties  map[string]interface{}
}

func (o *OwnerGetClientAppointmentsParams) GetIncludeCancelled() (value *bool) {
	if o == nil {
		return
	}
	return o.IncludeCancelled
}

func (o *OwnerGetClientAppointmentsParams) SetIncludeCancelled(includeCancelled *bool) {
	o.IncludeCancelled = includeCancelled
}

func (o *OwnerGetClientAppointmentsParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerGetClientAppointmentsParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerGetClientAppointmentsParams to string"
	}
	return string(jsonData)
}

func (o *OwnerGetClientAppointmentsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerGetClientAppointmentsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerGetClientAppointmentsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerGetClientAppointmentsParams) MarshalJSON() ([]byte, error) {
	type embed OwnerGetClientAppointmentsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
