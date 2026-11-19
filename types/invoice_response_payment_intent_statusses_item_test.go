package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInvoiceResponsePaymentIntentStatussesItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant InvoiceResponsePaymentIntentStatussesItem
		wire    string
	}
	cases := []enumCase{
		{variant: InvoiceResponsePaymentIntentStatussesItemCanceled, wire: "\"canceled\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemProcessing, wire: "\"processing\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemRequiresAction, wire: "\"requires_action\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemRequiresCapture, wire: "\"requires_capture\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemRequiresConfirmation, wire: "\"requires_confirmation\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemRequiresPaymentMethod, wire: "\"requires_payment_method\""},
		{variant: InvoiceResponsePaymentIntentStatussesItemSucceeded, wire: "\"succeeded\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded InvoiceResponsePaymentIntentStatussesItem
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
