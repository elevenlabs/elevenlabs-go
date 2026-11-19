package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The agent's data collection items that steer follow-up calls. Each defaults
// to the built-in item for its purpose; `null` turns it off.
type DataCollectionConfig struct {
	// The item holding when the contact asked to be called back.
	CallbackTime *string `json:"callback_time,omitempty" xml:"callback_time,omitempty" minLength:"1"`
	// The item holding another number the contact asked to be called on.
	RedirectPhoneNumber *string `json:"redirect_phone_number,omitempty" xml:"redirect_phone_number,omitempty" minLength:"1"`
	// The item that is true when the contact asked not to be called again.
	Stop *string `json:"stop,omitempty" xml:"stop,omitempty" minLength:"1"`
	// Copies data collection results into dynamic variables for later calls: each key is a data collection item ID and its value the variable's name.
	Mappings        map[string]string `json:"mappings,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DataCollectionConfig) GetCallbackTime() (value *string) {
	if d == nil {
		return
	}
	return d.CallbackTime
}

func (d *DataCollectionConfig) GetRedirectPhoneNumber() (value *string) {
	if d == nil {
		return
	}
	return d.RedirectPhoneNumber
}

func (d *DataCollectionConfig) GetStop() (value *string) {
	if d == nil {
		return
	}
	return d.Stop
}

func (d *DataCollectionConfig) GetMappings() (value map[string]string) {
	if d == nil {
		return
	}
	return d.Mappings
}

func (d *DataCollectionConfig) SetCallbackTime(callbackTime *string) {
	d.CallbackTime = callbackTime
}

func (d *DataCollectionConfig) SetRedirectPhoneNumber(redirectPhoneNumber *string) {
	d.RedirectPhoneNumber = redirectPhoneNumber
}

func (d *DataCollectionConfig) SetStop(stop *string) {
	d.Stop = stop
}

func (d *DataCollectionConfig) SetMappings(mappings map[string]string) {
	d.Mappings = mappings
}

func (d *DataCollectionConfig) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DataCollectionConfig) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataCollectionConfig to string"
	}
	return string(jsonData)
}

func (d *DataCollectionConfig) UnmarshalJSON(data []byte) error {
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

func (d *DataCollectionConfig) MarshalJSON() ([]byte, error) {
	type embed DataCollectionConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
