package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSipTrunkTransportEnumVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant SipTrunkTransportEnum
		wire    string
	}
	cases := []enumCase{
		{variant: SipTrunkTransportEnumAuto, wire: "\"auto\""},
		{variant: SipTrunkTransportEnumUdp, wire: "\"udp\""},
		{variant: SipTrunkTransportEnumTcp, wire: "\"tcp\""},
		{variant: SipTrunkTransportEnumTls, wire: "\"tls\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded SipTrunkTransportEnum
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

func TestSipTrunkTransportEnumVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant SipTrunkTransportEnum
		wire    string
	}
	cases := []enumCase{
		{variant: SipTrunkTransportEnumAuto, wire: "\"auto\""},
		{variant: SipTrunkTransportEnumUdp, wire: "\"udp\""},
		{variant: SipTrunkTransportEnumTcp, wire: "\"tcp\""},
		{variant: SipTrunkTransportEnumTls, wire: "\"tls\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded SipTrunkTransportEnum
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
