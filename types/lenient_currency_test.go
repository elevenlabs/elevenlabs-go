package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLenientCurrencyVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant LenientCurrency
		wire    string
	}
	cases := []enumCase{
		{variant: LenientCurrencyUsd, wire: "\"USD\""},
		{variant: LenientCurrencyEur, wire: "\"EUR\""},
		{variant: LenientCurrencyGbp, wire: "\"GBP\""},
		{variant: LenientCurrencyPln, wire: "\"PLN\""},
		{variant: LenientCurrencyCad, wire: "\"CAD\""},
		{variant: LenientCurrencyAud, wire: "\"AUD\""},
		{variant: LenientCurrencyChf, wire: "\"CHF\""},
		{variant: LenientCurrencySek, wire: "\"SEK\""},
		{variant: LenientCurrencyNok, wire: "\"NOK\""},
		{variant: LenientCurrencyDkk, wire: "\"DKK\""},
		{variant: LenientCurrencyIsk, wire: "\"ISK\""},
		{variant: LenientCurrencyCzk, wire: "\"CZK\""},
		{variant: LenientCurrencyHuf, wire: "\"HUF\""},
		{variant: LenientCurrencyRon, wire: "\"RON\""},
		{variant: LenientCurrencyBgn, wire: "\"BGN\""},
		{variant: LenientCurrencyJpy, wire: "\"JPY\""},
		{variant: LenientCurrencyCny, wire: "\"CNY\""},
		{variant: LenientCurrencyTwd, wire: "\"TWD\""},
		{variant: LenientCurrencyInr, wire: "\"INR\""},
		{variant: LenientCurrencyBrl, wire: "\"BRL\""},
		{variant: LenientCurrencyMxn, wire: "\"MXN\""},
		{variant: LenientCurrencyKrw, wire: "\"KRW\""},
		{variant: LenientCurrencySgd, wire: "\"SGD\""},
		{variant: LenientCurrencyHkd, wire: "\"HKD\""},
		{variant: LenientCurrencyNzd, wire: "\"NZD\""},
		{variant: LenientCurrencyZar, wire: "\"ZAR\""},
		{variant: LenientCurrencyAed, wire: "\"AED\""},
		{variant: LenientCurrencySar, wire: "\"SAR\""},
		{variant: LenientCurrencyQar, wire: "\"QAR\""},
		{variant: LenientCurrencyKwd, wire: "\"KWD\""},
		{variant: LenientCurrencyIls, wire: "\"ILS\""},
		{variant: LenientCurrencyTry, wire: "\"TRY\""},
		{variant: LenientCurrencyThb, wire: "\"THB\""},
		{variant: LenientCurrencyIdr, wire: "\"IDR\""},
		{variant: LenientCurrencyMyr, wire: "\"MYR\""},
		{variant: LenientCurrencyPhp, wire: "\"PHP\""},
		{variant: LenientCurrencyVnd, wire: "\"VND\""},
		{variant: LenientCurrencyBdt, wire: "\"BDT\""},
		{variant: LenientCurrencyRub, wire: "\"RUB\""},
		{variant: LenientCurrencyUah, wire: "\"UAH\""},
		{variant: LenientCurrencyClp, wire: "\"CLP\""},
		{variant: LenientCurrencyCop, wire: "\"COP\""},
		{variant: LenientCurrencyArs, wire: "\"ARS\""},
		{variant: LenientCurrencyPen, wire: "\"PEN\""},
		{variant: LenientCurrencyEgp, wire: "\"EGP\""},
		{variant: LenientCurrencyNgn, wire: "\"NGN\""},
		{variant: LenientCurrencyKes, wire: "\"KES\""},
		{variant: LenientCurrencyPkr, wire: "\"PKR\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded LenientCurrency
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
