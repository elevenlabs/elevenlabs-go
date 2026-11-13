package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OrderDeliverablesResponse struct {
	// The list of delivered files for the order. Empty if the order is not yet completed.
	Deliverables    []*DeliverableInfo `json:"deliverables" xml:"deliverables" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OrderDeliverablesResponse) GetDeliverables() (value []*DeliverableInfo) {
	if o == nil {
		return
	}
	return o.Deliverables
}

func (o *OrderDeliverablesResponse) SetDeliverables(deliverables []*DeliverableInfo) {
	o.Deliverables = deliverables
}

func (o *OrderDeliverablesResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrderDeliverablesResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrderDeliverablesResponse to string"
	}
	return string(jsonData)
}

func (o *OrderDeliverablesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler OrderDeliverablesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrderDeliverablesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrderDeliverablesResponse) MarshalJSON() ([]byte, error) {
	type embed OrderDeliverablesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
