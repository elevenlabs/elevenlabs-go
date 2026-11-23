package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAgentToolResponseAgentToolResponseStatusVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AgentToolResponseAgentToolResponseStatus
		wire    string
	}
	cases := []enumCase{
		{variant: AgentToolResponseAgentToolResponseStatusSuccess, wire: "\"success\""},
		{variant: AgentToolResponseAgentToolResponseStatusError, wire: "\"error\""},
		{variant: AgentToolResponseAgentToolResponseStatusBlocked, wire: "\"blocked\""},
		{variant: AgentToolResponseAgentToolResponseStatusSkipped, wire: "\"skipped\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AgentToolResponseAgentToolResponseStatus
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
