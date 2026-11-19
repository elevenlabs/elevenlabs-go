package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelSpan struct {
	TraceId           string           `json:"traceId" xml:"traceId" required:"true"`
	SpanId            string           `json:"spanId" xml:"spanId" required:"true"`
	ParentSpanId      *string          `json:"parentSpanId,omitempty" xml:"parentSpanId,omitempty"`
	Name              string           `json:"name" xml:"name" required:"true"`
	Kind              int              `json:"kind" xml:"kind" required:"true"`
	StartTimeUnixNano string           `json:"startTimeUnixNano" xml:"startTimeUnixNano" required:"true"`
	EndTimeUnixNano   string           `json:"endTimeUnixNano" xml:"endTimeUnixNano" required:"true"`
	Attributes        []*OTelAttribute `json:"attributes" xml:"attributes" required:"true"`
	Status            *OTelStatus      `json:"status" xml:"status" required:"true"`
	extraProperties   map[string]interface{}
}

func (o *OTelSpan) GetTraceId() (value string) {
	if o == nil {
		return
	}
	return o.TraceId
}

func (o *OTelSpan) GetSpanId() (value string) {
	if o == nil {
		return
	}
	return o.SpanId
}

func (o *OTelSpan) GetParentSpanId() (value *string) {
	if o == nil {
		return
	}
	return o.ParentSpanId
}

func (o *OTelSpan) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OTelSpan) GetKind() (value int) {
	if o == nil {
		return
	}
	return o.Kind
}

func (o *OTelSpan) GetStartTimeUnixNano() (value string) {
	if o == nil {
		return
	}
	return o.StartTimeUnixNano
}

func (o *OTelSpan) GetEndTimeUnixNano() (value string) {
	if o == nil {
		return
	}
	return o.EndTimeUnixNano
}

func (o *OTelSpan) GetAttributes() (value []*OTelAttribute) {
	if o == nil {
		return
	}
	return o.Attributes
}

func (o *OTelSpan) GetStatus() (value *OTelStatus) {
	if o == nil {
		return
	}
	return o.Status
}

func (o *OTelSpan) SetTraceId(traceId string) {
	o.TraceId = traceId
}

func (o *OTelSpan) SetSpanId(spanId string) {
	o.SpanId = spanId
}

func (o *OTelSpan) SetParentSpanId(parentSpanId *string) {
	o.ParentSpanId = parentSpanId
}

func (o *OTelSpan) SetName(name string) {
	o.Name = name
}

func (o *OTelSpan) SetKind(kind int) {
	o.Kind = kind
}

func (o *OTelSpan) SetStartTimeUnixNano(startTimeUnixNano string) {
	o.StartTimeUnixNano = startTimeUnixNano
}

func (o *OTelSpan) SetEndTimeUnixNano(endTimeUnixNano string) {
	o.EndTimeUnixNano = endTimeUnixNano
}

func (o *OTelSpan) SetAttributes(attributes []*OTelAttribute) {
	o.Attributes = attributes
}

func (o *OTelSpan) SetStatus(status *OTelStatus) {
	o.Status = status
}

func (o *OTelSpan) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelSpan) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelSpan to string"
	}
	return string(jsonData)
}

func (o *OTelSpan) UnmarshalJSON(data []byte) error {
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

func (o *OTelSpan) MarshalJSON() ([]byte, error) {
	type embed OTelSpan
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
