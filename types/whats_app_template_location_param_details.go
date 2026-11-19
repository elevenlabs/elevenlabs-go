package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateLocationParamDetails struct {
	Latitude        string `json:"latitude" xml:"latitude" required:"true"`
	Longitude       string `json:"longitude" xml:"longitude" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	Address         string `json:"address" xml:"address" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateLocationParamDetails) GetLatitude() (value string) {
	if w == nil {
		return
	}
	return w.Latitude
}

func (w *WhatsAppTemplateLocationParamDetails) GetLongitude() (value string) {
	if w == nil {
		return
	}
	return w.Longitude
}

func (w *WhatsAppTemplateLocationParamDetails) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WhatsAppTemplateLocationParamDetails) GetAddress() (value string) {
	if w == nil {
		return
	}
	return w.Address
}

func (w *WhatsAppTemplateLocationParamDetails) SetLatitude(latitude string) {
	w.Latitude = latitude
}

func (w *WhatsAppTemplateLocationParamDetails) SetLongitude(longitude string) {
	w.Longitude = longitude
}

func (w *WhatsAppTemplateLocationParamDetails) SetName(name string) {
	w.Name = name
}

func (w *WhatsAppTemplateLocationParamDetails) SetAddress(address string) {
	w.Address = address
}

func (w *WhatsAppTemplateLocationParamDetails) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateLocationParamDetails) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateLocationParamDetails to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateLocationParamDetails) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateLocationParamDetails
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateLocationParamDetails(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateLocationParamDetails) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateLocationParamDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
