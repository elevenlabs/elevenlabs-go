package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ResellerUpcomingPaymentsResponse struct {
	Payments        []*ResellerUpcomingPayment `json:"payments" xml:"payments" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResellerUpcomingPaymentsResponse) GetPayments() (value []*ResellerUpcomingPayment) {
	if r == nil {
		return
	}
	return r.Payments
}

func (r *ResellerUpcomingPaymentsResponse) SetPayments(payments []*ResellerUpcomingPayment) {
	r.Payments = payments
}

func (r *ResellerUpcomingPaymentsResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerUpcomingPaymentsResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerUpcomingPaymentsResponse to string"
	}
	return string(jsonData)
}

func (r *ResellerUpcomingPaymentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResellerUpcomingPaymentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResellerUpcomingPaymentsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerUpcomingPaymentsResponse) MarshalJSON() ([]byte, error) {
	type embed ResellerUpcomingPaymentsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
