package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTelephonyProviderVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TelephonyProvider
		wire    string
	}
	cases := []enumCase{
		{variant: TelephonyProviderTwilio, wire: "\"twilio\""},
		{variant: TelephonyProviderSipTrunk, wire: "\"sip_trunk\""},
		{variant: TelephonyProviderExotel, wire: "\"exotel\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TelephonyProvider
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

func TestTelephonyProviderVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TelephonyProvider
		wire    string
	}
	cases := []enumCase{
		{variant: TelephonyProviderTwilio, wire: "\"twilio\""},
		{variant: TelephonyProviderSipTrunk, wire: "\"sip_trunk\""},
		{variant: TelephonyProviderExotel, wire: "\"exotel\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TelephonyProvider
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

func TestTelephonyProviderVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TelephonyProvider
		wire    string
	}
	cases := []enumCase{
		{variant: TelephonyProviderTwilio, wire: "\"twilio\""},
		{variant: TelephonyProviderSipTrunk, wire: "\"sip_trunk\""},
		{variant: TelephonyProviderExotel, wire: "\"exotel\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TelephonyProvider
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
