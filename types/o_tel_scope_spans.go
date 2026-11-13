package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelScopeSpans struct {
	Scope           *OTelScope  `json:"scope" xml:"scope" required:"true"`
	Spans           []*OTelSpan `json:"spans" xml:"spans" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelScopeSpans) GetScope() (value *OTelScope) {
	if o == nil {
		return
	}
	return o.Scope
}

func (o *OTelScopeSpans) GetSpans() (value []*OTelSpan) {
	if o == nil {
		return
	}
	return o.Spans
}

func (o *OTelScopeSpans) SetScope(scope *OTelScope) {
	o.Scope = scope
}

func (o *OTelScopeSpans) SetSpans(spans []*OTelSpan) {
	o.Spans = spans
}

func (o *OTelScopeSpans) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelScopeSpans) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelScopeSpans to string"
	}
	return string(jsonData)
}

func (o *OTelScopeSpans) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OTelScopeSpans
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OTelScopeSpans(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelScopeSpans) MarshalJSON() ([]byte, error) {
	type embed OTelScopeSpans
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
