package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type RecipientDeliveryCounts struct {
	Sent            *int `json:"sent,omitempty" xml:"sent,omitempty"`
	Delivered       *int `json:"delivered,omitempty" xml:"delivered,omitempty"`
	Read            *int `json:"read,omitempty" xml:"read,omitempty"`
	Replied         *int `json:"replied,omitempty" xml:"replied,omitempty"`
	Failed          *int `json:"failed,omitempty" xml:"failed,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RecipientDeliveryCounts) GetSent() (value *int) {
	if r == nil {
		return
	}
	return r.Sent
}

func (r *RecipientDeliveryCounts) GetDelivered() (value *int) {
	if r == nil {
		return
	}
	return r.Delivered
}

func (r *RecipientDeliveryCounts) GetRead() (value *int) {
	if r == nil {
		return
	}
	return r.Read
}

func (r *RecipientDeliveryCounts) GetReplied() (value *int) {
	if r == nil {
		return
	}
	return r.Replied
}

func (r *RecipientDeliveryCounts) GetFailed() (value *int) {
	if r == nil {
		return
	}
	return r.Failed
}

func (r *RecipientDeliveryCounts) SetSent(sent *int) {
	r.Sent = sent
}

func (r *RecipientDeliveryCounts) SetDelivered(delivered *int) {
	r.Delivered = delivered
}

func (r *RecipientDeliveryCounts) SetRead(read *int) {
	r.Read = read
}

func (r *RecipientDeliveryCounts) SetReplied(replied *int) {
	r.Replied = replied
}

func (r *RecipientDeliveryCounts) SetFailed(failed *int) {
	r.Failed = failed
}

func (r *RecipientDeliveryCounts) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RecipientDeliveryCounts) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RecipientDeliveryCounts to string"
	}
	return string(jsonData)
}

func (r *RecipientDeliveryCounts) UnmarshalJSON(data []byte) error {
	type unmarshaler RecipientDeliveryCounts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecipientDeliveryCounts(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RecipientDeliveryCounts) MarshalJSON() ([]byte, error) {
	type embed RecipientDeliveryCounts
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
