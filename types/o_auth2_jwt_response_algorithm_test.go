package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOAuth2JwtResponseAlgorithmVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant OAuth2JwtResponseAlgorithm
		wire    string
	}
	cases := []enumCase{
		{variant: OAuth2JwtResponseAlgorithmHs256, wire: "\"HS256\""},
		{variant: OAuth2JwtResponseAlgorithmHs384, wire: "\"HS384\""},
		{variant: OAuth2JwtResponseAlgorithmHs512, wire: "\"HS512\""},
		{variant: OAuth2JwtResponseAlgorithmRs256, wire: "\"RS256\""},
		{variant: OAuth2JwtResponseAlgorithmRs384, wire: "\"RS384\""},
		{variant: OAuth2JwtResponseAlgorithmRs512, wire: "\"RS512\""},
		{variant: OAuth2JwtResponseAlgorithmPs256, wire: "\"PS256\""},
		{variant: OAuth2JwtResponseAlgorithmPs384, wire: "\"PS384\""},
		{variant: OAuth2JwtResponseAlgorithmPs512, wire: "\"PS512\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded OAuth2JwtResponseAlgorithm
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
