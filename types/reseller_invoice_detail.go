package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResellerInvoiceDetail struct {
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
	PayUrl          *string                `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	PdfUrl          *string                `json:"pdf_url,omitempty" xml:"pdf_url,omitempty"`
	Lines           []*ResellerInvoiceLine `json:"lines" xml:"lines" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResellerInvoiceDetail) GetInvoiceId() (value string) {
	if r == nil {
		return
	}
	return r.InvoiceId
}

func (r *ResellerInvoiceDetail) GetWorkspaceId() (value *string) {
	if r == nil {
		return
	}
	return r.WorkspaceId
}

func (r *ResellerInvoiceDetail) GetWorkspaceName() (value string) {
	if r == nil {
		return
	}
	return r.WorkspaceName
}

func (r *ResellerInvoiceDetail) GetStatus() (value ResellerInvoiceStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *ResellerInvoiceDetail) GetAmountCents() (value int) {
	if r == nil {
		return
	}
	return r.AmountCents
}

func (r *ResellerInvoiceDetail) GetCurrency() (value string) {
	if r == nil {
		return
	}
	return r.Currency
}

func (r *ResellerInvoiceDetail) GetCreatedUnix() (value int) {
	if r == nil {
		return
	}
	return r.CreatedUnix
}

func (r *ResellerInvoiceDetail) GetPaidUnix() (value *int) {
	if r == nil {
		return
	}
	return r.PaidUnix
}

func (r *ResellerInvoiceDetail) GetPayUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PayUrl
}

func (r *ResellerInvoiceDetail) GetPdfUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PdfUrl
}

func (r *ResellerInvoiceDetail) GetLines() (value []*ResellerInvoiceLine) {
	if r == nil {
		return
	}
	return r.Lines
}

func (r *ResellerInvoiceDetail) SetInvoiceId(invoiceId string) {
	r.InvoiceId = invoiceId
}

func (r *ResellerInvoiceDetail) SetWorkspaceId(workspaceId *string) {
	r.WorkspaceId = workspaceId
}

func (r *ResellerInvoiceDetail) SetWorkspaceName(workspaceName string) {
	r.WorkspaceName = workspaceName
}

func (r *ResellerInvoiceDetail) SetStatus(status ResellerInvoiceStatus) {
	r.Status = status
}

func (r *ResellerInvoiceDetail) SetAmountCents(amountCents int) {
	r.AmountCents = amountCents
}

func (r *ResellerInvoiceDetail) SetCurrency(currency string) {
	r.Currency = currency
}

func (r *ResellerInvoiceDetail) SetCreatedUnix(createdUnix int) {
	r.CreatedUnix = createdUnix
}

func (r *ResellerInvoiceDetail) SetPaidUnix(paidUnix *int) {
	r.PaidUnix = paidUnix
}

func (r *ResellerInvoiceDetail) SetPayUrl(payUrl *string) {
	r.PayUrl = payUrl
}

func (r *ResellerInvoiceDetail) SetPdfUrl(pdfUrl *string) {
	r.PdfUrl = pdfUrl
}

func (r *ResellerInvoiceDetail) SetLines(lines []*ResellerInvoiceLine) {
	r.Lines = lines
}

func (r *ResellerInvoiceDetail) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerInvoiceDetail) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerInvoiceDetail to string"
	}
	return string(jsonData)
}

func (r *ResellerInvoiceDetail) UnmarshalJSON(data []byte) error {
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

func (r *ResellerInvoiceDetail) MarshalJSON() ([]byte, error) {
	type embed ResellerInvoiceDetail
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
