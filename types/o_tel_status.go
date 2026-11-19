package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelStatus struct {
	Code            int     `json:"code" xml:"code" required:"true"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OTelStatus) GetCode() (value int) {
	if o == nil {
		return
	}
	return o.Code
}

func (o *OTelStatus) GetMessage() (value *string) {
	if o == nil {
		return
	}
	return o.Message
}

func (o *OTelStatus) SetCode(code int) {
	o.Code = code
}

func (o *OTelStatus) SetMessage(message *string) {
	o.Message = message
}

func (o *OTelStatus) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelStatus) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelStatus to string"
	}
	return string(jsonData)
}

func (o *OTelStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelStatus) MarshalJSON() ([]byte, error) {
	type embed OTelStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
