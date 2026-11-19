package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestReceptionAgentInfoResponseCurrencyVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ReceptionAgentInfoResponseCurrency
		wire    string
	}
	cases := []enumCase{
		{variant: ReceptionAgentInfoResponseCurrencyUsd, wire: "\"USD\""},
		{variant: ReceptionAgentInfoResponseCurrencyEur, wire: "\"EUR\""},
		{variant: ReceptionAgentInfoResponseCurrencyGbp, wire: "\"GBP\""},
		{variant: ReceptionAgentInfoResponseCurrencyPln, wire: "\"PLN\""},
		{variant: ReceptionAgentInfoResponseCurrencyCad, wire: "\"CAD\""},
		{variant: ReceptionAgentInfoResponseCurrencyAud, wire: "\"AUD\""},
		{variant: ReceptionAgentInfoResponseCurrencyChf, wire: "\"CHF\""},
		{variant: ReceptionAgentInfoResponseCurrencySek, wire: "\"SEK\""},
		{variant: ReceptionAgentInfoResponseCurrencyNok, wire: "\"NOK\""},
		{variant: ReceptionAgentInfoResponseCurrencyDkk, wire: "\"DKK\""},
		{variant: ReceptionAgentInfoResponseCurrencyIsk, wire: "\"ISK\""},
		{variant: ReceptionAgentInfoResponseCurrencyCzk, wire: "\"CZK\""},
		{variant: ReceptionAgentInfoResponseCurrencyHuf, wire: "\"HUF\""},
		{variant: ReceptionAgentInfoResponseCurrencyRon, wire: "\"RON\""},
		{variant: ReceptionAgentInfoResponseCurrencyBgn, wire: "\"BGN\""},
		{variant: ReceptionAgentInfoResponseCurrencyJpy, wire: "\"JPY\""},
		{variant: ReceptionAgentInfoResponseCurrencyCny, wire: "\"CNY\""},
		{variant: ReceptionAgentInfoResponseCurrencyTwd, wire: "\"TWD\""},
		{variant: ReceptionAgentInfoResponseCurrencyInr, wire: "\"INR\""},
		{variant: ReceptionAgentInfoResponseCurrencyBrl, wire: "\"BRL\""},
		{variant: ReceptionAgentInfoResponseCurrencyMxn, wire: "\"MXN\""},
		{variant: ReceptionAgentInfoResponseCurrencyKrw, wire: "\"KRW\""},
		{variant: ReceptionAgentInfoResponseCurrencySgd, wire: "\"SGD\""},
		{variant: ReceptionAgentInfoResponseCurrencyHkd, wire: "\"HKD\""},
		{variant: ReceptionAgentInfoResponseCurrencyNzd, wire: "\"NZD\""},
		{variant: ReceptionAgentInfoResponseCurrencyZar, wire: "\"ZAR\""},
		{variant: ReceptionAgentInfoResponseCurrencyAed, wire: "\"AED\""},
		{variant: ReceptionAgentInfoResponseCurrencySar, wire: "\"SAR\""},
		{variant: ReceptionAgentInfoResponseCurrencyQar, wire: "\"QAR\""},
		{variant: ReceptionAgentInfoResponseCurrencyKwd, wire: "\"KWD\""},
		{variant: ReceptionAgentInfoResponseCurrencyIls, wire: "\"ILS\""},
		{variant: ReceptionAgentInfoResponseCurrencyTry, wire: "\"TRY\""},
		{variant: ReceptionAgentInfoResponseCurrencyThb, wire: "\"THB\""},
		{variant: ReceptionAgentInfoResponseCurrencyIdr, wire: "\"IDR\""},
		{variant: ReceptionAgentInfoResponseCurrencyMyr, wire: "\"MYR\""},
		{variant: ReceptionAgentInfoResponseCurrencyPhp, wire: "\"PHP\""},
		{variant: ReceptionAgentInfoResponseCurrencyVnd, wire: "\"VND\""},
		{variant: ReceptionAgentInfoResponseCurrencyBdt, wire: "\"BDT\""},
		{variant: ReceptionAgentInfoResponseCurrencyRub, wire: "\"RUB\""},
		{variant: ReceptionAgentInfoResponseCurrencyUah, wire: "\"UAH\""},
		{variant: ReceptionAgentInfoResponseCurrencyClp, wire: "\"CLP\""},
		{variant: ReceptionAgentInfoResponseCurrencyCop, wire: "\"COP\""},
		{variant: ReceptionAgentInfoResponseCurrencyArs, wire: "\"ARS\""},
		{variant: ReceptionAgentInfoResponseCurrencyPen, wire: "\"PEN\""},
		{variant: ReceptionAgentInfoResponseCurrencyEgp, wire: "\"EGP\""},
		{variant: ReceptionAgentInfoResponseCurrencyNgn, wire: "\"NGN\""},
		{variant: ReceptionAgentInfoResponseCurrencyKes, wire: "\"KES\""},
		{variant: ReceptionAgentInfoResponseCurrencyPkr, wire: "\"PKR\""},
		{variant: ReceptionAgentInfoResponseCurrencyGel, wire: "\"GEL\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ReceptionAgentInfoResponseCurrency
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
