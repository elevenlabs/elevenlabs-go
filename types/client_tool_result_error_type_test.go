package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClientToolResultErrorTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ClientToolResultErrorType
		wire    string
	}
	cases := []enumCase{
		{variant: ClientToolResultErrorTypeUserRejected, wire: "\"user_rejected\""},
		{variant: ClientToolResultErrorTypeExternalServer, wire: "\"external_server\""},
		{variant: ClientToolResultErrorTypeExternalClient, wire: "\"external_client\""},
		{variant: ClientToolResultErrorTypeCustomerAuth, wire: "\"customer_auth\""},
		{variant: ClientToolResultErrorTypeClientTimeout, wire: "\"client_timeout\""},
		{variant: ClientToolResultErrorTypeUnknown, wire: "\"unknown\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ClientToolResultErrorType
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
