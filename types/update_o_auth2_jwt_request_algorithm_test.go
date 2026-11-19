package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpdateOAuth2JwtRequestAlgorithmVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant UpdateOAuth2JwtRequestAlgorithm
		wire    string
	}
	cases := []enumCase{
		{variant: UpdateOAuth2JwtRequestAlgorithmHs256, wire: "\"HS256\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmHs384, wire: "\"HS384\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmHs512, wire: "\"HS512\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmRs256, wire: "\"RS256\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmRs384, wire: "\"RS384\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmRs512, wire: "\"RS512\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmPs256, wire: "\"PS256\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmPs384, wire: "\"PS384\""},
		{variant: UpdateOAuth2JwtRequestAlgorithmPs512, wire: "\"PS512\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded UpdateOAuth2JwtRequestAlgorithm
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
