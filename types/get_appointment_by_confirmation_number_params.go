package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Look up an appointment by the booking confirmation number the caller quotes.
//
// The confirmation number is the 8-character code shown on the booking
// confirmation page after a “#“. Callers may read it back with or
// without the leading “#“ and with varied spacing; the tool normalizes
// the input and does a prefix match on the stored calendar item id.
type GetAppointmentByConfirmationNumberParams struct {
	extraProperties map[string]interface{}
}

func (g *GetAppointmentByConfirmationNumberParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAppointmentByConfirmationNumberParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAppointmentByConfirmationNumberParams to string"
	}
	return string(jsonData)
}

func (g *GetAppointmentByConfirmationNumberParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAppointmentByConfirmationNumberParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAppointmentByConfirmationNumberParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAppointmentByConfirmationNumberParams) MarshalJSON() ([]byte, error) {
	type embed GetAppointmentByConfirmationNumberParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
