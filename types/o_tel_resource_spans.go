package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelResourceSpans struct {
	Resource        *OTelResource     `json:"resource" xml:"resource" required:"true"`
	ScopeSpans      []*OTelScopeSpans `json:"scopeSpans" xml:"scopeSpans" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelResourceSpans) GetResource() (value *OTelResource) {
	if o == nil {
		return
	}
	return o.Resource
}

func (o *OTelResourceSpans) GetScopeSpans() (value []*OTelScopeSpans) {
	if o == nil {
		return
	}
	return o.ScopeSpans
}

func (o *OTelResourceSpans) SetResource(resource *OTelResource) {
	o.Resource = resource
}

func (o *OTelResourceSpans) SetScopeSpans(scopeSpans []*OTelScopeSpans) {
	o.ScopeSpans = scopeSpans
}

func (o *OTelResourceSpans) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelResourceSpans) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelResourceSpans to string"
	}
	return string(jsonData)
}

func (o *OTelResourceSpans) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OTelResourceSpans
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OTelResourceSpans(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelResourceSpans) MarshalJSON() ([]byte, error) {
	type embed OTelResourceSpans
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
