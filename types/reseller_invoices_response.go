package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResellerInvoicesResponse struct {
	Invoices        []*ResellerInvoiceSummary `json:"invoices" xml:"invoices" required:"true"`
	NextCursor      *string                   `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         *bool                     `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ResellerInvoicesResponse) GetInvoices() (value []*ResellerInvoiceSummary) {
	if r == nil {
		return
	}
	return r.Invoices
}

func (r *ResellerInvoicesResponse) GetNextCursor() (value *string) {
	if r == nil {
		return
	}
	return r.NextCursor
}

func (r *ResellerInvoicesResponse) GetHasMore() (value *bool) {
	if r == nil {
		return
	}
	return r.HasMore
}

func (r *ResellerInvoicesResponse) SetInvoices(invoices []*ResellerInvoiceSummary) {
	r.Invoices = invoices
}

func (r *ResellerInvoicesResponse) SetNextCursor(nextCursor *string) {
	r.NextCursor = nextCursor
}

func (r *ResellerInvoicesResponse) SetHasMore(hasMore *bool) {
	r.HasMore = hasMore
}

func (r *ResellerInvoicesResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerInvoicesResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerInvoicesResponse to string"
	}
	return string(jsonData)
}

func (r *ResellerInvoicesResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerInvoicesResponse) MarshalJSON() ([]byte, error) {
	type embed ResellerInvoicesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
