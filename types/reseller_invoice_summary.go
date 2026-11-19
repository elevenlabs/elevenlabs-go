package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One charge to the reseller's card, paying for one workspace's invoice.
type ResellerInvoiceSummary struct {
	InvoiceId string `json:"invoice_id" xml:"invoice_id" required:"true"`
	// The workspace this charge paid for.
	WorkspaceId   *string `json:"workspace_id" xml:"workspace_id" required:"true"`
	WorkspaceName string  `json:"workspace_name" xml:"workspace_name" required:"true"`
	// `open` while the card is still being retried; `void` once retries ran out and the workspace's plan was ended.
	Status ResellerInvoiceStatus `json:"status" xml:"status" required:"true"`
	// Invoice total, in the currency's minor unit.
	AmountCents int    `json:"amount_cents" xml:"amount_cents" required:"true"`
	Currency    string `json:"currency" xml:"currency" required:"true"`
	CreatedUnix int    `json:"created_unix" xml:"created_unix" required:"true"`
	PaidUnix    *int   `json:"paid_unix,omitempty" xml:"paid_unix,omitempty"`
	// Stripe's hosted payment page, only while open.
	PayUrl          *string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	PdfUrl          *string `json:"pdf_url,omitempty" xml:"pdf_url,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ResellerInvoiceSummary) GetInvoiceId() (value string) {
	if r == nil {
		return
	}
	return r.InvoiceId
}

func (r *ResellerInvoiceSummary) GetWorkspaceId() (value *string) {
	if r == nil {
		return
	}
	return r.WorkspaceId
}

func (r *ResellerInvoiceSummary) GetWorkspaceName() (value string) {
	if r == nil {
		return
	}
	return r.WorkspaceName
}

func (r *ResellerInvoiceSummary) GetStatus() (value ResellerInvoiceStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *ResellerInvoiceSummary) GetAmountCents() (value int) {
	if r == nil {
		return
	}
	return r.AmountCents
}

func (r *ResellerInvoiceSummary) GetCurrency() (value string) {
	if r == nil {
		return
	}
	return r.Currency
}

func (r *ResellerInvoiceSummary) GetCreatedUnix() (value int) {
	if r == nil {
		return
	}
	return r.CreatedUnix
}

func (r *ResellerInvoiceSummary) GetPaidUnix() (value *int) {
	if r == nil {
		return
	}
	return r.PaidUnix
}

func (r *ResellerInvoiceSummary) GetPayUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PayUrl
}

func (r *ResellerInvoiceSummary) GetPdfUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PdfUrl
}

func (r *ResellerInvoiceSummary) SetInvoiceId(invoiceId string) {
	r.InvoiceId = invoiceId
}

func (r *ResellerInvoiceSummary) SetWorkspaceId(workspaceId *string) {
	r.WorkspaceId = workspaceId
}

func (r *ResellerInvoiceSummary) SetWorkspaceName(workspaceName string) {
	r.WorkspaceName = workspaceName
}

func (r *ResellerInvoiceSummary) SetStatus(status ResellerInvoiceStatus) {
	r.Status = status
}

func (r *ResellerInvoiceSummary) SetAmountCents(amountCents int) {
	r.AmountCents = amountCents
}

func (r *ResellerInvoiceSummary) SetCurrency(currency string) {
	r.Currency = currency
}

func (r *ResellerInvoiceSummary) SetCreatedUnix(createdUnix int) {
	r.CreatedUnix = createdUnix
}

func (r *ResellerInvoiceSummary) SetPaidUnix(paidUnix *int) {
	r.PaidUnix = paidUnix
}

func (r *ResellerInvoiceSummary) SetPayUrl(payUrl *string) {
	r.PayUrl = payUrl
}

func (r *ResellerInvoiceSummary) SetPdfUrl(pdfUrl *string) {
	r.PdfUrl = pdfUrl
}

func (r *ResellerInvoiceSummary) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerInvoiceSummary) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerInvoiceSummary to string"
	}
	return string(jsonData)
}

func (r *ResellerInvoiceSummary) UnmarshalJSON(data []byte) error {
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

func (r *ResellerInvoiceSummary) MarshalJSON() ([]byte, error) {
	type embed ResellerInvoiceSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
