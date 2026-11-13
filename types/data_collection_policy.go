package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DataCollectionPolicy struct {
	CallbackTime        *string `json:"callback_time,omitempty" xml:"callback_time,omitempty"`
	RedirectPhoneNumber *string `json:"redirect_phone_number,omitempty" xml:"redirect_phone_number,omitempty"`
	Stop                *string `json:"stop,omitempty" xml:"stop,omitempty"`
	// Copies data collection results into dynamic variables for later calls: each key is a data collection item ID and its value the variable's name.
	Mappings        map[string]string `json:"mappings,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DataCollectionPolicy) GetCallbackTime() (value *string) {
	if d == nil {
		return
	}
	return d.CallbackTime
}

func (d *DataCollectionPolicy) GetRedirectPhoneNumber() (value *string) {
	if d == nil {
		return
	}
	return d.RedirectPhoneNumber
}

func (d *DataCollectionPolicy) GetStop() (value *string) {
	if d == nil {
		return
	}
	return d.Stop
}

func (d *DataCollectionPolicy) GetMappings() (value map[string]string) {
	if d == nil {
		return
	}
	return d.Mappings
}

func (d *DataCollectionPolicy) SetCallbackTime(callbackTime *string) {
	d.CallbackTime = callbackTime
}

func (d *DataCollectionPolicy) SetRedirectPhoneNumber(redirectPhoneNumber *string) {
	d.RedirectPhoneNumber = redirectPhoneNumber
}

func (d *DataCollectionPolicy) SetStop(stop *string) {
	d.Stop = stop
}

func (d *DataCollectionPolicy) SetMappings(mappings map[string]string) {
	d.Mappings = mappings
}

func (d *DataCollectionPolicy) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DataCollectionPolicy) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataCollectionPolicy to string"
	}
	return string(jsonData)
}

func (d *DataCollectionPolicy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DataCollectionPolicy) MarshalJSON() ([]byte, error) {
	type embed DataCollectionPolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
