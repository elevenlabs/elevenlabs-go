package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OTelTracePayload struct {
	ResourceSpans   []*OTelResourceSpans `json:"resourceSpans" xml:"resourceSpans" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelTracePayload) GetResourceSpans() (value []*OTelResourceSpans) {
	if o == nil {
		return
	}
	return o.ResourceSpans
}

func (o *OTelTracePayload) SetResourceSpans(resourceSpans []*OTelResourceSpans) {
	o.ResourceSpans = resourceSpans
}

func (o *OTelTracePayload) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelTracePayload) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelTracePayload to string"
	}
	return string(jsonData)
}

func (o *OTelTracePayload) UnmarshalJSON(data []byte) error {
	type unmarshaler OTelTracePayload
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OTelTracePayload(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelTracePayload) MarshalJSON() ([]byte, error) {
	type embed OTelTracePayload
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
